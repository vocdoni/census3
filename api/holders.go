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

type holdersHandler struct {
	scanner *service.HoldersScanner
}

func initHoldersHandler(currentApi *api.API, scanner *service.HoldersScanner) {
	handler := holdersHandler{scanner}

	currentApi.RegisterMethod("/tokens", "POST",
		api.MethodAccessTypePublic, handler.createToken)
	currentApi.RegisterMethod("/tokens/{address}/holders", "GET",
		api.MethodAccessTypePublic, handler.getHolders)
}

func (h *holdersHandler) createToken(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
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

func (h *holdersHandler) getHolders(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	addr := common.HexToAddress(ctx.URLParam("address"))
	th, err := h.scanner.GetHolders(addr)
	if err != nil {
		log.Errorf("error getting token with address %s: %s", addr, err)
		return ErrCantGetTokenHolders.Withf("error getting token with address %s", addr)
	}
	if th == nil {
		log.Errorf("token with address %s", addr)
		return ErrCantGetTokenHolders.Withf("token with address %s", addr)
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
