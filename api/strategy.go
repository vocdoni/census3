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
	"go.vocdoni.io/dvote/httprouter"
	api "go.vocdoni.io/dvote/httprouter/apirest"
	"go.vocdoni.io/dvote/log"
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

// getStrategies function handler returns the current registered strategies from
// the database. It returns a 204 response if any strategy is registered or a
// 500 error if something fails.
func (capi *census3API) getStrategies(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	internalCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	// TODO: Support for pagination
	// get strategies from the database
	rows, err := capi.sqlc.ListStrategies(internalCtx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNoStrategies
		}
		log.Errorw(ErrCantGetStrategies, err.Error())
		return ErrCantGetStrategies
	}
	if len(rows) == 0 {
		return ErrNoStrategies
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
	// get provided strategyID
	strategyID, err := strconv.Atoi(ctx.URLParam("strategyID"))
	if err != nil {
		log.Errorw(ErrMalformedStrategyID, err.Error())
		return ErrMalformedStrategyID
	}
	// get strategy from the database
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
	// parse strategy information
	strategy := GetStrategyResponse{
		ID:        uint64(strategyData.ID),
		Predicate: strategyData.Predicate,
		Tokens:    []GetStrategyToken{},
	}
	// get information of the strategy related tokens
	tokensData, err := capi.sqlc.TokensByStrategyID(internalCtx, strategyData.ID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Errorw(ErrCantGetTokens, err.Error())
		return ErrCantGetTokens
	}
	// parse and encode tokens information
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
	// get the tokenID provided
	tokenID := ctx.URLParam("tokenID")
	internalCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	// get strategies associated to the token provided
	rows, err := capi.sqlc.StrategiesByTokenID(internalCtx, common.HexToAddress(tokenID).Bytes())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNoStrategies
		}
		log.Errorw(ErrCantGetStrategies, err.Error())
		return ErrCantGetStrategies
	}
	if len(rows) == 0 {
		return ErrNoStrategies
	}
	// parse and encode strategies
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
