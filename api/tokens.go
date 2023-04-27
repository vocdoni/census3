package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
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
	capi.endpoint.RegisterMethod("/token", "GET",
		api.MethodAccessTypePublic, capi.getTokens)
	capi.endpoint.RegisterMethod("/token", "POST",
		api.MethodAccessTypePublic, capi.createToken)
	capi.endpoint.RegisterMethod("/token/{tokenID}", "GET",
		api.MethodAccessTypePublic, capi.getToken)
	capi.endpoint.RegisterMethod("/token/types", "GET",
		api.MethodAccessTypePublic, capi.getTokenTypes)
}

// getTokens function handler returns the registered tokens information from the
// database. It returns a 204 response if no tokens are registered or a 500
// error if something fails.
func (capi *census3API) getTokens(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	internalCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	// TODO: Support for pagination
	// get tokens from the database
	rows, err := capi.sqlc.PaginatedTokens(internalCtx, queries.PaginatedTokensParams{
		Limit:  -1,
		Offset: 0,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNoTokens
		}
		log.Errorw(ErrCantGetTokens, err.Error())
		return ErrCantGetTokens
	}
	if len(rows) == 0 {
		return ErrNoTokens
	}
	// parse and encode resulting tokens
	tokens := GetTokensResponse{Tokens: []GetTokenResponse{}}
	for _, tokenData := range rows {
		tokens.Tokens = append(tokens.Tokens, GetTokenResponse{
			ID:         common.BytesToAddress(tokenData.ID).String(),
			Type:       state.TokenType(int(tokenData.TypeID)).String(),
			Decimals:   uint64(tokenData.Decimals.Int32),
			StartBlock: uint64(tokenData.CreationBlock),
			Name:       tokenData.Name.String,
		})
	}
	res, err := json.Marshal(tokens)
	if err != nil {
		log.Errorw(ErrEncodeTokens, err.Error())
		return ErrEncodeTokens
	}
	return ctx.Send(res, api.HTTPstatusOK)
}

// createToken function creates a new token in the current database instance. It
// first gets the token information from the network and then stores it in the
// database. The new token created will be scanned from the block number
// provided as argument. It returns a 400 error if the provided inputs are
// wrong or empty or a 500 error if something fails.
func (capi *census3API) createToken(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	req := CreateTokenRequest{}
	if err := json.Unmarshal(msg.Data, &req); err != nil {
		log.Errorf("error unmarshalling token information: %s", err)
		return ErrMalformedToken.With("error unmarshalling token information")
	}
	tokenType := state.TokenTypeFromString(req.Type)
	addr := common.HexToAddress(req.ID)

	w3 := state.Web3{}
	internalCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := w3.Init(internalCtx, capi.web3, addr, tokenType); err != nil {
		log.Errorw(ErrInitializingWeb3, err.Error())
		return ErrInitializingWeb3
	}
	info, err := w3.GetTokenData()
	if err != nil {
		log.Errorw(ErrCantGetToken, err.Error())
		return ErrCantGetToken
	}
	var (
		name     = new(sql.NullString)
		symbol   = new(sql.NullString)
		decimals = new(sql.NullInt32)
	)
	if err := name.Scan(info.Name); err != nil {
		log.Errorw(ErrCantGetToken, err.Error())
		return ErrCantGetToken
	}
	if err := symbol.Scan(info.Symbol); err != nil {
		log.Errorw(ErrCantGetToken, err.Error())
		return ErrCantGetToken
	}
	if err := decimals.Scan(info.Decimals); err != nil {
		log.Errorw(ErrCantGetToken, err.Error())
		return ErrCantGetToken
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
// from the database. It returns a 400 error if the provided ID is wrong or
// empty, a 404 error if the token is not found or a 500 error if something
// fails.
func (capi *census3API) getToken(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	address := common.HexToAddress(ctx.URLParam("tokenID"))
	internalCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	tokenData, err := capi.sqlc.TokenByID(internalCtx, address.Bytes())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Errorw(ErrNotFoundToken, err.Error())
			return ErrNotFoundToken
		}
		log.Errorw(ErrCantGetToken, err.Error())
		return ErrCantGetToken
	}

	// TODO: Only for the MVP, consider to remove it
	tokenStrategies, err := capi.sqlc.PaginatedStrategiesByTokenID(internalCtx,
		queries.PaginatedStrategiesByTokenIDParams{
			TokenID: tokenData.ID,
			Limit:   -1,
			Offset:  0,
		})
	log.Info(tokenStrategies)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Errorw(ErrCantGetToken, err.Error())
		return ErrCantGetToken
	}
	defaultStrategyID := uint64(0)
	if len(tokenStrategies) > 0 {
		defaultStrategyID = uint64(tokenStrategies[0].ID)
	}

	res, err := json.Marshal(GetTokenResponse{
		ID:          address.String(),
		Type:        state.TokenType(int(tokenData.TypeID)).String(),
		Decimals:    uint64(tokenData.Decimals.Int32),
		StartBlock:  uint64(tokenData.CreationBlock),
		Name:        tokenData.Name.String,
		Symbol:      tokenData.Symbol.String,
		TotalSupply: new(big.Int).SetBytes(tokenData.TotalSupply),
		Status: GetTokenStatusResponse{
			AtBlock:  20000000,
			Synced:   true,
			Progress: 100,
		},
		// TODO: Only for the MVP, consider to remove it
		DefaultStrategy: defaultStrategyID,
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
