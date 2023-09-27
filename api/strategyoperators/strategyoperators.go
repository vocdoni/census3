package strategyoperators

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	queries "github.com/vocdoni/census3/db/sqlc"
	"github.com/vocdoni/census3/lexer"
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
		"tag":         ANDTag,
		"description": "logical operator that returns the common token holders between symbols with fixed balance to 1",
	},
	{
		"tag":         ANDSUMTag,
		"description": "logical operator that returns the common token holders between symbols with the sum of their balances on both tokens",
	},
	{
		"tag":         ANDMULTag,
		"description": "logical operator that returns the common token holders between symbols with the multiplication of their balances on both tokens",
	},
	{
		"tag":         ORTag,
		"description": "logical operator that returns the token holders of both symbols with fixed balance to 1",
	},
	{
		"tag":         ORSUMTag,
		"description": "logical operator that returns the token holders of both symbols with the sum of their balances on both tokens",
	},
	{
		"tag":         ORMULTag,
		"description": "logical operator that returns the token holders of both symbols with the multiplication of their balances on both tokens",
	},
}

type TokenInformation struct {
	ID         string
	ChainID    uint64
	MinBalance string
	Decimals   uint64
}

type StrategyIteration struct {
	decimals uint64
	Data     map[string]*big.Int
}

// StrategyOperators struct represents a custom set of predicate operators
// associated with a SQL database as data source. It brings access to SQL data
// inside the lexer evaluator operators.
type StrategyOperators struct {
	db         *queries.Queries
	tokensInfo map[string]*TokenInformation
}

// InitOperators function creates a new StrategyOperators struct with the db
// instance and info about tokens provided.
func InitOperators(db *queries.Queries, info map[string]*TokenInformation) *StrategyOperators {
	return &StrategyOperators{
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
func (op *StrategyOperators) tokenInfoBySymbol(symbol string) (common.Address, uint64, *big.Int, error) {
	tokenInfo, ok := op.tokensInfo[symbol]
	if !ok {
		return common.Address{}, 0, nil, fmt.Errorf("token symbol not found: %s", symbol)
	}
	minBalance := new(big.Int)
	if tokenInfo.MinBalance != "" {
		if _, ok := minBalance.SetString(tokenInfo.MinBalance, 10); !ok {
			return common.Address{}, 0, nil, fmt.Errorf("error decoding min balance for %s", symbol)
		}
	}
	return common.HexToAddress(tokenInfo.ID), tokenInfo.ChainID, minBalance, nil
}

// holdersBySymbol method queries to the database for the holders associated to
// the symbol provided. It calls to tokenInfoBySymbol first to get the token
// information by this symbol, and then queries the database.
func (op *StrategyOperators) holdersBySymbol(ctx context.Context, symbol string) (map[string]*big.Int, error) {
	// get token information by its symbol
	address, chainID, minBalance, err := op.tokenInfoBySymbol(symbol)
	if err != nil {
		return nil, err
	}
	// get token filtered information from the database
	rows, err := op.db.TokenHoldersByTokenIDAndChainIDAndMinBalance(ctx,
		queries.TokenHoldersByTokenIDAndChainIDAndMinBalanceParams{
			TokenID: address.Bytes(),
			ChainID: chainID,
			Balance: minBalance.Bytes(),
		})
	if err != nil {
		return nil, fmt.Errorf("error getting holders of token %s on chainID %d", symbol, chainID)
	}
	// decode the resulting addresses
	data := map[string]*big.Int{}
	for _, r := range rows {
		data[common.BytesToAddress(r.HolderID).String()] = new(big.Int).SetBytes(r.Balance)
	}
	return data, nil
}

// decimalsBySymbol method returns the decimals of the token identified by the
// symbol provided. If the token information does not exists, returns false.
func (op *StrategyOperators) decimalsBySymbol(symbol string) (uint64, bool) {
	if info, ok := op.tokensInfo[symbol]; ok {
		return info.Decimals, true
	}
	return 0, false
}
