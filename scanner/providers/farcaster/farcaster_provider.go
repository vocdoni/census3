package farcaster

import (
	"bytes"
	"context"
	"database/sql"
	"embed"
	"encoding/gob"
	"errors"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
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
	IdRegistryAddress        = "0x00000000Fc6c5F01Fc30151999387Bb99A9f489b"
	KeyRegistryAddress       = "0x00000000Fc1237824fb747aBDE0FF18990E59b7e"
	ChainID                  = 10
	defaultRecoveryAddress   = "0x00000000fcb080a4d6c39a9354da9eb9bc104cd7"
)

var (
	FarcasterIDRegistryType  = []byte{0}
	FarcasterKeyRegistryType = []byte{1}

	ErrUserAlreadyExists = errors.New("user already exists")
	VoidAddress          = common.Address{}
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
	AppKeys []common.Hash
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
	currentEndpoint, exists := p.endpoints.EndpointByChainID(ChainID)
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
	idRegistryAddress := common.HexToAddress(IdRegistryAddress)
	keyRegistryAddress := common.HexToAddress(KeyRegistryAddress)
	if p.contracts.idRegistry, err = fcir.NewFarcasterIDRegistry(idRegistryAddress, client); err != nil {
		return errors.Join(web3.ErrInitializingContract, fmt.Errorf("[FARCASTER ID REGISTRY] %s: %w", idRegistryAddress, err))
	}
	if p.contracts.keyRegistry, err = fckr.NewFarcasterKeyRegistry(keyRegistryAddress, client); err != nil {
		return errors.Join(web3.ErrInitializingContract, fmt.Errorf("[FARCASTER KEY REGISTRY] %s: %w", keyRegistryAddress, err))
	}

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
		"address IDRegistry", IdRegistryAddress,
		"type", p.TypeName(),
		"from", fromBlock,
		"to", toBlock)

	// read logs from the IDRegistry
	// iterate scanning the logs in the range of blocks until the last block is reached
	newRegisters, lastBlock, _, errLogsIDRegistry := p.ScanLogsIDRegistry(ctx, fromBlock, toBlock)
	if errLogsIDRegistry != nil {
		return nil, 0, fromBlock, false, nil, errLogsIDRegistry
	}

	// save new users registered on the database
	// from the logs of the IDRegistry we can obtain the user FID and the custody and recovery addresses
	usersDBData := make([]*FarcasterUserData, 0)
	for fid, to := range newRegisters {
		// if the user already exists in the database skip it
		if _, err := p.db.QueriesRO.GetUserByFID(ctx, fid.Uint64()); err == nil {
			continue
		} else if !errors.Is(err, sql.ErrNoRows) {
			return nil, 0, fromBlock, false, nil, err
		}

		// create a new user data and add for saving
		userData := &FarcasterUserData{
			FID:             fid,
			CustodyAddress:  to,
			RecoveryAddress: common.HexToAddress(defaultRecoveryAddress),
		}
		usersDBData = append(usersDBData, userData)
	}

	// update the database with the new users
	log.Debugf("Updating farcaster database with %d users", len(usersDBData))
	if err := p.updateFarcasterDB(ctx, usersDBData); err != nil {
		return nil, 0, fromBlock, false, nil, err
	}

	// read the logs from the KeyRegistry
	log.Infow("scan iteration",
		"address KeyRegistry", KeyRegistryAddress,
		"type", p.TypeName(),
		"from", fromBlock,
		"to", toBlock)

	// iterate scanning the logs in the range of blocks until the last block is reached
	// note that the scanning will be done using as toBlock the last block scanned that was
	// returned by the IDRegistry scanning process
	// that way we can be sure that the KeyRegistry is synced with the IDRegistry
	newKeys, lastBlock2, synced, errLogsKeyRegistry := p.ScanLogsKeyRegistry(ctx, fromBlock, lastBlock)
	if errLogsKeyRegistry != nil {
		return nil, 0, fromBlock, false, nil, errLogsKeyRegistry
	}

	// at this point we have the new registered users and the new registered app keys

	// get existing users from the database
	fidsFromDB, err := p.db.QueriesRO.ListUsers(ctx)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, 0, fromBlock, false, nil, err
	}

	usersCensusData := make(map[common.Address]*big.Int)
	if len(fidsFromDB) != 0 {
		usersDBDataPost := make([]*FarcasterUserData, 0)
		// iterate the users and update the database with the new app keys
		for _, fid := range fidsFromDB {
			// check if the user has new keys to add
			if keys, ok := newKeys[fid.Fid]; !ok {
				continue
			} else {
				userData := &FarcasterUserData{
					FID:             big.NewInt(int64(fid.Fid)),
					CustodyAddress:  common.BytesToAddress(fid.CustodyAddress),
					RecoveryAddress: common.BytesToAddress(fid.RecoveryAddress),
					Signer:          common.BytesToAddress(fid.Signer),
				}
				// if not app keys are found, set the signer to the first key
				alteredSignerKey := common.Address{}
				if len(fid.AppKeys) == 0 {
					// no need to hash as the value is already hashed because it is an indexed value in the event
					alteredSignerKey = common.BytesToAddress(keys[0].Key[:])
					userData.Signer = alteredSignerKey
				}
				// create key list
				k := make([]common.Hash, 0)
				for _, key := range keys {
					h := common.Hash{}
					h.SetBytes(key.KeyBytes)
					k = append(k, h)
				}
				// deserialize app keys since they are stored as bytes
				if len(fid.AppKeys) != 0 {
					appKeys, err := deserializeAppKeys(fid.AppKeys)
					if err != nil {
						return nil, 0, fromBlock, false, nil, err
					}
					userData.AppKeys = appKeys
				}
				// append the new keys to the user data
				userData.AppKeys = append(userData.AppKeys, k...)

				// append modified user data to the list for saving
				usersDBDataPost = append(usersDBDataPost, userData)
				// add the user to the map for storing on the scanner database
				usersCensusData[alteredSignerKey] = big.NewInt(1)
			}
		}
		// update the farcaster database with the new data
		log.Debugf("Updating farcaster database with %d users after key registry scan", len(usersDBDataPost))
		if err := p.updateFarcasterDB(ctx, usersDBDataPost); err != nil {
			return nil, 0, fromBlock, false, nil, err
		}
	}
	// NOTE: we are assuming that the key registry is synced if the id registry is synced
	// Return the smallest block for starting next iteration from there
	blockToReturn := 0
	if lastBlock >= lastBlock2 {
		blockToReturn = int(lastBlock2)
	} else {
		blockToReturn = int(lastBlock)
	}
	totalSupply, err := p.TotalSupply(nil)
	// if error getting total supply, get old supply from database
	if err != nil {
		log.Warnf("Error getting total supply: %s", err.Error())
		ts, err := p.db.QueriesRO.CountUsers(ctx)
		if err != nil {
			return nil, 0, fromBlock, false, nil, err
		}
		totalSupply = big.NewInt(int64(ts))
	}
	// usersCensusData is a map of signers and their balances set to 1 to indicate that the user exists
	// signer is the address calculated as keccak256(fid.AppKeys[0])
	return usersCensusData, uint64(len(usersCensusData)), uint64(blockToReturn), synced, totalSupply, nil
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
		_, errGetUser := qtx.GetUserByFID(internalCtx, userData.FID.Uint64())
		if errGetUser != nil {
			// if not exists create a new user
			if errors.Is(errGetUser, sql.ErrNoRows) {
				if _, err := qtx.CreateUser(internalCtx, queries.CreateUserParams{
					Fid: userData.FID.Uint64(),
					// Username:        userData.Username,
					CustodyAddress:  userData.CustodyAddress[:],
					RecoveryAddress: userData.RecoveryAddress[:],
					Signer:          make([]byte, 0),
					AppKeys:         make([]byte, 0),
				}); err != nil {
					if strings.Contains(err.Error(), "UNIQUE constraint failed") {
						return fmt.Errorf("cannot update farcaster db: %w", ErrUserAlreadyExists)
					}
					return fmt.Errorf("cannot update farcaster db: %w", err)
				}
			} else {
				return fmt.Errorf("cannot update farcaster db: %w", errGetUser)
			}
		} else { // if user exists update the user data
			// serialize app keys before saving
			serializedAppKeys := make([]byte, 0)
			if len(userData.AppKeys) != 0 {
				serializedAppKeys, err = serializeAppKeys(userData.AppKeys)
				if err != nil {
					return fmt.Errorf("cannot update farcaster db: %w", err)
				}
			}
			if _, err := qtx.UpdateUser(internalCtx, queries.UpdateUserParams{
				Fid: userData.FID.Uint64(),
				// Username:        userData.Username,
				CustodyAddress:  userData.CustodyAddress[:],
				RecoveryAddress: userData.RecoveryAddress[:],
				Signer:          userData.Signer[:],
				AppKeys:         serializedAppKeys,
			}); err != nil {
				return fmt.Errorf("cannot update farcaster db: %w", err)
			}
		}
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("cannot update farcaster db: %w", err)
	}
	return nil
}

