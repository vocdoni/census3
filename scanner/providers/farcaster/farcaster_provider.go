package farcaster

import (
	"bytes"
	"context"
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pressly/goose/v3"
	fcir "github.com/vocdoni/census3/contracts/farcaster/idRegistry"
	fckr "github.com/vocdoni/census3/contracts/farcaster/keyRegistry"
	"github.com/vocdoni/census3/scanner/providers"
	queries "github.com/vocdoni/census3/scanner/providers/farcaster/sqlc"
	"github.com/vocdoni/census3/scanner/providers/web3"
	"go.vocdoni.io/dvote/log"
)

//go:generate go run github.com/sqlc-dev/sqlc/cmd/sqlc@v1.23.0 generate

//go:embed migrations/*.sql
var migrationsFS embed.FS

const (
	idRegistryCreationBlock  = 111816351
	keyRegistryCreationBlock = 111816359
	idRegistryAddress        = "0x00000000Fc6c5F01Fc30151999387Bb99A9f489b"
	keyRegistryAddress       = "0x00000000Fc1237824fb747aBDE0FF18990E59b7e"
	chainID                  = 10
	defaultRecoveryAddress   = "0x00000000fcb080a4d6c39a9354da9eb9bc104cd7"
)

var (
	FarcasterIDRegistryType  = []byte{0}
	FarcasterKeyRegistryType = []byte{1}

	ErrUserAlreadyExists = errors.New("user already exists")
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
	endpoints web3.NetworkEndpoints
	client    *ethclient.Client
	db        *DB

	contracts        FarcasterContracts
	lastNetworkBlock uint64
}

type FarcasterUserData struct {
	// FID is the Farcaster ID
	FID *big.Int
	// Username is the username of the user
	Username string
	// CustodyAddress is the custody address of the user
	CustodyAddress common.Address
	// RecoveryAddress is the address used to recover the user's account
	RecoveryAddress common.Address
	// Signer is the hash of the ED25519 public key of the user
	// user to sign Farcaster verifiable messages
	Signer common.Address
	// AppKeys is the array of keys of the user found in the KeyRegistry
	AppKeys [][]byte
}

func (p *FarcasterProvider) Init(iconf any) error {
	// parse the config and set the endpoints
	conf, ok := iconf.(FarcasterProviderConf)
	if !ok {
		return errors.New("invalid config type, it must be Web3ProviderConfig")
	}
	p.endpoints = conf.Endpoints
	p.db = conf.DB
	p.contracts.idRegistrySynced.Store(false)
	p.contracts.keyRegistrySynced.Store(false)

	return nil
}

// SetRef sets the reference of the token desired to use to the provider. It
// receives a Web3ProviderRef struct with the address and chainID of the token
// to use. It connects to the endpoint and initializes the contract.
func (p *FarcasterProvider) SetRef(iref any) error {
	if p.endpoints == nil {
		return errors.New("endpoints not defined")
	}
	currentEndpoint, exists := p.endpoints.EndpointByChainID(chainID)
	if !exists {
		return errors.New("endpoint not found for the given chainID")
	}
	// connect to the endpoint
	client, err := currentEndpoint.GetClient(web3.DefaultMaxWeb3ClientRetries)
	if err != nil {
		return errors.Join(web3.ErrConnectingToWeb3Client, fmt.Errorf("[FARCASTER]: %w", err))
	}
	// set the client, parse the addresses and initialize the contracts
	p.client = client
	idRegistryAddress := common.HexToAddress(idRegistryAddress)
	keyRegistryAddress := common.HexToAddress(keyRegistryAddress)
	if p.contracts.idRegistry, err = fcir.NewFarcasterIDRegistry(idRegistryAddress, client); err != nil {
		return errors.Join(web3.ErrInitializingContract, fmt.Errorf("[FARCASTER ID REGISTRY] %s: %w", idRegistryAddress, err))
	}
	if p.contracts.keyRegistry, err = fckr.NewFarcasterKeyRegistry(keyRegistryAddress, client); err != nil {
		return errors.Join(web3.ErrInitializingContract, fmt.Errorf("[FARCASTER KEY REGISTRY] %s: %w", keyRegistryAddress, err))
	}
	p.lastNetworkBlock = 0
	p.contracts.idRegistrySynced.Store(false)
	p.contracts.keyRegistrySynced.Store(false)
	return nil
}

