package state

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestTokenTypeString(t *testing.T) {
	c := qt.New(t)
	for str, tt := range TokenTypeIntMap {
		c.Assert(tt.String(), qt.Equals, str)
	}
	var wrongTokenType TokenType = 1000
	c.Assert(wrongTokenType.String(), qt.Equals, TokenTypeStringMap[CONTRACT_TYPE_UNKNOWN])
}

func TestTokenTypeFromString(t *testing.T) {
	c := qt.New(t)
	for tt, str := range TokenTypeStringMap {
		c.Assert(TokenTypeFromString(str), qt.Equals, tt)
	}
	wrongTokenType := "wrongType"
	c.Assert(TokenTypeFromString(wrongTokenType), qt.Equals, CONTRACT_TYPE_UNKNOWN)
}
