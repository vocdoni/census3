package lexer

import (
	"fmt"
	"testing"

	qt "github.com/frankban/quicktest"
)

var (
	testEvalDataSources = map[string][]int{
		"powsOf2": {1, 2, 4, 8, 16, 32},
		"odd":     {2, 4, 6, 8, 10, 12},
		"even":    {1, 3, 5, 7, 9, 11},
		"primes":  {1, 2, 3, 5, 7, 11},
	}
	testEvalValidOperatorsTags   = []string{"AND", "OR"}
	testEvalInvalidOperatorsTags = []string{"AND", "NAND", "XOR"}
	testEvalOperators            = []*Operator[[]int]{
		{Tag: "AND", Fn: testAndOperator},
		{Tag: "OR", Fn: testOrOperator},
	}
	testEvalValidCases = map[string][]int{
		"odd AND even":                          {},
		"even AND primes":                       {1, 3, 5, 7, 11},
		"odd AND (even OR primes)":              {2},
		"(odd OR powsOf2) AND (even OR primes)": {1, 2},
	}
	testEvalInvalidCases = []string{
		"odd AND naturals",
		"odd NAND even",
		"even NAND primes",
		"odd NAND (even XOR primes)",
		"(odd XOR powsOf2) NAND (even XOR primes)",
	}
)

func testAndOperator(iter *Iteration[[]int]) ([]int, error) {
	tagA, dataA := iter.A()
	if dataA == nil {
		var ok bool
		dataA, ok = testEvalDataSources[tagA]
		if !ok {
			return nil, fmt.Errorf("unrecognised data")
		}
	}
	tagB, dataB := iter.B()
	if dataB == nil {
		var ok bool
		dataB, ok = testEvalDataSources[tagB]
		if !ok {
			return nil, fmt.Errorf("unrecognised data")
		}
	}

	res := []int{}
	for _, a := range dataA {
		exist := false
		for _, b := range dataB {
			if a == b {
				exist = true
				break
			}
		}
		if exist {
			res = append(res, a)
		}
	}
	return res, nil
}

func testOrOperator(iter *Iteration[[]int]) ([]int, error) {
	tagA, dataA := iter.A()
	if dataA == nil {
		var ok bool
		dataA, ok = testEvalDataSources[tagA]
		if !ok {
			return nil, fmt.Errorf("unrecognised data")
		}
	}
	tagB, dataB := iter.B()
	if dataB == nil {
		var ok bool
		dataB, ok = testEvalDataSources[tagB]
		if !ok {
			return nil, fmt.Errorf("unrecognised data")
		}
	}

	res := append([]int{}, dataB...)
	for _, a := range dataA {
		exist := false
		for _, b := range dataB {
			if a == b {
				exist = true
				break
			}
		}
		if !exist {
			res = append(res, a)
		}
	}
	return res, nil
}

func TestEval(t *testing.T) {
	c := qt.New(t)

	c.Run("valid evaluations", func(c *qt.C) {
		lx := NewLexer(testEvalValidOperatorsTags)
		for predicate, results := range testEvalValidCases {
			token, err := lx.Parse(predicate)
			c.Assert(err, qt.IsNil)

			res, err := NewEval[[]int](testEvalOperators).EvalToken(token, nil)
			c.Assert(err, qt.IsNil)
			c.Assert(res, qt.ContentEquals, results)
		}
	})

	c.Run("invalid evaluations", func(c *qt.C) {
		lx := NewLexer(testEvalInvalidOperatorsTags)
		for _, predicate := range testEvalInvalidCases {
			token, err := lx.Parse(predicate)
			c.Assert(err, qt.IsNil)

			_, err = NewEval[[]int](testEvalOperators).EvalToken(token, nil)
			c.Assert(err, qt.IsNotNil)
		}

		token, err := lx.Parse("even AND odd")
		c.Assert(err, qt.IsNil)
		// force undefined data
		token.Childs.Tokens[1] = NewLiteralToken("naturals")
		_, err = NewEval[[]int](testEvalOperators).EvalToken(token, nil)
		c.Assert(err, qt.IsNotNil)
	})
}
