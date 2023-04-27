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
	"go.vocdoni.io/dvote/log"
)

func (capi *census3API) initStrategiesHandlers() {
	capi.endpoint.RegisterMethod("/strategies", "GET",
		api.MethodAccessTypePublic, capi.getStrategies)
	capi.endpoint.RegisterMethod("/strategies/{strategyID}", "GET",
		api.MethodAccessTypePublic, capi.getStrategy)
	capi.endpoint.RegisterMethod("/strategies/token/{tokenID}", "GET",
		api.MethodAccessTypePublic, capi.getTokenStrategies)
}

// getStrategies function handler returns the current registered strategies from
// the database. It returns a 204 response if any strategy is registered or a
// 500 error if something fails.
func (capi *census3API) getStrategies(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	internalCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	// TODO: Support for pagination
	// get strategies from the database
	rows, err := capi.sqlc.PaginatedStrategies(internalCtx, queries.PaginatedStrategiesParams{
		Limit:  -1,
		Offset: 0,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNoStrategies
		}
		log.Errorw(ErrCantGetStrategies, err.Error())
		return ErrCantGetStrategies
	}
	// parse and encode the strategies
	strategies := GetStrategiesResponse{Strategies: []uint64{}}
	for _, strategy := range rows {
		strategies.Strategies = append(strategies.Strategies, uint64(strategy.ID))
	}
	res, err := json.Marshal(strategies)
	if err != nil {
		log.Errorw(ErrEncodeStrategies, err.Error())
		return ErrEncodeStrategies
	}
	return ctx.Send(res, api.HTTPstatusOK)
}

// getStrategy function handler return the information of the strategy
// indetified by the ID provided. It returns a 400 error if the provided ID is
// wrong or empty, a 404 error if the strategy is not found or a 500 error if
// something fails.
func (capi *census3API) getStrategy(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	strategyID, err := strconv.Atoi(ctx.URLParam("strategyID"))
	if err != nil {
		log.Errorw(ErrMalformedStrategyID, err.Error())
		return ErrMalformedStrategyID
	}
	internalCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	strategyData, err := capi.sqlc.StrategyByID(internalCtx, int64(strategyID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNotFoundStrategy
		}
		log.Errorw(ErrCantGetStrategy, err.Error())
		return ErrCantGetStrategy
	}
	strategy := GetStrategyResponse{
		ID:        uint64(strategyData.ID),
		Predicate: strategyData.Predicate,
		Tokens:    []GetStrategyToken{},
	}

	tokensData, err := capi.sqlc.TokensByStrategyID(internalCtx, queries.TokensByStrategyIDParams{
		StrategyID: strategyData.ID,
		Limit:      -1,
		Offset:     0,
	})
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Errorw(ErrCantGetTokens, err.Error())
		return ErrCantGetTokens
	}
	for _, tokenData := range tokensData {
		strategy.Tokens = append(strategy.Tokens, GetStrategyToken{
			ID:         common.BytesToAddress(tokenData.ID).String(),
			Name:       tokenData.Name.String,
			MinBalance: new(big.Int).SetBytes(tokenData.MinBalance),
			Method:     common.Bytes2Hex(tokenData.MethodHash),
		})
	}

	res, err := json.Marshal(strategy)
	if err != nil {
		log.Errorw(ErrEncodeStrategy, err.Error())
		return ErrEncodeStrategy
	}
	return ctx.Send(res, api.HTTPstatusOK)
}

// getTokenStrategies function handler returns the strategies that involves the
// token identified by the ID (token address) provided. It returns a 400 error
// if the provided ID is wrong or empty, a 204 response if the token has not any
// associated strategy or a 500 error if something fails.
func (capi *census3API) getTokenStrategies(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	tokenID := ctx.URLParam("tokenID")
	internalCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	rows, err := capi.sqlc.PaginatedStrategiesByTokenID(internalCtx, queries.PaginatedStrategiesByTokenIDParams{
		TokenID: common.HexToAddress(tokenID).Bytes(),
		Limit:   -1,
		Offset:  0,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNoStrategies
		}
		log.Errorw(ErrCantGetStrategies, err.Error())
		return ErrCantGetStrategies
	}

	strategies := GetStrategiesResponse{Strategies: []uint64{}}
	for _, tokenStrategy := range rows {
		strategies.Strategies = append(strategies.Strategies, uint64(tokenStrategy.ID))
	}

	res, err := json.Marshal(strategies)
	if err != nil {
		log.Errorw(ErrEncodeStrategies, err.Error())
		return ErrEncodeStrategies
	}
	return ctx.Send(res, api.HTTPstatusOK)
}