// SetLastBalances method is not implemented for Farcaster contracts.
func (p *FarcasterProvider) SetLastBalances(_ context.Context, _ []byte,
	_ map[common.Address]*big.Int, _ uint64,
) error {
	return nil
}

// SetLastBlockNumber sets the last block number of the token set in the
// provider. It is used to calculate the delta balances in the next call to
// HoldersBalances from the given from point in time. It helps to avoid
// GetBlockNumber calls to the provider.
func (p *FarcasterProvider) SetLastBlockNumber(blockNumber uint64) {
	p.lastNetworkBlock = blockNumber
}

// HoldersBalances returns the balances of the token holders for the current
// defined token (using SetRef method). It returns the balances of the holders
// for this token from the block number provided to the latest posible block
// number (chosen between the last block number of the network and the maximun
// number of blocks to scan). It calls to rangeOfLogs to get the logs of the
// token transfers in the range of blocks and then it iterates the logs to
// calculate the balances of the holders. It returns the balances, the number
// of new transfers, the last block scanned, if the provider is synced and an
// error if it exists.
func (p *FarcasterProvider) HoldersBalances(ctx context.Context, _ []byte, fromBlock uint64) (
	map[common.Address]*big.Int, uint64, uint64, bool, *big.Int, error,
) {
	// calculate the range of blocks to scan, by default take the last block
	// scanned and scan to the latest block, calculate the latest block if the
	// current last network block is not defined
	toBlock := p.lastNetworkBlock
	if toBlock == 0 {
		var err error
		toBlock, err = p.LatestBlockNumber(ctx, nil)
		if err != nil {
			return nil, 0, fromBlock, false, big.NewInt(0), err
		}
	}
	log.Infow("scan iteration",
		"address IDRegistry", idRegistryAddress,
		"address KeyRegistry", keyRegistryAddress,
		"type", p.TypeName(),
		"from", fromBlock,
		"to", toBlock)

	// iterate scanning the logs in the range of blocks until the last blockis reached
	newRegisters, lastBlock, synced, err := p.ScanLogsIDRegistry(ctx, fromBlock, toBlock)
	if err != nil {
		return nil, 0, fromBlock, false, nil, err
	}
	// at this point we have the new registered users
	// and we use the key registry to fetch the keys of each user
	usersCensusData := make(map[common.Address]*big.Int, 0)
	usersDBData := make([]*FarcasterUserData, 0)
	for custodyAddress, fid := range newRegisters {
		userData := &FarcasterUserData{
			FID:             fid,
			CustodyAddress:  custodyAddress,
			RecoveryAddress: common.HexToAddress(defaultRecoveryAddress),
		}
		// get signer key
		signer, err := p.contracts.keyRegistry.KeysOf(nil, fid, 1)
		if err != nil {
			return nil, 0, fromBlock, false, nil, err
		}

		// NOTE THAT WE ARE ASSUMING THAT THE SIGNER IS THE FIRST KEY
		//
		// NOTE THAT WE ARE NOT USING THE SIGNER AS IT IS RETURNED FROM THE CONTRACT
		// For compatibility with the current implementation and to be able to generate a census
		// we take the signer returned by the contract and it is hashed to generate a common.Address
		userData.Signer = common.Address(crypto.Keccak256(signer[0]))
		// store the other keys
		userData.AppKeys = append(userData.AppKeys, signer[1:]...)
		usersCensusData[userData.Signer] = big.NewInt(1) // weight for the census is 1 as there is one user

		usersDBData = append(usersDBData, userData)
	}

	// update farcaster database with the users data
	if err := p.updateFarcasterDB(ctx, usersDBData); err != nil {
		return nil, 0, fromBlock, false, nil, err
	}

	// here we return the hashed signer and the farcasterID in the map
	return usersCensusData, uint64(len(newRegisters)), lastBlock, synced, nil, nil
}

