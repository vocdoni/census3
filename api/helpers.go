package api

import (
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	queries "github.com/vocdoni/census3/db/sqlc"
	"github.com/vocdoni/census3/helpers/lexer"
	"github.com/vocdoni/census3/helpers/strategyoperators"
	"github.com/vocdoni/census3/scanner/providers/web3"
	"go.vocdoni.io/dvote/api/censusdb"
	"go.vocdoni.io/dvote/censustree"
	storagelayer "go.vocdoni.io/dvote/data"
	"go.vocdoni.io/dvote/httprouter"
	"go.vocdoni.io/dvote/types"
	"go.vocdoni.io/proto/build/go/models"
)

// holdersBatchSize is the number of holders to add to the census tree in a
// single batch. This is used to avoid adding all the holders at once, which
// could not allow to track the progress of the census creation. Creating the
// census in batches allows to update the progress of the census creation.
const holdersBatchSize = 1000

// paginationFromCtx extracts from the request and returns the page size,
// the database page size, the current cursor and the direction of that cursor.
// The page size is the number of elements of the page, the database page size
// is the number of elements of the page plus one, to get the previous and next
// cursors. The cursor and the direction are extracted from the request. If
// both cursors are provided, it returns an error.
func paginationFromCtx(ctx *httprouter.HTTPContext) (int32, int32, string, bool, error) {
	// define the initial page size by increasing the probvided value to get
	// the previous and next cursors
	pageSize := defaultPageSize
	dbPageSize := defaultPageSize + 1
	// if the page size is provided, use the provided value instead, increasing
	// it by 2 to get the previous and next cursors
	if strPageSize := ctx.Request.URL.Query().Get("pageSize"); strPageSize != "" {
		if intPageSize, err := strconv.Atoi(strPageSize); err == nil {
			if intPageSize < 0 {
				pageSize = -1
				dbPageSize = -1
			} else if intPageSize < int(defaultPageSize) {
				pageSize = int32(intPageSize)
				dbPageSize = int32(intPageSize) + 1
			}
		}
	}
	// get posible previous and next cursors
	prevCursor := ctx.Request.URL.Query().Get("prevCursor")
	nextCursor := ctx.Request.URL.Query().Get("nextCursor")
	// if both cursors are provided, return an error
	if prevCursor != "" && nextCursor != "" {
		return 0, 0, "", false, fmt.Errorf("both cursors provided, next and previous")
	}
	// by default go forward, if the previous cursor is provided, go backwards
	goForward := prevCursor == ""
	cursor := nextCursor
	if !goForward {
		cursor = prevCursor
	}
	// return the page size, the cursor and the direction
	return pageSize, dbPageSize, cursor, goForward, nil
}

// paginationToRequest returns the rows of the page, the next cursor and the
// previous cursor. If the rows size is the same as the database page size, the
// last element of the page is the next cursor, so it has to be removed from the
// rows. If the current page is the first one, the previous cursor is nil, and
// the rows are empty, because the first element is the cursor and there is
// include it in the following page. It uses generics to support any type of
// rows. The cursors will alwways be strings.
func paginationToRequest[T any](rows []T, dbPageSize int32, goForward bool) ([]T, *T, *T) {
	// if the rows are empty there is no results or next and previous cursor
	if len(rows) == 0 {
		return rows, nil, nil
	}
	// by default, the next cursor is the last element of the page, and the
	// previous cursor is the first element of the page
	nextCursor := &rows[len(rows)-1]
	prevCursor := &rows[0]
	// if the length of the rows is less than the maximun page size, there is
	// no next cursor, and all the rows are part of the page
	if len(rows) < int(dbPageSize)-1 {
		if len(rows) > 1 {
			return rows, nil, prevCursor
		}
		// if the rows has just one element, there is no next or previous cursor, so
		// if the direction is forward, the next cursor is nil, and if the direction
		// is backwards, the previous cursor is nil and the rows are empty, because
		// the first element is the cursor and there is include it in the following
		// page.
		if len(rows) == 1 {
			if goForward {
				nextCursor = nil
			} else {
				prevCursor = nil
				rows = []T{}
			}
		}
	}
	// if the page size is the same as the database page size, the last element
	// of the page is the next cursor, so it has to be removed from the rows
	if len(rows) == int(dbPageSize) {
		rows = rows[:len(rows)-1]
	}
	return rows, nextCursor, prevCursor
}

// CensusOptions envolves the required parameters to create and publish a
// census merkle tree
type CensusOptions struct {
	ID      uint64
	Type    models.Census_Type
	Holders map[common.Address]*big.Int
}

