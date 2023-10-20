package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	queries "github.com/vocdoni/census3/db/sqlc"
	"github.com/vocdoni/census3/lexer"
	"github.com/vocdoni/census3/strategyoperators"
	"go.vocdoni.io/dvote/httprouter"
	api "go.vocdoni.io/dvote/httprouter/apirest"
	"go.vocdoni.io/dvote/log"
)

func (capi *census3API) initStrategiesHandlers() error {
	if err := capi.endpoint.RegisterMethod("/strategies", "GET",
		api.MethodAccessTypePublic, capi.getStrategies); err != nil {
		return err
	}
	if err := capi.endpoint.RegisterMethod("/strategies", "POST",
		api.MethodAccessTypePublic, capi.createStrategy); err != nil {
		return err
	}
	if err := capi.endpoint.RegisterMethod("/strategies/import/{ipfsCID}", "POST",
		api.MethodAccessTypePublic, capi.importStrategy); err != nil {
		return err
	}
	if err := capi.endpoint.RegisterMethod("/strategies/{strategyID}", "GET",
		api.MethodAccessTypePublic, capi.getStrategy); err != nil {
		return err
	}
	if err := capi.endpoint.RegisterMethod("/strategies/token/{tokenID}", "GET",
		api.MethodAccessTypePublic, capi.getTokenStrategies); err != nil {
		return err
	}
	if err := capi.endpoint.RegisterMethod("/strategies/predicate/parse", "POST",
		api.MethodAccessTypePublic, capi.validateStrategyPredicate); err != nil {
		return err
	}
	if err := capi.endpoint.RegisterMethod("/strategies/predicate/operators", "GET",
		api.MethodAccessTypePublic, capi.supportedStrategyPredicateOperators); err != nil {
		return err
	}
	return nil
}

// getStrategies function handler returns the current registered strategies from
// the database. It returns a 204 response if any strategy is registered or a
// 500 error if something fails.
func (capi *census3API) getStrategies(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	internalCtx, cancel := context.WithTimeout(ctx.Request.Context(), getStrategiesTimeout)
	defer cancel()
	// create db transaction
	tx, err := capi.db.RO.BeginTx(internalCtx, nil)
	if err != nil {
		return ErrCantGetStrategies.WithErr(err)
	}
	defer func() {
		if err := tx.Rollback(); err != nil && !errors.Is(sql.ErrTxDone, err) {
			log.Errorw(err, "get strategies transaction rollback failed")
		}
	}()
	qtx := capi.db.QueriesRO.WithTx(tx)
	// get strategies associated to the token provided
	rows, err := qtx.ListStrategies(internalCtx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNoStrategies.WithErr(err)
		}
		return ErrCantGetStrategies.WithErr(err)
	}
	// parse and encode strategies
	strategies := GetStrategiesResponse{Strategies: []*GetStrategyResponse{}}
	for _, strategy := range rows {
		strategyResponse := &GetStrategyResponse{
			ID:        strategy.ID,
			Alias:     strategy.Alias,
			Predicate: strategy.Predicate,
			URI:       strategy.Uri,
			Tokens:    make(map[string]*StrategyToken),
		}
		strategyTokens, err := qtx.StrategyTokensByStrategyID(internalCtx, strategy.ID)
		if err != nil {
			return ErrCantGetStrategies.WithErr(err)
		}
		for _, strategyToken := range strategyTokens {
			if strategyToken.Symbol == "" {
				return ErrCantGetStrategies.With("invalid token symbol")
			}
			strategyResponse.Tokens[strategyToken.Symbol] = &StrategyToken{
				ID:         common.BytesToAddress(strategyToken.TokenID).String(),
				ChainID:    strategyToken.ChainID,
				MinBalance: new(big.Int).SetBytes(strategyToken.MinBalance).String(),
			}
		}
		strategies.Strategies = append(strategies.Strategies, strategyResponse)
	}
	res, err := json.Marshal(strategies)
	if err != nil {
		return ErrEncodeStrategies.WithErr(err)
	}
	return ctx.Send(res, api.HTTPstatusOK)
}