// updates farcaster database with the users data
func (p *FarcasterProvider) updateFarcasterDB(ctx context.Context, usersData []*FarcasterUserData) error {
	// init db transaction
	internalCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	tx, err := p.db.RW.BeginTx(internalCtx, nil)
	if err != nil {
		return fmt.Errorf("cannot update farcaster db: %w", err)
	}
	defer func() {
		if err := tx.Rollback(); err != nil && !errors.Is(sql.ErrTxDone, err) {
			log.Errorw(err, "farcaster transaction rollback failed")
		}
	}()

	qtx := p.db.QueriesRW.WithTx(tx)
	// iterate the users and update the database
	// if the user does not exist on the database create a new one
	// if it exists update the user data
	for _, userData := range usersData {
		// check if the user exists
		_, err := qtx.GetUserByFID(internalCtx, userData.FID.Uint64())
		// check if error is that the user does not exist
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("cannot update farcaster db: %w", err)
		}
		// if the user does not exist create a new one
		if errors.Is(err, sql.ErrNoRows) {
			if _, err := qtx.CreateUser(internalCtx, queries.CreateUserParams{
				Fid:             userData.FID.Uint64(),
				Username:        userData.Username,
				CustodyAddress:  userData.CustodyAddress[:],
				RecoveryAddress: userData.RecoveryAddress[:],
				Signer:          userData.Signer[:],
			}); err != nil {
				if strings.Contains(err.Error(), "UNIQUE constraint failed") {
					return fmt.Errorf("cannot update farcaster db: %w", ErrUserAlreadyExists)
				}
				return fmt.Errorf("cannot update farcaster db: %w", err)
			}
		}
		// if the user exists update the user data
		if _, err := qtx.UpdateUser(internalCtx, queries.UpdateUserParams{
			Fid:             userData.FID.Uint64(),
			Username:        userData.Username,
			CustodyAddress:  userData.CustodyAddress[:],
			RecoveryAddress: userData.RecoveryAddress[:],
			Signer:          userData.Signer[:],
		}); err != nil {
			return fmt.Errorf("cannot update farcaster db: %w", err)
		}
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("cannot update farcaster db: %w", err)
	}
	return nil
}

// ScanLogsIDRegistry scans the logs of the Farcaster ID Registry contract
func (p *FarcasterProvider) ScanLogsIDRegistry(ctx context.Context, fromBlock, toBlock uint64) (map[common.Address]*big.Int, uint64, bool, error) {
	startTime := time.Now()
	logs, lastBlock, synced, err := web3.RangeOfLogs(
		ctx,
		p.client,
		p.Address(FarcasterIDRegistryType),
		fromBlock,
		toBlock,
		web3.LOG_TOPIC_FARCASTER_REGISTER,
	)
	if err != nil {
		return nil, 0, false, err
	}
	// encode the number of new registers
	newFIDs := make(map[common.Address]*big.Int, 0)
	// iterate the logs and update the balances
	for _, currentLog := range logs {
		logData, err := p.contracts.idRegistry.ParseRegister(currentLog)
		if err != nil {
			return newFIDs, lastBlock, false, errors.Join(web3.ErrParsingTokenLogs, fmt.Errorf("[Farcaster ID Registry]: %w", err))
		}
		newFIDs[logData.To] = logData.Id
	}

	log.Infow("saving blocks",
		"count", len(newFIDs),
		"logs", len(logs),
		"blocks/s", 1000*float32(lastBlock-fromBlock)/float32(time.Since(startTime).Milliseconds()),
		"took", time.Since(startTime).Seconds(),
		"progress", fmt.Sprintf("%d%%", (fromBlock*100)/toBlock))

	p.contracts.idRegistrySynced.Store(synced)

	return newFIDs, lastBlock, synced, nil
}

// Close method is not implemented for Farcaster Key Registry.
func (p *FarcasterProvider) Close() error {
	p.client.Close()
	return nil
}

// IsExternal returns false because the provider is not an external API.
func (p *FarcasterProvider) IsExternal() bool {
	return false
}

// IsSynced returns true if the current state of the provider is synced. It also
// receives an external ID but it is not used by the provider.
// Checks that both Farcaster contracts are synced
func (p *FarcasterProvider) IsSynced(_ []byte) bool {
	return p.contracts.idRegistrySynced.Load() && p.contracts.keyRegistrySynced.Load()
}

