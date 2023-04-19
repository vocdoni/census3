package api

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/census3/contractstate"
	"github.com/vocdoni/census3/service"
	"go.vocdoni.io/dvote/httprouter"
	api "go.vocdoni.io/dvote/httprouter/apirest"
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
		return new(api.APIerror).WithErr(err).Send(ctx)
	}

	tType := contractstate.ContractTypeFromString(req.Type)
	tAddr := common.HexToAddress(req.Address)
	if err := h.scanner.AddToken(tAddr, tType, req.StartBlock); err != nil {
		return new(api.APIerror).WithErr(err)
	}
	return ctx.Send([]byte("Ok"), api.HTTPstatusOK)
}

func (h *holdersHandler) getHolders(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	addr := common.HexToAddress(ctx.URLParam("address"))
	th, err := h.scanner.GetHolders(addr)
	if err != nil {
		return new(api.APIerror).WithErr(err).Send(ctx)
	}

	holders := TokenHoldersResponse{Holders: []string{}}
	for _, addr := range th.Holders() {
		holders.Holders = append(holders.Holders, addr.Hex())
	}
	response, err := json.Marshal(holders)
	if err != nil {
		return new(api.APIerror).WithErr(err).Send(ctx)
	}
	return ctx.Send(response, api.HTTPstatusOK)
}
