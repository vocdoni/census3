package strategyoperators

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	queries "github.com/vocdoni/census3/db/sqlc"
	"github.com/vocdoni/census3/lexer"
)

const (
	// ANDTag constant contains the string that identifies the AND operator
	ANDTag = "AND"
	// ORTag constant contains the string that identifies the OR operator
	ORTag = "OR"
)

// ValidOperatorsTags variable contains the supported operator tags
var ValidOperatorsTags = []string{ANDTag, ORTag}

// ValidOperators variable contains the information of the supported operators
var ValidOperators = []map[string]string{
	{
		"tag":         ANDTag,
		"description": "logical operator that returns the common token holders between symbols with fixed balance to 1",
	},
	{
		"tag":         ORTag,
		"description": "logical operator that returns the token holders of both symbols with fixed balance to 1",
	},
}

type TokenInformation struct {
	ID         string
	ChainID    uint64
	MinBalance string
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
func (op *StrategyOperators) Map() []*lexer.Operator[map[string]string] {
	return []*lexer.Operator[map[string]string]{
		{Tag: ANDTag, Fn: op.AND},
		{Tag: ORTag, Fn: op.OR},
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

// AND method defines the AND operator, which returns the common items between
// the list of holders provided in each iteration. Like any definition of
// lexer.Operator, it receives an lexer.Iterarion struct, which helps to get
// both tokens symbols or the results of previous iterations.
func (op *StrategyOperators) AND(iter *lexer.Iteration[map[string]string]) (map[string]string, error) {
	interalCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	// get information about the current operation
	symbolA, dataA := iter.A()
	symbolB, dataB := iter.B()
	// no results for any component from previous operation, so get both data
	// from the database using an AND SQL query
	if len(dataA) == 0 && len(dataB) == 0 {
		// get both tokens information by their symbols
		addressA, chainIDA, minBalanceA, err := op.tokenInfoBySymbol(symbolA)
		if err != nil {
			return nil, err
		}
		addressB, chainIDB, minBalanceB, err := op.tokenInfoBySymbol(symbolB)
		if err != nil {
			return nil, err
		}
		// run the AND query
		rows, err := op.db.AndQueryHolders(interalCtx, queries.AndQueryHoldersParams{
			TokenIDA:    addressA.Bytes(),
			ChainIDA:    chainIDA,
			TokenIDB:    addressB.Bytes(),
			ChainIDB:    chainIDB,
			MinBalanceA: minBalanceA.Bytes(),
			MinBalanceB: minBalanceB.Bytes(),
		})
		if err != nil {
			return nil, fmt.Errorf("error getting holders of tokens %s (chainID: %d) and %s (chainID: %d)",
				symbolA, chainIDA, symbolB, chainIDB)
		}
		// decode the results and return them
		res := map[string]string{}
		for _, r := range rows {
			res[common.BytesToAddress(r).String()] = "1"
		}
		return res, nil
	}
	// if the dataA is empty (does not contains results of previous operarion),
	// fill its data with the records of the database
	if len(dataA) == 0 {
		// get token information by its symbol
		address, chainID, minBalance, err := op.tokenInfoBySymbol(symbolA)
		if err != nil {
			return nil, err
		}
		// get token filtered information from the database
		rows, err := op.db.TokenHoldersByTokenIDAndChainIDAndMinBalance(interalCtx,
			queries.TokenHoldersByTokenIDAndChainIDAndMinBalanceParams{
				TokenID: address.Bytes(),
				ChainID: chainID,
				Balance: minBalance.Bytes(),
			})
		if err != nil {
			return nil, fmt.Errorf("error getting holders of token %s on chainID %d", symbolA, chainID)
		}
		// decode the resulting addresses
		dataA = map[string]string{}
		for _, r := range rows {
			dataA[common.BytesToAddress(r).String()] = "1"
		}
	}
	// if the dataB is empty (does not contains results of previous operarion),
	// fill its data with the records of the database
	if len(dataB) == 0 {
		// get token information by its symbol
		address, chainID, minBalance, err := op.tokenInfoBySymbol(symbolB)
		if err != nil {
			return nil, err
		}
		// get token filtered information from the database
		rows, err := op.db.TokenHoldersByTokenIDAndChainIDAndMinBalance(interalCtx,
			queries.TokenHoldersByTokenIDAndChainIDAndMinBalanceParams{
				TokenID: address.Bytes(),
				ChainID: chainID,
				Balance: minBalance.Bytes(),
			})
		if err != nil {
			return nil, fmt.Errorf("error getting holders of token %s on chainID %d", symbolB, chainID)
		}
		// decode the resulting addresses
		dataB = map[string]string{}
		for _, r := range rows {
			dataB[common.BytesToAddress(r).String()] = "1"
		}
	}
	// when both data sources are filled, do the intersection of both lists.
	res := map[string]string{}
	for addressA, value := range dataA {
		for addressB := range dataB {
			if addressA == addressB {
				res[addressA] = value
				break
			}
		}
	}
	return res, nil
}

// OR method defines the OR operator, which returns a list with the items of
// both lists of holders provided in each iteration. Like any definition of
// lexer.Operator, it receives an lexer.Iterarion struct, which helps to get
// both tokens symbols or the results of previous iterations.
func (op *StrategyOperators) OR(iter *lexer.Iteration[map[string]string]) (map[string]string, error) {
	interalCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	// get information about the current operation
	symbolA, dataA := iter.A()
	symbolB, dataB := iter.B()
	// no results for any component from previous operation, so get both data
	// from the database using an OR SQL query
	if len(dataA) == 0 && len(dataB) == 0 {
		// get both tokens information by their symbols
		addressA, chainIDA, minBalanceA, err := op.tokenInfoBySymbol(symbolA)
		if err != nil {
			return nil, err
		}
		addressB, chainIDB, minBalanceB, err := op.tokenInfoBySymbol(symbolB)
		if err != nil {
			return nil, err
		}
		// run the OR query
		rows, err := op.db.OrQueryHolders(interalCtx, queries.OrQueryHoldersParams{
			TokenIDA:    addressA.Bytes(),
			ChainIDA:    chainIDA,
			TokenIDB:    addressB.Bytes(),
			ChainIDB:    chainIDB,
			MinBalanceA: minBalanceA.Bytes(),
			MinBalanceB: minBalanceB.Bytes(),
		})
		if err != nil {
			return nil, fmt.Errorf("error getting holders of tokens %s (chainID: %d) and %s (chainID: %d)",
				symbolA, chainIDA, symbolB, chainIDB)
		}
		// decode the results and return them
		res := map[string]string{}
		for _, r := range rows {
			res[common.BytesToAddress(r).String()] = "1"
		}
		return res, nil
	}
	// if the dataA is empty (does not contains results of previous operarion),
	// fill its data with the records of the database
	if len(dataA) == 0 {
		// get token information by its symbol
		address, chainID, minBalance, err := op.tokenInfoBySymbol(symbolA)
		if err != nil {
			return nil, err
		}
		// get token filtered information from the database
		rows, err := op.db.TokenHoldersByTokenIDAndChainIDAndMinBalance(interalCtx,
			queries.TokenHoldersByTokenIDAndChainIDAndMinBalanceParams{
				TokenID: address.Bytes(),
				ChainID: chainID,
				Balance: minBalance.Bytes(),
			})
		if err != nil {
			return nil, fmt.Errorf("error getting holders of token %s on chainID %d", symbolA, chainID)
		}
		// decode the resulting addresses
		dataA := map[string]string{}
		for _, r := range rows {
			dataA[common.BytesToAddress(r).String()] = "1"
		}
	}
	// if the dataB is empty (does not contains results of previous operarion),
	// fill its data with the records of the database
	if len(dataB) == 0 {
		// get token information by its symbol
		address, chainID, minBalance, err := op.tokenInfoBySymbol(symbolB)
		if err != nil {
			return nil, err
		}
		// get token filtered information from the database
		rows, err := op.db.TokenHoldersByTokenIDAndChainIDAndMinBalance(interalCtx,
			queries.TokenHoldersByTokenIDAndChainIDAndMinBalanceParams{
				TokenID: address.Bytes(),
				ChainID: chainID,
				Balance: minBalance.Bytes(),
			})
		if err != nil {
			return nil, fmt.Errorf("error getting holders of token %s on chainID %d", symbolB, chainID)
		}
		// decode the resulting addresses
		dataB := map[string]string{}
		for _, r := range rows {
			dataB[common.BytesToAddress(r).String()] = "1"
		}
	}
	// when both data sources are filled, do the union of both lists.
	res := dataA
	for addressB, value := range dataB {
		if _, ok := dataA[addressB]; !ok {
			res[addressB] = value
		}
	}
	return res, nil
}
