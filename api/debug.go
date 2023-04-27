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

// TODO: Only for the MVP, remove it.
func (capi *census3API) initDebugHandlers() {
	capi.endpoint.RegisterMethod("/token/{address}/holder", "GET",
		api.MethodAccessTypePublic, capi.getTokenHolders)
	capi.endpoint.RegisterMethod("/token/{address}/holder/count", "GET",
		api.MethodAccessTypePublic, capi.countHolders)
	capi.endpoint.RegisterMethod("/censu/{censusID}/check/{root}", "POST",
		api.MethodAccessTypePublic, capi.checkIPFSCensus)
}

// createDummyStrategy creates the default strategy for a given token. This
// basic strategy only includes the holders of the given token which have a
// balance positive balance (holder_balance > 0).
//
// TODO: Only for the MVP, remove it.
func (capi *census3API) createDummyStrategy(tokenID []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := capi.sqlc.CreateStategy(ctx, "test")
	if err != nil {
		return err
	}
	strategyID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	_, err = capi.sqlc.CreateStrategyToken(ctx, queries.CreateStrategyTokenParams{
		StrategyID: strategyID,
		TokenID:    tokenID,
		MinBalance: big.NewInt(0).Bytes(),
		MethodHash: []byte("test"),
	})
	return err
}

// getTokenHolders handler function gets the token holders states from the
// database, of the token identified by the contract address provided.
//
// TODO: Only for the MVP, remove it.
func (capi *census3API) getTokenHolders(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	ctx2, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// get token holders from the database
	addr := common.HexToAddress(ctx.URLParam("address"))
	dbHolders, err := capi.sqlc.TokenHoldersByTokenID(ctx2,
		queries.TokenHoldersByTokenIDParams{
			TokenID: addr.Bytes(),
			Limit:   -1,
			Offset:  0,
		})
	if err != nil {
		// if database does not contain any token holder for this token, return
		// no content, else return generic error.
		if errors.Is(sql.ErrNoRows, err) {
			log.Errorf("no holders found for address %s: %w", addr, err)
			return ctx.Send(nil, api.HTTPstatusNoContent)
		}
		log.Errorf("error getting token with address %s: %w", addr, err)
		return ErrCantGetTokenHolders.Withf("error getting token with address %s", addr)
	}
	// if no error but the results are empty, return no content
	if len(dbHolders) == 0 {
		log.Errorf("no holders found for address %s: %w", addr, err)
		return ctx.Send(nil, api.HTTPstatusNoContent)
	}
	// encode the response with the token holders addresses
	holders := TokenHoldersResponse{Holders: []string{}}
	for _, holder := range dbHolders {
		holders.Holders = append(holders.Holders, common.BytesToAddress(holder).Hex())
	}
	response, err := json.Marshal(holders)
	if err != nil {
		log.Errorf("error marshalling holder of %s: %w", addr, err)
		return ErrEncodeTokenHolders.Withf("error marshalling holder of %s", addr)
	}
	return ctx.Send(response, api.HTTPstatusOK)
}

// countHolders handler function returns the current number of token holders
// for the provided token address in the current database.
//
// TODO: Only for the MVP, remove it.
func (capi *census3API) countHolders(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	ctx2, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	addr := common.HexToAddress(ctx.URLParam("address"))
	numberOfHolders, err := capi.sqlc.CountTokenHoldersByTokenID(ctx2, addr.Bytes())
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			log.Errorf("no holders found for address %s: %w", addr, err)
			return ctx.Send(nil, api.HTTPstatusNoContent)
		}
		log.Errorf("error getting holders of %s: %w", addr, err)
		return ErrCantGetTokenHolders.Withf("token address: %s", addr)
	}
	// if no error but the results are empty, return no content
	if numberOfHolders == 0 {
		log.Errorf("no holders found for address %s: %w", addr, err)
		return ctx.Send(nil, api.HTTPstatusNoContent)
	}
	response, err := json.Marshal(struct {
		Count int64 `json:"count"`
	}{
		Count: numberOfHolders,
	})
	if err != nil {
		log.Errorf("error marshalling holder of %s: %s", addr, err)
		return ErrEncodeTokenHolders.Withf("token address: %s", addr)
	}
	return ctx.Send(response, api.HTTPstatusOK)
}

// TODO: Only for the MVP, remove it.
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
