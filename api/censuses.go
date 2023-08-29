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
		api.MethodAccessTypePublic, capi.createAndPublishCensus); err != nil {
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
		Published:  currentCensus.Published,
	})
	if err != nil {
		log.Errorw(err, "error encoding census")
		return ErrEncodeCensus
	}
	return ctx.Send(res, api.HTTPstatusOK)
}

// createAndPublishCensus handler creates a census tree based on the token
// holders of the tokens that are included in the given strategy. It recovers
// all the required information from the database, and then creates and publish
// the census merkle tree on IPFS. Then saves the resulting information of the
// census tree in the database and returns its ID.
//
// TODO: This handler is costly, specially for big censuses. It should be refactored to be a background task.
func (capi *census3API) createAndPublishCensus(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	// decode request
	req := &CreateCensusResquest{}
	if err := json.Unmarshal(msg.Data, req); err != nil {
		return ErrMalformedStrategyID
	}
	// get tokens associated to the strategy
	internalCtx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// get the tokens of the strategy provided and check them
	strategyTokens, err := capi.sqlc.TokensByStrategyID(internalCtx, int64(req.StrategyID))
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			log.Errorf("no strategy found for id %d: %s", req.StrategyID, err.Error())
			return ErrNoStrategyTokens
		}
		log.Errorf("error getting strategy with id %d: %s", req.StrategyID, err.Error())
		return ErrCantCreateCensus
	}
	if len(strategyTokens) == 0 {
		log.Errorf("no tokens for strategy %d", req.StrategyID)
		return ErrNoStrategyTokens
	}

	// get the maximun current census ID to calculate the next one, if any
	// census has been created yet, continue
	lastCensusID, err := capi.sqlc.LastCensusID(internalCtx)
	if err != nil && !errors.Is(sql.ErrNoRows, err) {
		log.Errorw(err, "error getting last census ID")
		return ErrCantCreateCensus
	}
	// compute the new censusId and censusType
	newCensusID := int(lastCensusID) + 1
	censusType := census.DefaultCensusType
	if req.Anonymous {
		censusType = census.AnonymousCensusType
	}
	// create the new census on database
	if _, err := capi.sqlc.CreateCensus(internalCtx, queries.CreateCensusParams{
		ID:         int64(newCensusID),
		StrategyID: int64(req.StrategyID),
		CensusType: int64(censusType),
		MerkleRoot: []byte{},
		Published:  false,
	}); err != nil {
		log.Errorw(err, "error saving the census on the database")
		return ErrCantCreateCensus
	}

	// create and publish census merkle tree in background
	go func(db *sql.DB, q *queries.Queries, cID int, tokens []queries.TokensByStrategyIDRow) {
		bgCtx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
		defer cancel()
		// begin a transaction for group sql queries
		tx, err := db.BeginTx(bgCtx, nil)
		if err != nil {
			log.Errorw(err, "error starting database")
			return
		}
		defer func() {
			if err := tx.Rollback(); err != nil && !errors.Is(sql.ErrTxDone, err) {
				log.Errorw(err, "holders transaction rollback failed")
			}
		}()
		qtx := q.WithTx(tx)

		// get holders associated to every strategy token, create a map to avoid
		// duplicates and count the sum of the balances to get the weight of the
		// census
		censusWeight := new(big.Int)
		strategyHolders := map[common.Address]*big.Int{}
		for _, token := range tokens {
			holders, err := qtx.TokenHoldersByTokenID(bgCtx, token.ID)
			if err != nil {
				if errors.Is(sql.ErrNoRows, err) {
					continue
				}
				log.Errorf("error getting token holders of %s: %v", common.BytesToAddress(token.ID), err)
				return
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

		// create a census tree and publish on IPFS
		def := census.NewCensusDefinition(cID, int(req.StrategyID), strategyHolders, req.Anonymous)
		newCensus, err := capi.censusDB.CreateAndPublish(def)
		if err != nil {
			log.Errorw(err, "error creating or publishing the census")
			return
		}
		// check if the census already exists using the merkle root of the generated
		// census
		_, err = qtx.CensusByMerkleRoot(bgCtx, newCensus.RootHash)
		if err == nil {
			log.Info("existing")
			if _, err := qtx.DeleteCensus(bgCtx, int64(cID)); err != nil {
				log.Errorw(err, "census already exists, error deleting new redundant census")
			}
			if err := tx.Commit(); err != nil {
				log.Errorw(err, "error commiting the census deletion on the database")
				return
			}
			return
		}
		if err != nil && !errors.Is(sql.ErrNoRows, err) {
			log.Errorw(err, "error checking if the generated census already exists")
			return
		}
		// save the new census in the SQL database
		sqlURI := new(sql.NullString)
		if err := sqlURI.Scan(newCensus.URI); err != nil {
			log.Errorw(err, "error saving the census on the database")
			return
		}
		sqlCensusSize := sql.NullInt32{}
		if err := sqlCensusSize.Scan(int64(len(strategyHolders))); err != nil {
			log.Errorw(err, "error encoding census size")
			return
		}
		sqlCensusWeight := sql.NullString{}
		if err := sqlCensusWeight.Scan(censusWeight.String()); err != nil {
			log.Errorw(err, "error encoding census size")
			return
		}
		_, err = qtx.UpdateCensus(bgCtx, queries.UpdateCensusParams{
			ID:         int64(newCensus.ID),
			MerkleRoot: newCensus.RootHash,
			Uri:        *sqlURI,
			Size:       sqlCensusSize,
			Weight:     sqlCensusWeight,
			Published:  true,
		})
		if err != nil {
			log.Errorw(err, "error saving the census on the database")
			return
		}
		if err := tx.Commit(); err != nil {
			log.Errorw(err, "error commiting the census on the database")
			return
		}
	}(capi.db, capi.sqlc, newCensusID, strategyTokens)

	// encoding the result and response it
	res, err := json.Marshal(CreateCensusResponse{
		CensusID: uint64(newCensusID),
	})
	if err != nil {
		log.Error("error marshalling census")
		return ErrEncodeCensus
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
