package api

import (
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
	newAPI.endpoint, err = api.NewAPI(&r, "/api")
	if err != nil {
		log.Errorw(err, "error starting the API")
		return err
	}

	newAPI.initTokenHandlers()
	newAPI.initHoldersHandlers()
	return nil
}
