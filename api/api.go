package api

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/vocdoni/tokenstate/service"
	"go.vocdoni.io/dvote/crypto/ethereum"
	"go.vocdoni.io/dvote/httprouter"
	api "go.vocdoni.io/dvote/httprouter/bearerstdapi"
)

type Reply struct {
	Contracts []string           `json:"contracts,omitempty"`
	Token     *service.TokenInfo `json:"token,omitempty"`
	Root      string             `json:"root,omitempty"`
	Data      []byte             `json:"data,omitempty"`
	Block     uint64             `json:"block,omitempty"`
	Ok        bool               `json:"ok"`
}

func Init(host string, port int32, signer *ethereum.SignKeys, scanner *service.Scanner) error {
	r := httprouter.HTTProuter{}
	err := r.Init(host, int(port))
	if err != nil {
		return err
	}
	endpoint, err := api.NewBearerStandardAPI(&r, "/api")
	ch := contractHandler{scanner: scanner}
	endpoint.RegisterMethod("/addContract/{contract}/{startBlock}", "GET",
		api.MethodAccessTypePublic, ch.addContract)
	endpoint.RegisterMethod("/snapshot/{contract}/child/{childContract}/block/{blockNum}", "GET",
		api.MethodAccessTypePublic, ch.snapshotChildContract)
	endpoint.RegisterMethod("/listContracts", "GET",
		api.MethodAccessTypePublic, ch.listContracts)
	endpoint.RegisterMethod("/getContract/{contract}", "GET",
		api.MethodAccessTypePublic, ch.getContract)
	endpoint.RegisterMethod("/balances/{contract}", "GET",
		api.MethodAccessTypePublic, ch.dumpBalances)
	endpoint.RegisterMethod("/rescan/{contract}", "GET",
		api.MethodAccessTypePublic, ch.rescan)
	endpoint.RegisterMethod("/root/{contract}", "GET",
		api.MethodAccessTypePublic, ch.root)
	endpoint.RegisterMethod("/root/{contract}/{blockNum}", "GET",
		api.MethodAccessTypePublic, ch.root)
	endpoint.RegisterMethod("/queueExport/{contract}", "GET",
		api.MethodAccessTypePublic, ch.exportTree)
	endpoint.RegisterMethod("/fetchExport/{contract}/{blockNum}", "GET",
		api.MethodAccessTypePublic, ch.fetchTree)
	return nil
}

type contractHandler struct {
	scanner *service.Scanner
}

func (ch *contractHandler) addContract(msg *api.BearerStandardAPIdata, ctx *httprouter.HTTPContext) error {
	startBlock := ctx.URLParam("startBlock")
	if startBlock == "" {
		startBlock = "0"
	}
	startBlockU64, err := strconv.ParseUint(startBlock, 10, 64)
	if err != nil {
		return err
	}
	if err = ch.scanner.AddContract(ctx.URLParam("contract"), startBlockU64); err != nil {
		return err
	}
	resp := Reply{Ok: true}
	data, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	return ctx.Send(data, api.HTTPstatusCodeOK)
}

func (ch *contractHandler) snapshotChildContract(msg *api.BearerStandardAPIdata, ctx *httprouter.HTTPContext) error {
	atBlock, err := strconv.ParseUint(ctx.URLParam("blockNum"), 10, 64)
	if err != nil {
		return err
	}
	if err := ch.scanner.ChildSnapshot(context.Background(),
		ctx.URLParam("contract"),
		ctx.URLParam("childContract"),
		atBlock); err != nil {
		return err
	}
	resp := Reply{Ok: true}
	data, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	return ctx.Send(data, api.HTTPstatusCodeOK)
}

func (ch *contractHandler) listContracts(msg *api.BearerStandardAPIdata, ctx *httprouter.HTTPContext) error {
	resp := &Reply{Ok: true}
	resp.Contracts = ch.scanner.ListContracts()
	data, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	return ctx.Send(data, api.HTTPstatusCodeOK)
}

func (ch *contractHandler) getContract(msg *api.BearerStandardAPIdata, ctx *httprouter.HTTPContext) error {
	var err error
	resp := &Reply{Ok: true}
	resp.Token, err = ch.scanner.GetContract(ctx.URLParam("contract"))
	if err != nil {
		return err
	}

	data, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	return ctx.Send(data, api.HTTPstatusCodeOK)
}

func (ch *contractHandler) root(msg *api.BearerStandardAPIdata, ctx *httprouter.HTTPContext) error {
	token, err := ch.scanner.GetContract(ctx.URLParam("contract"))
	if err != nil {
		return err
	}
	resp := &Reply{Ok: true}
	if ctx.URLParam("blockNum") == "0" || ctx.URLParam("blockNum") == "" {
		resp.Root = token.LastRoot
	} else {
		height, err := strconv.Atoi(ctx.URLParam("blockNum"))
		if err != nil {
			return err
		}
		root, err := ch.scanner.Root(token.Contract, uint64(height))
		if err != nil {
			return err
		}
		resp.Root = fmt.Sprintf("%x", root)
	}
	data, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	return ctx.Send(data, api.HTTPstatusCodeOK)
}

func (ch *contractHandler) dumpBalances(msg *api.BearerStandardAPIdata, ctx *httprouter.HTTPContext) error {
	balances, err := ch.scanner.Balances(ctx.URLParam("contract"))
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(balances, "", " ")
	if err != nil {
		return err
	}
	return ctx.Send(data, api.HTTPstatusCodeOK)
}

func (ch *contractHandler) exportTree(msg *api.BearerStandardAPIdata, ctx *httprouter.HTTPContext) error {
	var err error
	resp := &Reply{Ok: true}
	resp.Block, err = ch.scanner.QueueExport(ctx.URLParam("contract"))
	if err != nil {
		return err
	}
	data, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	return ctx.Send(data, api.HTTPstatusCodeOK)
}

func (ch *contractHandler) fetchTree(msg *api.BearerStandardAPIdata, ctx *httprouter.HTTPContext) error {
	var err error
	block, err := strconv.Atoi(ctx.URLParam("blockNum"))
	if err != nil {
		return err
	}
	resp := &Reply{Ok: true}
	resp.Data, err = ch.scanner.FetchExport(ctx.URLParam("contract"), uint64(block))
	if err != nil {
		return err
	}
	data, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	return ctx.Send(data, api.HTTPstatusCodeOK)
}

func (ch *contractHandler) rescan(msg *api.BearerStandardAPIdata, ctx *httprouter.HTTPContext) error {
	if err := ch.scanner.RescanContract(ctx.URLParam("contract")); err != nil {
		return err
	}
	resp := &Reply{Ok: true}
	data, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	return ctx.Send(data, api.HTTPstatusCodeOK)
}
