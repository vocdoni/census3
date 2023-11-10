package api

import (
	"encoding/json"
	"path/filepath"

	"github.com/vocdoni/census3/db"
	"github.com/vocdoni/census3/queue"
	"github.com/vocdoni/census3/service"
	"github.com/vocdoni/census3/state"
	"go.vocdoni.io/dvote/api/censusdb"
	storagelayer "go.vocdoni.io/dvote/data"
	"go.vocdoni.io/dvote/data/downloader"
	"go.vocdoni.io/dvote/data/ipfs"
	"go.vocdoni.io/dvote/data/ipfs/ipfsconnect"
	vocdoniDB "go.vocdoni.io/dvote/db"
	"go.vocdoni.io/dvote/db/metadb"
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
	ExtProviders  map[state.TokenType]service.HolderProvider
}

type census3API struct {
	conf         Census3APIConf
	db           *db.DB
	endpoint     *api.API
	censusDB     *censusdb.CensusDB
	queue        *queue.BackgroundQueue
	w3p          state.Web3Providers
	storage      storagelayer.Storage
	downloader   *downloader.Downloader
	extProviders map[state.TokenType]service.HolderProvider
}

func Init(db *db.DB, conf Census3APIConf) (*census3API, error) {
	newAPI := &census3API{
		conf:         conf,
		db:           db,
		w3p:          conf.Web3Providers,
		queue:        queue.NewBackgroundQueue(),
		extProviders: conf.ExtProviders,
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
	// init the database for the census trees
	censusesDB, err := metadb.New(vocdoniDB.TypePebble, filepath.Join(conf.DataDir, "censusdb"))
	if err != nil {
		return nil, err
	}
	// init the censusDB of the API
	if newAPI.censusDB = censusdb.NewCensusDB(censusesDB); err != nil {
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
