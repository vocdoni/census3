package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	gocid "github.com/ipfs/go-cid"
	queries "github.com/vocdoni/census3/db/sqlc"
	"github.com/vocdoni/census3/internal"
	"github.com/vocdoni/census3/internal/lexer"
	"github.com/vocdoni/census3/internal/roundedcensus"
	"github.com/vocdoni/census3/internal/strategyoperators"
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
	if err := capi.endpoint.RegisterMethod("/strategies/import/{cID}", "POST",
		api.MethodAccessTypePublic, capi.launchStrategyImport); err != nil {
		return err
	}
	if err := capi.endpoint.RegisterMethod("/strategies/import/queue/{queueID}", "GET",
		api.MethodAccessTypePublic, capi.enqueueImportStrategy); err != nil {
		return err
	}
	if err := capi.endpoint.RegisterMethod("/strategies/{strategyID}", "GET",
		api.MethodAccessTypePublic, capi.getStrategy); err != nil {
		return err
	}
	if err := capi.endpoint.RegisterMethod("/strategies/{strategyID}/holders", "GET",
		api.MethodAccessTypePublic, capi.listStrategyHolders); err != nil {
		return err
	}
	if err := capi.endpoint.RegisterMethod("/strategies/{strategyID}/estimation", "GET",
		api.MethodAccessTypePublic, capi.launchStrategyEstimation); err != nil {
		return err
	}
	if err := capi.endpoint.RegisterMethod("/strategies/{strategyID}/estimation/queue/{queueID}", "GET",
		api.MethodAccessTypePublic, capi.enqueueStrategyEstimation); err != nil {
		return err
	}
	if err := capi.endpoint.RegisterMethod("/strategies/token/{tokenID}", "GET",
		api.MethodAccessTypePublic, capi.getTokenStrategies); err != nil {
		return err
	}
	if err := capi.endpoint.RegisterMethod("/strategies/predicate/validate", "POST",
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
	// get pagination information from the request
	pageSize, dbPageSize, cursor, goForward, err := paginationFromCtx(ctx)
	if err != nil {
		return ErrMalformedPagination.WithErr(err)
	}
	iCursor := 0
	if cursor != "" {
		iCursor, err = strconv.Atoi(cursor)
		if err != nil {
			return ErrMalformedPagination.WithErr(err)
		}
	}
	internalCtx, cancel := context.WithTimeout(context.Background(), getStrategiesTimeout)
	defer cancel()
	// init db transaction
	tx, err := capi.db.RO.BeginTx(internalCtx, nil)
	if err != nil {
		return ErrCantGetStrategies.WithErr(err)
	}
	defer func() {
		if err := tx.Rollback(); err != nil && !errors.Is(sql.ErrTxDone, err) {
			log.Errorw(err, "create strategy transaction rollback failed")
		}
	}()
	qtx := capi.db.QueriesRW.WithTx(tx)
	// get the strategies from the database using the provided cursor, get the
	// following or previous page depending on the direction of the cursor
	var rows []queries.Strategy
	if goForward {
		rows, err = qtx.NextStrategiesPage(internalCtx, queries.NextStrategiesPageParams{
			PageCursor: uint64(iCursor),
			Limit:      dbPageSize,
		})
	} else {
		rows, err = qtx.PrevStrategiesPage(internalCtx, queries.PrevStrategiesPageParams{
			PageCursor: uint64(iCursor),
			Limit:      dbPageSize,
		})
	}
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
	strategiesResponse := GetStrategiesResponse{
		Strategies: []*GetStrategyResponse{},
		Pagination: &Pagination{PageSize: pageSize},
	}
	// get the next and previous cursors and add them to the response
	rows, nextCursorRow, prevCursorRow := paginationToRequest(rows, dbPageSize, cursor, goForward)
	if nextCursorRow != nil {
		strategiesResponse.Pagination.NextCursor = fmt.Sprint(nextCursorRow.ID)
	}
	if prevCursorRow != nil {
		strategiesResponse.Pagination.PrevCursor = fmt.Sprint(prevCursorRow.ID)
	}
	// parse and encode strategies
	for _, strategy := range rows {
		strategyResponse := &GetStrategyResponse{
			ID:        strategy.ID,
			Alias:     strategy.Alias,
			Predicate: strategy.Predicate,
			URI:       strategy.Uri,
			Tokens:    make(map[string]*StrategyToken),
		}
		strategyTokens, err := qtx.StrategyTokens(internalCtx, strategy.ID)
		if err != nil {
			return ErrCantGetStrategies.WithErr(err)
		}
		for _, strategyToken := range strategyTokens {
			if strategyToken.Symbol == "" {
				return ErrCantGetStrategies.With("invalid token symbol")
			}
			strategyResponse.Tokens[strategyToken.Symbol] = &StrategyToken{
				ID:           common.BytesToAddress(strategyToken.TokenID).String(),
				ChainID:      strategyToken.ChainID,
				MinBalance:   strategyToken.MinBalance,
				ChainAddress: strategyToken.ChainAddress,
				ExternalID:   strategyToken.ExternalID,
			}
		}
		strategiesResponse.Strategies = append(strategiesResponse.Strategies, strategyResponse)
	}
	res, err := json.Marshal(strategiesResponse)
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
		// if not use one
		minBalance := big.NewInt(1)
		if tokenData.MinBalance != "" {
			if _, ok := minBalance.SetString(tokenData.MinBalance, 10); !ok {
				return ErrEncodeStrategy.Withf("error with %s minBalance", symbol)
			}
		}
		// create the strategy_token in the database
		if _, err := qtx.CreateStrategyToken(internalCtx, queries.CreateStrategyTokenParams{
			StrategyID: uint64(strategyID),
			TokenID:    common.HexToAddress(tokenData.ID).Bytes(),
			MinBalance: minBalance.String(),
			ChainID:    tokenData.ChainID,
			ExternalID: tokenData.ExternalID,
		}); err != nil {
			return ErrCantCreateStrategy.WithErr(err)
		}
		// get the chain address of the token
		chainAddress, _ := capi.w3p.ChainAddress(tokenData.ChainID, tokenData.ID)
		req.Tokens[symbol].ChainAddress = chainAddress
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
	// update metrics
	internal.TotalNumberOfStrategies.Inc()
	internal.NewStrategiesByTime.Update(1)
	response, err := json.Marshal(map[string]any{"strategyID": strategyID})
	if err != nil {
		return ErrEncodeStrategy.WithErr(err)
	}
	return ctx.Send(response, api.HTTPstatusOK)
}

