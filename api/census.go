package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/census3/census"
	queries "github.com/vocdoni/census3/db/sqlc"
	"go.vocdoni.io/dvote/httprouter"
	api "go.vocdoni.io/dvote/httprouter/apirest"
	"go.vocdoni.io/dvote/log"
)

func (capi *census3API) initCensusHandlers() {
	capi.endpoint.RegisterMethod("/census/{censusID}", "GET",
		api.MethodAccessTypePublic, capi.getCensus)
	capi.endpoint.RegisterMethod("/census", "POST",
		api.MethodAccessTypePublic, capi.createAndPublishCensus)

	// TODO: Only for debug, remove it
	capi.endpoint.RegisterMethod("/census/{censusID}/check/{root}", "POST",
		api.MethodAccessTypePublic, capi.checkIPFSCensus)
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
	currentCensus, err := capi.sqlc.CensusByID(internalCtx, int64(censusID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNotFoundCensus
		}
		return ErrCantGetCensus
	}
	res, err := json.Marshal(GetCensusResponse{
		CensusID:   fmt.Sprint(censusID),
		StrategyID: fmt.Sprint(currentCensus.StrategyID),
		MerkleRoot: common.Bytes2Hex(currentCensus.MerkleRoot),
		URI:        "ipfs://" + currentCensus.Uri.String,
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
func (capi *census3API) createAndPublishCensus(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	// decode request
	req := &CreateCensusResquest{}
	if err := json.Unmarshal(msg.Data, req); err != nil {
		return ErrMalformedStrategyID
	}
	// get tokens associated to the strategy
	internalCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	strategyTokens, err := capi.sqlc.TokensByStrategyID(internalCtx, queries.TokensByStrategyIDParams{
		StrategyID: int64(req.StrategyID),
		Limit:      -1,
		Offset:     0,
	})
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			log.Errorf("no strategy found for id %d: %w", req.StrategyID, err)
			return ErrNotFoundStrategy
		}
		log.Errorf("error getting strategy with id %d: %w", req.StrategyID, err)
		return ErrCantGetStrategy
	}
	// get holders associated to every strategy token
	// create a map to avoid duplicates
	strategyHolders := map[common.Address]int{}
	for _, token := range strategyTokens {
		holders, err := capi.sqlc.TokenHoldersByTokenID(internalCtx, queries.TokenHoldersByTokenIDParams{
			TokenID: token.ID,
			Limit:   -1,
			Offset:  0,
		})
		if err != nil {
			if errors.Is(sql.ErrNoRows, err) {
				continue
			}
			return ErrCantGetTokenHolders.Withf("for the token with address %s",
				common.BytesToAddress(token.ID))
		}
		for _, bAddr := range holders {
			holderAddr := common.BytesToAddress(bAddr)
			if _, exists := strategyHolders[holderAddr]; !exists {
				strategyHolders[holderAddr] = 1
			}
		}
	}
	// get the maximun current census ID to calculate the next one
	lastCensusID, err := capi.sqlc.LastCensusID(internalCtx)
	if err != nil {
		log.Errorf("error getting last census ID, continue")
	}
	// create a census tree and publish on IPFS
	def := census.DefaultCensusDefinition(int(lastCensusID+1), int(req.StrategyID), strategyHolders)
	newCensus, err := capi.censusDB.CreateAndPublish(def)
	if err != nil {
		log.Errorw(err, "error creating or publishing the census")
		return ErrCantCreateCensus
	}
	// save the new census in the SQL database
	sqlURI := new(sql.NullString)
	if err := sqlURI.Scan(newCensus.URI); err != nil {
		log.Errorf("error saving the census on the database: %w", err)
		return ErrCantCreateCensus
	}
	_, err = capi.sqlc.CreateCensus(internalCtx, queries.CreateCensusParams{
		ID:         int64(newCensus.ID),
		StrategyID: int64(req.StrategyID),
		MerkleRoot: newCensus.RootHash,
		Uri:        *sqlURI,
	})
	if err != nil {
		log.Errorf("error saving the census on the database: %w", err)
		return ErrCantCreateCensus
	}
	// encoding the result and response it
	res, err := json.Marshal(CreateCensusResponse{
		CensusID: strconv.Itoa(newCensus.ID),
	})
	if err != nil {
		log.Error("error marshalling holders")
		return ErrEncodeStrategyHolders
	}
	return ctx.Send(res, api.HTTPstatusOK)
}

// TODO: Only for debug, remove it
func (capi *census3API) checkIPFSCensus(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	// get strategy and query about the tokens that it includes
	censusID, err := strconv.Atoi(ctx.URLParam("censusID"))
	root := ctx.URLParam("root")
	if err != nil {
		return ErrMalformedStrategyID
	}

	internalCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	currentCensus, err := capi.sqlc.CensusByID(internalCtx, int64(censusID))
	if err != nil {
		return ErrCantGetStrategy
	}

	censusDef := census.DefaultCensusDefinition(censusID, 0, nil)
	censusDef.URI = currentCensus.Uri.String
	if err := capi.censusDB.Check(censusDef, []byte(root)); err != nil {
		log.Error(err)
		return api.APIerror{Code: 5100, HTTPstatus: 500, Err: err}
	}

	return ctx.Send([]byte("ok"), api.HTTPstatusOK)
}
