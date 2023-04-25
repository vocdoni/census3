package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/census3/census"
	queries "github.com/vocdoni/census3/db/sqlc"
	"go.vocdoni.io/dvote/httprouter"
	api "go.vocdoni.io/dvote/httprouter/apirest"
	"go.vocdoni.io/dvote/log"
)

func (capi *census3API) initCensusHandlers() {
	capi.endpoint.RegisterMethod("/census", "POST",
		api.MethodAccessTypePublic, capi.createAndPublishCensus)
}

func (capi *census3API) createAndPublishCensus(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	// get strategy and query about the tokens that it includes
	strategyID, err := strconv.Atoi(ctx.URLParam("strategyID"))
	if err != nil {
		return ErrMalformedStrategyID
	}
	// get tokens associated to the strategy
	ctx2, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	strategyTokens, err := capi.sqlc.TokensByStrategyID(ctx2, queries.TokensByStrategyIDParams{
		StrategyID: int64(strategyID),
		Limit:      -1,
		Offset:     0,
	})
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			log.Errorf("no strategy found for id %d: %w", strategyID, err)
			return ErrNotFoundStrategy
		}
		log.Errorf("error getting strategy with id %d: %w", strategyID, err)
		return ErrCantGetStrategy
	}
	// get holders associated to every strategy token
	// create a map to avoid duplicates
	strategyHolders := map[common.Address]int{}
	for _, token := range strategyTokens {
		holders, err := capi.sqlc.TokenHoldersByTokenID(ctx2, queries.TokenHoldersByTokenIDParams{
			TokenID: token.ID,
			Limit:   -1,
			Offset:  0,
		})
		if err != nil {
			if errors.Is(sql.ErrNoRows, err) {
				continue
			}
			return ErrCantGetTokenHolders.Withf("for the token with address %s",
				common.BytesToAddress(token.ID))
		}
		for _, bAddr := range holders {
			holderAddr := common.BytesToAddress(bAddr)
			if _, exists := strategyHolders[holderAddr]; !exists {
				strategyHolders[holderAddr] = 1
			}
		}
	}
	// get the maximun current census ID to calculate the next one
	lastCensusID, err := capi.sqlc.LastCensusID(ctx2)
	if err != nil {
		log.Errorf("error getting last census ID, continue")
	}
	// create a census tree and publish on IPFS
	def := census.DefaultCensusDefinition(int(lastCensusID+1), strategyHolders)
	newCensus, err := capi.censusDB.CreateAndPublish(def)
	if err != nil {
		log.Errorf("error creating or publishing the census: %w", err)
		return ErrCantCreateCensus
	}
	// save the new census in the SQL database
	sqlURI := new(sql.NullString)
	if err := sqlURI.Scan(newCensus.URI); err != nil {
		log.Errorf("error saving the census on the database: %w", err)
		return ErrCantCreateCensus
	}
	_, err = capi.sqlc.CreateCensus(ctx2, queries.CreateCensusParams{
		ID:         int64(newCensus.ID),
		StrategyID: int64(strategyID),
		MerkleRoot: newCensus.Root,
		Uri:        *sqlURI,
	})
	if err != nil {
		log.Errorf("error saving the census on the database: %w", err)
		return ErrCantCreateCensus
	}
	// encoding the result and response it
	res, err := json.Marshal(CreateCensusResponse{
		CensusID: strconv.Itoa(newCensus.ID),
	})
	if err != nil {
		log.Error("error marshalling holders")
		return ErrEncodeStrategyHolders
	}
	return ctx.Send(res, api.HTTPstatusOK)
}
