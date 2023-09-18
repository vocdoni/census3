package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	queries "github.com/vocdoni/census3/db/sqlc"
	"github.com/vocdoni/census3/state"
	"go.vocdoni.io/dvote/httprouter"
	api "go.vocdoni.io/dvote/httprouter/apirest"
	"go.vocdoni.io/dvote/log"
)

func (capi *census3API) initTokenHandlers() error {
	if err := capi.endpoint.RegisterMethod("/token", "GET",
		api.MethodAccessTypePublic, capi.getTokens); err != nil {
		return err
	}
	if err := capi.endpoint.RegisterMethod("/token", "POST",
		api.MethodAccessTypePublic, capi.createToken); err != nil {
		return err
	}
	if err := capi.endpoint.RegisterMethod("/token/{tokenID}", "GET",
		api.MethodAccessTypePublic, capi.getToken); err != nil {
		return err
	}
	return capi.endpoint.RegisterMethod("/token/types", "GET",
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
	rows, err := capi.db.QueriesRO.ListTokens(internalCtx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNoTokens.WithErr(err)
		}
		return ErrCantGetTokens.WithErr(err)
	}
	if len(rows) == 0 {
		return ErrNoTokens
	}
	// parse and encode resulting tokens
	tokens := GetTokensResponse{Tokens: []GetTokensItem{}}
	for _, tokenData := range rows {
		tokenResponse := GetTokensItem{
			ID:         common.BytesToAddress(tokenData.ID).String(),
			Type:       state.TokenType(int(tokenData.TypeID)).String(),
			Name:       tokenData.Name.String,
			StartBlock: tokenData.CreationBlock.Int64,
			Tag:        tokenData.Tag.String,
			Symbol:     tokenData.Symbol.String,
		}
		tokens.Tokens = append(tokens.Tokens, tokenResponse)
	}
	res, err := json.Marshal(tokens)
	if err != nil {
		return ErrEncodeTokens.WithErr(err)
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
		return ErrMalformedToken.WithErr(err)
	}
	tokenType := state.TokenTypeFromString(req.Type)
	addr := common.HexToAddress(req.ID)
	// init web3 client to get the token information before register in the
	// database
	w3 := state.Web3{}
	internalCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// get correct web3 uri provider
	w3uri, exists := capi.w3p[req.ChainID]
	if !exists {
		return ErrChainIDNotSupported.With("chain ID not supported")
	}
	if err := w3.Init(internalCtx, w3uri, addr, tokenType); err != nil {
		log.Errorw(ErrInitializingWeb3, err.Error())
		return ErrInitializingWeb3.WithErr(err)
	}
	info, err := w3.TokenData()
	if err != nil {
		return ErrCantGetToken.WithErr(err)
	}
	var (
		name          = new(sql.NullString)
		symbol        = new(sql.NullString)
		creationBlock = new(sql.NullInt64)
		totalSupply   = new(big.Int)
		tag           = new(sql.NullString)
	)
	if err := name.Scan(info.Name); err != nil {
		return ErrCantGetToken.WithErr(err)
	}
	if err := symbol.Scan(info.Symbol); err != nil {
		return ErrCantGetToken.WithErr(err)
	}
	if info.TotalSupply != nil {
		totalSupply = info.TotalSupply
	}
	if req.Tag != "" {
		if err := tag.Scan(req.Tag); err != nil {
			return ErrCantGetToken.WithErr(err)
		}
	}
	_, err = capi.db.QueriesRW.CreateToken(internalCtx, queries.CreateTokenParams{
		ID:            info.Address.Bytes(),
		Name:          *name,
		Symbol:        *symbol,
		Decimals:      info.Decimals,
		TotalSupply:   totalSupply.Bytes(),
		CreationBlock: *creationBlock,
		TypeID:        uint64(tokenType),
		Synced:        false,
		Tag:           *tag,
		ChainID:       req.ChainID,
	})
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return ErrTokenAlreadyExists.WithErr(err)
		}
		return ErrCantCreateToken.WithErr(err)
	}
	// TODO: Only for the MVP, consider to remove it
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
	tokenData, err := capi.db.QueriesRO.TokenByID(internalCtx, address.Bytes())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNotFoundToken.WithErr(err)
		}
		return ErrCantGetToken.WithErr(err)
	}
	// TODO: Only for the MVP, consider to remove it
	tokenStrategies, err := capi.db.QueriesRO.StrategiesByTokenID(internalCtx, tokenData.ID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return ErrCantGetToken.WithErr(err)
	}
	defaultStrategyID := uint64(0)
	if len(tokenStrategies) > 0 {
		defaultStrategyID = tokenStrategies[0].ID
	}
	// get last block with token information
	atBlock, err := capi.db.QueriesRO.LastBlockByTokenID(internalCtx, address.Bytes())
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return ErrCantGetToken.WithErr(err)
		}
		atBlock = 0
	}
	// if the token is not synced, get the last block of the network to
	// calculate the current scan progress
	tokenProgress := 100
	if !tokenData.Synced {
		// get correct web3 uri provider
		w3uri, exists := capi.w3p[tokenData.ChainID]
		if !exists {
			return ErrChainIDNotSupported.With("chain ID not supported")
		}
		// get last block of the network, if something fails return progress 0
		w3 := state.Web3{}
		if err := w3.Init(internalCtx, w3uri, address, state.TokenType(tokenData.TypeID)); err != nil {
			return ErrInitializingWeb3.WithErr(err)
		}
		// fetch the last block header and calculate progress
		lastBlockNumber, err := w3.LatestBlockNumber(internalCtx)
		if err != nil {
			return ErrCantGetLastBlockNumber.WithErr(err)
		}
		tokenProgress = int(float64(atBlock) / float64(lastBlockNumber) * 100)
	}

	// get token holders count
	countHoldersCtx, cancel2 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel2()
	holders, err := capi.db.QueriesRO.CountTokenHoldersByTokenID(countHoldersCtx, address.Bytes())
	if err != nil {
		return ErrCantGetTokenCount.WithErr(err)
	}

	// build response
	tokenResponse := GetTokenResponse{
		ID:          address.String(),
		Type:        state.TokenType(int(tokenData.TypeID)).String(),
		Decimals:    tokenData.Decimals,
		Size:        uint64(holders),
		Name:        tokenData.Name.String,
		Symbol:      tokenData.Symbol.String,
		TotalSupply: new(big.Int).SetBytes(tokenData.TotalSupply).String(),
		Status: &GetTokenStatusResponse{
			AtBlock:  atBlock,
			Synced:   tokenData.Synced,
			Progress: tokenProgress,
		},
		// TODO: Only for the MVP, consider to remove it
		Tag:             tokenData.Tag.String,
		DefaultStrategy: defaultStrategyID,
		ChainID:         tokenData.ChainID,
	}
	if tokenData.CreationBlock.Valid {
		tokenResponse.StartBlock = uint64(tokenData.CreationBlock.Int64)
	}
	res, err := json.Marshal(tokenResponse)
	if err != nil {
		return ErrEncodeToken.WithErr(err)
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
		return ErrEncodeTokenTypes.WithErr(err)
	}
	return ctx.Send(res, api.HTTPstatusOK)
}
