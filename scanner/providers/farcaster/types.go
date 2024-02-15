package farcaster

import (
	"sync/atomic"

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
}

type FarcasterProvider struct {
	// web3
	endpoints        web3.NetworkEndpoints
	client           *ethclient.Client
	contracts        FarcasterContracts
	lastNetworkBlock uint64
	// db
	db *DB
}

type FarcasterUserData struct {
	FID uint64
}
