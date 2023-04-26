package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	queries "github.com/vocdoni/census3/db/sqlc"
	"github.com/vocdoni/census3/state"
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
		return ErrMalformedToken.With("error unmarshalling token information")
	}
	tokenType := state.TokenTypeFromString(req.Type)
	addr := common.HexToAddress(req.Address)

	w3 := state.Web3{}
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

	if err := capi.createDummyStrategy(info.Address.Bytes()); err != nil {
		log.Warn(err, "error creating dummy strategy for this token")
	}

	return ctx.Send([]byte("Ok"), api.HTTPstatusOK)
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
