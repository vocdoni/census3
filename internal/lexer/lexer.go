package lexer

import (
	"fmt"
	"slices"
)

// Lexer struct helps to parse a predicated with a predefined operators,
// generating a iterable struct to eval the it then.
type Lexer struct {
	ops []string
}

// NewLexer returns a initialized lexer with the operators tags provided
func NewLexer(ops []string) *Lexer {
	return &Lexer{ops: ops}
}

// SupportedOperator method return if the tag provided is a operator supported
// by the current lexer
func (l *Lexer) SupportedOperator(tag string) bool {
	return slices.Contains(l.ops, tag)
}

// Parse method tokenize the predicate provided a decode its groups levels
// recursively returning a Part
func (l *Lexer) Parse(predicate string) (*Token, error) {
	tokens := l.splitPredicate(predicate)
	group, _, err := l.parseTokens(0, tokens)
	if err != nil {
		return nil, fmt.Errorf("error parsing predicate: %w", err)
	}
	return group, nil
}

// splitPredicate function iterates over the characters of the provided predicate
// grouping them in tokens, including special characters and operators.
func (l *Lexer) splitPredicate(predicate string) []string {
	tokens := []string{}
	currentToken := ""
	commit := func() {
		if currentToken != "" {
			tokens = append(tokens, currentToken)
			currentToken = ""
		}
	}
	// iterate over predicate bytes
	for i := 0; i < len(predicate); i++ {
		switch bChar := predicate[i]; bChar {
		case bScape:
			// if backslash is found, include in the current token the following
			// char and continiu from it.
			currentToken += string(predicate[i+1])
			i++
			continue
		case bStartGroup, bEndGroup:
			// if the char is a bStartGroup or bEndGroup commit the current
			// token and include as single token
			commit()
			tokens = append(tokens, string(bChar))
		case bSpace:
			// if char is a space
			commit()
		default:
			currentToken += string(bChar)
		}
	}
	commit()
	return tokens
}

// parseTokens method parses the level provided, starting by the tokens also
// provided. It calls itself recursively when it finds a sub level group. It
// returns a token and the number of parts decoded (offset). If the token is a
// group, it also returns if the group is correctly closed. It checks that
// each Group and Token is formed correctly.
func (l *Lexer) parseTokens(level int, tokens []string) (*Token, int, error) {
	nTokens := len(tokens)
	// if no tokens provided or just two, return an error
	if nTokens == 0 || nTokens == 2 {
		return nil, -1, fmt.Errorf("bad formatted tokens")
	}
	// if only one token is provided and is not a special character return it
	// without error
	if nTokens == 1 && tokens[0] != startGroup && tokens[0] != endGroup {
		return NewLiteralToken(tokens[0]), 1, nil
	}
	// init a the resulting Token as a Group with and empty Group for the
	// current level
	token := NewGroupToken(NewEmptyGroup(level))
	// iterate over predicate parts provided
	for i := 0; i < len(tokens); i++ {
		part := tokens[i]
		if l.SupportedOperator(part) {
			// if the part is a supported operator, store it into the current
			// token
			if token.Childs.Operator != "" {
				return nil, -1, fmt.Errorf("duplicated operator at level %d", level)
			}
			token.Childs.Operator = part
		} else if part == startGroup {
			// if the part is a start of a group, call this method to decode
			// the next level with the same tokens skipping the current one,
			// then forward the number of tokens returned as offset and store
			// the resulting token
			childToken, offset, err := l.parseTokens(level+1, tokens[i+1:])
			if err != nil {
				return nil, offset, err
			}
			if err := token.Childs.AddToken(childToken); err != nil {
				return nil, -1, fmt.Errorf("error parsing level %d token: %w", level, err)
			}
			i += offset
		} else if part == endGroup {
			// if the part is the end of the group, return the current token
			// with the correct offset. If the current token has not 2 child
			// tokens or has not registered operator, return an error.
			if !token.Childs.Complete() {
				return nil, -1, fmt.Errorf("group at level %d no completed", level)
			}
			return token, i + 1, nil
		} else {
			// if the part is a literal token, store it into the current token
			// as child token
			if err := token.Childs.AddToken(NewLiteralToken(part)); err != nil {
				return nil, -1, fmt.Errorf("error parsing level %d token: %w", level, err)
			}
		}
	}
	// If the current token childs is not completed, return an error. Else if
	// return the current token and the numbers of parts provided as offset.
	if !token.Childs.Complete() {
		return nil, -1, fmt.Errorf("group at level %d no completed", level)
	}
	return token, len(tokens), nil
}
