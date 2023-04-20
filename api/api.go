package api

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/census3/contractstate"
	"github.com/vocdoni/census3/service"
	"go.vocdoni.io/dvote/crypto/ethereum"
	"go.vocdoni.io/dvote/httprouter"
	api "go.vocdoni.io/dvote/httprouter/apirest"
)

type Reply struct {
	Contracts []string           `json:"contracts,omitempty"`
	Token     *service.TokenInfo `json:"token,omitempty"`
	Root      string             `json:"root,omitempty"`
	Data      []byte             `json:"data,omitempty"`
	Block     uint64             `json:"block,omitempty"`
	Ok        bool               `json:"ok"`
}

func Init(host string, port int32, signer *ethereum.SignKeys, scanner *service.Scanner, holderScanner *service.HoldersScanner) error {
	r := httprouter.HTTProuter{}
	err := r.Init(host, int(port))
	if err != nil {
		return err
	}
	endpoint, err := api.NewAPI(&r, "/api")
	if err != nil {
		return err
	}
	ch := contractHandler{scanner: scanner}
	endpoint.RegisterMethod("/supportedContracts", "GET",
		api.MethodAccessTypePublic, ch.supportedContracts)
	// TODO: Start block not required, get the block where the contract was deployed.
	// Using a startblock can lead to errors on token holders and balances as some logs
	// can be missed.
	// To get the block where the contract was deployed use the GetTransactionReceipt JSONRPC call
	endpoint.RegisterMethod("/addContract/{contract}/{type}/{startBlock}", "GET",
		api.MethodAccessTypePublic, ch.addContract)
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

	// init token holders new methods
	initHoldersHandlers(endpoint, holderScanner)
	return nil
}

type contractHandler struct {
	scanner *service.Scanner
}

func (ch *contractHandler) supportedContracts(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	resp := Reply{Ok: true, Contracts: make([]string, 0)}
	for k := range contractstate.ContractTypeIntMap {
		resp.Contracts = append(resp.Contracts, k)
	}
	data, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	return ctx.Send(data, api.HTTPstatusOK)
}

func (ch *contractHandler) addContract(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	startBlock := ctx.URLParam("startBlock")
	if startBlock == "" {
		startBlock = "0"
	}
	startBlockU64, err := strconv.ParseUint(startBlock, 10, 64)
	if err != nil {
		return err
	}
	// get contract type from url param type
	cType := ctx.URLParam("type")
	if cType == "" {
		return fmt.Errorf("contract type not specified")
	}
	ccType := contractstate.ContractTypeFromString(cType)
	contract := common.HexToAddress(ctx.URLParam("contract"))
	if err = ch.scanner.AddContract(contract, ccType, startBlockU64); err != nil {
		return err
	}
	resp := Reply{Ok: true}
	data, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	return ctx.Send(data, api.HTTPstatusOK)
}

func (ch *contractHandler) listContracts(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	resp := &Reply{Ok: true}
	contracts := ch.scanner.ListContracts()
	for _, c := range contracts {
		resp.Contracts = append(resp.Contracts, c.Hex())
	}
	data, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	return ctx.Send(data, api.HTTPstatusOK)
}

func (ch *contractHandler) getContract(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	var err error
	resp := &Reply{Ok: true}
	resp.Token, err = ch.scanner.GetContract(common.HexToAddress(ctx.URLParam("contract")))
	if err != nil {
		return err
	}

	data, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	return ctx.Send(data, api.HTTPstatusOK)
}

func (ch *contractHandler) root(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	token, err := ch.scanner.GetContract(common.HexToAddress(ctx.URLParam("contract")))
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
		root, err := ch.scanner.Root(token.Address, uint64(height))
		if err != nil {
			return err
		}
		resp.Root = fmt.Sprintf("%x", root)
	}
	data, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	return ctx.Send(data, api.HTTPstatusOK)
}

func (ch *contractHandler) dumpBalances(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	balances, err := ch.scanner.Balances(common.HexToAddress(ctx.URLParam("contract")))
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(balances, "", " ")
	if err != nil {
		return err
	}
	return ctx.Send(data, api.HTTPstatusOK)
}

func (ch *contractHandler) exportTree(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	var err error
	resp := &Reply{Ok: true}
	resp.Block, err = ch.scanner.QueueExport(common.HexToAddress(ctx.URLParam("contract")))
	if err != nil {
		return err
	}
	data, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	return ctx.Send(data, api.HTTPstatusOK)
}

func (ch *contractHandler) fetchTree(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	var err error
	block, err := strconv.Atoi(ctx.URLParam("blockNum"))
	if err != nil {
		return err
	}
	resp := &Reply{Ok: true}
	resp.Data, err = ch.scanner.FetchExport(common.HexToAddress(ctx.URLParam("contract")), uint64(block))
	if err != nil {
		return err
	}
	data, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	return ctx.Send(data, api.HTTPstatusOK)
}

func (ch *contractHandler) rescan(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	if err := ch.scanner.RescanContract(common.HexToAddress(ctx.URLParam("contract"))); err != nil {
		return err
	}
	resp := &Reply{Ok: true}
	data, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	return ctx.Send(data, api.HTTPstatusOK)
}
