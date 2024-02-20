package farcaster

import (
	"bytes"
	"context"
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
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
	if err := p.storeNewRegisteredUsers(ctx, newRegisters, fromBlock); err != nil {
		return nil, 0, fromBlock, false, nil, fmt.Errorf("cannot store new registered users into farcaster DB %w", err)
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
	addedKeys, removedKeys, lastBlock2, synced, errLogsKeyRegistry := p.ScanLogsKeyRegistry(ctx, fromBlock, lastBlock)
	if errLogsKeyRegistry != nil {
		return nil, 0, fromBlock, false, nil, errLogsKeyRegistry
	}

	// at this point we have the new registered users, the added app keys and the removed ones

	// get existing users from the database
	fidList, err := p.db.QueriesRO.ListUsers(ctx)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, 0, fromBlock, false, nil, err
	}

	// add app keys to the database and get the added ones
	ak, err := p.addAppKeys(ctx, fidList, addedKeys)
	if err != nil {
		return nil, 0, fromBlock, false, nil, fmt.Errorf("cannot store new app keys %w", err)
	}

	// remove app keys from the database on get the removed ones
	rk, err := p.deleteAppKeys(ctx, fidList, removedKeys)
	if err != nil {
		return nil, 0, fromBlock, false, nil, fmt.Errorf("cannot delete app keys %w", err)
	}

	// filter app keys that will be included in the scanner database
	scannerKeys := make(map[common.Address]*big.Int)
	avoidKeys := make(map[common.Address]bool)
	for _, removedAppkey := range rk {
		for _, kr := range removedAppkey {
			keyAddress := common.BytesToAddress(kr[:])
			scannerKeys[keyAddress] = big.NewInt(-1)
			avoidKeys[keyAddress] = true
		}
	}
	for _, addedAppkey := range ak {
		for _, kr := range addedAppkey {
			keyAddress := common.BytesToAddress(kr[:])
			if _, ok := avoidKeys[keyAddress]; !ok {
				scannerKeys[keyAddress] = big.NewInt(1)
			}
		}
	}

	log.Debugf("returning %d keys from HolderBalances", len(scannerKeys))

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
	// scannerKeys is a map of keys and their balances set to 1 to indicate that the user app key exists
	return scannerKeys, uint64(len(scannerKeys)), uint64(blockToReturn), synced, totalSupply, nil
}

// ScanLogsIDRegistry scans the logs of the Farcaster ID Registry contract
func (p *FarcasterProvider) ScanLogsIDRegistry(ctx context.Context, fromBlock, toBlock uint64) (
	map[uint64]common.Address, uint64, bool, error,
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
	newFIDs := make(map[uint64]common.Address, 0)
	// iterate the logs and update the balances
	for _, currentLog := range logs {
		logData, err := p.contracts.idRegistry.ParseRegister(currentLog)
		if err != nil {
			return nil, 0, false, errors.Join(web3.ErrParsingTokenLogs, fmt.Errorf("[Farcaster ID Registry]: %w", err))
		}
		newFIDs[logData.Id.Uint64()] = logData.To
	}

	log.Infow("saving blocks",
		"count", len(newFIDs),
		"logs", len(logs),
		"blocks/s", 1000*float32(lastBlock-fromBlock)/float32(time.Since(startTime).Milliseconds()),
		"took", time.Since(startTime).Seconds(),
		"progress", fmt.Sprintf("%d%%", (fromBlock*100)/toBlock))

	p.contracts.idRegistrySynced.Store(synced)

	log.Debugf("found %d users registered in the id registry contract", len(newFIDs))

	return newFIDs, lastBlock, synced, nil
}

