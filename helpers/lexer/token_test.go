package lexer

import (
	"fmt"
	"testing"

	qt "github.com/frankban/quicktest"
)

const (
	testOperatorAND = "AND"
	testOperatorOR  = "OR"
)

func TestToken(t *testing.T) {
	c := qt.New(t)

	c.Run("invalid token", func(c *qt.C) {
		it := &Token{}
		c.Assert(it.IsLiteral(), qt.IsTrue)
		c.Assert(it.String(), qt.Equals, "")
		it.Type = TokenTypeGroup
		c.Assert(it.IsLiteral(), qt.IsFalse)
		c.Assert(it.String(), qt.Equals, "")
	})

	c.Run("new literal token", func(c *qt.C) {
		tl := NewLiteralToken("test")
		c.Assert(tl.Type, qt.Equals, TokenTypeLiteral)
		c.Assert(tl.IsLiteral(), qt.IsTrue)
		c.Assert(tl.String(), qt.Equals, "test")
	})

	c.Run("new group token", func(c *qt.C) {
		tg := NewGroupToken(nil)
		c.Assert(tg.Childs.Complete(), qt.IsFalse)
		c.Assert(tg.Type, qt.Equals, TokenTypeGroup)
		c.Assert(tg.IsLiteral(), qt.IsFalse)
		c.Assert(fmt.Sprint(tg.Childs.ID), qt.Equals, tg.Literal)
		c.Assert(tg.String(), qt.Equals, "")

		tg.Childs.Operator = testOperatorAND
		c.Assert(tg.Childs.AddToken(NewLiteralToken("ETH")), qt.IsNil)
		c.Assert(tg.Childs.AddToken(NewLiteralToken("BTC")), qt.IsNil)
		c.Assert(tg.String(), qt.Equals, "(ETH AND BTC)")
		c.Assert(tg.Childs.Complete(), qt.IsTrue)
		// full group
		c.Assert(tg.Childs.AddToken(NewLiteralToken("BTC")), qt.IsNotNil)
		// full group
		c.Assert(tg.Childs.AddToken(NewLiteralToken("USDC")), qt.IsNotNil)
	})

	c.Run("token childs literals", func(c *qt.C) {
		c.Assert(NewGroupToken(nil).AllLiterals(), qt.ContentEquals, []string{})
		c.Assert((&Token{Type: TokenTypeGroup}).AllLiterals(), qt.IsNil)
		c.Assert(NewLiteralToken("test").AllLiterals(), qt.ContentEquals, []string{"test"})

		subChildGroup := NewEmptyGroup(2)
		subChildGroup.Operator = testOperatorOR
		c.Assert(subChildGroup.AddToken(NewLiteralToken("BTC")), qt.IsNil)
		c.Assert(subChildGroup.AddToken(NewLiteralToken("ETH")), qt.IsNil)

		childGroup := NewEmptyGroup(1)
		childGroup.Operator = testOperatorOR
		c.Assert(childGroup.AddToken(NewLiteralToken("ETH")), qt.IsNil)
		c.Assert(childGroup.AddToken(NewGroupToken(subChildGroup)), qt.IsNil)

		token := NewGroupToken(NewEmptyGroup(0))
		token.Childs.Operator = testOperatorAND
		c.Assert(token.Childs.AddToken(NewLiteralToken("ANT")), qt.IsNil)
		c.Assert(token.Childs.AddToken(NewGroupToken(childGroup)), qt.IsNil)

		literals := token.AllLiterals()
		c.Assert(literals, qt.ContentEquals, []string{"ANT", "ETH", "BTC"})
	})

	c.Run("token childs group", func(c *qt.C) {
		c.Assert(NewLiteralToken("test").AllGroups(), qt.IsNil)

		token := NewGroupToken(NewEmptyGroup(12))
		tChilds := token.AllGroups()
		c.Assert(tChilds, qt.HasLen, 1)
		c.Assert(tChilds[0].ID, qt.Equals, token.Childs.ID)

		childGroup := NewEmptyGroup(1)
		childGroup.Operator = testOperatorOR
		c.Assert(childGroup.AddToken(NewLiteralToken("ETH")), qt.IsNil)
		c.Assert(childGroup.AddToken(NewLiteralToken("BTC")), qt.IsNil)

		token.Childs.Operator = testOperatorAND
		c.Assert(token.Childs.AddToken(NewLiteralToken("ANT")), qt.IsNil)
		c.Assert(token.Childs.AddToken(NewGroupToken(childGroup)), qt.IsNil)
		tChilds = token.AllGroups()
		c.Assert(tChilds, qt.HasLen, 2)
		c.Assert(tChilds[0].ID, qt.Equals, token.Childs.ID)
		c.Assert(tChilds[1].ID, qt.Equals, childGroup.ID)
	})
}

func TestGroup(t *testing.T) {
	c := qt.New(t)

	c.Run("new group", func(c *qt.C) {
		group := NewEmptyGroup(12)
		c.Assert(group.Level, qt.Equals, 12)
		c.Assert(group.Operator, qt.Equals, "")
		c.Assert(group.Tokens, qt.HasLen, 0)

		altGroup := NewEmptyGroup(12)
		c.Assert(group.ID, qt.Not(qt.Equals), altGroup.ID)
	})
}

func TestScapeTokenSymbol(t *testing.T) {
	c := qt.New(t)

	symbol := "A (B)"
	expected := "A\\ \\(B\\)"
	c.Assert(ScapeTokenSymbol(symbol), qt.Equals, expected)
}
