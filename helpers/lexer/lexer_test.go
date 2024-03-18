package lexer

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

var (
	testLexerOps             = []string{"AND", "OR"}
	testLexerValidPredicates = map[string]string{
		"single":       "Monkey$\\ Token",
		"simple":       "Monkey$\\ Token AND ETH",
		"repeated-and": "ETH AND ETH",
		"repeated-or":  "ETH OR ETH",
		"nested":       "Monkey$\\ Token AND (ETH OR BTC)",
		"deep-nested":  "(Monkey$\\ Token AND ANT) AND (ETH OR (USDC AND BTC))",
	}
	testLexerInvalidPredicates = map[string]string{
		"single":              "ETH AND",
		"nested":              "Monkey$\\ Token AND (ETH OR",
		"deep-nested":         "(Monkey$\\ Token AND ANT)) AND (ETH OR (USDC AND BTC))",
		"malformed-simple":    "ETH AND BTC AND Monkey$\\ Token",
		"malformed-nested":    "ETH AND (Monkey$\\ Token AND ANT BTC)",
		"operator-as-literal": "AND OR ETH",
		"no-operator":         "ETH BTC ANT",
	}
	testLexerExpectedTokens = map[string]*Token{
		// Monkey$\\ Token
		"single": {Type: TokenTypeLiteral, Literal: "Monkey$ Token"},
		// Monkey$\\ Token AND ETH
		"simple": {
			Type: TokenTypeGroup,
			Childs: &Group{
				ID:       0,
				Level:    0,
				Operator: "AND",
				Tokens: []*Token{
					&Token{Type: TokenTypeLiteral, Literal: "Monkey$ Token"},
					&Token{Type: TokenTypeLiteral, Literal: "ETH"},
				},
				firstToken:  "Monkey$ Token",
				secondToken: "ETH",
			},
		},
		"repeated-and": {
			Type: TokenTypeGroup,
			Childs: &Group{
				ID:       0,
				Level:    0,
				Operator: "AND",
				Tokens: []*Token{
					&Token{Type: TokenTypeLiteral, Literal: "ETH"},
					&Token{Type: TokenTypeLiteral, Literal: "ETH"},
				},
				firstToken:  "ETH",
				secondToken: "ETH",
			},
		},
		"repeated-or": {
			Type: TokenTypeGroup,
			Childs: &Group{
				ID:       0,
				Level:    0,
				Operator: "OR",
				Tokens: []*Token{
					&Token{Type: TokenTypeLiteral, Literal: "ETH"},
					&Token{Type: TokenTypeLiteral, Literal: "ETH"},
				},
				firstToken:  "ETH",
				secondToken: "ETH",
			},
		},
		// Monkey$\\ Token AND (ETH OR BTC)
		"nested": {
			Type: TokenTypeGroup,
			Childs: &Group{
				ID:       0,
				Level:    0,
				Operator: "AND",
				Tokens: []*Token{
					&Token{Type: TokenTypeLiteral, Literal: "Monkey$ Token"},
					&Token{
						Type: TokenTypeGroup,
						Childs: &Group{
							ID:       1,
							Level:    1,
							Operator: "OR",
							Tokens: []*Token{
								&Token{Type: TokenTypeLiteral, Literal: "ETH"},
								&Token{Type: TokenTypeLiteral, Literal: "BTC"},
							},
							firstToken:  "ETH",
							secondToken: "BTC",
						},
					},
				},
				firstToken:  "Monkey$ Token",
				secondToken: "1",
			},
		},
		// (Monkey$\\ Token AND ANT) AND (ETH OR (USDC AND BTC))
		"deep-nested": {
			Type: TokenTypeGroup,
			Childs: &Group{
				ID:       0,
				Level:    0,
				Operator: "AND",
				Tokens: []*Token{
					&Token{
						Type: TokenTypeGroup,
						Childs: &Group{
							ID:       1,
							Level:    1,
							Operator: "AND",
							Tokens: []*Token{
								&Token{Type: TokenTypeLiteral, Literal: "Monkey$ Token"},
								&Token{Type: TokenTypeLiteral, Literal: "ANT"},
							},
							firstToken:  "Monkey$ Token",
							secondToken: "ANT",
						},
					},
					&Token{
						Type: TokenTypeGroup,
						Childs: &Group{
							ID:       2,
							Level:    1,
							Operator: "OR",
							Tokens: []*Token{
								&Token{Type: TokenTypeLiteral, Literal: "ETH"},
								&Token{
									Type: TokenTypeGroup,
									Childs: &Group{
										ID:       3,
										Level:    1,
										Operator: "AND",
										Tokens: []*Token{
											&Token{Type: TokenTypeLiteral, Literal: "USDC"},
											&Token{Type: TokenTypeLiteral, Literal: "BTC"},
										},
										firstToken:  "USDC",
										secondToken: "BTC",
									},
								},
							},
							firstToken:  "ETH",
							secondToken: "3",
						},
					},
				},
				firstToken:  "1",
				secondToken: "2",
			},
		},
	}
	testLexerExpectedValidSplitted = map[string][]string{
		// Monkey$\\ Token
		"single": {"Monkey$ Token"},
		// Monkey$\\ Token AND ETH
		"simple": {"Monkey$ Token", "AND", "ETH"},
		// ETH AND ETH
		"repeated-and": {"ETH", "AND", "ETH"},
		// ETH OR ETH
		"repeated-or": {"ETH", "OR", "ETH"},
		// Monkey$\\ Token AND (ETH OR BTC)
		"nested": {"Monkey$ Token", "AND", "(", "ETH", "OR", "BTC", ")"},
		// (Monkey$\\ Token AND ANT) AND (ETH OR (USDC AND BTC)
		"deep-nested": {"(", "Monkey$ Token", "AND", "ANT", ")", "AND", "(", "ETH", "OR", "(", "USDC", "AND", "BTC", ")", ")"},
	}
	testLexerExpectedInvalidSplitted = map[string][]string{
		// "Monkey$ Token"
		"single": {"ETH", "AND"},
		// "Monkey$\\ Token AND (ETH OR"
		"nested": {"Monkey$ Token", "AND", "(", "ETH", "OR"},
		// "(Monkey$\\ Token AND ANT)) AND (ETH OR (USDC AND BTC))"
		"deep-nested": {"(", "Monkey$ Token", "AND", "ANT", ")", ")", "AND", "(", "ETH", "OR", "(", "USDC", "AND", "BTC", ")", ")"},
		// "ETH AND BTC AND Monkey$\\ Token"
		"malformed-simple": {"ETH", "AND", "BTC", "AND", "Monkey$ Token"},
		// "ETH AND (Monkey$\\ Token AND ANT BTC)"
		"malformed-nested": {"ETH", "AND", "(", "Monkey$ Token", "AND", "ANT", "BTC", ")"},
		// "AND OR ETH"
		"operator-as-literal": {"AND", "OR", "ETH"},
		// "ETH BTC ANT"
		"no-operator": {"ETH", "BTC", "ANT"},
	}
)