func (capi *census3API) launchStrategyImport(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	// get the cID from the url
	ipfsCID := ctx.URLParam("cID")
	if ipfsCID == "" {
		return ErrMalformedStrategy.With("no IPFS cID provided")
	}
	// check if the cID is valid
	if cid, err := gocid.Decode(ipfsCID); err != nil {
		return ErrMalformedStrategy.WithErr(err)
	} else {
		ipfsCID = cid.String()
	}
	// import the strategy from IPFS in background generating a queueID
	queueID := capi.queue.Enqueue()
	go func() {
		internalCtx, cancel := context.WithTimeout(context.Background(), getStrategyTimeout)
		defer cancel()
		// check if the strategy exists in the database using the IPFS URI
		ipfsURI := fmt.Sprintf("%s%s", capi.downloader.RemoteStorage.URIprefix(), ipfsCID)
		exists, err := capi.db.QueriesRO.ExistsStrategyByURI(internalCtx, ipfsURI)
		if err != nil {
			if ok := capi.queue.Update(queueID, true, nil, err); !ok {
				log.Errorf("error updating import strategy queue %s", queueID)
			}
			return
		}
		if exists {
			err := ErrCantImportStrategy.WithErr(fmt.Errorf("strategy already exists"))
			if ok := capi.queue.Update(queueID, true, nil, err); !ok {
				log.Errorf("error updating import strategy queue %s", queueID)
			}
			return
		}
		// download the strategy from IPFS and import it into the database if it
		// does not exist
		capi.downloader.AddToQueue(ipfsURI, func(_ string, dump []byte) {
			strategyID, err := capi.importStrategyDump(ipfsURI, dump)
			if err != nil {
				if ok := capi.queue.Update(queueID, true, nil, err); !ok {
					log.Errorf("error updating import strategy queue %s", queueID)
				}
				return
			}
			queueData := map[string]any{"strategyID": strategyID}
			if ok := capi.queue.Update(queueID, true, queueData, nil); !ok {
				log.Errorf("error updating import strategy queue %s", queueID)
			}
		}, true)
	}()
	// encode and send the queueID
	res, err := json.Marshal(QueueResponse{QueueID: queueID})
	if err != nil {
		return ErrEncodeStrategy.WithErr(err)
	}
	return ctx.Send(res, api.HTTPstatusOK)
}

