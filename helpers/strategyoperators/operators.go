package strategyoperators

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	queries "github.com/vocdoni/census3/db/sqlc"
	"github.com/vocdoni/census3/helpers/lexer"
)

// AND method returns a AND operator function that can be used in a strategy
// evaluation. The AND operator returns the common token holders between symbols
// database information or previous operations results. It applies a fixed
// balance to 1 to indicate the membership.
func (op *StrategyOperators) AND() func(iter *lexer.Iteration[*StrategyIteration]) (*StrategyIteration, error) {
	return func(iter *lexer.Iteration[*StrategyIteration]) (*StrategyIteration, error) {
		data, decimals, err := op.andOperator(iter)
		if err != nil {
			return nil, err
		}
		return &StrategyIteration{
			decimals: decimals,
			Data:     membershipCombinator(data),
		}, nil
	}
}

// AND_SUM method returns a AND operator function that can be used in a
// strategy evaluation. The AND operator returns the common token holders
// between symbols database information or previous operations results. It
// applies sum between holder balances on both tokens, normalized them to
// the max number of decimals.
func (op *StrategyOperators) AND_SUM() func(iter *lexer.Iteration[*StrategyIteration]) (*StrategyIteration, error) {
	return func(iter *lexer.Iteration[*StrategyIteration]) (*StrategyIteration, error) {
		data, decimals, err := op.andOperator(iter)
		if err != nil {
			return nil, err
		}
		return &StrategyIteration{
			decimals: decimals,
			Data:     sumBalancesCombinator(data),
		}, nil
	}
}

// AND_MUL method returns a AND operator function that can be used in a
// strategy evaluation. The AND operator returns the common token holders
// between symbols database information or previous operations results. It
// applies multiplication between holder balances on both tokens, normalized
// them to the max number of decimals.
func (op *StrategyOperators) AND_MUL() func(iter *lexer.Iteration[*StrategyIteration]) (*StrategyIteration, error) {
	return func(iter *lexer.Iteration[*StrategyIteration]) (*StrategyIteration, error) {
		data, decimals, err := op.andOperator(iter)
		if err != nil {
			return nil, err
		}
		return &StrategyIteration{
			decimals: decimals,
			Data:     mulBalancesCombinator(data, decimals, true),
		}, nil
	}
}

// ON method returns a ON operator function that can be used in a strategy
// evaluation. The ON operator returns the common and not common token holders
// between symbols database information or previous operations results. It
// applies a fixed balance to 1 to indicate the membership.
func (op *StrategyOperators) OR() func(iter *lexer.Iteration[*StrategyIteration]) (*StrategyIteration, error) {
	return func(iter *lexer.Iteration[*StrategyIteration]) (*StrategyIteration, error) {
		data, decimals, err := op.orOperator(iter)
		if err != nil {
			return nil, err
		}
		return &StrategyIteration{
			decimals: decimals,
			Data:     membershipCombinator(data),
		}, nil
	}
}

// ON method returns a ON operator function that can be used in a strategy
// evaluation. The ON operator returns the common and not common token holders
// between symbols database information or previous operations results. It
// applies sum between holder balances on both tokens, normalized them to
// the max number of decimals.
func (op *StrategyOperators) OR_SUM() func(iter *lexer.Iteration[*StrategyIteration]) (*StrategyIteration, error) {
	return func(iter *lexer.Iteration[*StrategyIteration]) (*StrategyIteration, error) {
		data, decimals, err := op.orOperator(iter)
		if err != nil {
			return nil, err
		}
		return &StrategyIteration{
			decimals: decimals,
			Data:     sumBalancesCombinator(data),
		}, nil
	}
}

// ON method returns a ON operator function that can be used in a strategy
// evaluation. The ON operator returns the common and not common token holders
// between symbols database information or previous operations results. It
// applies multiplication between holder balances on both tokens, normalized
// them to the max number of decimals.
func (op *StrategyOperators) OR_MUL() func(iter *lexer.Iteration[*StrategyIteration]) (*StrategyIteration, error) {
	return func(iter *lexer.Iteration[*StrategyIteration]) (*StrategyIteration, error) {
		data, decimals, err := op.orOperator(iter)
		if err != nil {
			return nil, err
		}
		return &StrategyIteration{
			decimals: decimals,
			Data:     mulBalancesCombinator(data, decimals, false),
		}, nil
	}
}

