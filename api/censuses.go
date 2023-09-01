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
	"github.com/vocdoni/census3/census"
	queries "github.com/vocdoni/census3/db/sqlc"
	"go.vocdoni.io/dvote/httprouter"
	api "go.vocdoni.io/dvote/httprouter/apirest"
	"go.vocdoni.io/dvote/log"
)

func (capi *census3API) initCensusHandlers() error {
	if err := capi.endpoint.RegisterMethod("/census/{censusID}", "GET",
		api.MethodAccessTypePublic, capi.getCensus); err != nil {
		return err
	}
	if err := capi.endpoint.RegisterMethod("/census", "POST",
		api.MethodAccessTypePublic, capi.launchCensusCreation); err != nil {
		return err
	}
	if err := capi.endpoint.RegisterMethod("/census/queue/{queueID}", "GET",
		api.MethodAccessTypePublic, capi.getEnqueueCensus); err != nil {
		return err
	}
	return capi.endpoint.RegisterMethod("/census/strategy/{strategyID}", "GET",
		api.MethodAccessTypePublic, capi.getStrategyCensuses)
}

// getCensus handler responses with the information regarding of the census
// requested by its ID.
func (capi *census3API) getCensus(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	censusID, err := strconv.Atoi(ctx.URLParam("censusID"))
	if err != nil {
		return ErrMalformedCensusID
	}
	internalCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// begin a transaction for group sql queries
	tx, err := capi.db.BeginTx(internalCtx, nil)
	if err != nil {
		log.Errorw(err, "error starting database")
		return ErrCantGetCensus
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Errorw(err, "holders transaction rollback failed")
		}
	}()
	qtx := capi.sqlc.WithTx(tx)
	currentCensus, err := qtx.CensusByID(internalCtx, int64(censusID))
	if err != nil {
		log.Errorw(err, "error getting census from database")
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNotFoundCensus
		}
		return ErrCantGetCensus
	}
	chainID, err := qtx.ChainID(internalCtx)
	if err != nil {
		log.Errorw(err, "error getting chainID")
		return ErrCantGetCensus
	}
	censusSize := int32(0)
	if currentCensus.Size.Valid {
		censusSize = currentCensus.Size.Int32
	}
	censusWeight := []byte{}
	if currentCensus.Weight.Valid {
		censusWeight = []byte(currentCensus.Weight.String)
	}
	res, err := json.Marshal(GetCensusResponse{
		CensusID:   uint64(censusID),
		StrategyID: uint64(currentCensus.StrategyID),
		MerkleRoot: common.Bytes2Hex(currentCensus.MerkleRoot),
		URI:        "ipfs://" + currentCensus.Uri.String,
		Size:       int32(censusSize),
		Weight:     new(big.Int).SetBytes(censusWeight).String(),
		ChainID:    uint64(chainID),
		Anonymous:  currentCensus.CensusType == int64(census.AnonymousCensusType),
	})
	if err != nil {
		log.Errorw(err, "error encoding census")
		return ErrEncodeCensus
	}
	return ctx.Send(res, api.HTTPstatusOK)
}

// launchCensusCreation handler parses the census creation request, enqueues it
// and starts the creation process, then returns the queue identifier of that
// process to support tracking it. When the process ends updates the queue item
// with the resulting status or error into the queue.
func (capi *census3API) launchCensusCreation(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	// decode request
	req := &CreateCensusResquest{}
	if err := json.Unmarshal(msg.Data, req); err != nil {
		return ErrMalformedStrategyID
	}
	// create and publish census merkle tree in background
	queueID := capi.queue.Enqueue()
	go func(req *CreateCensusResquest) {
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
	}(req)
	// encoding the result and response it
	res, err := json.Marshal(CreateCensusResponse{
		QueueID: queueID,
	})
	if err != nil {
		log.Error("error marshalling census")
		return ErrEncodeCensus
	}
	return ctx.Send(res, api.HTTPstatusOK)
}

