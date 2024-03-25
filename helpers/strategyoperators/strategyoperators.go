package strategyoperators

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	queries "github.com/vocdoni/census3/db/sqlc"
	"github.com/vocdoni/census3/helpers/lexer"
)

const (
	// ANDTag constant contains the string that identifies the AND operator
	ANDTag    = "AND"
	ANDSUMTag = "AND:sum"
	ANDMULTag = "AND:mul"
	// ORTag constant contains the string that identifies the OR operator
	ORTag    = "OR"
	ORSUMTag = "OR:sum"
	ORMULTag = "OR:mul"
)

// ValidOperatorsTags variable contains the supported operator tags
var ValidOperatorsTags = []string{
	ANDTag,
	ANDSUMTag,
	ANDMULTag,
	ORTag,
	ORSUMTag,
	ORMULTag,
}

// ValidOperators variable contains the information of the supported operators
var ValidOperators = []map[string]string{
	{
		"tag": ANDTag,
		"description": "AND logical operator that returns the common token " +
			"holders between symbols with fixed balance to 1",
	},
	{
		"tag": ANDSUMTag,
		"description": "AND:sum logical operator that returns the common token " +
			"holders between symbols with the sum of their balances on both tokens",
	},
	{
		"tag": ANDMULTag,
		"description": "AND:mul logical operator that returns the common token " +
			"holders between symbols with the multiplication of their balances on both tokens",
	},
	{
		"tag": ORTag,
		"description": "OR logical operator that returns the token holders " +
			"of both symbols with fixed balance to 1",
	},
	{
		"tag": ORSUMTag,
		"description": "OR:sum logical operator that returns the token holders " +
			"of both symbols with the sum of their balances on both tokens",
	},
	{
		"tag": ORMULTag,
		"description": "OR:mul logical operator that returns the token holders " +
			"of both symbols with the multiplication of their balances on both tokens",
	},
}

// TokenInformation struct represents the information of a token that is used
// by strategy operators in a predicate evaluation.
type TokenInformation struct {
	ID         string
	ChainID    uint64
	MinBalance string
	Decimals   uint64
	ExternalID string
}

// StrategyIteration struct represents the data that is passed to the operators
// in every iteration of the evaluation process. It contains the data of the
// token holders of the current symbol or previous iterarion result and number
// the decimals of the balances.
type StrategyIteration struct {
	decimals uint64
	Data     map[string]*big.Int
}

// StrategyOperators struct represents a custom set of predicate operators
// associated with a SQL database as data source. It brings access to SQL data
// inside the lexer evaluator operators.
type StrategyOperators struct {
	ctx        context.Context
	db         *queries.Queries
	tokensInfo map[string]*TokenInformation
}

// InitOperators function creates a new StrategyOperators struct with the db
// instance and info about tokens provided.
func InitOperators(ctx context.Context, db *queries.Queries, info map[string]*TokenInformation) *StrategyOperators {
	return &StrategyOperators{
		ctx:        ctx,
		db:         db,
		tokensInfo: info,
	}
}

// Map method return the current operators in a map, associated with theirs
// operator tag.
func (op *StrategyOperators) Map() []*lexer.Operator[*StrategyIteration] {
	return []*lexer.Operator[*StrategyIteration]{
		{Tag: ANDTag, Fn: op.AND()},
		{Tag: ANDSUMTag, Fn: op.AND_SUM()},
		{Tag: ANDMULTag, Fn: op.AND_MUL()},
		{Tag: ORTag, Fn: op.OR()},
		{Tag: ORSUMTag, Fn: op.OR_SUM()},
		{Tag: ORMULTag, Fn: op.OR_MUL()},
	}
}

// tokenInfoBySymbol method checks if the current token information includes
// the information related to the token identified by the symbol provided. It
// also decodes the address of the token and the min balance (by default 0). If
// it does not contains any related token information or the decoding process
// fails, returns an error.
func (op *StrategyOperators) tokenInfoBySymbol(symbol string) (common.Address, uint64, *big.Int, string, error) {
	tokenInfo, ok := op.tokensInfo[symbol]
	if !ok {
		return common.Address{}, 0, nil, "", fmt.Errorf("token symbol not found: %s", symbol)
	}
	minBalance := new(big.Int)
	if tokenInfo.MinBalance != "" {
		if _, ok := minBalance.SetString(tokenInfo.MinBalance, 10); !ok {
			return common.Address{}, 0, nil, "", fmt.Errorf("error decoding min balance for %s", symbol)
		}
	}
	return common.HexToAddress(tokenInfo.ID), tokenInfo.ChainID, minBalance, tokenInfo.ExternalID, nil
}

// decimalsBySymbol method returns the decimals of the token identified by the
// symbol provided. If the token information does not exists, returns false.
func (op *StrategyOperators) decimalsBySymbol(symbol string) (uint64, bool) {
	if info, ok := op.tokensInfo[symbol]; ok {
		return info.Decimals, true
	}
	return 0, false
}

// holdersBySymbol method queries to the database for the holders associated to
// the symbol provided. It calls to tokenInfoBySymbol first to get the token
// information by this symbol, and then queries the database.
func (op *StrategyOperators) holdersBySymbol(ctx context.Context, symbol string) (map[string]*big.Int, error) {
	// get token information by its symbol
	address, chainID, minBalance, externalID, err := op.tokenInfoBySymbol(symbol)
	if err != nil {
		return nil, err
	}
	// get token filtered information from the database
	rows, err := op.db.TokenHoldersByTokenIDAndChainIDAndMinBalance(ctx,
		queries.TokenHoldersByTokenIDAndChainIDAndMinBalanceParams{
			TokenID:    address.Bytes(),
			ChainID:    chainID,
			Balance:    minBalance.String(),
			ExternalID: externalID,
		})
	if err != nil {
		return nil, fmt.Errorf("error getting holders of token %s on chainID %d", symbol, chainID)
	}
	if len(rows) == 0 {
		return nil, fmt.Errorf("no holders for the token %s on chainID %d", symbol, chainID)
	}
	// decode the resulting addresses
	data := map[string]*big.Int{}
	for _, r := range rows {
		balance, ok := new(big.Int).SetString(r.Balance, 10)
		if !ok {
			return nil, fmt.Errorf("error decoding balance of token %s on chainID %d", symbol, chainID)
		}
		data[common.BytesToAddress(r.HolderID).String()] = balance
	}
	return data, nil
}
