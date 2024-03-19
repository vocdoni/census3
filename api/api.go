package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/vocdoni/census3/db"
	"github.com/vocdoni/census3/db/annotations"
	queries "github.com/vocdoni/census3/db/sqlc"
	"github.com/vocdoni/census3/helpers/queue"
	"github.com/vocdoni/census3/scanner"
	"github.com/vocdoni/census3/scanner/providers"
	"github.com/vocdoni/census3/scanner/providers/web3"
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
	"go.vocdoni.io/dvote/types"
)

type Census3APIConf struct {
	MainCtx         context.Context
	Hostname        string
	Port            int
	DataDir         string
	GroupKey        string
	AdminToken      string
	ScannerCooldown time.Duration
}

type census3API struct {
	conf       Census3APIConf
	db         *db.DB
	scanner    *scanner.Scanner
	endpoint   *api.API
	censusDB   *censusdb.CensusDB
	queue      *queue.BackgroundQueue
	storage    storagelayer.Storage
	downloader *downloader.Downloader
	cache      *lru.Cache[CacheKey, any]
	router     *httprouter.HTTProuter
}

func Init(db *db.DB, scanner *scanner.Scanner, conf Census3APIConf) (*census3API, error) {
	cache, err := lru.New[CacheKey, any](apiCacheSize)
	if err != nil {
		return nil, err
	}
	newAPI := &census3API{
		conf:    conf,
		db:      db,
		scanner: scanner,
		queue:   queue.NewBackgroundQueue(),
		cache:   cache,
		router:  &httprouter.HTTProuter{},
	}
	// get the current chainID
	log.Infow("starting API",
		"networks", scanner.Networks().String(),
		"providers", scanner.SupportedTypes())

	// create a new http router with the hostname and port provided in the conf
	if err = newAPI.router.Init(conf.Hostname, conf.Port); err != nil {
		return nil, err
	}
	// expose metrics endpoint
	newAPI.router.ExposePrometheusEndpoint("/metrics")

	// init API using the http router created
	if newAPI.endpoint, err = api.NewAPI(newAPI.router, "/api"); err != nil {
		return nil, err
	}
	// set admin token
	newAPI.endpoint.SetAdminToken(conf.AdminToken)
	// init the IPFS service and the storage layer and connect them
	newAPI.storage = new(ipfs.Handler)
	if err = newAPI.storage.Init(&types.DataStore{Datadir: conf.DataDir}); err != nil {
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
	newAPI.censusDB = censusdb.NewCensusDB(censusesDB)
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
	if err := newAPI.initHoldersHandlers(); err != nil {
		return nil, err
	}
	return newAPI, nil
}

func (capi *census3API) Stop() error {
	capi.downloader.Stop()
	capi.cache.Purge()
	if err := capi.storage.Stop(); err != nil {
		return err
	}
	return capi.db.Close()
}

func (capi *census3API) initAPIHandlers() error {
	if err := capi.endpoint.RegisterMethod("/info", http.MethodGet,
		api.MethodAccessTypePublic, capi.getAPIInfo); err != nil {
		return err
	}
	if err := capi.endpoint.RegisterMethod("/db/import", http.MethodPost,
		api.MethodAccessTypeAdmin, capi.importDatabase); err != nil {
		return err
	}
	if err := capi.endpoint.RegisterMethod("/db/export", http.MethodGet,
		api.MethodAccessTypeAdmin, capi.exportDatabase); err != nil {
		return err
	}
	return nil
}

func (capi *census3API) getAPIInfo(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	info := &APIInfo{
		SupportedChains: []SupportedChain{},
	}
	for _, provider := range capi.scanner.Networks() {
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

func (capi *census3API) importDatabase(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	go func(dbDump []byte) {
		log.Infow("importing database", "size", len(dbDump)/1024/1024)
		// backup scanner networks and providers to stop it and restore it after the import
		networks := capi.scanner.Networks()
		providers := []providers.HolderProvider{}
		for _, provider := range capi.scanner.HolderProviders() {
			providers = append(providers, provider)
		}
		log.Debug("closing database and stopping scanner")
		// stop the scanner and close the database
		if err := capi.db.Close(); err != nil {
			log.Error("error closing database")
			return
		}
		capi.scanner.Stop()
		// overwrite the database file with the temp file, remove the old database and create a new one
		log.Debug("overwriting database")
		if err := os.RemoveAll(filepath.Join(capi.conf.DataDir, "census3.sql")); err != nil {
			log.Errorw(err, "error removing old database")
			return
		}
		// open the database
		database, err := db.Init(capi.conf.DataDir, "census3.sql")
		if err != nil {
			log.Errorw(err, "error opening database")
			return
		}
		if err := database.Import(context.Background(), dbDump); err != nil {
			log.Errorw(err, "error importing database")
			return
		}
		capi.db = database
		// restore the database and the scanner and start the scanner
		capi.scanner = scanner.NewScanner(database, networks, capi.conf.ScannerCooldown)
		if err := capi.scanner.SetProviders(providers...); err != nil {
			log.Errorw(err, "error setting providers")
			return
		}
		log.Debug("starting scanner")
		go capi.scanner.Start(capi.conf.MainCtx)
		log.Debugw("scanner started",
			"networks", capi.scanner.Networks().String(),
			"providers", capi.scanner.SupportedTypes())
	}(msg.Data)
	return ctx.Send([]byte("Ok"), api.HTTPstatusOK)

}

func (capi *census3API) exportDatabase(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	dbDump, err := capi.db.Export(context.Background())
	if err != nil {
		log.Errorw(err, "error exporting database")
		return ErrDatabaseExport.WithErr(err)
	}
	ctx.SetHeader("Content-Disposition", "attachment; filename=census3.sql")
	ctx.SetHeader("Content-Length", "application/octet-stream")
	return ctx.Send(dbDump, api.HTTPstatusOK)
}

// CreateInitialTokens creates the tokens defined in the file provided in the
// tokensPath if it is defined. This function is used to create the initial
// tokens of the census3 database. It read the tokens file, parse it and create
// the tokens in the database. It also creates the default token strategy for
// each token. The tokens file must be a json file with the following format:
// [
//
//	{
//	  "ID": "0x0000000000000000000000000000000000000001"
//	  "chainID": "token name",
//	  "externalID": "token symbol",
//	  "type": "erc20",
//	},
//	...
//
// ]
func (capi *census3API) CreateInitialTokens(tokensPath string) error {
	// skip if the tokens file is not defined
	if tokensPath == "" {
		return nil
	}
	// read the tokens file
	content, err := os.ReadFile(tokensPath)
	if err != nil {
		return err
	}
	// parse the tokens file
	tokens := []GetTokenResponse{}
	if err := json.Unmarshal(content, &tokens); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	// create the tokens
	tx, err := capi.db.RW.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err := tx.Rollback(); err != nil && !errors.Is(sql.ErrTxDone, err) {
			log.Errorw(err, "create token transaction rollback failed")
		}
	}()
	qtx := capi.db.QueriesRW.WithTx(tx)
	for _, token := range tokens {
		// if something fails getting the token information, skip it
		// if something fails interacting with the database, return the error

		// get the correct holder provider for the token type
		tokenType := providers.TokenTypeID(token.Type)
		provider, exists := capi.scanner.HolderProvider(tokenType)
		if !exists {
			log.Warnw("token type provided in initial list not supported, check provider is set. SKIPPING...",
				"tokenID", token.ID,
				"chainID", token.ChainID,
				"externalID", token.ExternalID,
				"type", token.Type)
			continue
		}
		if !provider.IsExternal() {
			if err := provider.SetRef(web3.Web3ProviderRef{
				HexAddress: token.ID,
				ChainID:    token.ChainID,
			}); err != nil {
				return ErrInitializingWeb3.WithErr(err)
			}
		}
		// get token information from the external provider
		address := provider.Address([]byte(token.ExternalID))
		name, err := provider.Name([]byte(token.ExternalID))
		if err != nil {
			log.Warnw("can't get token name",
				"tokenID", token.ID,
				"chainID", token.ChainID,
				"externalID", token.ExternalID,
				"error", err)
			continue
		}
		symbol, err := provider.Symbol([]byte(token.ExternalID))
		if err != nil {
			log.Warnw("can't get token symbol",
				"tokenID", token.ID,
				"chainID", token.ChainID,
				"externalID", token.ExternalID,
				"error", err)
			continue
		}
		decimals, err := provider.Decimals([]byte(token.ExternalID))
		if err != nil {
			log.Warnw("can't get token decimals",
				"tokenID", token.ID,
				"chainID", token.ChainID,
				"externalID", token.ExternalID,
				"error", err)
			continue
		}
		totalSupply, err := provider.TotalSupply([]byte(token.ExternalID))
		if err != nil {
			log.Warnw("can't get token total supply",
				"tokenID", token.ID,
				"chainID", token.ChainID,
				"externalID", token.ExternalID,
				"error", err)
			continue
		}
		// get the chain address for the token based on the chainID and tokenID
		chainAddress, ok := capi.scanner.Networks().ChainAddress(token.ChainID, address.String())
		if !ok {
			log.Warnw("can't get chain address", "chainID", token.ChainID, "tokenID", token.ID)
			continue
		}
		iconURI, err := provider.IconURI([]byte(token.ExternalID))
		if err != nil {
			log.Warnw("can't get token icon URI",
				"tokenID", token.ID,
				"chainID", token.ChainID,
				"externalID", token.ExternalID,
				"error", err)
			continue
		}
		// create the token in the database
		addr := common.HexToAddress(token.ID)
		_, err = qtx.CreateToken(ctx, queries.CreateTokenParams{
			ID:            addr.Bytes(),
			Name:          name,
			Symbol:        symbol,
			Decimals:      decimals,
			TotalSupply:   annotations.BigInt(totalSupply.String()),
			CreationBlock: int64(token.StartBlock),
			TypeID:        providers.TokenTypeID(token.Type),
			Synced:        false,
			Tags:          token.Tags,
			ChainID:       token.ChainID,
			ChainAddress:  chainAddress,
			ExternalID:    token.ExternalID,
			IconUri:       iconURI,
			LastBlock:     int64(token.StartBlock),
		})
		if err != nil {
			if !strings.Contains(err.Error(), "UNIQUE constraint failed") {
				log.Errorf("error creating token: %s", err)
			}
			continue
		}
		strategyID, err := capi.createDefaultTokenStrategy(ctx, qtx,
			addr, token.ChainID, chainAddress, symbol, token.ExternalID)
		if err != nil {
			log.Errorf("error creating default token strategy: %s", err)
			continue
		}
		if _, err := qtx.UpdateTokenDefaultStrategy(ctx, queries.UpdateTokenDefaultStrategyParams{
			ID:              addr.Bytes(),
			DefaultStrategy: uint64(strategyID),
			ChainID:         token.ChainID,
			ExternalID:      token.ExternalID,
		}); err != nil {
			log.Errorf("error updating token default strategy: %s", err)
			continue
		}
		log.Infow("token created", "tokenID", token.ID, "chainID", token.ChainID, "externalID", token.ExternalID)
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
