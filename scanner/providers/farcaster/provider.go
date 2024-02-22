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

	// timeouts
	TooManyReqeustsTimeout  = 30 * time.Second
	IterationCooldown       = 15 * time.Second
	IterationSyncedCooldown = 60 * time.Second
)

func (p *FarcasterProvider) Init(iconf any) error {
	// parse the config and set the endpoints
	conf, ok := iconf.(FarcasterProviderConf)
	if !ok {
		return errors.New("invalid config type, it must be Web3ProviderConfig")
	}
	if conf.Endpoints == nil {
		return errors.New("endpoints not defined")
	}
	p.endpoints = conf.Endpoints
	p.db = conf.DB
	// set vars to sync the contracts in background
	p.lastNetworkBlock.Store(0)
	p.currentScannerHolders = make(map[common.Address]*big.Int)
	// set the contracts vars to nil
	p.contracts.idRegistrySynced.Store(false)
	p.contracts.keyRegistrySynced.Store(false)

	// check if latests blocks are stored in the database, if they exist, set
	// them in the provider, if not, set them to 0 in both places, the provider
	// and the database. By default, the last block is the creation block of the
	// key registry, because in the gap between the creation of the ID and Key
	// registries, there are no logs to scan.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	idRegistryLastBlock, err := p.db.QueriesRO.LastBlock(ctx, IdRegistryAddress)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("cannot get last block from farcaster DB %w", err)
		}
		// create the row in the database
		if _, err := p.db.QueriesRW.InsertLatestBlock(ctx, queries.InsertLatestBlockParams{
			Contract:    IdRegistryAddress,
			BlockNumber: keyRegistryCreationBlock,
		}); err != nil {
			return fmt.Errorf("cannot create last block in farcaster DB %w", err)
		}
		idRegistryLastBlock = keyRegistryCreationBlock
	}
	keyRegistryLastBlock, err := p.db.QueriesRO.LastBlock(ctx, KeyRegistryAddress)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("cannot get last block from farcaster DB %w", err)
		}
		// create the row in the database
		if _, err := p.db.QueriesRW.InsertLatestBlock(ctx, queries.InsertLatestBlockParams{
			Contract:    KeyRegistryAddress,
			BlockNumber: keyRegistryCreationBlock,
		}); err != nil {
			return fmt.Errorf("cannot create last block in farcaster DB %w", err)
		}
		keyRegistryLastBlock = keyRegistryCreationBlock
	}
	// choose as last block the smallest between the id and key registry last blocks
	lastBlock := idRegistryLastBlock
	if keyRegistryLastBlock < idRegistryLastBlock {
		lastBlock = keyRegistryLastBlock
	}
	p.contracts.lastBlock.Store(uint64(lastBlock))
	// init the web3 client and contracts
	currentEndpoint, exists := p.endpoints.EndpointByChainID(ChainID)
	if !exists {
		return errors.New("endpoint not found for the given chainID")
	}
	// connect to the endpoint and set the client
	p.client, err = currentEndpoint.GetClient(web3.DefaultMaxWeb3ClientRetries)
	if err != nil {
		return errors.Join(web3.ErrConnectingToWeb3Client, fmt.Errorf("[FARCASTER]: %w", err))
	}
	// parse the addresses and initialize the contracts
	idRegistryAddress := common.HexToAddress(IdRegistryAddress)
	keyRegistryAddress := common.HexToAddress(KeyRegistryAddress)
	if p.contracts.idRegistry, err = fcir.NewFarcasterIDRegistry(idRegistryAddress, p.client); err != nil {
		return errors.Join(web3.ErrInitializingContract, fmt.Errorf("[FARCASTER ID REGISTRY] %s: %w", idRegistryAddress, err))
	}
	if p.contracts.keyRegistry, err = fckr.NewFarcasterKeyRegistry(keyRegistryAddress, p.client); err != nil {
		return errors.Join(web3.ErrInitializingContract, fmt.Errorf("[FARCASTER KEY REGISTRY] %s: %w", keyRegistryAddress, err))
	}
	p.contracts.idRegistrySynced.Store(false)
	p.contracts.keyRegistrySynced.Store(false)
	// start the internal scanner
	p.scannerCtx, p.cancelScanner = context.WithCancel(context.Background())
	go p.initInternalScanner()
	return nil
}

// SetRef sets the reference of the token desired to use to the provider. It
// receives a Web3ProviderRef struct with the address and chainID of the token
// to use. It connects to the endpoint and initializes the contract.
func (p *FarcasterProvider) SetRef(_ any) error {
	return nil
}

// SetLastBalances method is not implemented for Farcaster contracts.
func (p *FarcasterProvider) SetLastBalances(_ context.Context, _ []byte,
	currentScannerBalances map[common.Address]*big.Int, _ uint64,
) error {
	p.currentScannerHoldersMtx.Lock()
	defer p.currentScannerHoldersMtx.Unlock()
	p.currentScannerHolders = make(map[common.Address]*big.Int)
	for k, v := range currentScannerBalances {
		p.currentScannerHolders[k] = new(big.Int).Set(v)
	}
	return nil
}

