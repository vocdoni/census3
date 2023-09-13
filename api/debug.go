package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"go.vocdoni.io/dvote/httprouter"
	api "go.vocdoni.io/dvote/httprouter/apirest"
	"go.vocdoni.io/dvote/log"
)

// TODO: Only for the MVP, remove it.
func (capi *census3API) initDebugHandlers() error {
	if err := capi.endpoint.RegisterMethod("/debug/token/{address}/holders", "GET",
		api.MethodAccessTypePublic, capi.getTokenHolders); err != nil {
		return err
	}
	if err := capi.endpoint.RegisterMethod("/debug/token/{address}/holders/count", "GET",
		api.MethodAccessTypePublic, capi.countHolders); err != nil {
		return err
	}
	return nil
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
	dbHolders, err := capi.db.QueriesRO.TokenHoldersByTokenID(ctx2, addr.Bytes())
	if err != nil {
		// if database does not contain any token holder for this token, return
		// no content, else return generic error.
		if errors.Is(sql.ErrNoRows, err) {
			log.Errorf("no holders found for address %s: %s", addr, err.Error())
			return ctx.Send(nil, api.HTTPstatusNoContent)
		}
		log.Errorf("error getting token with address %s: %s", addr, err.Error())
		return ErrCantGetTokenHolders.Withf("error getting token with address %s", addr)
	}
	// if no error but the results are empty, return no content
	if len(dbHolders) == 0 {
		log.Errorf("no holders found for address %s", addr)
		return ctx.Send(nil, api.HTTPstatusNoContent)
	}
	// encode the response with the token holders addresses
	holders := TokenHoldersResponse{Holders: map[string]string{}}
	for _, holder := range dbHolders {
		addr := common.BytesToAddress(holder.ID).String()
		balance := new(big.Int).SetBytes(holder.Balance)
		holders.Holders[addr] = balance.String()
	}
	response, err := json.Marshal(holders)
	if err != nil {
		log.Errorf("error marshalling holder of %s: %s", addr, err.Error())
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
	numberOfHolders, err := capi.db.QueriesRO.CountTokenHoldersByTokenID(ctx2, addr.Bytes())
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			log.Errorf("no holders found for address %s: %s", addr, err.Error())
			return ctx.Send(nil, api.HTTPstatusNoContent)
		}
		log.Errorf("error getting holders of %s: %s", addr, err.Error())
		return ErrCantGetTokenHolders.Withf("token address: %s", addr)
	}
	// if no error but the results are empty, return no content
	if numberOfHolders == 0 {
		log.Errorf("no holders found for address %s: %s", addr, err.Error())
		return ctx.Send(nil, api.HTTPstatusNoContent)
	}
	response, err := json.Marshal(struct {
		Count int64 `json:"count"`
	}{
		Count: numberOfHolders,
	})
	if err != nil {
		log.Errorf("error marshalling holder of %s: %s", addr, err.Error())
		return ErrEncodeTokenHolders.Withf("token address: %s", addr)
	}
	return ctx.Send(response, api.HTTPstatusOK)
}