// ScanLogsKeyRegistry scans the logs of the Farcaster Key Registry contract
func (p *FarcasterProvider) ScanLogsKeyRegistry(ctx context.Context, fromBlock, toBlock uint64) (
	map[uint64][][]byte, map[uint64][][]byte, uint64, bool, error,
) {
	startTime := time.Now()
	logs, lastBlock, synced, err := web3.RangeOfLogs(
		ctx,
		p.client,
		p.Address(FarcasterKeyRegistryType),
		fromBlock,
		toBlock,
		web3.LOG_TOPIC_FARCASTER_ADDKEY,
		web3.LOG_TOPIC_FARCASTER_REMOVEKEY,
	)
	if err != nil {
		return nil, nil, 0, false, err
	}
	addedKeys := make(map[uint64][][]byte, 0)
	removedKeys := make(map[uint64][][]byte, 0)
	for _, currentLog := range logs {
		switch currentLog.Topics[0].Hex()[2:] {
		case web3.LOG_TOPIC_FARCASTER_ADDKEY:
			logData, err := p.contracts.keyRegistry.ParseAdd(currentLog)
			if err != nil {
				return nil, nil, 0, false, errors.Join(web3.ErrParsingTokenLogs, fmt.Errorf("[Farcaster Key Registry]: %w", err))
			}
			// note that logData.Key is the Keccak256 of logData.KeyBytes because logData.Key is an indexed EVM event value
			fid := logData.Fid.Uint64()
			addedKeys[fid] = append(addedKeys[fid], logData.Key[:])
		case web3.LOG_TOPIC_FARCASTER_REMOVEKEY:
			logData, err := p.contracts.keyRegistry.ParseRemove(currentLog)
			if err != nil {
				return nil, nil, 0, false, errors.Join(web3.ErrParsingTokenLogs, fmt.Errorf("[Farcaster Key Registry]: %w", err))
			}
			fid := logData.Fid.Uint64()
			removedKeys[fid] = append(removedKeys[fid], logData.Key[:])
		default:
			return nil, nil, 0, false, fmt.Errorf("unknown log topic")
		}
	}
	log.Infow("saving blocks",
		"count", len(addedKeys)+len(removedKeys),
		"logs", len(logs),
		"blocks/s", 1000*float32(lastBlock-fromBlock)/float32(time.Since(startTime).Milliseconds()),
		"took", time.Since(startTime).Seconds(),
		"progress", fmt.Sprintf("%d%%", (fromBlock*100)/toBlock),
	)
	p.contracts.keyRegistrySynced.Store(synced)
	log.Debugf("found %d keys added for users registered in the key registry contract", len(addedKeys))
	log.Debugf("found %d keys removed for users registered in the key registry contract", len(removedKeys))
	return addedKeys, removedKeys, lastBlock, synced, nil
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
	// if the contract type is defined and it is the Key Registry, return the
	// Key Registry address
	if bytes.Equal(contractType, FarcasterKeyRegistryType) {
		return common.HexToAddress(KeyRegistryAddress)
	}
	// by default return the ID Registry address
	return common.HexToAddress(IdRegistryAddress)
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
	// if the contract type is defined and it is the ID Registry, return the
	// creation block of the ID Registry
	if bytes.Equal(contractType, FarcasterIDRegistryType) {
		return p.idRegistryCreationBlock(), nil
	}
	// if the contract type is defined and it is the Key Registry, return the
	// creation block of the Key Registry
	if bytes.Equal(contractType, FarcasterKeyRegistryType) {
		return p.keyRegistryCreationBlock(), nil
	}
	// by default return the creation block of the key registry
	return p.keyRegistryCreationBlock(), nil
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

// CensusKeys method returns the holders and balances provided transformed. The
// Farcaster resolve the FID of the provided addresses, grouping them by FID and
// returning the balances of the FID.
func (p *FarcasterProvider) CensusKeys(data map[common.Address]*big.Int) (map[common.Address]*big.Int, error) {
	return nil, nil
}

func (p *FarcasterProvider) storeNewRegisteredUsers(
	ctx context.Context, newRegisters map[uint64]common.Address, fromBlock uint64,
) error {
	usersDBData := make([]FarcasterUserData, 0)
	for fid := range newRegisters {
		_, err := p.db.QueriesRO.GetUserByFID(ctx, fid)
		if err == nil { // if the user already exists in the database skip it
			continue
		}
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("cannot get user by fid %w", err)
		}
		userData := FarcasterUserData{
			FID: fid,
		}
		usersDBData = append(usersDBData, userData)
	}
	// update the database with the new users
	log.Debugf("Updating farcaster database with %d users after id registry scan", len(usersDBData))
	if err := p.updateFarcasterDB(ctx, usersDBData); err != nil {
		return fmt.Errorf("cannot update farcaster DB %w", err)
	}
	return nil
}

func (p *FarcasterProvider) addAppKeys(
	ctx context.Context, fidList []uint64, addedKeys map[uint64][][]byte,
) (map[uint64][][]byte, error) {
	if len(fidList) == 0 {
		return nil, nil
	}
	ak := make(map[uint64][][]byte, 0)
	for _, fid := range fidList {
		keys, ok := addedKeys[fid]
		if !ok {
			continue
		}
		for _, key := range keys {
			exists, err := p.db.QueriesRO.CheckFidAppKeyExists(ctx, queries.CheckFidAppKeyExistsParams{
				Fid:    fid,
				AppKey: key[:],
			})
			if err != nil && !errors.Is(err, sql.ErrNoRows) {
				return nil, fmt.Errorf("cannot get fid app keys %w", err)
			} else if exists {
				continue
			}
			h := common.Hash{}
			h.SetBytes(key)
			// create ref for each fid and key on fid_appkeys table
			if err := p.createFidAppKey(ctx, fid, h); err != nil {
				return nil, fmt.Errorf("cannot create fid app key %w", err)
			}
			ak[fid] = append(ak[fid], key)
		}
	}
	return ak, nil
}

func (p *FarcasterProvider) deleteAppKeys(
	ctx context.Context, fidList []uint64, deletedKeys map[uint64][][]byte,
) (map[uint64][][]byte, error) {
	if len(fidList) == 0 {
		return nil, nil
	}
	dk := make(map[uint64][][]byte, 0)
	for _, fid := range fidList {
		keys, ok := deletedKeys[fid]
		if !ok {
			continue
		}
		for _, key := range keys {
			exists, err := p.db.QueriesRO.CheckFidAppKeyExists(ctx, queries.CheckFidAppKeyExistsParams{
				Fid:    fid,
				AppKey: key[:],
			})
			if err != nil && !errors.Is(err, sql.ErrNoRows) {
				return nil, fmt.Errorf("cannot get fid app keys %w", err)
			} else if exists {
				h := common.Hash{}
				h.SetBytes(key)
				if err := p.deleteFidAppKey(ctx, fid, h); err != nil {
					return nil, fmt.Errorf("cannot delete fid app key %w", err)
				}
				dk[fid] = append(dk[fid], key)
			}
		}
	}
	return dk, nil
}
