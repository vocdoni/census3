package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/census3/census"
	queries "github.com/vocdoni/census3/db/sqlc"
	"go.vocdoni.io/dvote/httprouter"
	api "go.vocdoni.io/dvote/httprouter/apirest"
	"go.vocdoni.io/dvote/log"
)

func (capi *census3API) initCensusHandlers() error {
	if err := capi.endpoint.RegisterMethod("/censuses/{censusID}", "GET",
		api.MethodAccessTypePublic, capi.getCensus); err != nil {
		return err
	}
	if err := capi.endpoint.RegisterMethod("/censuses", "POST",
		api.MethodAccessTypePublic, capi.createAndPublishCensus); err != nil {
		return err
	}
	return capi.endpoint.RegisterMethod("/censuses/strategy/{strategyID}", "GET",
		api.MethodAccessTypePublic, capi.getStrategyCensuses)
}

// getCensus handler responses with the information regarding of the census
// requested by its ID.
func (capi *census3API) getCensus(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	censusID, err := strconv.Atoi(ctx.URLParam("censusID"))
	if err != nil {
		return ErrMalformedCensusID
	}
	internalCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	currentCensus, err := capi.db.QueriesRO.CensusByID(internalCtx, int64(censusID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNotFoundCensus
		}
		return ErrCantGetCensus
	}
	res, err := json.Marshal(GetCensusResponse{
		CensusID:   uint64(censusID),
		StrategyID: uint64(currentCensus.StrategyID),
		MerkleRoot: common.Bytes2Hex(currentCensus.MerkleRoot),
		URI:        "ipfs://" + currentCensus.Uri.String,
		Size:       int32(currentCensus.Size),
		Weight:     new(big.Int).SetBytes(currentCensus.Weight).String(),
		Anonymous:  currentCensus.CensusType == int64(census.AnonymousCensusType),
	})
	if err != nil {
		return ErrEncodeCensus
	}
	return ctx.Send(res, api.HTTPstatusOK)
}