// CreateAndPublishCensus function creates a new census tree based on the
// options provided and publishes it to IPFS. It needs to persist it temporaly
// into a internal trees database. It returns the root of the tree, the IPFS
// URI and the tree dump.
func CreateAndPublishCensus(db *censusdb.CensusDB, storage storagelayer.Storage,
	opts CensusOptions, progressCh chan float64,
) (types.HexBytes, string, []byte, error) {
	bID := make([]byte, 8)
	binary.LittleEndian.PutUint64(bID, opts.ID)
	ref, err := db.New(bID, opts.Type, "", nil, censustree.DefaultMaxLevels)
	if err != nil {
		return nil, "", nil, err
	}
	// encode the holders
	holdersAddresses, holdersValues := [][]byte{}, [][]byte{}
	for addr, balance := range opts.Holders {
		key := addr.Bytes()[:censustree.DefaultMaxKeyLen]
		if opts.Type != anonymousCensusType {
			if key, err = ref.Tree().Hash(addr.Bytes()); err != nil {
				return nil, "", nil, err
			}
		}
		holdersAddresses = append(holdersAddresses, key[:censustree.DefaultMaxKeyLen])
		value := ref.Tree().BigIntToBytes(balance)
		holdersValues = append(holdersValues, value)
	}
	// add the holders to the census tree
	db.Lock()
	defer db.Unlock()
	// add the holders addresses and values to the census tree
	if len(holdersAddresses) < holdersBatchSize {
		if _, err := ref.Tree().AddBatch(holdersAddresses, holdersValues); err != nil {
			return nil, "", nil, err
		}
	} else {
		// iterate over holders in batches to add them to the tree secuentally
		// and update the progress if the channel is provided
		for i := 0; i < len(holdersAddresses); i += holdersBatchSize {
			end := i + holdersBatchSize
			if end > len(holdersAddresses) {
				end = len(holdersAddresses)
			}
			if _, err := ref.Tree().AddBatch(holdersAddresses[i:end], holdersValues[i:end]); err != nil {
				return nil, "", nil, err
			}
			if progressCh != nil {
				progressCh <- float64(end) / float64(len(holdersAddresses)) * 100
			}
		}
	}
	// get the root of the tree
	root, err := ref.Tree().Root()
	if err != nil {
		return nil, "", nil, err
	}
	// get the tree dump
	data, err := ref.Tree().Dump()
	if err != nil {
		return nil, "", nil, err
	}
	// generate the tree dump
	dump, err := censusdb.BuildExportDump(root, data, opts.Type, censustree.DefaultMaxLevels)
	if err != nil {
		return nil, "", nil, err
	}
	// publish it on IPFS
	ctx, cancel := context.WithTimeout(context.Background(), publishCensusTimeout)
	defer cancel()
	uri, err := storage.Publish(ctx, dump)
	if err != nil {
		return nil, "", nil, err
	}
	if err := db.Del(bID); err != nil {
		return nil, "", nil, err
	}
	return root, uri, dump, nil
}

// InnerCensusID generates a unique identifier by concatenating the BlockNumber, StrategyID,
// and a numerical representation of the Anonymous flag from a CreateCensusRequest struct.
// The BlockNumber and StrategyID are concatenated as they are, and the Anonymous flag is
// represented as 1 for true and 0 for false. This concatenated string is then converted
// to a uint64 to create a unique identifier.
func InnerCensusID(blockNumber, strategyID uint64, anonymous bool) uint64 {
	// Convert the boolean to a uint64: 1 for true, 0 for false
	var anonymousUint uint64
	if anonymous {
		anonymousUint = 1
	}
	// Concatenate the three values as strings
	concatenated := fmt.Sprintf("%d%d%d", blockNumber, strategyID, anonymousUint)
	// Convert the concatenated string back to a uint64
	result, err := strconv.ParseUint(concatenated, 10, 64)
	if err != nil {
		panic(err)
	}
	if result > math.MaxInt64 {
		panic("overflow")
	}
	return result
}

