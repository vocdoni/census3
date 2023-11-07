package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	queries "github.com/vocdoni/census3/db/sqlc"
	"github.com/vocdoni/census3/lexer"
	"github.com/vocdoni/census3/state"
	"go.vocdoni.io/dvote/httprouter"
	api "go.vocdoni.io/dvote/httprouter/apirest"
	"go.vocdoni.io/dvote/log"
)

func (capi *census3API) initTokenHandlers() error {
	if err := capi.endpoint.RegisterMethod("/tokens", "GET",
		api.MethodAccessTypePublic, capi.getTokens); err != nil {
		return err
	}
	if err := capi.endpoint.RegisterMethod("/tokens", "POST",
		api.MethodAccessTypePublic, capi.createToken); err != nil {
		return err
	}
	if err := capi.endpoint.RegisterMethod("/tokens/{tokenID}", "GET",
		api.MethodAccessTypePublic, capi.getToken); err != nil {
		return err
	}
	if err := capi.endpoint.RegisterMethod("/tokens/{tokenID}/holders/{holderID}", "GET",
		api.MethodAccessTypePublic, capi.isTokenHolder); err != nil {
		return err
	}
	return capi.endpoint.RegisterMethod("/tokens/types", "GET",
		api.MethodAccessTypePublic, capi.getTokenTypes)
}

