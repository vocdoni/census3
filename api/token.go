package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/census3/contractstate"
	queries "github.com/vocdoni/census3/db/sqlc"
	"go.vocdoni.io/dvote/httprouter"
	api "go.vocdoni.io/dvote/httprouter/apirest"
	"go.vocdoni.io/dvote/log"
)

func (capi *census3API) initTokenHandlers() {
	capi.endpoint.RegisterMethod("/tokens", "POST", api.MethodAccessTypePublic,
		capi.createToken)
}

// createToken function creates a new token in the current database instance. It
// first gets the token information from the network and then stores it in the
// database. The new token created will be scanned from the block number
// provided as argument.
func (capi *census3API) createToken(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	req := CreateTokenRequest{}
	if err := json.Unmarshal(msg.Data, &req); err != nil {
		log.Errorf("error unmarshalling token information: %s", err)
		return ErrTokenMalformed.With("error unmarshalling token information")
	}

	tokenType := contractstate.ContractTypeFromString(req.Type)
	addr := common.HexToAddress(req.Address)

	w3 := contractstate.Web3{}
	ctx2, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := w3.Init(ctx2, capi.web3, addr, tokenType); err != nil {
		return err
	}
	info, err := w3.GetTokenData()
	if err != nil {
		log.Errorw(err, "error getting token contract data")
		return err
	}
	var (
		name     = new(sql.NullString)
		symbol   = new(sql.NullString)
		decimals = new(sql.NullInt32)
	)
	if err := name.Scan(info.Name); err != nil {
		return err
	}
	if err := symbol.Scan(info.Symbol); err != nil {
		return err
	}
	if err := decimals.Scan(info.Decimals); err != nil {
		return err
	}
	_, err = capi.sqlc.CreateToken(ctx2, queries.CreateTokenParams{
		ID:            info.Address.Bytes(),
		Name:          *name,
		Symbol:        *symbol,
		Decimals:      *decimals,
		TotalSupply:   info.TotalSupply.Bytes(),
		CreationBlock: int64(req.StartBlock),
		TypeID:        int64(tokenType),
	})
	if err != nil {
		log.Errorw(err, "error creating token on the database")
		return ErrCantCreateToken.Withf("error creating token with address %s", addr)
	}

	return ctx.Send([]byte("Ok"), api.HTTPstatusOK)
}