// andOperator method returns the common token holders between symbols from the
// token information in the database or previous operations results. It applies
// a balance normalization to the max number of decimals and also returns the
// max number of decimals.
func (op *StrategyOperators) andOperator(iter *lexer.Iteration[*StrategyIteration]) (map[string][2]*big.Int, uint64, error) {
	interalCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	// get information about the current operation
	symbolA, dataA := iter.A()
	symbolB, dataB := iter.B()
	// no results for any component from previous operation, so get both data
	// from the database using an AND SQL query
	if dataA == nil && dataB == nil {
		data, err := op.andHoldersDBOperator(interalCtx, symbolA, symbolB)
		if err != nil {
			return nil, 0, err
		}
		// get tokens decimals
		aDecimals, ok := op.decimalsBySymbol(symbolA)
		if !ok {
			return nil, 0, fmt.Errorf("token decimals not found: %s", symbolA)
		}
		bDecimals, ok := op.decimalsBySymbol(symbolB)
		if !ok {
			return nil, 0, fmt.Errorf("token decimals not found: %s", symbolA)
		}
		// normalize balances and get the comma places moved
		data, commaPlaces, ok := op.normalizeHolderBalances(data, aDecimals, bDecimals)
		if !ok {
			return nil, 0, fmt.Errorf("error normalizing balances of %s and %s", symbolA, symbolB)
		}
		return data, commaPlaces, nil
	}
	aDecimals, bDecimals := uint64(0), uint64(0)
	// if the dataA is empty (does not contains results of previous operarion),
	// fill its data with the records of the database
	if dataA == nil {
		dataA = &StrategyIteration{}
		// get holders by token symbol
		var err error
		if dataA.Data, err = op.holdersBySymbol(interalCtx, symbolA); err != nil {
			return nil, 0, err
		}
		// get token decimals
		var ok bool
		aDecimals, ok = op.decimalsBySymbol(symbolA)
		if !ok {
			return nil, 0, fmt.Errorf("token decimals not found: %s", symbolA)
		}
	} else {
		aDecimals = dataA.decimals
	}
	// if the dataB is empty (does not contains results of previous operarion),
	// fill its data with the records of the database
	if dataB == nil {
		dataB = &StrategyIteration{}
		// get holders by token symbol
		var err error
		if dataB.Data, err = op.holdersBySymbol(interalCtx, symbolB); err != nil {
			return nil, 0, err
		}
		// get token decimals
		var ok bool
		bDecimals, ok = op.decimalsBySymbol(symbolB)
		if !ok {
			return nil, 0, fmt.Errorf("token decimals not found: %s", symbolB)
		}
	} else {
		bDecimals = dataB.decimals
	}
	// when both data sources are filled, do the intersection of both lists.
	data := intersection(dataA.Data, dataB.Data)
	// normalize balances and get the comma places moved
	data, commaPlaces, ok := op.normalizeHolderBalances(data, aDecimals, bDecimals)
	if !ok {
		return nil, 0, fmt.Errorf("error normalizing balances of %s and %s", symbolA, symbolB)
	}
	return data, commaPlaces, nil
}

// orOperator method returns the common and not common token holders between
// symbols from the token information in the database or previous operations results. It
// applies a balance normalization to the max number of decimals and also
// returns the max number of decimals.
func (op *StrategyOperators) orOperator(iter *lexer.Iteration[*StrategyIteration]) (map[string][2]*big.Int, uint64, error) {
	interalCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	// get information about the current operation
	symbolA, dataA := iter.A()
	symbolB, dataB := iter.B()
	// no results for any component from previous operation, so get both data
	// from the database using an OR SQL query
	if dataA == nil && dataB == nil {
		data, err := op.orHoldersDBOperator(interalCtx, symbolA, symbolB)
		if err != nil {
			return nil, 0, err
		}
		// get tokens decimals
		aDecimals, ok := op.decimalsBySymbol(symbolA)
		if !ok {
			return nil, 0, fmt.Errorf("token decimals not found: %s", symbolA)
		}
		bDecimals, ok := op.decimalsBySymbol(symbolB)
		if !ok {
			return nil, 0, fmt.Errorf("token decimals not found: %s", symbolA)
		}
		// normalize balances and get the comma places moved
		data, commaPlaces, ok := op.normalizeHolderBalances(data, aDecimals, bDecimals)
		if !ok {
			return nil, 0, fmt.Errorf("error normalizing balances of %s and %s", symbolA, symbolB)
		}
		return data, commaPlaces, nil
	}
	aDecimals, bDecimals := uint64(0), uint64(0)
	// if the dataA is empty (does not contains results of previous operarion),
	// fill its data with the records of the database
	if dataA == nil {
		dataA = &StrategyIteration{}
		// get holders by token symbol
		var err error
		if dataA.Data, err = op.holdersBySymbol(interalCtx, symbolA); err != nil {
			return nil, 0, err
		}
		// get token decimals
		var ok bool
		aDecimals, ok = op.decimalsBySymbol(symbolA)
		if !ok {
			return nil, 0, fmt.Errorf("token decimals not found: %s", symbolA)
		}
	} else {
		aDecimals = dataA.decimals
	}
	// if the dataB is empty (does not contains results of previous operarion),
	// fill its data with the records of the database
	if dataB == nil {
		dataB = &StrategyIteration{}
		// get holders by token symbol
		var err error
		if dataB.Data, err = op.holdersBySymbol(interalCtx, symbolB); err != nil {
			return nil, 0, err
		}
		// get token decimals
		var ok bool
		bDecimals, ok = op.decimalsBySymbol(symbolB)
		if !ok {
			return nil, 0, fmt.Errorf("token decimals not found: %s", symbolB)
		}
	} else {
		bDecimals = dataB.decimals
	}
	// when both data sources are filled, do the combination of both lists.
	data := combination(dataA.Data, dataB.Data)
	// normalize balances and get the comma places moved
	data, commaPlaces, ok := op.normalizeHolderBalances(data, aDecimals, bDecimals)
	if !ok {
		return nil, 0, fmt.Errorf("error normalizing balances of %s and %s", symbolA, symbolB)
	}
	return data, commaPlaces, nil
}