// getTokens function handler returns the registered tokens information from the
// database. It returns a 204 response if no tokens are registered or a 500
// error if something fails.
func (capi *census3API) getTokens(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	// get pagination information from the request
	pageSize, dbPageSize, cursor, goForward, err := paginationFromCtx(ctx)
	if err != nil {
		return ErrMalformedPagination.WithErr(err)
	}
	// if there is a cursor, decode it to bytes
	bCursor := []byte{}
	if cursor != "" {
		bCursor = common.HexToAddress(cursor).Bytes()
	}
	// init context with timeout and database transaction
	internalCtx, cancel := context.WithTimeout(context.Background(), getTokensTimeout)
	defer cancel()
	tx, err := capi.db.RO.BeginTx(internalCtx, nil)
	if err != nil {
		return ErrCantGetTokens.WithErr(err)
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Errorw(err, "error rolling back tokens transaction")
		}
	}()
	qtx := capi.db.QueriesRO.WithTx(tx)
	// get the tokens from the database using the provided cursor, get the
	// following or previous page depending on the direction of the cursor
	var rows []queries.Token
	if goForward {
		rows, err = qtx.NextTokensPage(internalCtx, queries.NextTokensPageParams{
			PageCursor: bCursor,
			Limit:      dbPageSize,
		})
	} else {
		rows, err = qtx.PrevTokensPage(internalCtx, queries.PrevTokensPageParams{
			PageCursor: bCursor,
			Limit:      dbPageSize,
		})
	}
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNoTokens.WithErr(err)
		}
		return ErrCantGetTokens.WithErr(err)
	}
	if len(rows) == 0 {
		return ErrNoTokens
	}
	// init response struct with the initial pagination information and empty
	// list of tokens
	tokensResponse := GetTokensResponse{
		Tokens:     []GetTokensItem{},
		Pagination: &Pagination{PageSize: pageSize},
	}
	rows, nextCursorRow, prevCursorRow := paginationToRequest(rows, dbPageSize, cursor, goForward)
	if nextCursorRow != nil {
		tokensResponse.Pagination.NextCursor = common.BytesToAddress(nextCursorRow.ID).String()
	}
	if prevCursorRow != nil {
		tokensResponse.Pagination.PrevCursor = common.BytesToAddress(prevCursorRow.ID).String()
	}
	// parse results from database to the response format
	for _, tokenData := range rows {
		tokensResponse.Tokens = append(tokensResponse.Tokens, GetTokensItem{
			ID:           common.BytesToAddress(tokenData.ID).String(),
			Type:         state.TokenType(int(tokenData.TypeID)).String(),
			Name:         tokenData.Name,
			StartBlock:   tokenData.CreationBlock,
			Tags:         tokenData.Tags,
			Symbol:       tokenData.Symbol,
			ChainID:      tokenData.ChainID,
			ChainAddress: tokenData.ChainAddress,
			ExternalID:   tokenData.ExternalID,
		})
	}
	// encode the response and send it
	res, err := json.Marshal(tokensResponse)
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
	internalCtx, cancel := context.WithTimeout(ctx.Request.Context(), createTokenTimeout)
	defer cancel()

	var info *state.TokenData
	tokenType := state.TokenTypeFromString(req.Type)
	if provider, exists := capi.extProviders[tokenType]; exists {
		// get token information from the external provider
		address, err := provider.Address(internalCtx, []byte(req.ExternalID))
		if err != nil {
			return ErrCantGetToken.WithErr(err)
		}
		name, err := provider.Name(internalCtx, []byte(req.ExternalID))
		if err != nil {
			return ErrCantGetToken.WithErr(err)
		}
		symbol, err := provider.Symbol(internalCtx, []byte(req.ExternalID))
		if err != nil {
			return ErrCantGetToken.WithErr(err)
		}
		decimals, err := provider.Decimals(internalCtx, []byte(req.ExternalID))
		if err != nil {
			return ErrCantGetToken.WithErr(err)
		}
		totalSupply, err := provider.TotalSupply(internalCtx, []byte(req.ExternalID))
		if err != nil {
			return ErrCantGetToken.WithErr(err)
		}
		// build token information struct with the data from the external
		// provider
		info = &state.TokenData{
			Type:        tokenType,
			Address:     address,
			Name:        name,
			Symbol:      symbol,
			Decimals:    decimals,
			TotalSupply: totalSupply,
		}
	} else {
		addr := common.HexToAddress(req.ID)
		// init web3 client to get the token information before register in the
		// database
		w3 := state.Web3{}
		// get correct web3 uri provider
		w3URI, exists := capi.w3p.URIByChainID(req.ChainID)
		if !exists {
			return ErrChainIDNotSupported.With("chain ID not supported")
		}
		// init web3 client to get the token information
		err := w3.Init(internalCtx, w3URI, addr, tokenType)
		if err != nil {
			return ErrInitializingWeb3.WithErr(err)
		}
		// get token information from the web3 client
		if info, err = w3.TokenData(); err != nil {
			return ErrCantGetToken.WithErr(err)
		}
	}
	// init db transaction
	tx, err := capi.db.RW.BeginTx(internalCtx, nil)
	if err != nil {
		return ErrCantCreateStrategy.WithErr(err)
	}
	defer func() {
		if err := tx.Rollback(); err != nil && !errors.Is(sql.ErrTxDone, err) {
			log.Errorw(err, "create strategy transaction rollback failed")
		}
	}()
	// get the chain address for the token based on the chainID and tokenID
	chainAddress, ok := capi.w3p.ChainAddress(req.ChainID, info.Address.String())
	if !ok {
		return ErrChainIDNotSupported.Withf("chainID: %d, tokenID: %s", req.ChainID, req.ID)
	}
	totalSupply := big.NewInt(0).Bytes()
	if info.TotalSupply != nil {
		totalSupply = info.TotalSupply.Bytes()
	}
	qtx := capi.db.QueriesRW.WithTx(tx)
	_, err = qtx.CreateToken(internalCtx, queries.CreateTokenParams{
		ID:            info.Address.Bytes(),
		Name:          info.Name,
		Symbol:        info.Symbol,
		Decimals:      info.Decimals,
		TotalSupply:   totalSupply,
		CreationBlock: 0,
		TypeID:        uint64(tokenType),
		Synced:        false,
		Tags:          req.Tags,
		ChainID:       req.ChainID,
		ChainAddress:  chainAddress,
		ExternalID:    req.ExternalID,
	})
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return ErrTokenAlreadyExists.WithErr(err)
		}
		return ErrCantCreateToken.WithErr(err)
	}
	// create a default strategy to support censuses over the holders of this
	// single token
	res, err := qtx.CreateStategy(internalCtx, queries.CreateStategyParams{
		Alias:     fmt.Sprintf("Default strategy for token %s", info.Symbol),
		Predicate: lexer.ScapeTokenSymbol(info.Symbol),
	})
	if err != nil {
		return err
	}
	strategyID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	if _, err := qtx.CreateStrategyToken(internalCtx, queries.CreateStrategyTokenParams{
		StrategyID: uint64(strategyID),
		TokenID:    info.Address.Bytes(),
		ChainID:    req.ChainID,
		MinBalance: big.NewInt(0).Bytes(),
	}); err != nil {
		return ErrCantGetToken.WithErr(err)
	}
	if _, err := qtx.UpdateTokenDefaultStrategy(internalCtx, queries.UpdateTokenDefaultStrategyParams{
		ID:              info.Address.Bytes(),
		DefaultStrategy: uint64(strategyID),
		ChainID:         req.ChainID,
		ExternalID:      req.ExternalID,
	}); err != nil {
		return ErrCantGetToken.WithErr(err)
	}
	if err := tx.Commit(); err != nil {
		return ErrCantGetToken.WithErr(err)
	}
	return ctx.Send([]byte("Ok"), api.HTTPstatusOK)
}