// createAndPublishCensus method creates a census tree based on the token
// holders of the tokens that are included in the given strategy. It recovers
// all the required information from the database, and then creates and publish
// the census merkle tree on IPFS. Then saves the resulting information of the
// census tree in the database.
func (capi *census3API) createAndPublishCensus(req *CreateCensusResquest, qID string) (int, error) {
	bgCtx, cancel := context.WithTimeout(context.Background(), censusCreationTimeout)
	defer cancel()
	// begin a transaction for group sql queries
	tx, err := capi.db.BeginTx(bgCtx, nil)
	if err != nil {
		log.Errorw(err, "error starting database")
		return -1, ErrCantCreateCensus.With("error starting database")
	}
	defer func() {
		if err := tx.Rollback(); err != nil && !errors.Is(sql.ErrTxDone, err) {
			log.Errorw(err, "holders transaction rollback failed")
		}
	}()
	qtx := capi.sqlc.WithTx(tx)
	// get the tokens of the strategy provided and check them
	strategyTokens, err := qtx.TokensByStrategyID(bgCtx, int64(req.StrategyID))
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			log.Errorf("no strategy found for id %d: %s", req.StrategyID, err.Error())
			return -1, ErrNoStrategyTokens.With("no strategy found")
		}
		log.Errorf("error getting strategy with id %d: %s", req.StrategyID, err.Error())
		return -1, ErrCantCreateCensus.With("error getting strategy")
	}
	if len(strategyTokens) == 0 {
		log.Errorf("no tokens for strategy %d", req.StrategyID)
		return -1, ErrNoStrategyTokens.With("no tokens for strategy")
	}

	// get the maximun current census ID to calculate the next one, if any
	// census has been created yet, continue
	lastCensusID, err := qtx.LastCensusID(bgCtx)
	if err != nil && !errors.Is(sql.ErrNoRows, err) {
		log.Errorw(err, "error getting last census ID")
		return -1, ErrCantCreateCensus.With("error getting last census ID")
	}
	// compute the new censusId and censusType
	newCensusID := int(lastCensusID) + 1
	censusType := census.DefaultCensusType
	if req.Anonymous {
		censusType = census.AnonymousCensusType
	}
	// get holders associated to every strategy token, create a map to avoid
	// duplicates and count the sum of the balances to get the weight of the
	// census
	censusWeight := new(big.Int)
	strategyHolders := map[common.Address]*big.Int{}
	for _, token := range strategyTokens {
		holders, err := qtx.TokenHoldersByTokenID(bgCtx, token.ID)
		if err != nil {
			if errors.Is(sql.ErrNoRows, err) {
				continue
			}
			log.Errorf("error getting token holders of %s: %v", common.BytesToAddress(token.ID), err)
			return -1, ErrCantCreateCensus.With("error getting token holders")
		}
		for _, holder := range holders {
			holderAddr := common.BytesToAddress(holder.ID)
			holderBalance := new(big.Int).SetBytes(holder.Balance)
			if _, exists := strategyHolders[holderAddr]; !exists {
				strategyHolders[holderAddr] = holderBalance
				censusWeight = new(big.Int).Add(censusWeight, holderBalance)
			}
		}
	}
	if len(strategyHolders) == 0 {
		log.Errorf("no holders for strategy '%d'", req.StrategyID)
		return -1, ErrNotFoundTokenHolders.With("no holders for strategy")
	}
	// create a census tree and publish on IPFS
	def := census.NewCensusDefinition(newCensusID, int(req.StrategyID), strategyHolders, req.Anonymous)
	newCensus, err := capi.censusDB.CreateAndPublish(def)
	if err != nil {
		log.Errorf("error creating or publishing the census: %v", err)
		return -1, ErrCantCreateCensus.With("error creating or publishing the census")
	}
	// check if the census already exists using the merkle root of the generated
	// census
	currentCensus, err := qtx.CensusByMerkleRoot(bgCtx, newCensus.RootHash)
	if err == nil {
		return int(currentCensus.ID), ErrCensusAlreadyExists
	}
	if err != nil && !errors.Is(sql.ErrNoRows, err) {
		log.Errorf("error checking if the generated census already exists: %v", err)
		return -1, ErrCantCreateCensus.With("error checking if the generated census already exists")
	}
	// save the new census in the SQL database
	sqlURI := new(sql.NullString)
	if err := sqlURI.Scan(newCensus.URI); err != nil {
		log.Errorf("error encoding census uri: %v", err)
		return -1, ErrCantCreateCensus.With("error encoding census uri")
	}
	sqlCensusSize := sql.NullInt32{}
	if err := sqlCensusSize.Scan(int64(len(strategyHolders))); err != nil {
		log.Errorf("error encoding census size: %v", err)
		return -1, ErrCantCreateCensus.With("error encoding census size")
	}
	sqlCensusWeight := sql.NullString{}
	if err := sqlCensusWeight.Scan(censusWeight.String()); err != nil {
		log.Errorf("error encoding census weight: %v", err)
		return -1, ErrCantCreateCensus.With("error encoding census weight")
	}
	_, err = qtx.CreateCensus(bgCtx, queries.CreateCensusParams{
		ID:         int64(newCensus.ID),
		StrategyID: int64(req.StrategyID),
		CensusType: int64(censusType),
		MerkleRoot: newCensus.RootHash,
		Uri:        *sqlURI,
		Size:       sqlCensusSize,
		Weight:     sqlCensusWeight,
		QueueID:    qID,
	})
	if err != nil {
		log.Errorf("error saving the census on the database: %v", err)
		return -1, ErrCantCreateCensus.With("error saving the census on the database")
	}
	if err := tx.Commit(); err != nil {
		log.Errorf("error committing the census on the database: %v", err)
		return -1, ErrCantCreateCensus.With("error committing the census on the database")
	}
	return newCensus.ID, nil
}

