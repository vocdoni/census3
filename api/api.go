package api

import (
	"database/sql"

	"github.com/vocdoni/census3/census"
	queries "github.com/vocdoni/census3/db/sqlc"
	"go.vocdoni.io/dvote/httprouter"
	api "go.vocdoni.io/dvote/httprouter/apirest"
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
	// create a new http router with the hostname and port provided in the conf
	r := httprouter.HTTProuter{}
	err := r.Init(conf.Hostname, conf.Port)
	if err != nil {
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
	return nil
}