// normalizeHolderBalances method normalizes the balances for a holder
// balances map. It also returns the comma places moved to normalize the
// balances. If the token decimals are not found, returns false.
func (op *StrategyOperators) normalizeHolderBalances(
	data map[string][2]*big.Int,
	aDecimals, bDecimals uint64,
) (map[string][2]*big.Int, uint64, bool) {
	// normalize balances and get the comma places moved
	var commaPlaces uint64
	for address, balances := range data {
		var nBalanceA, nBalanceB *big.Int
		nBalanceA, nBalanceB, commaPlaces = normalize(balances[0], balances[1], aDecimals, bDecimals)
		data[address] = [2]*big.Int{nBalanceA, nBalanceB}
	}
	return data, commaPlaces, true
}

// addHoldersDBOperator method queries to the database for the holders
// associated to the symbols provided. It calls to tokenInfoBySymbol first to
// get the token information by this symbol, and then queries the database. To
// get the common holders between tokens, it uses an AND SQL query operator,
// filetered by the minimun balances and chain id associated to each token. It
// returns a map with the holders addresses as keys and the balances of both
// tokens as values.
func (op *StrategyOperators) andHoldersDBOperator(ctx context.Context,
	symbolA, symbolB string,
) (map[string][2]*big.Int, error) {
	// get both tokens information by their symbols
	addressA, chainIDA, minBalanceA, externalIDA, err := op.tokenInfoBySymbol(symbolA)
	if err != nil {
		return nil, err
	}
	addressB, chainIDB, minBalanceB, externalIDB, err := op.tokenInfoBySymbol(symbolB)
	if err != nil {
		return nil, err
	}
	// run the AND query
	rows, err := op.db.ANDOperator(ctx, queries.ANDOperatorParams{
		TokenIDA:    addressA.Bytes(),
		ChainIDA:    chainIDA,
		ExternalIDA: externalIDA,
		TokenIDB:    addressB.Bytes(),
		ChainIDB:    chainIDB,
		ExternalIDB: externalIDB,
		MinBalanceA: minBalanceA.String(),
		MinBalanceB: minBalanceB.String(),
	})
	if err != nil {
		return nil, fmt.Errorf(
			"error getting holders of %s (chainID: %d, externalID: %s) and %s (chainID: %d, externalID: %s)",
			symbolA, chainIDA, externalIDA, symbolB, chainIDB, externalIDB)
	}
	if len(rows) == 0 {
		return nil, fmt.Errorf(
			"no holders found for %s (chainID: %d, externalID: %s) and %s (chainID: %d, externalID: %s)",
			symbolA, chainIDA, externalIDA, symbolB, chainIDB, externalIDB)
	}
	// decode the results and return them
	data := make(map[string][2]*big.Int)
	for _, r := range rows {
		balanceA, ok := new(big.Int).SetString(r.BalanceA.(string), 10)
		if !ok {
			return nil, fmt.Errorf("error decoding balanceA: %s", r.BalanceA)
		}
		balanceB, ok := new(big.Int).SetString(r.BalanceB.(string), 10)
		if !ok {
			return nil, fmt.Errorf("error decoding balanceB: %s", r.BalanceB)
		}
		data[common.BytesToAddress(r.HolderID).String()] = [2]*big.Int{balanceA, balanceB}
	}
	return data, nil
}

