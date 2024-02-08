package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/vocdoni/census3/db"
	"github.com/vocdoni/census3/db/annotations"
	queries "github.com/vocdoni/census3/db/sqlc"
	"github.com/vocdoni/census3/internal/queue"
	"github.com/vocdoni/census3/scanner/providers"
	farcaster "github.com/vocdoni/census3/scanner/providers/farcaster"
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
	Hostname        string
	Port            int
	DataDir         string
	GroupKey        string
	Web3Providers   web3.NetworkEndpoints
	HolderProviders map[uint64]providers.HolderProvider
	AdminToken      string
}

type census3API struct {
	conf            Census3APIConf
	db              *db.DB
	farcasterdb     *farcaster.DB
	endpoint        *api.API
	censusDB        *censusdb.CensusDB
	queue           *queue.BackgroundQueue
	w3p             web3.NetworkEndpoints
	storage         storagelayer.Storage
	downloader      *downloader.Downloader
	holderProviders map[uint64]providers.HolderProvider
	cache           *lru.Cache[CacheKey, any]
}

func Init(db *db.DB, farcasterdb *farcaster.DB, conf Census3APIConf) (*census3API, error) {
	cache, err := lru.New[CacheKey, any](apiCacheSize)
	if err != nil {
		return nil, err
	}
	newAPI := &census3API{
		conf:            conf,
		db:              db,
		farcasterdb:     farcasterdb,
		w3p:             conf.Web3Providers,
		queue:           queue.NewBackgroundQueue(),
		holderProviders: conf.HolderProviders,
		cache:           cache,
	}
	// get the current chainID
	log.Infow("starting API", "web3Providers", conf.Web3Providers.String())

	// create a new http router with the hostname and port provided in the conf
	r := httprouter.HTTProuter{}
	if err = r.Init(conf.Hostname, conf.Port); err != nil {
		return nil, err
	}
	// expose metrics endpoint
	r.ExposePrometheusEndpoint("/metrics")

	// init API using the http router created
	if newAPI.endpoint, err = api.NewAPI(&r, "/api"); err != nil {
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
func (capi *census3API) CreateInitialTokens(tokensPath string, farcaster bool) error {
	// check if farcaster activated
	if farcaster {
		if err := capi.storeFarcaster(); err != nil {
			return fmt.Errorf("error creating farcaster: %w", err)
		}
	}

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
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
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
		// get the correct holder provider for the token type
		tokenType := providers.TokenTypeID(token.Type)
		provider, exists := capi.holderProviders[tokenType]
		if !exists {
			return ErrCantCreateCensus.With("token type not supported")
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
			return ErrCantGetToken.WithErr(err)
		}
		symbol, err := provider.Symbol([]byte(token.ExternalID))
		if err != nil {
			return ErrCantGetToken.WithErr(err)
		}
		decimals, err := provider.Decimals([]byte(token.ExternalID))
		if err != nil {
			return ErrCantGetToken.WithErr(err)
		}
		totalSupply, err := provider.TotalSupply([]byte(token.ExternalID))
		if err != nil {
			return ErrCantGetToken.WithErr(err)
		}
		// get the chain address for the token based on the chainID and tokenID
		chainAddress, ok := capi.w3p.ChainAddress(token.ChainID, address.String())
		if !ok {
			return ErrChainIDNotSupported.Withf("chainID: %d, tokenID: %s", token.ChainID, token.ID)
		}
		iconURI, err := provider.IconURI([]byte(token.ExternalID))
		if err != nil {
			return ErrCantGetToken.WithErr(err)
		}

		addr := common.HexToAddress(token.ID)
		_, err = qtx.CreateToken(ctx, queries.CreateTokenParams{
			ID:            addr.Bytes(),
			Name:          name,
			Symbol:        symbol,
			Decimals:      decimals,
			TotalSupply:   annotations.BigInt(totalSupply.String()),
			CreationBlock: 0,
			TypeID:        providers.TokenTypeID(token.Type),
			Synced:        false,
			Tags:          token.Tags,
			ChainID:       token.ChainID,
			ChainAddress:  chainAddress,
			ExternalID:    token.ExternalID,
			IconUri:       iconURI,
		})
		if err != nil {
			if strings.Contains(err.Error(), "UNIQUE constraint failed") {
				return nil
			}
			return err
		}
		strategyID, err := capi.createDefaultTokenStrategy(ctx, qtx,
			addr, token.ChainID, chainAddress, symbol, token.ExternalID)
		if err != nil {
			return err
		}
		if _, err := qtx.UpdateTokenDefaultStrategy(ctx, queries.UpdateTokenDefaultStrategyParams{
			ID:              addr.Bytes(),
			DefaultStrategy: uint64(strategyID),
			ChainID:         token.ChainID,
			ExternalID:      token.ExternalID,
		}); err != nil {
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (capi *census3API) storeFarcaster() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
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

	// get the correct holder provider for the token type
	tokenType := providers.TokenTypeID(providers.CONTRACT_NAME_FARCASTER)
	provider, exists := capi.holderProviders[tokenType]
	if !exists {
		return ErrCantCreateCensus.With("token type not supported")
	}

	if err := provider.SetRef(web3.Web3ProviderRef{
		HexAddress: farcaster.IdRegistryAddress,
		ChainID:    farcaster.ChainID,
	}); err != nil {
		return ErrInitializingWeb3.WithErr(err)
	}

	totalSupply, err := provider.TotalSupply(nil)
	if err != nil {
		return ErrCantGetToken.WithErr(err)
	}

	name, _ := provider.Name(nil)
	creationBlock, _ := provider.CreationBlock(context.TODO(), farcaster.FarcasterIDRegistryType)
	// get the chain address for the token based on the chainID and tokenID
	chainAddress, ok := capi.w3p.ChainAddress(provider.ChainID(), farcaster.IdRegistryAddress)
	if !ok {
		return ErrChainIDNotSupported.Withf("chainID: %d, tokenID: %s", provider.ChainID(), farcaster.IdRegistryAddress)
	}

	addr := provider.Address(farcaster.FarcasterIDRegistryType)
	_, err = qtx.CreateToken(ctx, queries.CreateTokenParams{
		ID:            addr.Bytes(),
		Name:          name,
		TotalSupply:   annotations.BigInt(totalSupply.String()),
		CreationBlock: int64(creationBlock),
		TypeID:        tokenType,
		Synced:        false,
		ChainID:       provider.ChainID(),
		ChainAddress:  chainAddress,
		LastBlock:     int64(creationBlock),
	})
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return nil
		}
		return err
	}
	strategyID, err := capi.createDefaultTokenStrategy(ctx, qtx,
		common.HexToAddress(farcaster.IdRegistryAddress), provider.ChainID(), chainAddress, "", "")
	if err != nil {
		return err
	}
	if _, err := qtx.UpdateTokenDefaultStrategy(ctx, queries.UpdateTokenDefaultStrategyParams{
		ID:              addr.Bytes(),
		DefaultStrategy: uint64(strategyID),
		ChainID:         provider.ChainID(),
		ExternalID:      "",
	}); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