// createAndPublishCensus handler creates a census tree based on the token
// holders of the tokens that are included in the given strategy. It recovers
// all the required information from the database, and then creates and publish
// the census merkle tree on IPFS. Then saves the resulting information of the
// census tree in the database and returns its ID.
//
// TODO: This handler is costly, specially for big censuses. It should be refactored to be a background task.
func (capi *census3API) createAndPublishCensus(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	// decode request
	req := &CreateCensusResquest{}
	if err := json.Unmarshal(msg.Data, req); err != nil {
		return ErrMalformedStrategyID
	}
	// get tokens associated to the strategy
	internalCtx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// begin a transaction for group sql queries
	tx, err := capi.db.RW.BeginTx(internalCtx, nil)
	if err != nil {
		return ErrCantCreateCensus
	}
	defer func() {
		if err := tx.Rollback(); err != nil && !errors.Is(sql.ErrTxDone, err) {
			log.Errorw(err, "holders transaction rollback failed")
		}
	}()
	qtx := capi.db.QueriesRW.WithTx(tx)

	strategyTokens, err := qtx.TokensByStrategyID(internalCtx, int64(req.StrategyID))
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			log.Errorf("no strategy found for id %d: %s", req.StrategyID, err.Error())
			return ErrNotFoundStrategy
		}
		log.Errorf("error getting strategy with id %d: %s", req.StrategyID, err.Error())
		return ErrCantGetStrategy
	}
	// get holders associated to every strategy token, create a map to avoid
	// duplicates and count the sum of the balances to get the weight of the
	// census
	censusWeight := new(big.Int)
	strategyHolders := map[common.Address]*big.Int{}
	for _, token := range strategyTokens {
		holders, err := qtx.TokenHoldersByTokenID(internalCtx, token.ID)
		if err != nil {
			if errors.Is(sql.ErrNoRows, err) {
				continue
			}
			log.Errorw(err, "error getting token holders")
			return ErrCantGetTokenHolders.Withf("for the token with address %s",
				common.BytesToAddress(token.ID))
		}
		for _, holder := range holders {
			holderAddr := common.BytesToAddress(holder.ID)
			holderBalance := new(big.Int).SetBytes(holder.Balance)
			if _, exists := strategyHolders[holderAddr]; !exists {
				strategyHolders[holderAddr] = holderBalance
				censusWeight = new(big.Int).Add(censusWeight, holderBalance)
			}
		}
	}
	// get the maximun current census ID to calculate the next one, if any
	// census has been created yet, continue
	lastCensusID, err := qtx.LastCensusID(internalCtx)
	if err != nil && !errors.Is(sql.ErrNoRows, err) {
		log.Errorw(err, "error getting last census ID")
		return ErrCantCreateCensus
	}
	// create a census tree and publish on IPFS
	def := census.NewCensusDefinition(int(lastCensusID+1), int(req.StrategyID), strategyHolders, req.Anonymous)
	newCensus, err := capi.censusDB.CreateAndPublish(def)
	if err != nil {
		log.Errorw(err, "error creating or publishing the census")
		return ErrCantCreateCensus
	}
	// check if the census already exists using the merkle root of the generated
	// census
	existingCensus, err := qtx.CensusByMerkleRoot(internalCtx, newCensus.RootHash)
	if err == nil {
		// encoding the result and response it
		res, err := json.Marshal(CreateCensusResponse{
			CensusID: uint64(existingCensus.ID),
		})
		if err != nil {
			log.Error("error marshalling holders")
			return ErrEncodeStrategyHolders
		}
		return ctx.Send(res, api.HTTPstatusOK)
	}
	if err != nil && !errors.Is(sql.ErrNoRows, err) {
		log.Errorw(err, "error checking if the generated census already exists")
		return ErrCantCreateCensus
	}
	// save the new census in the SQL database
	sqlURI := new(sql.NullString)
	if err := sqlURI.Scan(newCensus.URI); err != nil {
		log.Errorw(err, "error saving the census on the database")
		return ErrCantCreateCensus
	}
	_, err = qtx.CreateCensus(internalCtx, queries.CreateCensusParams{
		ID:         int64(newCensus.ID),
		StrategyID: int64(req.StrategyID),
		MerkleRoot: newCensus.RootHash,
		Uri:        *sqlURI,
		Size:       int64(len(strategyHolders)),
		Weight:     censusWeight.Bytes(),
		CensusType: int64(def.Type),
	})
	if err != nil {
		log.Errorw(err, "error saving the census on the database")
		return ErrCantCreateCensus
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	// encoding the result and response it
	res, err := json.Marshal(CreateCensusResponse{
		CensusID: uint64(newCensus.ID),
	})
	if err != nil {
		log.Error("error marshalling holders")
		return ErrEncodeStrategyHolders
	}
	return ctx.Send(res, api.HTTPstatusOK)
}

// getStrategyCensuses function handler returns the censuses that had been
// generated with the strategy identified by the ID provided.
func (capi *census3API) getStrategyCensuses(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	// get strategy ID
	strategyID, err := strconv.Atoi(ctx.URLParam("strategyID"))
	if err != nil {
		return ErrMalformedCensusID
	}
	// get censuses by this strategy ID
	internalCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rows, err := capi.db.QueriesRO.CensusByStrategyID(internalCtx, int64(strategyID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNotFoundCensus
		}
		return ErrCantGetCensus
	}
	// parse and encode response
	censuses := GetCensusesResponse{Censuses: []uint64{}}
	for _, censusInfo := range rows {
		censuses.Censuses = append(censuses.Censuses, uint64(censusInfo.ID))
	}
	res, err := json.Marshal(censuses)
	if err != nil {
		log.Errorw(ErrEncodeCensuses, err.Error())
		return ErrEncodeCensuses
	}
	return ctx.Send(res, api.HTTPstatusOK)
}