// getEnqueueCensus handler returns the current status of the queue item
// identified by the ID provided. If it not exists it returns that the census
// is not found. Else if the census exists and has been successfully created, it
// will be included into the response. If not, the response only will include
// if it is done or not and the resulting error.
func (capi *census3API) getEnqueueCensus(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
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
	queueItem := QueueItemResponse{
		Done:  done,
		Error: err,
	}
	// check if it is not finished or some error occurred
	if done && err == nil {
		// if everything is ok, get the census information an return it
		internalCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		censusID, ok := data["censusID"].(int)
		if !ok {
			log.Errorf("no census id registered on queue item: %v", err)
			return ErrCantGetCensus
		}

		// get the census from the database by queue_id
		currentCensus, err := capi.sqlc.CensusByID(internalCtx, int64(censusID))
		if err != nil {
			log.Errorf("error getting census by queue id: %v", err)
			return ErrCantGetCensus
		}
		// get current chain id
		chainID, err := capi.sqlc.ChainID(internalCtx)
		if err != nil {
			log.Errorw(err, "error getting chainID")
			return ErrCantGetCensus
		}
		// get values for optional parameters
		censusSize := int32(0)
		if currentCensus.Size.Valid {
			censusSize = currentCensus.Size.Int32
		}
		censusWeight := []byte{}
		if currentCensus.Weight.Valid {
			censusWeight = []byte(currentCensus.Weight.String)
		}
		// encode census
		queueItem.Census = &GetCensusResponse{
			CensusID:   uint64(currentCensus.ID),
			StrategyID: uint64(currentCensus.StrategyID),
			MerkleRoot: common.Bytes2Hex(currentCensus.MerkleRoot),
			URI:        "ipfs://" + currentCensus.Uri.String,
			Size:       censusSize,
			Weight:     new(big.Int).SetBytes(censusWeight).String(),
			ChainID:    uint64(chainID),
			Anonymous:  currentCensus.CensusType == int64(census.AnonymousCensusType),
		}
		// remove the item from the queue
		capi.queue.Dequeue(queueID)
	}
	// encode item response and send it
	res, err := json.Marshal(queueItem)
	if err != nil {
		log.Errorw(ErrEncodeQueueItem, err.Error())
		return ErrEncodeQueueItem
	}
	return ctx.Send(res, api.HTTPstatusOK)
}

// getStrategyCensuses function handler returns the censuses that had been
// generated with the strategy identified by the ID provided.
func (capi *census3API) getStrategyCensuses(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	// get strategy ID
	strategyID, err := strconv.Atoi(ctx.URLParam("strategyID"))
	if err != nil {
		return ErrMalformedCensusID
	}
	// get censuses by this strategy ID
	internalCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rows, err := capi.sqlc.CensusByStrategyID(internalCtx, int64(strategyID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNotFoundCensus
		}
		return ErrCantGetCensus
	}
	// parse and encode response
	censuses := GetCensusesResponse{Censuses: []uint64{}}
	for _, censusInfo := range rows {
		censuses.Censuses = append(censuses.Censuses, uint64(censusInfo.ID))
	}
	res, err := json.Marshal(censuses)
	if err != nil {
		log.Errorw(ErrEncodeCensuses, err.Error())
		return ErrEncodeCensuses
	}
	return ctx.Send(res, api.HTTPstatusOK)
}
