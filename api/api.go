package api

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/vocdoni/census3/census"
	queries "github.com/vocdoni/census3/db/sqlc"
	"go.vocdoni.io/dvote/httprouter"
	api "go.vocdoni.io/dvote/httprouter/apirest"
	"go.vocdoni.io/dvote/log"
)

type Census3APIConf struct {
	Hostname string
	Port     int
	DataDir  string
	Web3URI  string
	GroupKey string
}

type census3API struct {
	conf     Census3APIConf
	web3     string
	db       *sql.DB
	sqlc     *queries.Queries
	endpoint *api.API
	censusDB *census.CensusDB
}

func Init(db *sql.DB, q *queries.Queries, conf Census3APIConf) error {
	newAPI := &census3API{
		conf: conf,
		web3: conf.Web3URI,
		db:   db,
		sqlc: q,
	}
	// get the current chainID
	chainID, err := newAPI.setupChainID()
	if err != nil {
		log.Fatal(err)
	}
	log.Infow("starting API", "chainID", chainID, "web3", conf.Web3URI)

	// create a new http router with the hostname and port provided in the conf
	r := httprouter.HTTProuter{}
	if err = r.Init(conf.Hostname, conf.Port); err != nil {
		return err
	}
	// init API using the http router created
	if newAPI.endpoint, err = api.NewAPI(&r, "/api"); err != nil {
		return err
	}
	// init the census DB
	if newAPI.censusDB, err = census.NewCensusDB(conf.DataDir, conf.GroupKey); err != nil {
		return err
	}
	// init handlers
	if err := newAPI.initTokenHandlers(); err != nil {
		return err
	}
	if err := newAPI.initCensusHandlers(); err != nil {
		return err
	}
	if err := newAPI.initStrategiesHandlers(); err != nil {
		return err
	}
	// TODO: Only for the MVP, remove it.
	if err := newAPI.initDebugHandlers(); err != nil {
		return err
	}
	return nil
}

// setup function gets the chainID from the web3 uri and checks if it is
// registered in the database. If it is registered, the function compares both
// values and panics if they are not the same. If it is not registered, the
// function stores it.
func (capi *census3API) setupChainID() (int64, error) {
	web3client, err := ethclient.Dial(capi.web3)
	if err != nil {
		return -1, fmt.Errorf("error dialing to the web3 endpoint: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// get the chainID from the web3 endpoint
	chainID, err := web3client.ChainID(ctx)
	if err != nil {
		return -1, fmt.Errorf("error getting the chainID from the web3 endpoint: %w", err)
	}
	// get the current chainID from the database
	currentChainID, err := capi.sqlc.ChainID(ctx)
	if err != nil {
		// if it not exists register the value received from the web3 endpoint
		if errors.Is(err, sql.ErrNoRows) {
			_, err := capi.sqlc.SetChainID(ctx, chainID.Int64())
			if err != nil {
				return -1, fmt.Errorf("error setting the chainID in the database: %w", err)
			}
			return chainID.Int64(), nil
		}
		return -1, fmt.Errorf("error getting chainID from the database: %w", err)
	}
	// compare both values
	if currentChainID != chainID.Int64() {
		return -1, fmt.Errorf("received chainID is not the same that registered one: %w", err)
	}
	return currentChainID, nil
}
