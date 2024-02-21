package farcaster

import (
	"context"
	"math/big"
	"sync"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	fcir "github.com/vocdoni/census3/contracts/farcaster/idRegistry"
	fckr "github.com/vocdoni/census3/contracts/farcaster/keyRegistry"
	"github.com/vocdoni/census3/scanner/providers/web3"
)

type FarcasterProviderConf struct {
	Endpoints web3.NetworkEndpoints
	DB        *DB
}

type FarcasterContracts struct {
	keyRegistry       *fckr.FarcasterKeyRegistry
	idRegistry        *fcir.FarcasterIDRegistry
	idRegistrySynced  atomic.Bool
	keyRegistrySynced atomic.Bool
	lastBlock         atomic.Uint64
}

type FarcasterProvider struct {
	// web3
	endpoints        web3.NetworkEndpoints
	client           *ethclient.Client
	contracts        FarcasterContracts
	lastNetworkBlock atomic.Uint64
	// db
	db *DB
	// iteration vars
	currentScannerHolders    map[common.Address]*big.Int
	currentScannerHoldersMtx sync.Mutex
	scannerCtx               context.Context
	cancelScanner            context.CancelFunc
}

type FarcasterUserData struct {
	FID uint64
}