// Address returns the address of the Farcaster contract given the contract type.
func (p *FarcasterProvider) Address(contractType []byte) common.Address {
	if bytes.Equal(contractType, FarcasterIDRegistryType) {
		return common.HexToAddress(idRegistryAddress)
	} else if bytes.Equal(contractType, FarcasterKeyRegistryType) {
		return common.HexToAddress(keyRegistryAddress)
	}
	return common.Address{}
}

// Type returns the type of the current token set in the provider.
func (p *FarcasterProvider) Type() uint64 {
	return providers.CONTRACT_TYPE_FARCASTER
}

// TypeName returns the type name of the current token set in the provider.
func (p *FarcasterProvider) TypeName() string {
	return providers.TokenTypeName(providers.CONTRACT_TYPE_FARCASTER)
}

// ChainID returns the chain ID where the Farcaster contracts are deployed.
func (p *FarcasterProvider) ChainID() uint64 {
	return chainID
}

// Name is not implemented for Farcaster contracts.
func (p *FarcasterProvider) Name(_ []byte) (string, error) {
	return "Farcaster Key Registry", nil
}

// Symbol is not implemented for Farcaster contracts.
func (p *FarcasterProvider) Symbol(_ []byte) (string, error) {
	return "", nil
}

// Decimals is not implemented for Farcaster contracts.
func (p *FarcasterProvider) Decimals(_ []byte) (uint64, error) {
	return 0, nil
}

// TotalSupply returns the total number of Farcaster users found in the IDRegistry
// by calling idCountCounter method.
func (p *FarcasterProvider) TotalSupply(_ []byte) (*big.Int, error) {
	return p.contracts.idRegistry.IdCounter(nil)
}

// BalanceOf is not implemented for Farcaster contracts.
func (p *FarcasterProvider) BalanceOf(addr common.Address, _ []byte) (*big.Int, error) {
	return big.NewInt(1), nil
}

// BalanceAt is not implemented for Farcaster contracts.
func (p *FarcasterProvider) BalanceAt(ctx context.Context, addr common.Address,
	_ []byte, blockNumber uint64,
) (*big.Int, error) {
	return big.NewInt(0), nil
}

// BlockTimestamp returns the timestamp of the given block number for the
// current token set in the provider. It gets the timestamp from the client.
func (p *FarcasterProvider) BlockTimestamp(ctx context.Context, blockNumber uint64) (string, error) {
	blockHeader, err := p.client.HeaderByNumber(ctx, new(big.Int).SetUint64(blockNumber))
	if err != nil {
		return "", err
	}
	return time.Unix(int64(blockHeader.Time), 0).Format(web3.TimeLayout), nil
}

// BlockRootHash returns the root hash of the given block number for the current
// farcaster contracts set in the provider. It gets the root hash from the client.
func (p *FarcasterProvider) BlockRootHash(ctx context.Context, blockNumber uint64) ([]byte, error) {
	blockHeader, err := p.client.HeaderByNumber(ctx, new(big.Int).SetInt64(int64(blockNumber)))
	if err != nil {
		return nil, err
	}
	return blockHeader.Root.Bytes(), nil
}

// LatestBlockNumber returns the latest block number of the farcaster contracs set
// in the provider. It gets the latest block number from the client. It also
// receives an external ID but it is not used by the provider.
func (p *FarcasterProvider) LatestBlockNumber(ctx context.Context, _ []byte) (uint64, error) {
	lastBlockHeader, err := p.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return 0, err
	}
	return lastBlockHeader.Number.Uint64(), nil
}

// CreationBlock is not implemented for Farcaster contracts.
func (p *FarcasterProvider) CreationBlock(ctx context.Context, contractType []byte) (uint64, error) {
	if bytes.Equal(contractType, FarcasterIDRegistryType) {
		return p.idRegistryCreationBlock(), nil
	}
	if bytes.Equal(contractType, FarcasterKeyRegistryType) {
		return p.keyRegistryCreationBlock(), nil
	}
	return 0, fmt.Errorf("unknown type")
}

// IDRegistryCreationBlock returns the creation block of the Farcaster ID Registry
func (p *FarcasterProvider) idRegistryCreationBlock() uint64 {
	return idRegistryCreationBlock
}