// addHoldersDBOperator method queries to the database for the holders
// associated to the symbols provided. It calls to tokenInfoBySymbol first to
// get the token information by this symbol, and then queries the database. To
// get the holders of both tokens, it uses an OR SQL query operator, filetered
// by the minimun balances and chain id associated to each token. It returns a
// map with the holders addresses as keys and the balances of both tokens as
// values.
func (op *StrategyOperators) orHoldersDBOperator(ctx context.Context,
	symbolA, symbolB string,
) (map[string][2]*big.Int, error) {
	// get both tokens information by their symbols
	addressA, chainIDA, minBalanceA, externalIDA, err := op.tokenInfoBySymbol(symbolA)
	if err != nil {
		return nil, err
	}
	addressB, chainIDB, minBalanceB, externalIDB, err := op.tokenInfoBySymbol(symbolB)
	if err != nil {
		return nil, err
	}
	// run the AND query
	rows, err := op.db.OROperator(ctx, queries.OROperatorParams{
		TokenIDA:    addressA.Bytes(),
		ChainIDA:    chainIDA,
		ExternalIDA: externalIDA,
		TokenIDB:    addressB.Bytes(),
		ChainIDB:    chainIDB,
		ExternalIDB: externalIDB,
		MinBalanceA: minBalanceA.String(),
		MinBalanceB: minBalanceB.String(),
	})
	if err != nil {
		return nil, fmt.Errorf(
			"error getting holders of %s (chainID: %d, externalID: %s) and %s (chainID: %d, externalID: %s)",
			symbolA, chainIDA, externalIDA, symbolB, chainIDB, externalIDB)
	}
	if len(rows) == 0 {
		return nil, fmt.Errorf(
			"no holders found for %s (chainID: %d, externalID: %s) and %s (chainID: %d, externalID: %s)",
			symbolA, chainIDA, externalIDA, symbolB, chainIDB, externalIDB)
	}
	// decode the results and return them
	data := make(map[string][2]*big.Int)
	for _, r := range rows {
		balanceA, ok := new(big.Int).SetString(r.BalanceA.(string), 10)
		if !ok {
			return nil, fmt.Errorf("error decoding balanceA: %s", r.BalanceA)
		}
		balanceB, ok := new(big.Int).SetString(r.BalanceB.(string), 10)
		if !ok {
			return nil, fmt.Errorf("error decoding balanceB: %s", r.BalanceB)
		}
		data[common.BytesToAddress(r.HolderID).String()] = [2]*big.Int{balanceA, balanceB}
	}
	return data, nil
}

// intersection method returns the common token holders between two maps of
// holders balances. It returns a map with the common holders addresses as keys
// and the balances of both tokens as values. It uses a buffered channel to
// process the results in parallel.
func intersection(dataA, dataB map[string]*big.Int) map[string][2]*big.Int {
	data := make(map[string][2]*big.Int)
	var mu sync.Mutex // to safely update 'data'
	type result struct {
		address            string
		balanceA, balanceB *big.Int
	}
	resultsChan := make(chan result)
	var wg sync.WaitGroup
	for addressA, balanceA := range dataA {
		wg.Add(1)
		go func(addressA string, balanceA *big.Int) {
			defer wg.Done()
			if balanceB, exists := dataB[addressA]; exists {
				resultsChan <- result{address: addressA, balanceA: balanceA, balanceB: balanceB}
			}
		}(addressA, balanceA)
	}
	// close resultsChan when all goroutines are done.
	go func() {
		wg.Wait()
		close(resultsChan)
	}()
	// collect results from the channel.
	for r := range resultsChan {
		mu.Lock()
		data[r.address] = [2]*big.Int{r.balanceA, r.balanceB}
		mu.Unlock()
	}
	return data
}

// combination method returns the common and not common token holders between
// two maps of holders balances. It returns a map with the holders addresses as
// keys and the balances of both tokens as values. It uses a buffered channel to
// process the results in parallel.
func combination(dataA, dataB map[string]*big.Int) map[string][2]*big.Int {
	// init result data with dataA
	data := make(map[string][2]*big.Int)
	for addressA, balanceA := range dataA {
		data[addressA] = [2]*big.Int{balanceA, nil}
	}
	type update struct {
		address  string
		balanceB *big.Int
	}
	// create a buffered channel based on size of dataB to process updates in
	// parallel
	updates := make(chan update, len(dataB))
	var wg sync.WaitGroup
	for addressB, balanceB := range dataB {
		wg.Add(1)
		go func(addressB string, balanceB *big.Int) {
			defer wg.Done()
			updates <- update{address: addressB, balanceB: balanceB}
		}(addressB, balanceB)
	}
	// close the updates channel once all goroutines are finished
	go func() {
		wg.Wait()
		close(updates)
	}()
	// collect updates from the channel and apply them to data
	for update := range updates {
		if balanceA, ok := data[update.address]; ok {
			// If the address exists in data, update the second balance
			data[update.address] = [2]*big.Int{balanceA[0], update.balanceB}
		} else {
			// If the address doesn't exist in data, add a new entry
			data[update.address] = [2]*big.Int{nil, update.balanceB}
		}
	}
	return data
}
