package census

import (
	"context"
	"database/sql"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	queries "github.com/vocdoni/census3/db/sqlc"
	"github.com/vocdoni/census3/lexer"
	"github.com/vocdoni/census3/state"
	"github.com/vocdoni/census3/strategyoperators"
)

// CalculateStrategyHolders function returns the holders of a strategy and the
// total weight of the census. It also returns the total block number of the
// census, which is the sum of the strategy block number or the last block
// number of every token chain id. To calculate the census holders, it uses the
// supplied predicate to filter the token holders using a lexer and evaluator.
// The evaluator uses the strategy operators to evaluate the predicate which
// uses the database queries to get the token holders and their balances, and
// combines them.
func CalculateStrategyHolders(ctx context.Context, qdb *queries.Queries, w3p state.Web3Providers,
	id uint64, predicate string,
) (map[common.Address]*big.Int, *big.Int, uint64, error) {
	// init some variables to get computed in the following steps
	censusWeight := new(big.Int)
	strategyHolders := map[common.Address]*big.Int{}
	// parse the predicate
	lx := lexer.NewLexer(strategyoperators.ValidOperatorsTags)
	validPredicate, err := lx.Parse(predicate)
	if err != nil {
		return nil, nil, 0, err
	}
	// get strategy tokens from the database
	strategyTokens, err := qdb.TokensByStrategyID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil, 0, err
		}
		return nil, nil, 0, err
	}
	// any census strategy is identified by id created from the concatenation of
	// the block number, the strategy id and the anonymous flag. The creation of
	// censuses on specific block is not supported yet, so we need to get the
	// last block of every token chain id to sum them and get the total block
	// number, used to create the census id.
	totalTokensBlockNumber := uint64(0)
	for _, token := range strategyTokens {
		w3uri, exists := w3p[token.ChainID]
		if !exists {
			return nil, nil, 0, err
		}
		w3 := state.Web3{}
		if err := w3.Init(ctx, w3uri.URI, common.BytesToAddress(token.ID), state.TokenType(token.TypeID)); err != nil {
			return nil, nil, 0, err
		}
		currentBlockNumber, err := w3.LatestBlockNumber(ctx)
		if err != nil {
			return nil, nil, 0, err
		}
		totalTokensBlockNumber += currentBlockNumber
	}
	// if the current predicate is a literal, just query about its holders. If
	// it is a complex predicate, create a evaluator and evaluate the predicate
	if validPredicate.IsLiteral() {
		// get the strategy holders from the database
		holders, err := qdb.TokenHoldersByStrategyID(ctx, id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil, totalTokensBlockNumber, nil
			}
			return nil, nil, totalTokensBlockNumber, err
		}
		// parse holders addresses and balances
		for _, holder := range holders {
			holderAddr := common.BytesToAddress(holder.HolderID)
			holderBalance := new(big.Int).SetBytes(holder.Balance)
			if _, exists := strategyHolders[holderAddr]; !exists {
				strategyHolders[holderAddr] = holderBalance
				censusWeight = new(big.Int).Add(censusWeight, holderBalance)
			}
		}
	} else {
		// parse token information
		tokensInfo := map[string]*strategyoperators.TokenInformation{}
		for _, token := range strategyTokens {
			tokensInfo[token.Symbol] = &strategyoperators.TokenInformation{
				ID:         common.BytesToAddress(token.ID).String(),
				ChainID:    token.ChainID,
				MinBalance: new(big.Int).SetBytes(token.MinBalance).String(),
				Decimals:   token.Decimals,
			}
		}
		// init the operators and the predicate evaluator
		operators := strategyoperators.InitOperators(qdb, tokensInfo)
		eval := lexer.NewEval[*strategyoperators.StrategyIteration](operators.Map())
		// execute the evaluation of the predicate
		res, err := eval.EvalToken(validPredicate)
		if err != nil {
			return nil, nil, totalTokensBlockNumber, err
		}
		// parse the evaluation results
		for address, value := range res.Data {
			strategyHolders[common.HexToAddress(address)] = value
			censusWeight = new(big.Int).Add(censusWeight, value)
		}
	}
	// if no holders found, return an error
	if len(strategyHolders) == 0 {
		return nil, nil, totalTokensBlockNumber, nil
	}
	return strategyHolders, censusWeight, totalTokensBlockNumber, nil
}
