package api

import (
	"github.com/vocdoni/multirpc/endpoint"
	"github.com/vocdoni/multirpc/router"
	"github.com/vocdoni/multirpc/transports"
	"github.com/vocdoni/tokenstate/service"
	"gitlab.com/vocdoni/go-dvote/crypto/ethereum"
)

type TokenAPI struct {
	ID        string             `json:"request"`
	Method    string             `json:"method,omitempty"`
	Contract  string             `json:"contract,omitempty"`
	Contracts []string           `json:"contracts,omitempty"`
	Timestamp int32              `json:"timestamp"`
	Error     string             `json:"error,omitempty"`
	Block     int64              `json:"block,omitempty"`
	Root      string             `json:"root,omitempty"`
	Proof     string             `json:"proof,omitempty"`
	Address   string             `json:"address,omitempty"`
	Valid     bool               `json:"valid,omitempty"`
	Ok        bool               `json:"ok,omitempty"`
	Token     *service.TokenInfo `json:"token,omitempty"`
}

func (ta *TokenAPI) GetID() string {
	return ta.ID
}

func (ta *TokenAPI) SetID(id string) {
	ta.ID = id
}

func (ta *TokenAPI) SetTimestamp(ts int32) {
	ta.Timestamp = ts
}

func (ta *TokenAPI) SetError(e string) {
	ta.Error = e
}

func (ta *TokenAPI) GetMethod() string {
	return ta.Method
}

func NewAPI() transports.MessageAPI {
	return &TokenAPI{}
}

func Init(host string, port int32, signer *ethereum.SignKeys, scanner *service.Scanner) error {
	ep := endpoint.HTTPWSEndPoint{}
	ep.SetOption("listenHost", host)
	ep.SetOption("listenPort", port)
	listener := make(chan transports.Message)
	if err := ep.Init(listener); err != nil {
		return err
	}
	transportMap := make(map[string]transports.Transport)
	transportMap[ep.ID()] = ep.Transport()
	r := router.NewRouter(listener, transportMap, signer, NewAPI)
	if err := r.Transports[ep.ID()].AddNamespace("/api"); err != nil {
		return err
	}

	th := tokenHandler{scanner: scanner}

	r.AddHandler("ping", "/api", ping, false, true)
	if err := r.AddHandler("addContract", "/api", th.addContract, false, true); err != nil {
		return err
	}
	if err := r.AddHandler("listContracts", "/api", th.listContracts, false, true); err != nil {
		return err
	}
	if err := r.AddHandler("getContract", "/api", th.getContract, false, true); err != nil {
		return err
	}
	go r.Route()
	return nil
}

type tokenHandler struct {
	scanner *service.Scanner
}

func ping(rr router.RouterRequest) {
	msg := &TokenAPI{}
	rr.Send(router.BuildReply(msg, rr))
}

func (th *tokenHandler) addContract(rr router.RouterRequest) {
	reqmsg := rr.Message.(*TokenAPI)
	msg := &TokenAPI{}
	msg.ID = rr.Id
	err := th.scanner.AddContract(reqmsg.Contract)
	if err != nil {
		msg.Error = err.Error()
		msg.Ok = false
		rr.Send(router.BuildReply(msg, rr))
	} else {
		msg.Ok = true
		rr.Send(router.BuildReply(msg, rr))
	}
}

func (th *tokenHandler) listContracts(rr router.RouterRequest) {
	msg := &TokenAPI{}
	msg.Contracts = th.scanner.ListContracts()
	msg.Ok = true
	rr.Send(router.BuildReply(msg, rr))
}

func (th *tokenHandler) getContract(rr router.RouterRequest) {
	var err error
	reqmsg := rr.Message.(*TokenAPI)
	msg := &TokenAPI{}
	msg.Token, err = th.scanner.GetContract(reqmsg.Contract)
	if err != nil {
		msg.Ok = false
		msg.Error = err.Error()
		rr.Send(router.BuildReply(msg, rr))
	} else {
		msg.Ok = true
		rr.Send(router.BuildReply(msg, rr))
	}
}