// KeyRegistryCreationBlock returns the creation block of the Farcaster Key Registry
func (p *FarcasterProvider) keyRegistryCreationBlock() uint64 {
	return keyRegistryCreationBlock
}

// IconURI method is not implemented for Farcaster Key Registry tokens.
func (p *FarcasterProvider) IconURI(_ []byte) (string, error) {
	return "", nil
}

// Return the custody address of a given FarcasterID
func (p *FarcasterProvider) CustodyOf(fid *big.Int) (common.Address, error) {
	return p.contracts.idRegistry.CustodyOf(nil, fid)
}

// Return the ID of a given custody address
func (p *FarcasterProvider) IdOf(custody common.Address) (*big.Int, error) {
	return p.contracts.idRegistry.IdOf(nil, custody)
}

// Verifies a given FID signature
func (p *FarcasterProvider) VerifyFIDSignature(custodyAddress common.Address,
	fid *big.Int,
	digest [32]byte,
	signature []byte,
) (bool, error) {
	return p.contracts.idRegistry.VerifyFidSignature(nil, custodyAddress, fid, digest, signature)
}

// Return an array of all active keys for a given fid
func (p *FarcasterProvider) KeysOf(fid *big.Int) ([][]byte, error) {
	return p.contracts.keyRegistry.KeysOf(nil, fid, 1) // 1 is the default key state
}

// DB

// DB struct abstact a safe connection with the database using sqlc queries,
// sqlite as a database engine and go-sqlite3 as a driver.
type DB struct {
	RW *sql.DB
	RO *sql.DB

	QueriesRW *queries.Queries
	QueriesRO *queries.Queries
}

// Close function stops all internal connections to the database
func (db *DB) CloseDB() error {
	if err := db.RW.Close(); err != nil {
		return err
	}
	return db.RO.Close()
}

// Init function starts a database using the data path provided as argument. It
// opens two different connections, one for read only, and another for read and
// write, with different configurations, optimized for each use case.
func InitDB(dataDir string, dbName string) (*DB, error) {
	if dbName == "" {
		return nil, fmt.Errorf("database name is required")
	}
	dbFile := filepath.Join(dataDir, dbName)
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		if err := os.MkdirAll(dataDir, os.ModePerm); err != nil {
			return nil, fmt.Errorf("error creating a new database file: %w", err)
		}
	}
	// sqlite doesn't support multiple concurrent writers.
	// For that reason, rwDB is limited to one open connection.
	// Per https://github.com/mattn/go-sqlite3/issues/1022#issuecomment-1067353980,
	// we use WAL to allow multiple concurrent readers at the same time.
	rwDB, err := sql.Open("sqlite3", fmt.Sprintf("file:%s?mode=rwc&_journal_mode=wal&_txlock=immediate&_synchronous=normal", dbFile))
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}
	rwDB.SetMaxOpenConns(1)
	rwDB.SetMaxIdleConns(2)
	rwDB.SetConnMaxIdleTime(10 * time.Minute)
	rwDB.SetConnMaxLifetime(time.Hour)

	roDB, err := sql.Open("sqlite3", fmt.Sprintf("file:%s?mode=ro&_journal_mode=wal", dbFile))
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}
	// Increasing these numbers can allow for more queries to run concurrently,
	// but it also increases the memory used by sqlite and our connection pool.
	// Most read-only queries we run are quick enough, so a small number seems OK.
	roDB.SetMaxOpenConns(10)
	roDB.SetMaxIdleConns(20)
	roDB.SetConnMaxIdleTime(5 * time.Minute)
	roDB.SetConnMaxLifetime(time.Hour)

	// get census3 goose migrations and setup for sqlite3
	if err := goose.SetDialect("sqlite3"); err != nil {
		return nil, fmt.Errorf("error setting up driver for sqlite: %w", err)
	}
	goose.SetBaseFS(migrationsFS)
	// perform goose up
	if err := goose.Up(rwDB, "migrations"); err != nil {
		return nil, fmt.Errorf("error during goose up: %w", err)
	}
	// init sqlc
	return &DB{
		RW:        rwDB,
		RO:        roDB,
		QueriesRW: queries.New(rwDB),
		QueriesRO: queries.New(roDB),
	}, nil
}