func (capi *census3API) importStrategyDump(ipfsURI string, dump []byte) (uint64, error) {
	// init the internal context
	internalCtx, cancel := context.WithTimeout(context.Background(), importStrategyTimeout)
	defer cancel()
	// decode strategy
	importedStrategy := GetStrategyResponse{}
	if err := json.Unmarshal(dump, &importedStrategy); err != nil {
		return 0, ErrCantImportStrategy.WithErr(err)
	}
	// check if the strategy includes any token
	if len(importedStrategy.Tokens) == 0 {
		return 0, ErrCantImportStrategy.With("the imported strategy does not include any token")
	}
	log.Debugw("importing strategy", "strategy", importedStrategy)
	// init db transaction
	tx, err := capi.db.RW.BeginTx(internalCtx, nil)
	if err != nil {
		return 0, ErrCantCreateStrategy.WithErr(err)
	}
	defer func() {
		if err := tx.Rollback(); err != nil && !errors.Is(sql.ErrTxDone, err) {
			log.Errorw(err, "create strategy transaction rollback failed")
		}
	}()
	qtx := capi.db.QueriesRW.WithTx(tx)
	// ensure that all the tokens included in the strategy are registered in
	// the database
	for _, token := range importedStrategy.Tokens {
		exists, err := qtx.ExistsTokenByChainIDAndExternalID(internalCtx,
			queries.ExistsTokenByChainIDAndExternalIDParams{
				ID:         common.HexToAddress(token.ID).Bytes(),
				ChainID:    token.ChainID,
				ExternalID: token.ExternalID,
			})
		if err != nil {
			return 0, ErrCantGetToken.WithErr(err)
		}
		if !exists {
			return 0, ErrNotFoundToken.Withf("the imported strategy includes a no registered token: %s", token.ID)
		}
	}
	// check the strategy predicate
	lx := lexer.NewLexer(strategyoperators.ValidOperatorsTags)
	if _, err := lx.Parse(importedStrategy.Predicate); err != nil {
		return 0, ErrInvalidStrategyPredicate.With("the imported strategy includes a invalid predicate")
	}
	// create the strategy to get the ID and then create the strategy tokens
	result, err := qtx.CreateStategy(internalCtx, queries.CreateStategyParams{
		Alias:     importedStrategy.Alias,
		Predicate: importedStrategy.Predicate,
		Uri:       ipfsURI,
	})
	if err != nil {
		return 0, ErrCantCreateStrategy.WithErr(err)
	}
	strategyID, err := result.LastInsertId()
	if err != nil {
		return 0, ErrCantCreateStrategy.WithErr(err)
	}
	// iterate over the token included in the strategy and create them in the
	// database
	for symbol, token := range importedStrategy.Tokens {
		// decode the min balance for the current token if it is provided,
		// if not use zero
		minBalance := new(big.Int)
		if token.MinBalance != "" {
			if _, ok := minBalance.SetString(token.MinBalance, 10); !ok {
				return 0, ErrEncodeStrategy.Withf("error with %s minBalance", symbol)
			}
		}
		// create the strategy token in the database
		if _, err := qtx.CreateStrategyToken(internalCtx, queries.CreateStrategyTokenParams{
			StrategyID: uint64(strategyID),
			TokenID:    common.HexToAddress(token.ID).Bytes(),
			MinBalance: minBalance.String(),
			ChainID:    token.ChainID,
			ExternalID: token.ExternalID,
		}); err != nil {
			return 0, ErrCantCreateStrategy.WithErr(err)
		}
	}
	// commit the transaction and return the strategyID
	if err := tx.Commit(); err != nil {
		return 0, ErrCantCreateStrategy.WithErr(err)
	}
	internal.TotalNumberOfStrategies.Inc()
	internal.NewStrategiesByTime.Update(1)
	return uint64(strategyID), nil
}

