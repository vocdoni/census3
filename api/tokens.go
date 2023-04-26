package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"github.com/ethereum/go-ethereum/common"
	queries "github.com/vocdoni/census3/db/sqlc"
	"github.com/vocdoni/census3/state"
	"go.vocdoni.io/dvote/httprouter"
	api "go.vocdoni.io/dvote/httprouter/apirest"
	"go.vocdoni.io/dvote/log"
)

func (capi *census3API) initTokenHandlers() {
	capi.endpoint.RegisterMethod("/tokens", "GET",
		api.MethodAccessTypePublic, capi.getTokens)
	capi.endpoint.RegisterMethod("/tokens", "POST",
		api.MethodAccessTypePublic, capi.createToken)
	capi.endpoint.RegisterMethod("/tokens/{tokenID}", "GET",
		api.MethodAccessTypePublic, capi.getToken)
	capi.endpoint.RegisterMethod("/tokens/types", "GET",
		api.MethodAccessTypePublic, capi.getTokenTypes)
}

func (capi *census3API) getTokens(msg *api.APIdata, ctx *httprouter.HTTPContext) error { return nil }

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
	internalCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := w3.Init(internalCtx, capi.web3, addr, tokenType); err != nil {
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
	_, err = capi.sqlc.CreateToken(internalCtx, queries.CreateTokenParams{
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

// getToken function handler returns the information of the given token address
// from the database.
func (capi *census3API) getToken(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	address := common.HexToAddress(ctx.URLParam("tokenID"))
	internalCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	tokenData, err := capi.sqlc.TokenByID(internalCtx, address.Bytes())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Errorw(ErrUnknownToken, err.Error())
			return ErrUnknownToken
		}
		log.Errorw(ErrCantGetToken, err.Error())
		return ErrCantGetToken
	}

	res, err := json.Marshal(TokenResponse{
		ID:         address.String(),
		Type:       state.TokenType(int(tokenData.TypeID)).String(),
		Decimals:   int(tokenData.Decimals.Int32),
		StartBlock: uint64(tokenData.CreationBlock),
		Name:       tokenData.Name.String,
	})
	if err != nil {
		return ErrEncodeToken
	}
	return ctx.Send(res, api.HTTPstatusOK)
}

// getTokenTypes handler returns the list of string names of the currently
// supported types of token contracts.
func (capi *census3API) getTokenTypes(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	supportedTypes := []string{}
	for _, supportedType := range state.TokenTypeStringMap {
		supportedTypes = append(supportedTypes, supportedType)
	}

	res, err := json.Marshal(TokenTypesResponse{supportedTypes})
	if err != nil {
		return ErrEncodeTokenTypes
	}
	return ctx.Send(res, api.HTTPstatusOK)
}
