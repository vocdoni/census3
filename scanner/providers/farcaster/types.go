package farcaster

import (
	"math/big"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	fcir "github.com/vocdoni/census3/contracts/farcaster/idRegistry"
	fckr "github.com/vocdoni/census3/contracts/farcaster/keyRegistry"
	"github.com/vocdoni/census3/scanner/providers/web3"
)

type FarcasterProviderConf struct {
	APIEndpoint string
	APICooldown time.Duration
	AccessToken string
	Endpoints   web3.NetworkEndpoints
	DB          *DB
}

type FarcasterContracts struct {
	keyRegistry       *fckr.FarcasterKeyRegistry
	idRegistry        *fcir.FarcasterIDRegistry
	idRegistrySynced  atomic.Bool
	keyRegistrySynced atomic.Bool
}

type FarcasterProvider struct {
	// api
	apiEndpoint string
	apiCooldown time.Duration
	accessToken string
	// web3
	endpoints        web3.NetworkEndpoints
	client           *ethclient.Client
	contracts        FarcasterContracts
	lastNetworkBlock uint64
	// db
	db *DB
}

type FarcasterUserData struct {
	// FID is the Farcaster ID
	FID *big.Int
	// Username is the username of the user
	Username string
	// Signer is the signer address of the user
	Signer common.Hash
	// CustodyAddress is the custody address of the user
	CustodyAddress common.Address
	// RecoveryAddress is the address used to recover the user's account
	RecoveryAddress common.Address
	// AppKeys is the array of keys of the user found in the KeyRegistry
	AppKeys []common.Hash
	// LinkedEVM is the array of linked EVM addresses to the user
	LinkedEVM []common.Address
}

type KRLogData struct {
	Key      common.Hash
	KeyBytes []byte
}