func (capi *census3API) enqueueImportStrategy(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	// parse queueID from url
	queueID := ctx.URLParam("queueID")
	if queueID == "" {
		return ErrMalformedStrategyQueueID
	}
	// try to get and check if the strategy is in the queue
	exists, done, data, err := capi.queue.Done(queueID)
	if !exists {
		return ErrNotFoundStrategy.Withf("the ID %s does not exist in the queue", queueID)
	}
	// init the queue response
	queueStrategy := ImportStrategyQueueResponse{
		Done:  done,
		Error: err,
	}
	// check if it is not finished or some error occurred
	if done && err == nil {
		// if everything is ok, get the census information an return it
		internalCtx, cancel := context.WithTimeout(ctx.Request.Context(), enqueueStrategyImportTimeout)
		defer cancel()
		strategyID, ok := data["strategyID"].(uint64)
		if !ok {
			log.Errorf("no strategy id registered on queue item")
			return ErrCantGetStrategy
		}
		// get strategy from the database
		strategyData, err := capi.db.QueriesRO.StrategyByID(internalCtx, strategyID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return ErrNotFoundStrategy.WithErr(err)
			}
			return ErrCantGetStrategy.WithErr(err)
		}
		// encode census
		queueStrategy.Strategy = &GetStrategyResponse{
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
			queueStrategy.Strategy.Tokens[tokenData.Symbol] = &StrategyToken{
				ID:           common.BytesToAddress(tokenData.ID).String(),
				ChainAddress: tokenData.ChainAddress,
				MinBalance:   tokenData.MinBalance,
				ChainID:      tokenData.ChainID,
			}
		}
		// remove the item from the queue
		capi.queue.Dequeue(queueID)
	}
	// encode item response and send it
	res, err := json.Marshal(queueStrategy)
	if err != nil {
		return ErrEncodeQueueItem.WithErr(err)
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
	tokensData, err := capi.db.QueriesRO.StrategyTokens(internalCtx, strategyData.ID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return ErrCantGetTokens.WithErr(err)
	}
	// parse and encode tokens information
	for _, tokenData := range tokensData {
		strategy.Tokens[tokenData.Symbol] = &StrategyToken{
			ID:           common.BytesToAddress(tokenData.TokenID).String(),
			ChainAddress: tokenData.ChainAddress,
			MinBalance:   tokenData.MinBalance,
			ChainID:      tokenData.ChainID,
			ExternalID:   tokenData.ExternalID,
		}
	}
	res, err := json.Marshal(strategy)
	if err != nil {
		return ErrEncodeStrategy.WithErr(err)
	}
	return ctx.Send(res, api.HTTPstatusOK)
}

// listStrategyHolders function handler returns the list of the holders of the
// strategy ID provided. It returns a 400 error if the provided
func (capi *census3API) listStrategyHolders(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	// get provided strategyID
	iStrategyID, err := strconv.Atoi(ctx.URLParam("strategyID"))
	if err != nil {
		return ErrMalformedStrategyID.WithErr(err)
	}
	strategyID := uint64(iStrategyID)
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
	// get token information from the database
	internalCtx, cancel := context.WithTimeout(ctx.Request.Context(), getStrategyHoldersTimeout)
	defer cancel()
	tx, err := capi.db.RO.BeginTx(internalCtx, nil)
	if err != nil {
		return ErrCantGetStrategyHolders.WithErr(err)
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Errorw(err, "error rolling back tokens transaction")
		}
	}()
	qtx := capi.db.QueriesRO.WithTx(tx)
	strategy, err := qtx.StrategyByID(internalCtx, strategyID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNotFoundStrategy.WithErr(err)
		}
		return ErrCantGetStrategyHolders.WithErr(err)
	}
	// check predicate
	lx := lexer.NewLexer(strategyoperators.ValidOperatorsTags)
	validatedPredicate, err := lx.Parse(strategy.Predicate)
	if err != nil {
		return ErrInvalidStrategyPredicate.WithErr(err)
	}
	if !validatedPredicate.IsLiteral() {
		return ErrInvalidStrategyPredicate.With("this endpoint only supports single token predicates")
	}
	// get the tokens from the database using the provided cursor, get the
	// following or previous page depending on the direction of the cursor
	var rows []queries.TokenHolder
	if goForward {
		rows, err = qtx.NextStrategyTokenHoldersPage(internalCtx, queries.NextStrategyTokenHoldersPageParams{
			PageCursor: bCursor,
			Limit:      dbPageSize,
			StrategyID: strategyID,
		})
	} else {
		rows, err = qtx.NextStrategyTokenHoldersPage(internalCtx, queries.NextStrategyTokenHoldersPageParams{
			PageCursor: bCursor,
			Limit:      dbPageSize,
			StrategyID: strategyID,
		})
	}
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrCantGetStrategyHolders.WithErr(err)
		}
		return ErrCantGetStrategyHolders.WithErr(err)
	}
	if len(rows) == 0 {
		return ErrNotFoundTokenHolders
	}
	// init response struct with the initial pagination information and empty
	// list of holders
	holdersResponse := GetStrategyHoldersResponse{
		Holders:    make(map[string]string),
		Pagination: &Pagination{PageSize: pageSize},
	}
	// get the next and previous cursors and add them to the response
	rows, nextCursorRow, prevCursorRow := paginationToRequest(rows, dbPageSize, cursor, goForward)
	if nextCursorRow != nil {
		holdersResponse.Pagination.NextCursor = common.BytesToAddress(nextCursorRow.HolderID).String()
	}
	if prevCursorRow != nil {
		holdersResponse.Pagination.PrevCursor = common.BytesToAddress(prevCursorRow.HolderID).String()
	}
	// parse and encode holders
	for _, holder := range rows {
		holdersResponse.Holders[common.BytesToAddress(holder.HolderID).String()] = holder.Balance
	}
	// encode and send the response
	res, err := json.Marshal(holdersResponse)
	if err != nil {
		return ErrEncodeTokenHolders.WithErr(err)
	}
	return ctx.Send(res, api.HTTPstatusOK)
}