func (capi *census3API) createStrategy(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	internalCtx, cancel := context.WithTimeout(ctx.Request.Context(), createDummyStrategyTimeout)
	defer cancel()

	req := CreateStrategyRequest{}
	if err := json.Unmarshal(msg.Data, &req); err != nil {
		return ErrMalformedStrategy.WithErr(err)
	}
	if req.Predicate == "" || req.Alias == "" {
		return ErrMalformedStrategy.With("no predicate or alias provided")
	}
	// check predicate
	lx := lexer.NewLexer(strategyoperators.ValidOperatorsTags)
	validatedPredicate, err := lx.Parse(req.Predicate)
	if err != nil {
		return ErrInvalidStrategyPredicate.WithErr(err)
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
	qtx := capi.db.QueriesRW.WithTx(tx)
	// create the strategy to get the ID and then create the strategy tokens
	// with it
	result, err := qtx.CreateStategy(internalCtx, queries.CreateStategyParams{
		Alias:     req.Alias,
		Predicate: req.Predicate,
	})
	if err != nil {
		return ErrCantCreateStrategy.WithErr(err)
	}
	strategyID, err := result.LastInsertId()
	if err != nil {
		return ErrCantCreateStrategy.WithErr(err)
	}
	// iterate over the token symbols included in the predicate
	for _, symbol := range validatedPredicate.AllLiterals() {
		// check if the request includes the token information
		tokenData, ok := req.Tokens[symbol]
		if !ok {
			return ErrNoEnoughtStrategyTokens.Withf("undefined ID and chainID for symbol %s", symbol)
		}
		// check if the token exists in the database
		exists, err := qtx.ExistsTokenByChainID(internalCtx, queries.ExistsTokenByChainIDParams{
			ID:      common.HexToAddress(tokenData.ID).Bytes(),
			ChainID: tokenData.ChainID,
		})
		if err != nil {
			return ErrCantCreateStrategy.WithErr(err)
		}
		if !exists {
			return ErrNotFoundToken.Withf("the token with symbol %s not found", symbol)
		}
		// decode the min balance for the current token if it is provided,
		// if not use zero
		minBalance := new(big.Int)
		if tokenData.MinBalance != "" {
			if _, ok := minBalance.SetString(tokenData.MinBalance, 10); !ok {
				return ErrEncodeStrategy.Withf("error with %s minBalance", symbol)
			}
		}
		// create the strategy_token in the database
		if _, err := qtx.CreateStrategyToken(internalCtx, queries.CreateStrategyTokenParams{
			StrategyID: uint64(strategyID),
			TokenID:    common.HexToAddress(tokenData.ID).Bytes(),
			MinBalance: minBalance.Bytes(),
		}); err != nil {
			return ErrCantCreateStrategy.WithErr(err)
		}
	}
	// encode and compose final strategy data using the response of GET
	// strategy endpoint
	strategyDump, err := json.Marshal(GetStrategyResponse{
		ID:        uint64(strategyID),
		Alias:     req.Alias,
		Predicate: req.Predicate,
		Tokens:    req.Tokens,
	})
	if err != nil {
		return ErrEncodeStrategy.WithErr(err)
	}
	// publish the strategy to IPFS and update the database
	uri, err := capi.storage.Publish(internalCtx, strategyDump)
	if err != nil {
		return ErrCantCreateStrategy.WithErr(err)
	}
	if _, err := qtx.UpdateStrategyIPFSUri(internalCtx, queries.UpdateStrategyIPFSUriParams{
		ID:  uint64(strategyID),
		Uri: capi.storage.URIprefix() + uri,
	}); err != nil {
		return ErrCantCreateStrategy.WithErr(err)
	}
	// commit the transaction and return the strategyID
	if err := tx.Commit(); err != nil {
		return ErrCantCreateStrategy.WithErr(err)
	}
	response, err := json.Marshal(map[string]any{"strategyID": strategyID})
	if err != nil {
		return ErrEncodeStrategy.WithErr(err)
	}
	return ctx.Send(response, api.HTTPstatusOK)
}

// importStrategy function handler imports a strategy from IPFS and stores it
// into the database. It returns a 400 error if the provided IPFS CID is wrong
// or empty, a 500 error if something fails. It returns the strategyID of the
// imported strategy.
func (capi *census3API) importStrategy(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	// get the ipfsCID from the url
	ipfsCID := ctx.URLParam("ipfsCID")
	if ipfsCID == "" {
		return ErrMalformedStrategy.With("no ipfsCID provided")
	}
	// init the internal context
	internalCtx, cancel := context.WithTimeout(ctx.Request.Context(), importStrategyTimeout)
	defer cancel()
	// get the strategy from IPFS and decode it
	dump, err := capi.storage.Retrieve(internalCtx, ipfsCID, 0)
	if err != nil {
		return ErrCantImportStrategy.WithErr(err)
	}
	importedStrategy := GetStrategyResponse{}
	if err := json.Unmarshal(dump, &importedStrategy); err != nil {
		return ErrCantImportStrategy.WithErr(err)
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
	qtx := capi.db.QueriesRW.WithTx(tx)
	// create the strategy to get the ID and then create the strategy tokens
	result, err := qtx.CreateStategy(internalCtx, queries.CreateStategyParams{
		Alias:     importedStrategy.Alias,
		Predicate: importedStrategy.Predicate,
		Uri:       importedStrategy.URI,
	})
	if err != nil {
		return ErrCantCreateStrategy.WithErr(err)
	}
	strategyID, err := result.LastInsertId()
	if err != nil {
		return ErrCantCreateStrategy.WithErr(err)
	}
	// iterate over the token included in the predicate and create them in the
	// database
	for symbol, token := range importedStrategy.Tokens {
		// decode the min balance for the current token if it is provided,
		// if not use zero
		minBalance := new(big.Int)
		if token.MinBalance != "" {
			if _, ok := minBalance.SetString(token.MinBalance, 10); !ok {
				return ErrEncodeStrategy.Withf("error with %s minBalance", symbol)
			}
		}
		// create the strategy token in the database
		if _, err := qtx.CreateStrategyToken(internalCtx, queries.CreateStrategyTokenParams{
			StrategyID: importedStrategy.ID,
			TokenID:    common.HexToAddress(token.ID).Bytes(),
			MinBalance: minBalance.Bytes(),
			ChainID:    token.ChainID,
		}); err != nil {
			return ErrCantCreateStrategy.WithErr(err)
		}
	}
	// commit the transaction and return the strategyID
	if err := tx.Commit(); err != nil {
		return ErrCantCreateStrategy.WithErr(err)
	}
	response, err := json.Marshal(map[string]any{"strategyID": strategyID})
	if err != nil {
		return ErrEncodeStrategy.WithErr(err)
	}
	return ctx.Send(response, api.HTTPstatusOK)
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
	internalCtx, cancel := context.WithTimeout(ctx.Request.Context(), getStrategyTimeout)
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
		Alias:     strategyData.Alias,
		Predicate: strategyData.Predicate,
		URI:       strategyData.Uri,
		Tokens:    map[string]*StrategyToken{},
	}
	// get information of the strategy related tokens
	tokensData, err := capi.db.QueriesRO.TokensByStrategyID(internalCtx, strategyData.ID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return ErrCantGetTokens.WithErr(err)
	}
	// parse and encode tokens information
	for _, tokenData := range tokensData {
		strategy.Tokens[tokenData.Symbol] = &StrategyToken{
			ID:         common.BytesToAddress(tokenData.ID).String(),
			MinBalance: new(big.Int).SetBytes(tokenData.MinBalance).String(),
			ChainID:    tokenData.ChainID,
		}
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
	internalCtx, cancel := context.WithTimeout(ctx.Request.Context(), getTokensStrategyTimeout)
	defer cancel()
	// get the tokenID provided
	tokenID := ctx.URLParam("tokenID")
	// create db transaction
	tx, err := capi.db.RO.BeginTx(internalCtx, nil)
	if err != nil {
		return ErrCantGetStrategies.WithErr(err)
	}
	qtx := capi.db.QueriesRO.WithTx(tx)
	// get strategies associated to the token provided
	rows, err := qtx.StrategiesByTokenID(internalCtx, common.HexToAddress(tokenID).Bytes())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNoStrategies.WithErr(err)
		}
		return ErrCantGetStrategies.WithErr(err)
	}
	// parse and encode strategies
	strategies := GetStrategiesResponse{Strategies: []*GetStrategyResponse{}}
	for _, strategy := range rows {
		strategyResponse := &GetStrategyResponse{
			ID:        strategy.ID,
			Alias:     strategy.Alias,
			Predicate: strategy.Predicate,
			URI:       strategy.Uri,
			Tokens:    make(map[string]*StrategyToken),
		}
		strategyTokens, err := qtx.StrategyTokensByStrategyID(internalCtx, strategy.ID)
		if err != nil {
			return ErrCantGetStrategies.WithErr(err)
		}
		for _, strategyToken := range strategyTokens {
			if strategyToken.Symbol == "" {
				return ErrCantGetStrategies.With("invalid token symbol")
			}
			strategyResponse.Tokens[strategyToken.Symbol] = &StrategyToken{
				ID:         common.BytesToAddress(strategyToken.TokenID).String(),
				ChainID:    strategyToken.ChainID,
				MinBalance: new(big.Int).SetBytes(strategyToken.MinBalance).String(),
			}
		}
		strategies.Strategies = append(strategies.Strategies, strategyResponse)
	}
	res, err := json.Marshal(strategies)
	if err != nil {
		return ErrEncodeStrategies.WithErr(err)
	}
	return ctx.Send(res, api.HTTPstatusOK)
}