// ScanLogsIDRegistry scans the logs of the Farcaster ID Registry contract
func (p *FarcasterProvider) ScanLogsIDRegistry(ctx context.Context, fromBlock, toBlock uint64) (
	map[*big.Int]common.Address, uint64, bool, error,
) {
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
	newFIDs := make(map[*big.Int]common.Address, 0)
	// iterate the logs and update the balances
	for _, currentLog := range logs {
		logData, err := p.contracts.idRegistry.ParseRegister(currentLog)
		if err != nil {
			return newFIDs, lastBlock, false, errors.Join(web3.ErrParsingTokenLogs, fmt.Errorf("[Farcaster ID Registry]: %w", err))
		}
		newFIDs[logData.Id] = logData.To
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

type KRLogData struct {
	Key      common.Hash
	KeyBytes []byte
}

// ScanLogsKeyRegistry scans the logs of the Farcaster Key Registry contract
func (p *FarcasterProvider) ScanLogsKeyRegistry(ctx context.Context, fromBlock, toBlock uint64) (
	map[uint64][]KRLogData, uint64, bool, error,
) {
	startTime := time.Now()
	logs, lastBlock, synced, err := web3.RangeOfLogs(
		ctx,
		p.client,
		p.Address(FarcasterKeyRegistryType),
		fromBlock,
		toBlock,
		web3.LOG_TOPIC_FARCASTER_ADDKEY,
	)
	if err != nil {
		return nil, 0, false, err
	}
	// encode the number of new registers
	newKeys := make(map[uint64][]KRLogData, 0)
	// iterate the logs and update the balances
	for _, currentLog := range logs {
		logData, err := p.contracts.keyRegistry.ParseAdd(currentLog)
		if err != nil {
			return newKeys, lastBlock, false, errors.Join(web3.ErrParsingTokenLogs, fmt.Errorf("[Farcaster Key Registry]: %w", err))
		}
		nld := KRLogData{
			Key:      logData.Key,
			KeyBytes: logData.KeyBytes[:],
		}
		// note that logData.Key is the Keccak256 of logData.KeyBytes because logData.Key is an indexed EVM event value
		newKeys[logData.Fid.Uint64()] = append(newKeys[logData.Fid.Uint64()], nld)
	}

	log.Infow("saving blocks",
		"count", len(newKeys),
		"logs", len(logs),
		"blocks/s", 1000*float32(lastBlock-fromBlock)/float32(time.Since(startTime).Milliseconds()),
		"took", time.Since(startTime).Seconds(),
		"progress", fmt.Sprintf("%d%%", (fromBlock*100)/toBlock))

	p.contracts.keyRegistrySynced.Store(synced)

	return newKeys, lastBlock, synced, nil
}

// Close method is not implemented for Farcaster Key Registry.
func (p *FarcasterProvider) Close() error {
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
		return common.HexToAddress(IdRegistryAddress)
	} else if bytes.Equal(contractType, FarcasterKeyRegistryType) {
		return common.HexToAddress(KeyRegistryAddress)
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
	return ChainID
}

// Name returns a predefined name for convenience.
func (p *FarcasterProvider) Name(_ []byte) (string, error) {
	return providers.CONTRACT_NAME_FARCASTER, nil
}

// Symbol is not implemented for Farcaster contracts.
func (p *FarcasterProvider) Symbol(_ []byte) (string, error) {
	return "FC", nil
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

// LatestBlockNumber returns the latest block number of the farcaster contracts set
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

// CensusKeys method returns the holders and balances provided transformed. The
// Farcaster resolve the FID of the provided addresses, grouping them by FID and
// returning the balances of the FID.
func (p *FarcasterProvider) CensusKeys(data map[common.Address]*big.Int) (map[common.Address]*big.Int, error) {
	// TODO: modify to the provider implementation
	return data, nil
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

// Encode/Decode app_keys
func serializeAppKeys(data []common.Hash) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func deserializeAppKeys(serializedData []byte) ([]common.Hash, error) {
	var data []common.Hash
	buf := bytes.NewBuffer(serializedData)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

/*
// given a FID returns the user data stored in the farcaster database
func (p *FarcasterProvider) getUserByFID(ctx context.Context, fid uint64) *FarcasterUserData {
	user, err := p.db.QueriesRO.GetUserByFID(ctx, fid)
	// if error is sql.ErrNoRows, the user does not exist print log.Warn
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Fatalf("Cannot read from DB %w", err)
		return nil
	}
	if errors.Is(err, sql.ErrNoRows) {
		log.Warnf("User with FID %d not found", fid)
		return nil
	}
	fud := &FarcasterUserData{
		FID:             big.NewInt(int64(user.Fid)),
		Username:        user.Username,
		CustodyAddress:  common.BytesToAddress(user.CustodyAddress),
		RecoveryAddress: common.BytesToAddress(user.RecoveryAddress),
		Signer:          common.BytesToAddress(user.Signer),
	}
	// deserialize app keys since they are stored as bytes
	if len(user.AppKeys) != 0 {
		appKeys, err := deserializeAppKeys(user.AppKeys)
		if err != nil {
			log.Fatalf("Cannot deserialize app keys %w", err)
			return nil
		}
		fud.AppKeys = appKeys
	}
	return fud
}
*/