// SetLastBlockNumber sets the last block number of the token set in the
// provider. It is used to calculate the delta balances in the next call to
// HoldersBalances from the given from point in time. It helps to avoid
// GetBlockNumber calls to the provider.
func (p *FarcasterProvider) SetLastBlockNumber(blockNumber uint64) {
	p.lastNetworkBlock.Store(blockNumber)
}

// HoldersBalances method ignores every param provided unless the context. It
// returns the difference between the current holders and the last holders
// scanned by the provider. It also returns the number of logs scanned (same to
// the number of new holders), the last block scanned, if the provider is synced
// and the total supply of the token. It gets the current holders from the
// internal database and the current holders from the scanner and calculates the
// partial holders.
func (p *FarcasterProvider) HoldersBalances(ctx context.Context, _ []byte, fromBlock uint64) (
	map[common.Address]*big.Int, uint64, uint64, bool, *big.Int, error,
) {
	// check if both contracts are synced
	isSynced := p.contracts.idRegistrySynced.Load() && p.contracts.keyRegistrySynced.Load()
	// get current holders from internal db
	appKeys, err := p.db.QueriesRO.ListAppKeys(ctx)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, 0, fromBlock, isSynced, nil, fmt.Errorf("cannot get app keys from farcaster DB %s", err.Error())
	}
	currentHolders := make(map[common.Address]*big.Int)
	for _, appKey := range appKeys {
		currentHolders[common.BytesToAddress(appKey)] = big.NewInt(1)
	}
	totalSupply := big.NewInt(int64(len(currentHolders)))
	// get the current holders from the scanner
	p.currentScannerHoldersMtx.Lock()
	currentScannerHolders := make(map[common.Address]*big.Int)
	for k, v := range p.currentScannerHolders {
		currentScannerHolders[k] = new(big.Int).Set(v)
	}
	p.currentScannerHoldersMtx.Unlock()
	resultingHolders := providers.CalcPartialHolders(currentScannerHolders, currentHolders)
	return resultingHolders, uint64(len(resultingHolders)), p.contracts.lastBlock.Load(), isSynced, totalSupply, nil
}