func TestNewLexer(t *testing.T) {
	c := qt.New(t)
	c.Run("initialization", func(c *qt.C) {
		c.Assert(NewLexer(testLexerOps).ops, qt.CmpEquals(), testLexerOps)
	})
	c.Run("supported operator", func(c *qt.C) {
		lx := NewLexer(testLexerOps)
		// supported
		c.Assert(lx.SupportedOperator("AND"), qt.IsTrue)
		// not supported
		c.Assert(lx.SupportedOperator("XOR"), qt.IsFalse)
	})
}

func TestParse(t *testing.T) {
	c := qt.New(t)
	lx := NewLexer(testLexerOps)

	c.Run("valid predicates", func(c *qt.C) {
		for _, predicate := range testLexerValidPredicates {
			_, err := lx.Parse(predicate)
			c.Assert(err, qt.IsNil, qt.Commentf(predicate))
		}
	})
	c.Run("invalid predicates", func(c *qt.C) {
		for _, predicate := range testLexerInvalidPredicates {
			_, err := lx.Parse(predicate)
			c.Assert(err, qt.IsNotNil, qt.Commentf(predicate))
		}
	})
	c.Run("valid predicates format", func(c *qt.C) {
		for name, predicate := range testLexerValidPredicates {
			parsedToken, err := lx.Parse(predicate)
			c.Assert(err, qt.IsNil)
			expectedToken := testLexerExpectedTokens[name]
			c.Assert(parsedToken.Equals(expectedToken), qt.IsTrue,
				qt.Commentf("predicate %s, expected %s, got %s", name, expectedToken, parsedToken))
		}
	})
}

func Test_splitPredicate(t *testing.T) {
	c := qt.New(t)
	lx := NewLexer(testLexerOps)

	c.Run("valid predicates", func(c *qt.C) {
		for i, predicate := range testLexerValidPredicates {
			currentRes := lx.splitPredicate(predicate)
			expectedRes := testLexerExpectedValidSplitted[i]
			comment := qt.Commentf("%s [%s]: %v != %v", i, predicate, currentRes, expectedRes)
			c.Assert(currentRes, qt.HasLen, len(expectedRes), comment)
			c.Assert(currentRes, qt.ContentEquals, expectedRes, comment)
		}
	})

	c.Run("invalid predicates", func(c *qt.C) {
		for i, predicate := range testLexerInvalidPredicates {
			currentRes := lx.splitPredicate(predicate)
			expectedRes := testLexerExpectedInvalidSplitted[i]
			comment := qt.Commentf("%s: %v != %v", i, currentRes, expectedRes)
			c.Assert(currentRes, qt.HasLen, len(expectedRes), comment)
			c.Assert(currentRes, qt.ContentEquals, expectedRes, comment)
		}
	})
}

func Test_parseTokens(t *testing.T) {
	c := qt.New(t)
	lx := NewLexer(testLexerOps)

	c.Run("valid predicates", func(c *qt.C) {
		for i, predicate := range testLexerValidPredicates {
			predicateTokens := lx.splitPredicate(predicate)
			parsedToken, _, err := lx.parseTokens(0, predicateTokens)
			c.Assert(err, qt.IsNil, qt.Commentf(predicate))
			expectedToken := testLexerExpectedTokens[i]
			c.Assert(parsedToken.Equals(expectedToken), qt.IsTrue)
		}
	})

	c.Run("invalid predicates", func(c *qt.C) {
		for _, predicate := range testLexerInvalidPredicates {
			predicateTokens := lx.splitPredicate(predicate)
			_, _, err := lx.parseTokens(0, predicateTokens)
			c.Assert(err, qt.IsNotNil, qt.Commentf(predicate))
		}
	})
}