// launchStrategyEstimation function handler enqueues the estimation of the
// size of the strategy identified by the ID provided. It returns an error if
// the provided ID is wrong or if the queue item cannot be encoded.
func (capi *census3API) launchStrategyEstimation(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	// get provided strategyID
	iStrategyID, err := strconv.Atoi(ctx.URLParam("strategyID"))
	if err != nil {
		return ErrMalformedStrategyID.WithErr(err)
	}
	strategyID := uint64(iStrategyID)
	// get anonymous from query params and decode it as boolean
	anonymous := ctx.Request.URL.Query().Get("anonymous") == "true"
	// import the strategy from IPFS in background generating a queueID
	queueID := capi.queue.Enqueue()
	go func() {
		size, accuracy, err := capi.estimateStrategySizeAndAccuracy(strategyID, anonymous)
		if err != nil {
			if ok := capi.queue.Update(queueID, true, nil, err); !ok {
				log.Errorf("error updating import strategy queue %s", queueID)
			}
			return
		}
		queueData := map[string]any{
			"size":               uint64(size),
			"timeToCreateCensus": TimeToCreateCensus(uint64(size)),
			"accuracy":           accuracy,
		}
		if ok := capi.queue.Update(queueID, true, queueData, nil); !ok {
			log.Errorf("error updating import strategy queue %s", queueID)
		}
	}()
	// encode and send the queueID
	res, err := json.Marshal(QueueResponse{QueueID: queueID})
	if err != nil {
		return ErrEncodeStrategy.WithErr(err)
	}
	return ctx.Send(res, api.HTTPstatusOK)
}