// CalculateStrategyHolders function returns the holders of a strategy and the
// total weight of the census. It also returns the total block number of the
// census, which is the sum of the strategy block number or the last block
// number of every token chain id. To calculate the census holders, it uses the
// supplied predicate to filter the token holders using a lexer and evaluator.
// The evaluator uses the strategy operators to evaluate the predicate which
// uses the database queries to get the token holders and their balances, and
// combines them.
func (capi *census3API) CalculateStrategyHolders(ctx context.Context,
	predicate string, tokens map[string]*StrategyToken, progressCh chan float64,
) (map[common.Address]*big.Int, *big.Int, uint64, error) {
	// TODO: write a benchmark and try to optimize this function

	// init some variables to get computed in the following steps
	censusWeight := new(big.Int)
	strategyHolders := map[common.Address]*big.Int{}
	// parse the predicate
	lx := lexer.NewLexer(strategyoperators.ValidOperatorsTags)
	validPredicate, err := lx.Parse(predicate)
	if err != nil {
		return nil, nil, 0, err
	}
	// any census strategy is identified by id created from the concatenation of
	// the block number, the strategy id and the anonymous flag. The creation of
	// censuses on specific block is not supported yet, so we need to get the
	// last block of every token chain id to sum them and get the total block
	// number, used to create the census id.
	totalTokensBlockNumber := uint64(0)
	// init token information to get the token holders if the predicate is not a
	// literal, and to get the latest block number of every token chain id
	tokensInfo := map[string]*strategyoperators.TokenInformation{}
	for _, t := range tokens {
		token, err := capi.db.QueriesRO.GetToken(ctx, queries.GetTokenParams{
			ID:         common.HexToAddress(t.ID).Bytes(),
			ChainID:    t.ChainID,
			ExternalID: t.ExternalID,
		})
		if err != nil {
			return nil, nil, 0, err
		}
		tokensInfo[token.Symbol] = &strategyoperators.TokenInformation{
			ID:         common.BytesToAddress(token.ID).String(),
			ChainID:    token.ChainID,
			MinBalance: t.MinBalance,
			Decimals:   token.Decimals,
			ExternalID: token.ExternalID,
		}
		provider, exists := capi.holderProviders[token.TypeID]
		if !exists {
			return nil, nil, 0, fmt.Errorf("provider not found for token type id %d", token.TypeID)
		}
		if !provider.IsExternal() {
			if err := provider.SetRef(web3.Web3ProviderRef{
				HexAddress: t.ID,
				ChainID:    t.ChainID,
			}); err != nil {
				return nil, nil, 0, err
			}
		}
		currentBlockNumber, err := provider.LatestBlockNumber(ctx, []byte(token.ExternalID))
		if err != nil {
			return nil, nil, 0, err
		}
		totalTokensBlockNumber += currentBlockNumber
	}
	// if the current predicate is a literal, just query about its holders. If
	// it is a complex predicate, create a evaluator and evaluate the predicate
	if validPredicate.IsLiteral() {
		token := tokens[validPredicate.String()]
		// get the strategy holders from the database
		holders, err := capi.db.QueriesRO.TokenHoldersByMinBalance(ctx,
			queries.TokenHoldersByMinBalanceParams{
				TokenID:    common.HexToAddress(token.ID).Bytes(),
				ChainID:    token.ChainID,
				Balance:    token.MinBalance,
				ExternalID: token.ExternalID,
			})
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil, totalTokensBlockNumber, nil
			}
			return nil, nil, totalTokensBlockNumber, err
		}
		// parse holders addresses and balances
		for _, holder := range holders {
			holderAddr := common.BytesToAddress(holder.HolderID)
			holderBalance, ok := new(big.Int).SetString(holder.Balance, 10)
			if !ok {
				return nil, nil, totalTokensBlockNumber, fmt.Errorf("error decoding balance of holder %s", holderAddr.String())
			}
			if _, exists := strategyHolders[holderAddr]; !exists {
				strategyHolders[holderAddr] = holderBalance
				censusWeight = new(big.Int).Add(censusWeight, holderBalance)
			}
		}
	} else {
		// init the operators and the predicate evaluator
		operators := strategyoperators.InitOperators(capi.conf.MainCtx, capi.db.QueriesRO, tokensInfo)
		eval := lexer.NewEval[*strategyoperators.StrategyIteration](operators.Map())
		// execute the evaluation of the predicate
		res, err := eval.EvalToken(validPredicate, progressCh)
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

// TimeToCreateCensus function returns the estimated time to create a census
// based on the number of holders. It uses a linear regression to estimate the
// time to create a census based on the number of holders. The regression was
// calculated using the census3 data (11/24/2023).
func TimeToCreateCensus(size uint64) uint64 {
	// The equation would be:
	// 	seconds = m × size + c
	// Where:
	// 	* seconds is the estimated time to create a census
	// 	* size is the number of holders of the census
	// 	* m is the slope of the line (coefficient).
	// 	* c is the y-intercept (constant).
	// Based on the census3 data (11/24/2023), the value of m and c are:
	// 	* m = 0.0008017991071149796
	// 	* c = -1.1262389976474412
	// Based on the census3 data (3/11/2024), the value of m and c are:
	// 	* m = 0.00020543644248930586
	// 	* c = -0.1809418921100489
	// Based on the census3 data (3/13/2024), the value of m and c are:
	// 	* m = 0.000256
	// 	* c = -0.3035

	// To reproduce the constants, use the following python snippet:
	// 	import numpy as np
	//	A = np.array([...])
	//	B = np.array([...])
	//	m, c = np.polyfit(A, B, 1)
	m := 0.000256
	c := -0.3035
	seconds := m*float64(size) + c
	if seconds < 0 {
		seconds = 1
	}
	// add 5 seconds to the estimated time to create a census as a safety margin
	return uint64((seconds + 5) * 1000) // milliseconds
}

type CacheKey [16]byte

// EncCacheKey encodes the key to a string to be used as a map key, it uses md5
// to ensure the key is always the same length
func EncCacheKey(key string) CacheKey {
	return md5.Sum([]byte(key))
}
