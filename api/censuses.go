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
	queries "github.com/vocdoni/census3/db/sqlc"
	"github.com/vocdoni/census3/internal"
	"github.com/vocdoni/census3/roundedcensus"
	"go.vocdoni.io/dvote/httprouter"
	api "go.vocdoni.io/dvote/httprouter/apirest"
	"go.vocdoni.io/dvote/log"
	"go.vocdoni.io/dvote/types"
)

func (capi *census3API) initCensusHandlers() error {
	if err := capi.endpoint.RegisterMethod("/censuses/{censusID}", "GET",
		api.MethodAccessTypePublic, capi.getCensus); err != nil {
		return err
	}
	if err := capi.endpoint.RegisterMethod("/censuses", "POST",
		api.MethodAccessTypePublic, capi.launchCensusCreation); err != nil {
		return err
	}
	if err := capi.endpoint.RegisterMethod("/censuses/queue/{queueID}", "GET",
		api.MethodAccessTypePublic, capi.enqueueCensus); err != nil {
		return err
	}
	return capi.endpoint.RegisterMethod("/censuses/strategy/{strategyID}", "GET",
		api.MethodAccessTypePublic, capi.getStrategyCensuses)
}

// getCensus handler responses with the information regarding of the census
// requested by its ID.
func (capi *census3API) getCensus(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	iCensusID, err := strconv.Atoi(ctx.URLParam("censusID"))
	if err != nil {
		return ErrMalformedCensusID
	}
	censusID := uint64(iCensusID)
	internalCtx, cancel := context.WithTimeout(ctx.Request.Context(), getCensusTimeout)
	defer cancel()
	currentCensus, err := capi.db.QueriesRO.CensusByID(internalCtx, censusID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNotFoundCensus.WithErr(err)
		}
		return ErrCantGetCensus.WithErr(err)
	}
	censusWeight := []byte{}
	if currentCensus.Weight.Valid {
		censusWeight = []byte(currentCensus.Weight.String)
	}
	res, err := json.Marshal(GetCensusResponse{
		CensusID:   censusID,
		StrategyID: currentCensus.StrategyID,
		MerkleRoot: types.HexBytes(currentCensus.MerkleRoot),
		URI:        capi.downloader.RemoteStorage.URIprefix() + currentCensus.Uri.String,
		Size:       currentCensus.Size,
		Weight:     new(big.Int).SetBytes(censusWeight).String(),
		Anonymous:  currentCensus.CensusType == uint64(anonymousCensusType),
		Accuracy:   currentCensus.Accuracy,
	})
	if err != nil {
		return ErrEncodeCensus.WithErr(err)
	}
	return ctx.Send(res, api.HTTPstatusOK)
}

// launchCensusCreation handler parses the census creation request, enqueues it
// and starts the creation process, then returns the queue identifier of that
// process to support tracking it. When the process ends updates the queue item
// with the resulting status or error into the queue.
func (capi *census3API) launchCensusCreation(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	// decode request
	req := &CreateCensusRequest{}
	if err := json.Unmarshal(msg.Data, req); err != nil {
		return ErrMalformedStrategyID.WithErr(err)
	}
	// create and publish census merkle tree in background
	queueID := capi.queue.Enqueue()
	go func() {
		censusID, err := capi.createAndPublishCensus(req, queueID)
		if err != nil && !errors.Is(ErrCensusAlreadyExists, err) {
			if ok := capi.queue.Update(queueID, true, nil, err); !ok {
				log.Errorf("error updating census queue process with error: %v", err)
			}
			return
		}
		queueData := map[string]any{"censusID": censusID}
		if ok := capi.queue.Update(queueID, true, queueData, nil); !ok {
			log.Errorf("error updating census queue process with error")
		}
	}()
	// encoding the result and response it
	res, err := json.Marshal(QueueResponse{
		QueueID: queueID,
	})
	if err != nil {
		return ErrEncodeCensus.WithErr(err)
	}
	return ctx.Send(res, api.HTTPstatusOK)
}

