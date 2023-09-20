package api

import (
	queries "github.com/vocdoni/census3/db/sqlc"
	"github.com/vocdoni/census3/lexer"
)

const (
	ANDTag = "AND"
	ORTag  = "OR"
)

var ValidOperatorsTags = []string{ANDTag, ORTag}

type StrategyOperators struct {
	db         *queries.Queries
	tokensInfo []*StrategyToken
}

func InitOperators(db *queries.Queries, info []*StrategyToken) *StrategyOperators {
	return &StrategyOperators{
		db:         db,
		tokensInfo: info,
	}
}

func (op *StrategyOperators) AND(iter *lexer.Iteration[map[string]string]) (map[string]string, error) {
	return nil, nil
}

func (op *StrategyOperators) OR(iter *lexer.Iteration[map[string]string]) (map[string]string, error) {
	return nil, nil
}

func (op *StrategyOperators) Map() []*lexer.Operator[map[string]string] {
	return []*lexer.Operator[map[string]string]{
		{Tag: ANDTag, Fn: op.AND},
		{Tag: ORTag, Fn: op.OR},
	}
}
