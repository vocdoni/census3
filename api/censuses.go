package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
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
		api.MethodAccessTypePublic, capi.enqueueCensus); err != nil {
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
	currentCensus, err := capi.db.QueriesRO.CensusByID(internalCtx, int64(censusID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNotFoundCensus.WithErr(err)
		}
		return ErrCantGetCensus.WithErr(err)
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
		CensusID:   uint32(censusID),
		StrategyID: uint32(currentCensus.StrategyID),
		MerkleRoot: common.Bytes2Hex(currentCensus.MerkleRoot),
		URI:        "ipfs://" + currentCensus.Uri.String,
		Size:       censusSize,
		Weight:     new(big.Int).SetBytes(censusWeight).String(),
		Anonymous:  currentCensus.CensusType == int64(census.AnonymousCensusType),
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
	res, err := json.Marshal(CreateCensusResponse{
		QueueID: queueID,
	})
	if err != nil {
		return ErrEncodeCensus.WithErr(err)
	}
	return ctx.Send(res, api.HTTPstatusOK)
}

// innerCensusID generates a unique identifier by concatenating the BlockNumber, StrategyID,
// and a numerical representation of the Anonymous flag from a CreateCensusRequest struct.
// The BlockNumber and StrategyID are concatenated as they are, and the Anonymous flag is
// represented as 1 for true and 0 for false. This concatenated string is then converted
// to a uint64 to create a unique identifier.
func innerCensusID(c *CreateCensusRequest) uint64 {
	// Convert the boolean to a uint32: 1 for true, 0 for false
	var anonymousUint uint32
	if c.Anonymous {
		anonymousUint = 1
	}

	// Concatenate the three values as strings
	concatenated := fmt.Sprintf("%d%d%d", c.BlockNumber, c.StrategyID, anonymousUint)

	// Convert the concatenated string back to a uint64
	result, err := strconv.ParseUint(concatenated, 10, 64)
	if err != nil {
		panic(err)
	}

	return result
}

// createAndPublishCensus method creates a census tree based on the token
// holders of the tokens that are included in the given strategy. It recovers
// all the required information from the database, and then creates and publish
// the census merkle tree on IPFS. Then saves the resulting information of the
// census tree in the database.
func (capi *census3API) createAndPublishCensus(req *CreateCensusRequest, qID string) (uint64, error) {
	bgCtx, cancel := context.WithTimeout(context.Background(), censusCreationTimeout)
	defer cancel()
	// begin a transaction for group sql queries
	tx, err := capi.db.RW.BeginTx(bgCtx, nil)
	if err != nil {
		return 0, ErrCantCreateCensus.WithErr(err)
	}
	defer func() {
		if err := tx.Rollback(); err != nil && !errors.Is(sql.ErrTxDone, err) {
			log.Errorw(err, "holders transaction rollback failed")
		}
	}()
	qtx := capi.db.QueriesRW.WithTx(tx)

	strategyTokens, err := qtx.TokensByStrategyID(bgCtx, int64(req.StrategyID))
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return 0, ErrNoStrategyTokens.WithErr(err)
		}
		return 0, ErrCantCreateCensus.WithErr(err)
	}
	if len(strategyTokens) == 0 {
		return 0, ErrNoStrategyTokens.WithErr(err)
	}

	// compute the new censusId and censusType
	newCensusID := innerCensusID(req)

	// check if the census already exists
	_, err = qtx.CensusByID(bgCtx, int64(newCensusID))
	if err != nil {
		if !errors.Is(sql.ErrNoRows, err) {
			return 0, ErrCantCreateCensus.WithErr(err)
		}
	} else {
		return 0, ErrCensusAlreadyExists.Withf("census %d already exists", newCensusID)
	}

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
			return 0, ErrCantGetTokenHolders.WithErr(err)
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
		return 0, ErrNotFoundTokenHolders.With("no holders for strategy")
	}

	// create a census tree and publish on IPFS
	def := census.NewCensusDefinition(newCensusID, req.StrategyID, strategyHolders, req.Anonymous)
	newCensus, err := capi.censusDB.CreateAndPublish(def)
	if err != nil {
		return 0, ErrCantCreateCensus.WithErr(err)
	}

	// save the new census in the SQL database
	sqlURI := new(sql.NullString)
	if err := sqlURI.Scan(newCensus.URI); err != nil {
		return 0, ErrCantCreateCensus.WithErr(err)
	}
	sqlCensusSize := sql.NullInt32{}
	if err := sqlCensusSize.Scan(int64(len(strategyHolders))); err != nil {
		return 0, ErrCantCreateCensus.WithErr(err)
	}
	sqlCensusWeight := sql.NullString{}
	if err := sqlCensusWeight.Scan(censusWeight.String()); err != nil {
		return 0, ErrCantCreateCensus.WithErr(err)
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
		return 0, ErrCantCreateCensus.WithErr(err)
	}
	if err := tx.Commit(); err != nil {
		return 0, ErrCantCreateCensus.WithErr(err)
	}
	return newCensus.ID, nil
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
		internalCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		censusID, ok := data["censusID"].(int)
		if !ok {
			log.Errorf("no census id registered on queue item")
			return ErrCantGetCensus
		}

		// get the census from the database by queue_id
		currentCensus, err := capi.db.QueriesRO.CensusByID(internalCtx, int64(censusID))
		if err != nil {
			return ErrCantGetCensus.WithErr(err)
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
		queueCensus.Census = &GetCensusResponse{
			CensusID:   uint32(currentCensus.ID),
			StrategyID: uint32(currentCensus.StrategyID),
			MerkleRoot: common.Bytes2Hex(currentCensus.MerkleRoot),
			URI:        "ipfs://" + currentCensus.Uri.String,
			Size:       censusSize,
			Weight:     new(big.Int).SetBytes(censusWeight).String(),
			Anonymous:  currentCensus.CensusType == int64(census.AnonymousCensusType),
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
	internalCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rows, err := capi.db.QueriesRO.CensusByStrategyID(internalCtx, int64(strategyID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNotFoundCensus.WithErr(err)
		}
		return ErrCantGetCensus.WithErr(err)
	}
	// parse and encode response
	censuses := GetCensusesResponse{Censuses: []uint32{}}
	for _, censusInfo := range rows {
		censuses.Censuses = append(censuses.Censuses, uint32(censusInfo.ID))
	}
	res, err := json.Marshal(censuses)
	if err != nil {
		return ErrEncodeCensuses.WithErr(err)
	}
	return ctx.Send(res, api.HTTPstatusOK)
}
