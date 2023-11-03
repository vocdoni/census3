package api

import (
	"encoding/json"

	"github.com/vocdoni/census3/census"
	"github.com/vocdoni/census3/db"
	"github.com/vocdoni/census3/queue"
	"github.com/vocdoni/census3/state"
	storagelayer "go.vocdoni.io/dvote/data"
	"go.vocdoni.io/dvote/data/downloader"
	"go.vocdoni.io/dvote/data/ipfs"
	"go.vocdoni.io/dvote/data/ipfs/ipfsconnect"
	"go.vocdoni.io/dvote/httprouter"
	api "go.vocdoni.io/dvote/httprouter/apirest"
	"go.vocdoni.io/dvote/log"
)

type Census3APIConf struct {
	Hostname      string
	Port          int
	DataDir       string
	GroupKey      string
	Web3Providers state.Web3Providers
}

type census3API struct {
	conf       Census3APIConf
	db         *db.DB
	endpoint   *api.API
	censusDB   *census.CensusDB
	queue      *queue.BackgroundQueue
	w3p        state.Web3Providers
	storage    storagelayer.Storage
	downloader *downloader.Downloader
}

func Init(db *db.DB, conf Census3APIConf) (*census3API, error) {
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
		return nil, err
	}
	// init API using the http router created
	if newAPI.endpoint, err = api.NewAPI(&r, "/api"); err != nil {
		return nil, err
	}
	// init the IPFS service and the storage layer and connect them
	ipfsConfig := storagelayer.IPFSNewConfig(conf.DataDir)
	newAPI.storage, err = storagelayer.Init(storagelayer.IPFS, ipfsConfig)
	if err != nil {
		return nil, err
	}
	var ipfsConn *ipfsconnect.IPFSConnect
	if len(conf.GroupKey) > 0 {
		ipfsConn = ipfsconnect.New(conf.GroupKey, newAPI.storage.(*ipfs.Handler))
		ipfsConn.Start()
	}
	// init the downloader using the storage layer
	newAPI.downloader = downloader.NewDownloader(newAPI.storage)
	newAPI.downloader.Start()
	// init the census DB using the storage layer
	if newAPI.censusDB, err = census.NewCensusDB(conf.DataDir, newAPI.storage); err != nil {
		return nil, err
	}
	// init handlers
	if err := newAPI.initAPIHandlers(); err != nil {
		return nil, err
	}
	if err := newAPI.initTokenHandlers(); err != nil {
		return nil, err
	}
	if err := newAPI.initCensusHandlers(); err != nil {
		return nil, err
	}
	if err := newAPI.initStrategiesHandlers(); err != nil {
		return nil, err
	}
	return newAPI, nil
}

func (capi *census3API) Stop() error {
	capi.downloader.Stop()
	if err := capi.storage.Stop(); err != nil {
		return err
	}
	return nil
}

func (capi *census3API) initAPIHandlers() error {
	return capi.endpoint.RegisterMethod("/info", "GET",
		api.MethodAccessTypePublic, capi.getAPIInfo)
}

func (capi *census3API) getAPIInfo(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	info := &APIInfo{
		SupportedChains: []SupportedChain{},
	}
	for _, provider := range capi.w3p {
		info.SupportedChains = append(info.SupportedChains, SupportedChain{
			ChainID:   provider.ChainID,
			ShortName: provider.ShortName,
			Name:      provider.Name,
		})
	}
	res, err := json.Marshal(info)
	if err != nil {
		log.Errorw(err, "error encoding api info")
		return ErrEncodeAPIInfo
	}
	return ctx.Send(res, api.HTTPstatusOK)
}
