package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	queries "github.com/vocdoni/census3/db/sqlc"
	"go.vocdoni.io/dvote/httprouter"
	api "go.vocdoni.io/dvote/httprouter/apirest"
)

func (capi *census3API) initStrategiesHandlers() error {
	if err := capi.endpoint.RegisterMethod("/strategy/", "GET",
		api.MethodAccessTypePublic, capi.getStrategies); err != nil {
		return err
	}
	if err := capi.endpoint.RegisterMethod("/strategy/{strategyID}", "GET",
		api.MethodAccessTypePublic, capi.getStrategy); err != nil {
		return err
	}
	return capi.endpoint.RegisterMethod("/strategy/token/{tokenID}", "GET",
		api.MethodAccessTypePublic, capi.getTokenStrategies)
}

// createDummyStrategy creates the default strategy for a given token. This
// basic strategy only includes the holders of the given token which have a
// balance positive balance (holder_balance > 0).
//
// TODO: Only for the MVP, remove it.
func (capi *census3API) createDummyStrategy(tokenID []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := capi.db.QueriesRW.CreateStategy(ctx, "test")
	if err != nil {
		return err
	}
	strategyID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	_, err = capi.db.QueriesRW.CreateStrategyToken(ctx, queries.CreateStrategyTokenParams{
		StrategyID: uint64(strategyID),
		TokenID:    tokenID,
		MinBalance: big.NewInt(0).Bytes(),
		MethodHash: []byte("test"),
	})
	return err
}

// getStrategies function handler returns the current registered strategies from
// the database. It returns a 204 response if any strategy is registered or a
// 500 error if something fails.
func (capi *census3API) getStrategies(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	internalCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	// TODO: Support for pagination
	// get strategies from the database
	rows, err := capi.db.QueriesRO.ListStrategies(internalCtx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNoStrategies.WithErr(err)
		}
		return ErrCantGetStrategies.WithErr(err)
	}
	if len(rows) == 0 {
		return ErrNoStrategies
	}
	// parse and encode the strategies
	strategies := GetStrategiesResponse{Strategies: []uint64{}}
	for _, strategy := range rows {
		strategies.Strategies = append(strategies.Strategies, strategy.ID)
	}
	res, err := json.Marshal(strategies)
	if err != nil {
		return ErrEncodeStrategies.WithErr(err)
	}
	return ctx.Send(res, api.HTTPstatusOK)
}

// getStrategy function handler return the information of the strategy
// indetified by the ID provided. It returns a 400 error if the provided ID is
// wrong or empty, a 404 error if the strategy is not found or a 500 error if
// something fails.
func (capi *census3API) getStrategy(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	// get provided strategyID
	iStrategyID, err := strconv.Atoi(ctx.URLParam("strategyID"))
	if err != nil {
		return ErrMalformedStrategyID.WithErr(err)
	}
	strategyID := uint64(iStrategyID)
	// get strategy from the database
	internalCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	strategyData, err := capi.db.QueriesRO.StrategyByID(internalCtx, strategyID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNotFoundStrategy.WithErr(err)
		}
		return ErrCantGetStrategy.WithErr(err)
	}
	// parse strategy information
	strategy := GetStrategyResponse{
		ID:        strategyData.ID,
		Predicate: strategyData.Predicate,
		Tokens:    []GetStrategyToken{},
	}
	// get information of the strategy related tokens
	tokensData, err := capi.db.QueriesRO.TokensByStrategyID(internalCtx, strategyData.ID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return ErrCantGetTokens.WithErr(err)
	}
	// parse and encode tokens information
	for _, tokenData := range tokensData {
		strategy.Tokens = append(strategy.Tokens, GetStrategyToken{
			ID:         common.BytesToAddress(tokenData.ID).String(),
			Name:       tokenData.Name.String,
			MinBalance: new(big.Int).SetBytes(tokenData.MinBalance).String(),
			Method:     common.Bytes2Hex(tokenData.MethodHash),
		})
	}
	res, err := json.Marshal(strategy)
	if err != nil {
		return ErrEncodeStrategy.WithErr(err)
	}
	return ctx.Send(res, api.HTTPstatusOK)
}

// getTokenStrategies function handler returns the strategies that involves the
// token identified by the ID (token address) provided. It returns a 400 error
// if the provided ID is wrong or empty, a 204 response if the token has not any
// associated strategy or a 500 error if something fails.
func (capi *census3API) getTokenStrategies(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	// get the tokenID provided
	tokenID := ctx.URLParam("tokenID")
	internalCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	// get strategies associated to the token provided
	rows, err := capi.db.QueriesRO.StrategiesByTokenID(internalCtx, common.HexToAddress(tokenID).Bytes())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNoStrategies.WithErr(err)
		}
		return ErrCantGetStrategies.WithErr(err)
	}
	if len(rows) == 0 {
		return ErrNoStrategies
	}
	// parse and encode strategies
	strategies := GetStrategiesResponse{Strategies: []uint64{}}
	for _, tokenStrategy := range rows {
		strategies.Strategies = append(strategies.Strategies, tokenStrategy.ID)
	}
	res, err := json.Marshal(strategies)
	if err != nil {
		return ErrEncodeStrategies.WithErr(err)
	}
	return ctx.Send(res, api.HTTPstatusOK)
}
