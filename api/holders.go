package api

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/census3/contractstate"
	"github.com/vocdoni/census3/service"
	"go.vocdoni.io/dvote/httprouter"
	api "go.vocdoni.io/dvote/httprouter/apirest"
	"go.vocdoni.io/dvote/log"
)

// holdersHandlers struct envolves an initializated HolderScanner.
// TODO: replace ot with an struct to envolves all scanners or other services,
// this will help to call other scanners and services functions.
type holdersHandlers struct {
	scanner *service.HoldersScanner
}

// initHoldersHandlers function registers the endpoints related with the token
// holders such us get a token holders lists of a token. It uses the given
// initialized API and HoldersScanner service.
func initHoldersHandlers(currentApi *api.API, scanner *service.HoldersScanner) {
	handler := holdersHandlers{scanner}
	// TODO: Move the createToken endpoint to other api package file.
	currentApi.RegisterMethod("/tokens", "POST",
		api.MethodAccessTypePublic, handler.createToken)
	currentApi.RegisterMethod("/tokens/{address}/holders", "GET",
		api.MethodAccessTypePublic, handler.getHolders)
	currentApi.RegisterMethod("/tokens/{address}/holders/count", "GET",
		api.MethodAccessTypePublic, handler.countHolders)
}

// TODO: Move this handler to other api package file.
func (h *holdersHandlers) createToken(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	req := CreateTokenRequest{}
	if err := json.Unmarshal(msg.Data, &req); err != nil {
		log.Errorf("error unmarshalling token information: %s", err)
		return ErrTokenMalformed.With("error unmarshalling token information")
	}

	tType := contractstate.ContractTypeFromString(req.Type)
	tAddr := common.HexToAddress(req.Address)
	if err := h.scanner.AddToken(tAddr, tType, req.StartBlock); err != nil {
		log.Errorf("error creating token with address %s: %s", tAddr, err)
		return ErrCantCreateToken.Withf("error creating token with address %s", tAddr)
	}
	return ctx.Send([]byte("Ok"), api.HTTPstatusOK)
}

// getHolders handler function receives the token contract address from the
// request and tries to get its token holders. If something fails, the given
// token not exists or has not registered token holders, it will returns and api
// error.
func (h *holdersHandlers) getHolders(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	addr := common.HexToAddress(ctx.URLParam("address"))
	th, err := h.scanner.GetTokenHolders(addr)
	if err != nil {
		log.Errorf("error getting token with address %s: %s", addr, err)
		return ErrCantGetTokenHolders.Withf("error getting token with address %s", addr)
	}
	if th == nil {
		log.Errorf("token with address %s", addr)
		return ErrUnknownToken.Withf("token with address %s", addr)
	}

	holders := TokenHoldersResponse{Holders: []string{}}
	for _, addr := range th.Holders() {
		holders.Holders = append(holders.Holders, addr.Hex())
	}
	if len(holders.Holders) == 0 {
		log.Errorf("no holders found for %s", addr)
		return ErrNoFoundTokenHolders.Withf("no holders found for %s", addr)
	}
	response, err := json.Marshal(holders)
	if err != nil {
		log.Errorf("error marshalling holder of %s: %s", addr, err)
		return ErrEncodeTokenHolders.Withf("error marshalling holder of %s", addr)
	}
	return ctx.Send(response, api.HTTPstatusOK)
}

func (h *holdersHandlers) countHolders(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	addr := common.HexToAddress(ctx.URLParam("address"))
	count, err := h.scanner.GetNumberOfTokenHolders(addr)
	if err != nil {
		log.Errorf("error getting token with address %s: %s", addr, err)
		return ErrCantGetTokenHolders.Withf("token address: %s", addr)
	}
	if count == 0 {
		log.Errorf("token with address %s", addr)
		return ErrNoFoundTokenHolders.Withf("token address: %s", addr)
	}

	response, err := json.Marshal(struct {
		Count int64 `json:"count"`
	}{
		Count: count,
	})
	if err != nil {
		log.Errorf("error marshalling holder of %s: %s", addr, err)
		return ErrEncodeTokenHolders.Withf("token address: %s", addr)
	}
	return ctx.Send(response, api.HTTPstatusOK)
}
