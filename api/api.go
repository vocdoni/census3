package api

import (
	"encoding/json"

	"github.com/vocdoni/census3/census"
	"github.com/vocdoni/census3/db"
	"github.com/vocdoni/census3/queue"
	"go.vocdoni.io/dvote/httprouter"
	api "go.vocdoni.io/dvote/httprouter/apirest"
	"go.vocdoni.io/dvote/log"
)

type Census3APIConf struct {
	Hostname      string
	Port          int
	DataDir       string
	GroupKey      string
	Web3Providers map[uint64]string
}

type census3API struct {
	conf     Census3APIConf
	db       *db.DB
	endpoint *api.API
	censusDB *census.CensusDB
	queue    *queue.BackgroundQueue
	w3p      map[uint64]string
}

func Init(db *db.DB, conf Census3APIConf) error {
	newAPI := &census3API{
		conf:  conf,
		db:    db,
		w3p:   conf.Web3Providers,
		queue: queue.NewBackgroundQueue(),
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
	if err := newAPI.initAPIHandlers(); err != nil {
		return err
	}
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

func (capi *census3API) initAPIHandlers() error {
	return capi.endpoint.RegisterMethod("/info", "GET",
		api.MethodAccessTypePublic, capi.getAPIInfo)
}

func (capi *census3API) getAPIInfo(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	chainIDs := []uint64{}
	for chainID := range capi.w3p {
		chainIDs = append(chainIDs, chainID)
	}
	info := map[string]any{"chainIDs": chainIDs}
	res, err := json.Marshal(info)
	if err != nil {
		log.Errorw(err, "error encoding api info")
		return ErrEncodeAPIInfo
	}
	return ctx.Send(res, api.HTTPstatusOK)
}
