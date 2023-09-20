package lexer

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
)

// TokenType custom type defines the Token type attribute.
type TokenType int

const (
	// TokenTypeLiteral is the type value assigned to literal tokens
	TokenTypeLiteral TokenType = iota
	// TokenTypeGroup is the type value assigned to group tokens
	TokenTypeGroup
)

// Token struct envolves any type of token (literal or group).
type Token struct {
	Type    TokenType
	Literal string
	Childs  *Group
}

// MarshalJSON method helps to encode the current token into a JSON with just
// the important information, discarding internal information.
func (t *Token) MarshalJSON() ([]byte, error) {
	if t.IsLiteral() {
		// if it is a literal token include just the token symbol
		return json.Marshal(map[string]any{"literal": t.Literal})
	}
	// if it is a group token, include just the child tokens
	return json.Marshal(map[string]any{"childs": t.Childs})
}

// IsLiteral method returns if the current token is a literal token (or a group
// token).
func (t *Token) IsLiteral() bool {
	return t.Type == TokenTypeLiteral
}

// String function returns the string version of the current token. It can be
// used as predicated input to generate the same token.
func (t *Token) String() string {
	// if it is literal envolves between literals delimiters
	if t.IsLiteral() && t.Literal != "" {
		return fmt.Sprintf("'%s'", t.Literal)
	}
	// if it is group and it has not childs, return empty
	if t.Childs == nil || len(t.Childs.Tokens) < 2 {
		return ""
	}
	// return string group
	return fmt.Sprintf("(%s %s %s)",
		t.Childs.Tokens[t.Childs.firstToken],
		t.Childs.Operator,
		t.Childs.Tokens[t.Childs.secondToken],
	)
}

// Equals method return if the a token is equal to b token. It compares the
// between both tokens string versions.
func (a *Token) Equals(b *Token) bool {
	return a.String() == b.String()
}

// AllGroups method returns all the group tokens from the current token to the
// last child, including it. The method calls itself recursively if the current
// token is a group token.
func (t *Token) AllGroups() []*Group {
	if t.IsLiteral() || t.Childs == nil {
		return nil
	}
	res := []*Group{t.Childs}
	for _, t := range t.Childs.Tokens {
		if !t.IsLiteral() {
			res = append(res, t.AllGroups()...)
		}
	}
	return res
}

// AllLiterals method returns all the literal tokens from the current token to
// the last child, including it. The method calls itself recursively if the
// current token is a group token.
func (t *Token) AllLiterals() []string {
	if t.IsLiteral() {
		return []string{t.Literal}
	}
	if t.Childs == nil {
		return nil
	}
	res := []string{}
	commit := func(i string) {
		exists := false
		for _, r := range res {
			if r == i {
				exists = true
				break
			}
		}
		if !exists {
			res = append(res, i)
		}
	}

	for _, t := range t.Childs.Tokens {
		if t.IsLiteral() {
			commit(t.Literal)
		} else {
			for _, l := range t.AllLiterals() {
				commit(l)
			}
		}
	}
	return res
}

// NewLiteralToken function returns a new literal Token inflated with the
// provided value.
func NewLiteralToken(literal string) *Token {
	return &Token{
		Literal: literal,
		Type:    TokenTypeLiteral,
	}
}

// NewGroupToken function returns a new group Token inflated with the
// provided group. If it is nil, the function inflate it with a emptye group at
// level 0. It set the resulting Token Literal attribute as the string group ID.
func NewGroupToken(group *Group) *Token {
	if group == nil {
		group = NewEmptyGroup(0)
	}
	return &Token{
		Literal: fmt.Sprint(group.ID),
		Childs:  group,
		Type:    TokenTypeGroup,
	}
}

// Group struct envolves a group of tokens, identified by its ID. Any group
// will contain an operator tag, the level of the group, and the tokens of the
// group.
type Group struct {
	ID          int
	Operator    string
	Level       int
	Tokens      map[string]*Token
	firstToken  string
	secondToken string
}

// NewEmptyGroup function returns an empty Group for the level provided. The
// group ID is randomly generated, and it is lower than or equal to maxGroupID.
func NewEmptyGroup(level int) *Group {
	return &Group{
		ID:     rand.Intn(maxGroupID),
		Level:  level,
		Tokens: make(map[string]*Token),
	}
}

// MarshalJSON method helps to encode the current group into a JSON with just
// the important information, discarding internal information.
func (g *Group) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"operator": g.Operator,
		"tokens": []*Token{
			g.Tokens[g.firstToken],
			g.Tokens[g.secondToken],
		},
	})
}

// AddToken method assign the token provided to the current group. If the group
// already contains the provided token, skip without error. If the group is
// already completed (already has two tokens), raises an error. The method
// assigns the provided token as first token if the group has not one. If it
// already has a first token, the provided token will be the second one.
func (g *Group) AddToken(t *Token) error {
	if _, ok := g.Tokens[t.Literal]; ok {
		return nil
	}
	if len(g.Tokens) == 2 {
		return fmt.Errorf("current group already has two tokens")
	}
	if g.firstToken == "" {
		g.firstToken = t.Literal
	} else {
		g.secondToken = t.Literal
	}
	g.Tokens[t.Literal] = t
	return nil
}

// Complete method returns if the current group is already completed which means
// that it has an operator and two tokens (first and second one).
func (g *Group) Complete() bool {
	return g.Operator != "" && g.firstToken != "" && g.secondToken != "" && len(g.Tokens) == 2
}

func ScapeTokenSymbol(symbol string) string {
	symbol = strings.ReplaceAll(symbol, space, scape+space)
	symbol = strings.ReplaceAll(symbol, startGroup, scape+startGroup)
	return strings.ReplaceAll(symbol, endGroup, scape+endGroup)
}