// createAndPublishCensus method creates a census tree based on the token
// holders of the tokens that are included in the given strategy. It recovers
// all the required information from the database, and then creates and publish
// the census merkle tree on IPFS. Then saves the resulting information of the
// census tree in the database.
func (capi *census3API) createAndPublishCensus(req *CreateCensusRequest, qID string) (uint64, error) {
	internalCtx, cancel := context.WithTimeout(context.Background(), createAndPublishCensusTimeout)
	defer cancel()
	// begin a transaction for group sql queries
	tx, err := capi.db.RW.BeginTx(internalCtx, nil)
	if err != nil {
		return 0, ErrCantCreateCensus.WithErr(err)
	}
	defer func() {
		if err := tx.Rollback(); err != nil && !errors.Is(sql.ErrTxDone, err) {
			log.Errorw(err, "holders transaction rollback failed")
		}
	}()
	qtx := capi.db.QueriesRW.WithTx(tx)

	strategy, err := qtx.StrategyByID(internalCtx, req.StrategyID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrNotFoundStrategy.WithErr(err)
		}
		return 0, ErrCantCreateCensus.WithErr(err)
	}
	if strategy.Predicate == "" {
		return 0, ErrInvalidStrategyPredicate.With("empty predicate")
	}
	// init some variables to get computed in the following steps
	strategyHolders, censusWeight, totalTokensBlockNumber, err := CalculateStrategyHolders(
		internalCtx, capi.db.QueriesRO, capi.w3p, req.StrategyID, strategy.Predicate)
	if err != nil {
		return 0, ErrEvalStrategyPredicate.WithErr(err)
	}
	if len(strategyHolders) == 0 {
		return 0, ErrNoStrategyHolders
	}
	// compute the new censusId and censusType
	newCensusID := InnerCensusID(totalTokensBlockNumber, req.StrategyID, req.Anonymous)
	// check if the census already exists
	_, err = qtx.CensusByID(internalCtx, newCensusID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return 0, ErrCantCreateCensus.WithErr(err)
		}
	} else {
		// return the censusID to get the census information if it already
		// exists
		return newCensusID, ErrCensusAlreadyExists
	}
	// check the censusType
	censusType := defaultCensusType
	finalAccuracy := 100.0
	if req.Anonymous {
		censusType = anonymousCensusType
		// Round census balances using the default configuration
		censusParticipants := []*roundedcensus.Participant{}
		for address, balance := range strategyHolders {
			censusParticipants = append(censusParticipants, &roundedcensus.Participant{
				Address: address.String(),
				Balance: balance,
			})
		}
		res, accuracy, err := roundedcensus.GroupAndRoundCensus(censusParticipants, roundedcensus.DefaultGroupsConfig)
		if err != nil {
			log.Errorf("min accuracy couldn't be reached for strategy %d: %.2f: %v", req.StrategyID, accuracy, err)
		}
		// update the census holders with the rounded balances and the final accuracy
		strategyHolders = map[common.Address]*big.Int{}
		for _, participant := range res {
			strategyHolders[common.HexToAddress(participant.Address)] = participant.Balance
		}
		finalAccuracy = accuracy
	}
	// create a census tree and publish on IPFS
	root, uri, _, err := CreateAndPublishCensus(capi.censusDB, capi.storage, CensusOptions{
		ID:      newCensusID,
		Type:    censusType,
		Holders: strategyHolders,
	})
	if err != nil {
		return 0, ErrCantCreateCensus.WithErr(err)
	}
	// save the new census in the SQL database
	sqlURI := &sql.NullString{}
	if err := sqlURI.Scan(uri); err != nil {
		return 0, ErrCantCreateCensus.WithErr(err)
	}
	sqlCensusSize := &sql.NullInt64{}
	if err := sqlCensusSize.Scan(int64(len(strategyHolders))); err != nil {
		return 0, ErrCantCreateCensus.WithErr(err)
	}
	sqlCensusWeight := &sql.NullString{}
	if err := sqlCensusWeight.Scan(censusWeight.String()); err != nil {
		return 0, ErrCantCreateCensus.WithErr(err)
	}
	_, err = qtx.CreateCensus(internalCtx, queries.CreateCensusParams{
		ID:         newCensusID,
		StrategyID: req.StrategyID,
		CensusType: uint64(censusType),
		MerkleRoot: []byte(root),
		Uri:        *sqlURI,
		Size:       uint64(sqlCensusSize.Int64),
		Weight:     *sqlCensusWeight,
		QueueID:    qID,
		Accuracy:   finalAccuracy,
	})
	if err != nil {
		return 0, ErrCantCreateCensus.WithErr(err)
	}
	if err := tx.Commit(); err != nil {
		return 0, ErrCantCreateCensus.WithErr(err)
	}
	// update metrics
	internal.TotalNumberOfCensuses.Inc()
	internal.NumberOfCensusesByType.GetOrCreateCounter(fmt.Sprintf("%s%d",
		internal.NumberOfCensusesByTypePrefix, censusType,
	)).Inc()
	return newCensusID, nil
}

