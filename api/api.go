package api

import (
	"database/sql"

	"github.com/vocdoni/census3/census"
	queries "github.com/vocdoni/census3/db/sqlc"
	"go.vocdoni.io/dvote/httprouter"
	api "go.vocdoni.io/dvote/httprouter/apirest"
	"go.vocdoni.io/dvote/log"
)

type Census3APIConf struct {
	Hostname      string
	Port          int
	DataDir       string
	GroupKey      string
	Web3Providers map[int64]string
}

type census3API struct {
	conf     Census3APIConf
	db       *sql.DB
	sqlc     *queries.Queries
	endpoint *api.API
	censusDB *census.CensusDB
	w3p      map[int64]string
}

func Init(db *sql.DB, q *queries.Queries, conf Census3APIConf) error {
	newAPI := &census3API{
		conf: conf,
		db:   db,
		sqlc: q,
		w3p:  conf.Web3Providers,
	}
	// get the current chainID
	log.Infow("starting API", "chainID-web3Providers", conf.Web3Providers)

	// create a new http router with the hostname and port provided in the conf
	var err error
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