// Close method is not implemented for Farcaster Key Registry.
func (p *FarcasterProvider) Close() error {
	p.cancelScanner()
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

// IDRegistryCreationBlock returns the creation block of the Farcaster ID Registry
func (p *FarcasterProvider) idRegistryCreationBlock() uint64 {
	return idRegistryCreationBlock
}

// KeyRegistryCreationBlock returns the creation block of the Farcaster Key Registry
func (p *FarcasterProvider) keyRegistryCreationBlock() uint64 {
	return keyRegistryCreationBlock
}

// initInternalScanner method initializes the internal scanner of the provider.
// In each iteration, it launches a new scanIteration and sleeps for a while
// before the next iteration. It also updates the last block scanned and the
// sync vars of the provider.
func (p *FarcasterProvider) initInternalScanner() {
	for {
		select {
		case <-p.scannerCtx.Done():
			return
		default:
			lastBlock, idrSynced, krSynced, err := p.scanIteration(p.scannerCtx)
			if err != nil {
				log.Errorf("error scanning iteration: %s", err.Error())
				continue
			}
			p.contracts.lastBlock.Store(lastBlock)
			log.Debugw("scanning iteration finished, sleeping...",
				"lastBlock", lastBlock,
				"idRegistrySynced", idrSynced,
				"keyRegistrySynced", krSynced)
			// if the IDRegistry and KeyRegistry are synced, sleep for 10 seconds
			// before the next iteration
			if idrSynced && krSynced {
				time.Sleep(IterationSyncedCooldown)
			} else {
				time.Sleep(IterationCooldown)
			}
		}
	}
}

// scanIteration method scans the logs of the Farcaster ID and Key Registries
// contracts. It returns the last block scanned, if the ID Registry is synced,
// if the Key Registry is synced and an error if it exists. It calls to
// scanLogsIDRegistry and scanLogsKeyRegistry to get the logs of the contracts
// and then it stores the new registered users and the added and removed app
// keys in the database. It also updates the sync vars of the provider.
func (p *FarcasterProvider) scanIteration(ctx context.Context) (uint64, bool, bool, error) {
	isIDRegistrySynced := p.contracts.idRegistrySynced.Load()
	isKeyRegistrySynced := p.contracts.keyRegistrySynced.Load()
	// calculate the range of blocks to scan, by default take the last block
	// scanned and scan to the latest block, calculate the latest block if the
	// current last network block is not defined
	fromBlock := p.contracts.lastBlock.Load()
	toBlock := p.lastNetworkBlock.Load()
	if toBlock == 0 {
		var err error
		toBlock, err = p.LatestBlockNumber(ctx, nil)
		if err != nil {
			return fromBlock, isIDRegistrySynced, isKeyRegistrySynced,
				fmt.Errorf("cannot get latest block number %s", err.Error())
		}
	}
	log.Infow("scan iteration",
		"address IDRegistry", IdRegistryAddress,
		"type", p.TypeName(),
		"from", fromBlock,
		"to", toBlock)
	// read logs from the IDRegistry
	// iterate scanning the logs in the range of blocks until the last block is reached
	newRegisters, idrLastBlock, idrSynced, err := p.scanLogsIDRegistry(ctx, fromBlock, toBlock)
	if err != nil {
		if !errors.Is(err, web3.ErrTooManyRequests) {
			return fromBlock, isIDRegistrySynced, isKeyRegistrySynced,
				fmt.Errorf("cannot scan logs from IDRegistry %s", err.Error())
		}
		// if the error is about too many requests, sleep 30 seconds and
		// return an error to retry
		log.Debug("too many requests, sleeping for a while...")
		time.Sleep(TooManyReqeustsTimeout)
	}

	// save new users registered on the database
	// from the logs of the IDRegistry we can obtain the user FID and the custody and recovery addresses
	if err := p.storeNewRegisteredUsers(ctx, newRegisters, fromBlock); err != nil {
		return fromBlock, isIDRegistrySynced, isKeyRegistrySynced,
			fmt.Errorf("cannot store new registered users into farcaster DB %s", err.Error())
	}
	// read the logs from the KeyRegistry
	log.Infow("scan iteration",
		"address KeyRegistry", KeyRegistryAddress,
		"type", p.TypeName(),
		"from", fromBlock,
		"to", toBlock)
	// iterate scanning the logs in the range of blocks until the last block is
	// reached note that the scanning will be done using as toBlock the last
	// block scanned that was returned by the IDRegistry scanning process, that
	// way we can be sure that the KeyRegistry is synced with the IDRegistry
	addedKeys, removedKeys, krLastBlock, krSynced, err := p.scanLogsKeyRegistry(ctx, fromBlock, idrLastBlock)
	if err != nil {
		if !errors.Is(err, web3.ErrTooManyRequests) {
			return fromBlock, isIDRegistrySynced, isKeyRegistrySynced,
				fmt.Errorf("cannot scan logs from KeyRegistry %s", err.Error())
		}
		// if the error is about too many requests, sleep 30 seconds and
		// return an error to retry
		log.Debug("too many requests, sleeping for a while...")
		time.Sleep(TooManyReqeustsTimeout)
	}
	// at this point we have the new registered users, the added app keys and
	// the removed ones

	// get existing users from the database
	fidList, err := p.db.QueriesRO.ListUsers(ctx)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fromBlock, isIDRegistrySynced, isKeyRegistrySynced,
			fmt.Errorf("cannot get users from farcaster DB %s", err.Error())
	}
	// add app keys to the database and get the added ones
	if err := p.addAppKeys(ctx, fidList, addedKeys); err != nil {
		return fromBlock, isIDRegistrySynced, isKeyRegistrySynced,
			fmt.Errorf("cannot store new app keys %s", err.Error())
	}
	// remove app keys from the database on get the removed ones
	if err := p.deleteAppKeys(ctx, fidList, removedKeys); err != nil {
		return fromBlock, isIDRegistrySynced, isKeyRegistrySynced,
			fmt.Errorf("cannot delete app keys %s", err.Error())
	}
	// update sync vars and the last block scanned in the provider database
	p.contracts.idRegistrySynced.Store(idrSynced)
	p.contracts.keyRegistrySynced.Store(krSynced)
	if _, err := p.db.QueriesRW.SetLastBlock(ctx, queries.SetLastBlockParams{
		Contract:    IdRegistryAddress,
		BlockNumber: int64(idrLastBlock),
	}); err != nil {
		return fromBlock, isIDRegistrySynced, isKeyRegistrySynced,
			fmt.Errorf("cannot update last block in farcaster DB %s", err.Error())
	}
	if _, err := p.db.QueriesRW.SetLastBlock(ctx, queries.SetLastBlockParams{
		Contract:    KeyRegistryAddress,
		BlockNumber: int64(krLastBlock),
	}); err != nil {
		return fromBlock, isIDRegistrySynced, isKeyRegistrySynced,
			fmt.Errorf("cannot update last block in farcaster DB %s", err.Error())
	}
	// update the last block scanned with the smallest between the last block
	// of the IDRegistry and the KeyRegistry
	lastBlock := idrLastBlock
	if krLastBlock < idrLastBlock {
		lastBlock = krLastBlock
	}
	return lastBlock, idrSynced, krSynced, nil
}

// scanLogsIDRegistry scans the logs of the Farcaster ID Registry contract
func (p *FarcasterProvider) scanLogsIDRegistry(ctx context.Context, fromBlock, toBlock uint64) (
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
	if err != nil && !errors.Is(err, web3.ErrTooManyRequests) {
		return nil, fromBlock, false, err
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
	return newFIDs, lastBlock, synced, nil
}

// scanLogsKeyRegistry scans the logs of the Farcaster Key Registry contract
func (p *FarcasterProvider) scanLogsKeyRegistry(ctx context.Context, fromBlock, toBlock uint64) (
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
	if err != nil && !errors.Is(err, web3.ErrTooManyRequests) {
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