// enqueueCensus handler returns the current status of the queue item
// identified by the ID provided. If it not exists it returns that the census
// is not found. Else if the census exists and has been successfully created, it
// will be included into the response. If not, the response only will include
// if it is done or not and the resulting error.
func (capi *census3API) enqueueCensus(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	queueID := ctx.URLParam("queueID")
	if queueID == "" {
		return ErrMalformedCensusQueueID
	}
	// try to get and check if the census is in the queue
	exists, done, data, err := capi.queue.Done(queueID)
	if !exists {
		return ErrNotFoundCensus.Withf("the ID %s does not exist in the queue", queueID)
	}
	// init queue item response
	queueCensus := CensusQueueResponse{
		Done:  done,
		Error: err,
	}
	// check if it is not finished or some error occurred
	if done && err == nil {
		// if everything is ok, get the census information an return it
		internalCtx, cancel := context.WithTimeout(ctx.Request.Context(), enqueueCensusCreationTimeout)
		defer cancel()
		censusID, ok := data["censusID"].(uint64)
		if !ok {
			log.Errorf("no census id registered on queue item")
			return ErrCantGetCensus
		}
		// get the census from the database by queue_id
		currentCensus, err := capi.db.QueriesRO.CensusByID(internalCtx, censusID)
		if err != nil {
			return ErrCantGetCensus.WithErr(err)
		}
		// get values for optional parameters
		if !currentCensus.Weight.Valid {
			return ErrCantGetCensus.With("invalid census weight")
		}
		censusWeight, ok := new(big.Int).SetString(currentCensus.Weight.String, 10)
		if !ok {
			return ErrCantGetCensus.With("invalid census weight")
		}
		// encode census
		queueCensus.Census = &GetCensusResponse{
			CensusID:   currentCensus.ID,
			StrategyID: currentCensus.StrategyID,
			MerkleRoot: types.HexBytes(currentCensus.MerkleRoot),
			URI:        capi.downloader.RemoteStorage.URIprefix() + currentCensus.Uri.String,
			Size:       currentCensus.Size,
			Weight:     censusWeight.String(),
			Anonymous:  currentCensus.CensusType == uint64(anonymousCensusType),
			Accuracy:   currentCensus.Accuracy,
		}
		// remove the item from the queue
		capi.queue.Dequeue(queueID)
	}
	// encode item response and send it
	res, err := json.Marshal(queueCensus)
	if err != nil {
		return ErrEncodeQueueItem.WithErr(err)
	}
	return ctx.Send(res, api.HTTPstatusOK)
}

// getStrategyCensuses function handler returns the censuses that had been
// generated with the strategy identified by the ID provided.
func (capi *census3API) getStrategyCensuses(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	// get strategy ID
	strategyID, err := strconv.Atoi(ctx.URLParam("strategyID"))
	if err != nil {
		return ErrMalformedCensusID.WithErr(err)
	}
	// get censuses by this strategy ID
	internalCtx, cancel := context.WithTimeout(ctx.Request.Context(), getStrategyCensusesTimeout)
	defer cancel()
	rows, err := capi.db.QueriesRO.CensusByStrategyID(internalCtx, uint64(strategyID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNotFoundCensus.WithErr(err)
		}
		return ErrCantGetCensus.WithErr(err)
	}
	// parse and encode response
	censuses := GetCensusesResponse{Censuses: []*GetCensusResponse{}}
	for _, censusInfo := range rows {
		// get values for optional parameters
		if !censusInfo.Weight.Valid {
			return ErrCantGetCensus.With("invalid census weight")
		}
		censusWeight, ok := new(big.Int).SetString(censusInfo.Weight.String, 10)
		if !ok {
			return ErrCantGetCensus.With("invalid census weight")
		}

		censuses.Censuses = append(censuses.Censuses, &GetCensusResponse{
			CensusID:   censusInfo.ID,
			StrategyID: censusInfo.StrategyID,
			MerkleRoot: types.HexBytes(censusInfo.MerkleRoot),
			URI:        capi.downloader.RemoteStorage.URIprefix() + censusInfo.Uri.String,
			Size:       censusInfo.Size,
			Weight:     censusWeight.String(),
			Anonymous:  censusInfo.CensusType == uint64(anonymousCensusType),
			Accuracy:   censusInfo.Accuracy,
		})
	}
	res, err := json.Marshal(censuses)
	if err != nil {
		return ErrEncodeCensuses.WithErr(err)
	}
	return ctx.Send(res, api.HTTPstatusOK)
}