// estimateStrategySizeAndAccuracy function returns the estimated size of the
// strategy identified by the ID provided. The size is calculated using the strategy
// predicate. This method should be executed in background.
func (capi *census3API) estimateStrategySizeAndAccuracy(strategyID uint64, anonymous bool) (int, float64, error) {
	internalCtx, cancel := context.WithTimeout(context.Background(), createAndPublishCensusTimeout)
	defer cancel()
	// check if the strategy size is already cached
	cacheKey := EncCacheKey(fmt.Sprintf("strategy%d-%t", strategyID, anonymous))
	if rawData, exists := capi.cache.Get(cacheKey); exists {
		data, ok := rawData.(map[string]any)
		if !ok {
			return 0, 0, ErrCantGetStrategy.With("invalid cache data")
		}
		size, ok := data["size"].(int)
		if !ok {
			return 0, 0, ErrCantGetStrategy.With("invalid cache data")
		}
		accuracy, ok := data["accuracy"].(float64)
		if !ok {
			return 0, 0, ErrCantGetStrategy.With("invalid cache data")
		}
		return size, accuracy, nil
	}
	// begin a transaction for group sql queries
	tx, err := capi.db.RW.BeginTx(internalCtx, nil)
	if err != nil {
		return 0, 0, ErrCantGetStrategy.WithErr(err)
	}
	defer func() {
		if err := tx.Rollback(); err != nil && !errors.Is(sql.ErrTxDone, err) {
			log.Errorw(err, "holders transaction rollback failed")
		}
	}()
	strategy, err := capi.db.QueriesRO.StrategyByID(internalCtx, strategyID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, 0, ErrNotFoundStrategy.WithErr(err)
		}
		return 0, 0, ErrCantGetStrategy.WithErr(err)
	}
	if strategy.Predicate == "" {
		return 0, 0, ErrInvalidStrategyPredicate.With("empty predicate")
	}
	// calculate the strategy holders
	strategyHolders, _, _, err := CalculateStrategyHolders(internalCtx,
		capi.db.QueriesRO, capi.holderProviders, strategyID, strategy.Predicate)
	if err != nil {
		return 0, 0, ErrEvalStrategyPredicate.WithErr(err)
	}
	if size := len(strategyHolders); size != 0 {
		accuracy := 100.0
		// if the census is anonymous, group and round the census to get the
		// final accuracy, to do that we need to convert the strategy holders
		// map to a slice of census participants
		if anonymous {
			censusParticipants := []*roundedcensus.Participant{}
			for address, balance := range strategyHolders {
				censusParticipants = append(censusParticipants, &roundedcensus.Participant{
					Address: address.Hex(),
					Balance: balance,
				})
			}
			// calculate the accuracy of the census and return it with the size
			_, accuracy, _ = roundedcensus.GroupAndRoundCensus(censusParticipants, roundedcensus.DefaultGroupsConfig)
		}
		// cache the strategy if the estimated size is greater than the
		// threshold
		if size > strategyHoldersCacheThreshold {
			capi.cache.Add(cacheKey, map[string]any{
				"size":     size,
				"accuracy": accuracy,
			})
		}
		// return the size and resulting accuracy
		return size, accuracy, nil
	}
	return 0, 0, ErrNoStrategyHolders
}

// enqueueStrategyEstimation function handler dequeues the estimation of the
// size of the strategy identified by the queue ID provided. It returns an error
// if something fails with the queue item. The queue item, when the background
// size estimation ends, will contain the size of the strategy or the error
// occurred during the estimation.
func (capi *census3API) enqueueStrategyEstimation(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	// parse queueID from url
	queueID := ctx.URLParam("queueID")
	if queueID == "" {
		return ErrMalformedStrategyQueueID
	}
	// try to get and check if the strategy is in the queue
	exists, done, data, err := capi.queue.Done(queueID)
	if !exists {
		return ErrNotFoundStrategy.Withf("the ID %s does not exist in the queue", queueID)
	}
	// init the queue response
	queueStrategy := GetStrategyEstimationResponse{
		Done:  done,
		Error: err,
	}
	// check if it is not finished or some error occurred
	if done && err == nil {
		queueStrategy.Estimation = &StrategyEstimation{
			Size:               data["size"].(uint64),
			TimeToCreateCensus: data["timeToCreateCensus"].(uint64),
			Accuracy:           data["accuracy"].(float64),
		}
		// remove the item from the queue
		capi.queue.Dequeue(queueID)
	}
	// encode item response and send it
	res, err := json.Marshal(queueStrategy)
	if err != nil {
		return ErrEncodeQueueItem.WithErr(err)
	}
	return ctx.Send(res, api.HTTPstatusOK)
}

// getTokenStrategies function handler returns the strategies that involves the
// token identified by the ID (token address) provided. It returns a 400 error
// if the provided ID is wrong or empty, a 204 response if the token has not any
// associated strategy or a 500 error if something fails.
func (capi *census3API) getTokenStrategies(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
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
	internalCtx, cancel := context.WithTimeout(ctx.Request.Context(), getTokensStrategyTimeout)
	defer cancel()
	// create db transaction
	tx, err := capi.db.RO.BeginTx(internalCtx, nil)
	if err != nil {
		return ErrCantGetStrategies.WithErr(err)
	}
	qtx := capi.db.QueriesRO.WithTx(tx)
	// get strategies associated to the token provided
	rows, err := qtx.StrategiesByTokenIDAndChainIDAndExternalID(internalCtx,
		queries.StrategiesByTokenIDAndChainIDAndExternalIDParams{
			TokenID:    address.Bytes(),
			ChainID:    uint64(chainID),
			ExternalID: externalID,
		})
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
				ID:           common.BytesToAddress(strategyToken.ID).String(),
				ChainAddress: strategyToken.ChainAddress,
				ChainID:      strategyToken.ChainID,
				MinBalance:   strategyToken.MinBalance,
				ExternalID:   strategyToken.ExternalID,
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