// validateStrategyPredicate function handler returns if the provided strategy
// predicate is valid and well-formatted. If the predicate is valid the handler
// returns a parsed version of the predicate as a JSON.
func (capi *census3API) validateStrategyPredicate(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	req := CreateStrategyRequest{}
	if err := json.Unmarshal(msg.Data, &req); err != nil {
		return ErrMalformedStrategy.WithErr(err)
	}
	if req.Predicate == "" {
		return ErrMalformedStrategy.With("no predicate provided")
	}

	lx := lexer.NewLexer(strategyoperators.ValidOperatorsTags)
	resultingToken, err := lx.Parse(req.Predicate)
	if err != nil {
		return ErrInvalidStrategyPredicate.WithErr(err)
	}
	res, err := json.Marshal(map[string]any{"result": resultingToken})
	if err != nil {
		return ErrEncodeValidPredicate.WithErr(err)
	}
	return ctx.Send(res, api.HTTPstatusOK)
}

// supportedStrategyPredicateOperators function handler returns the information
// of the current supported operators to build strategy predicates.
func (capi *census3API) supportedStrategyPredicateOperators(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	res, err := json.Marshal(map[string]any{
		"operators": strategyoperators.ValidOperators,
	})
	if err != nil {
		return ErrEncodeStrategyPredicateOperators.WithErr(err)
	}
	return ctx.Send(res, api.HTTPstatusOK)
}