// getToken function handler returns the information of the given token address
// from the database. It returns a 400 error if the provided ID is wrong or
// empty, a 404 error if the token is not found or a 500 error if something
// fails.
func (capi *census3API) getToken(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	// get contract address from the tokenID query param and decode check if
	// it is provided, if not return an error
	strAddress := ctx.URLParam("tokenID")
	if strAddress == "" {
		return ErrMalformedToken.With("tokenID is required")
	}
	address := common.HexToAddress(strAddress)
	// get chainID from query params and decode it as integer, if it's not
	// provided or it's not a valid integer return an error
	strChainID := ctx.Request.URL.Query().Get("chainID")
	if strChainID == "" {
		return ErrMalformedChainID.With("chainID is required")
	}
	chainID, err := strconv.Atoi(strChainID)
	if err != nil {
		return ErrMalformedChainID.WithErr(err)
	} else if chainID < 0 {
		return ErrMalformedChainID.With("chainID must be a positive number")
	}
	// get externalID from query params and decode it as string, it is optional
	// so if it's not provided continue
	externalID := ctx.Request.URL.Query().Get("externalID")
	// get token information from the database
	internalCtx, cancel := context.WithTimeout(ctx.Request.Context(), getTokenTimeout)
	defer cancel()
	tokenData, err := capi.db.QueriesRO.TokenByIDAndChainIDAndExternalID(internalCtx,
		queries.TokenByIDAndChainIDAndExternalIDParams{
			ID:         address.Bytes(),
			ChainID:    uint64(chainID),
			ExternalID: externalID,
		})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNotFoundToken.WithErr(err)
		}
		return ErrCantGetToken.WithErr(err)
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
		w3URI, exists := capi.w3p.URIByChainID(tokenData.ChainID)
		if !exists {
			return ErrChainIDNotSupported.With("chain ID not supported")
		}
		// get last block of the network, if something fails return progress 0
		w3 := state.Web3{}
		if err := w3.Init(internalCtx, w3URI, address, state.TokenType(tokenData.TypeID)); err != nil {
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
	holders, err := capi.db.QueriesRO.CountTokenHoldersByTokenID(internalCtx, address.Bytes())
	if err != nil {
		return ErrCantGetTokenCount.WithErr(err)
	}
	// build response
	tokenResponse := GetTokenResponse{
		ID:          address.String(),
		Type:        state.TokenType(int(tokenData.TypeID)).String(),
		Decimals:    tokenData.Decimals,
		Size:        uint64(holders),
		Name:        tokenData.Name,
		Symbol:      tokenData.Symbol,
		TotalSupply: new(big.Int).SetBytes(tokenData.TotalSupply).String(),
		StartBlock:  uint64(tokenData.CreationBlock),
		Status: &GetTokenStatusResponse{
			AtBlock:  atBlock,
			Synced:   tokenData.Synced,
			Progress: tokenProgress,
		},
		Tags: tokenData.Tags,
		// TODO: Only for the MVP, consider to remove it
		DefaultStrategy: tokenData.DefaultStrategy,
		ChainID:         tokenData.ChainID,
		ChainAddress:    tokenData.ChainAddress,
		ExternalID:      tokenData.ExternalID,
	}
	res, err := json.Marshal(tokenResponse)
	if err != nil {
		return ErrEncodeToken.WithErr(err)
	}
	return ctx.Send(res, api.HTTPstatusOK)
}

func (capi *census3API) isTokenHolder(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	// get contract address from the tokenID query param and decode check if
	// it is provided, if not return an error
	strAddress := ctx.URLParam("tokenID")
	if strAddress == "" {
		return ErrMalformedToken.With("tokenID is required")
	}
	address := common.HexToAddress(strAddress)
	// get chainID from query params and decode it as integer, if it's not
	// provided or it's not a valid integer return an error
	strChainID := ctx.Request.URL.Query().Get("chainID")
	if strChainID == "" {
		return ErrMalformedChainID.With("chainID is required")
	}
	chainID, err := strconv.Atoi(strChainID)
	if err != nil {
		return ErrMalformedChainID.WithErr(err)
	} else if chainID < 0 {
		return ErrMalformedChainID.With("chainID must be a positive number")
	}
	// get holder address from the holderID query param and decode check if
	// it is provided, if not return an error
	strHolderID := ctx.URLParam("holderID")
	if strHolderID == "" {
		return ErrMalformedHolder.With("holderID is required")
	}
	holderID := common.HexToAddress(strHolderID)
	internalCtx, cancel := context.WithTimeout(ctx.Request.Context(), getTokenTimeout)
	defer cancel()

	exists, err := capi.db.QueriesRO.ExistTokenHolder(internalCtx, queries.ExistTokenHolderParams{
		TokenID:  address.Bytes(),
		HolderID: holderID.Bytes(),
		ChainID:  uint64(chainID),
	})
	if err != nil {
		return ErrCantGetTokenHolders.WithErr(err)
	}
	return ctx.Send([]byte(strconv.FormatBool(exists)), api.HTTPstatusOK)
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
