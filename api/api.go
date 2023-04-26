package api

import (
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
}

type census3API struct {
	conf     Census3APIConf
	web3     string
	sqlc     *queries.Queries
	endpoint *api.API
	censusDB *census.CensusDB
}

func Init(db *queries.Queries, conf Census3APIConf) error {
	newAPI := &census3API{
		conf: conf,
		web3: conf.Web3URI,
		sqlc: db,
	}
	// create a new http router with the hostname and port provided in the conf
	r := httprouter.HTTProuter{}
	err := r.Init(conf.Hostname, conf.Port)
	if err != nil {
		log.Errorw(err, "error creating http router")
		return err
	}
	// init API using the http router created
	if newAPI.endpoint, err = api.NewAPI(&r, "/api"); err != nil {
		log.Errorw(err, "error starting the API")
		return err
	}
	// init the census DB
	if newAPI.censusDB, err = census.NewCensusDB(conf.DataDir); err != nil {
		log.Errorw(err, "error starting census database")
	}
	newAPI.initTokenHandlers()
	newAPI.initCensusHandlers()
	// TODO: Only for the MVP, remove it.
	newAPI.initDebugHandlers()
	return nil
}
