package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"github.com/ethereum/go-ethereum/common"
	queries "github.com/vocdoni/census3/db/sqlc"
	"go.vocdoni.io/dvote/httprouter"
	api "go.vocdoni.io/dvote/httprouter/apirest"
	"go.vocdoni.io/dvote/log"
)

// initHoldersHandlers function registers the endpoints related with the token
// holders such us get a token holders lists of a token. It uses the given
// initialized API.
func (capi *census3API) initHoldersHandlers() {
	// TODO: Move the createToken endpoint to other api package file.
	capi.endpoint.RegisterMethod("/tokens/{address}/holders", "GET",
		api.MethodAccessTypePublic, capi.getTokenHolders)
	capi.endpoint.RegisterMethod("/tokens/{address}/holders/count", "GET",
		api.MethodAccessTypePublic, capi.countHolders)
}

// getTokenHolders handler function gets the token holders states from the
// database, of the token identified by the contract address provided.
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
		// not found, else return generic error.
		if errors.Is(sql.ErrNoRows, err) {
			log.Errorf("no holders found for address %s: %w", addr, err)
			return ErrNoFoundTokenHolders.Withf("no holders found for address %s", addr)
		}
		log.Errorf("error getting token with address %s: %w", addr, err)
		return ErrCantGetTokenHolders.Withf("error getting token with address %s", addr)
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
// TODO: Only for debug, consider to delete
func (capi *census3API) countHolders(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	ctx2, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	addr := common.HexToAddress(ctx.URLParam("address"))
	numberOfHolders, err := capi.sqlc.CountTokenHoldersByTokenID(ctx2, addr.Bytes())
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			log.Errorf("no holders found for address %s: %w", addr, err)
			return ErrNoFoundTokenHolders.Withf("token address: %s", addr)
		}
		log.Errorf("error getting holders of %s: %w", addr, err)
		return ErrCantGetTokenHolders.Withf("token address: %s", addr)
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
