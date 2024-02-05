package web3

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	fckr "github.com/vocdoni/census3/contracts/farcaster/idRegistry"
	"github.com/vocdoni/census3/scanner/providers"
	"go.vocdoni.io/dvote/log"
)

type FarcasterIDProvider struct {
	endpoints NetworkEndpoints
	client    *ethclient.Client

	contract         *fckr.FarcasterIDRegistry
	address          common.Address
	chainID          uint64
	name             string
	symbol           string
	decimals         uint64
	totalSupply      *big.Int
	creationBlock    uint64
	lastNetworkBlock uint64
	synced           atomic.Bool
}

func (p *FarcasterIDProvider) Init(iconf any) error {
	// parse the config and set the endpoints
	conf, ok := iconf.(Web3ProviderConfig)
	if !ok {
		return errors.New("invalid config type, it must be Web3ProviderConfig")
	}
	p.endpoints = conf.Endpoints
	p.synced.Store(false)
	// set the reference if the address and chainID are defined in the config
	if conf.HexAddress != "" && conf.ChainID > 0 {
		return p.SetRef(Web3ProviderRef{
			HexAddress: conf.HexAddress,
			ChainID:    conf.ChainID,
		})
	}
	return nil
}

// SetRef sets the reference of the token desired to use to the provider. It
// receives a Web3ProviderRef struct with the address and chainID of the token
// to use. It connects to the endpoint and initializes the contract.
func (p *FarcasterIDProvider) SetRef(iref any) error {
	if p.endpoints == nil {
		return errors.New("endpoints not defined")
	}
	ref, ok := iref.(Web3ProviderRef)
	if !ok {
		return errors.New("invalid ref type, it must be Web3ProviderRef")
	}
	currentEndpoint, exists := p.endpoints.EndpointByChainID(ref.ChainID)
	if !exists {
		return errors.New("endpoint not found for the given chainID")
	}
	// connect to the endpoint
	client, err := currentEndpoint.GetClient(DefaultMaxWeb3ClientRetries)
	if err != nil {
		return errors.Join(ErrConnectingToWeb3Client, fmt.Errorf("[FARCASTER ID REGISTRY] %s: %w", ref.HexAddress, err))
	}
	// set the client, parse the address and initialize the contract
	p.client = client
	address := common.HexToAddress(ref.HexAddress)
	if p.contract, err = fckr.NewFarcasterIDRegistry(address, client); err != nil {
		return errors.Join(ErrInitializingContract, fmt.Errorf("[FARCASTER ID REGISTRY] %s: %w", p.address, err))
	}
	// reset the internal attributes
	p.address = address
	p.chainID = ref.ChainID
	p.name = ""
	p.symbol = ""
	p.decimals = 0
	p.totalSupply = nil
	p.creationBlock = 0
	p.lastNetworkBlock = 0
	p.synced.Store(false)
	return nil
}

// SetLastBalances method is not implemented for Farcaster key registry, it already
// calculate the partial balances from logs without comparing with the previous
// balances.
func (p *FarcasterIDProvider) SetLastBalances(_ context.Context, _ []byte,
	_ map[common.Address]*big.Int, _ uint64,
) error {
	return nil
}

// SetLastBlockNumber sets the last block number of the token set in the
// provider. It is used to calculate the delta balances in the next call to
// HoldersBalances from the given from point in time. It helps to avoid
// GetBlockNumber calls to the provider.
func (p *FarcasterIDProvider) SetLastBlockNumber(blockNumber uint64) {
	p.lastNetworkBlock = blockNumber
}

// HoldersBalances returns the registries on the FarcasterIDRegistry for the current
// defined farcaster ID registry contract (using SetRef method). It returns the farcaster id of
// new registered users for this farcaster registry contract from the block number provided to the latest posible block
// number (chosen between the last block number of the network and the maximun
// number of blocks to scan). It calls to rangeOfLogs to get the logs of the
// registries in the range of blocks and then it iterates the logs to
// calculate to get the registries with the recovery address and the farcasterID.
// It returns the farcasterIDs, the number of new registries, the last block scanned
// if the provider is synced and an error if it exists.
//
// NOTE that map[common.Address]*big.Int is used to store the farcasterID of each recovery address
func (p *FarcasterIDProvider) HoldersBalances(ctx context.Context, _ []byte, fromBlock uint64) (
	map[common.Address]*big.Int, uint64, uint64, bool, error,
) {
	// calculate the range of blocks to scan, by default take the last block
	// scanned and scan to the latest block, calculate the latest block if the
	// current last network block is not defined
	toBlock := p.lastNetworkBlock
	if toBlock == 0 {
		var err error
		toBlock, err = p.LatestBlockNumber(ctx, nil)
		if err != nil {
			return nil, 0, fromBlock, false, err
		}
	}
	log.Infow("scan iteration",
		"address", p.address,
		"type", p.TypeName(),
		"from", fromBlock,
		"to", toBlock)
	// iterate scanning the logs in the range of blocks until the last block
	// is reached
	startTime := time.Now()
	logs, lastBlock, synced, err := rangeOfLogs(ctx, p.client, p.address, fromBlock, toBlock, LOG_TOPIC_ERC20_TRANSFER)
	if err != nil {
		return nil, 0, fromBlock, false, err
	}
	// encode the number of new registries
	newRegistries := uint64(len(logs))
	registries := make(map[common.Address]*big.Int)
	// iterate the logs and update the registries
	for _, currentLog := range logs {
		logData, err := p.contract.FarcasterIDRegistryFilterer.ParseRegister(currentLog)
		if err != nil {
			return nil, newRegistries, lastBlock, false, errors.Join(ErrParsingTokenLogs, fmt.Errorf("[Farcaster ID Registry] %s: %w", p.address, err))
		}
		// update registries
		if registry, ok := registries[logData.Recovery]; ok {
			registries[logData.Recovery] = new(big.Int).Add(registry, logData.Id)
		}
	}
	log.Infow("saving blocks",
		"count", len(registries),
		"logs", len(logs),
		"blocks/s", 1000*float32(lastBlock-fromBlock)/float32(time.Since(startTime).Milliseconds()),
		"took", time.Since(startTime).Seconds(),
		"progress", fmt.Sprintf("%d%%", (fromBlock*100)/toBlock))
	p.synced.Store(synced)
	return registries, newRegistries, lastBlock, synced, nil
}

// Close method is not implemented for Farcaster Key Registry.
func (p *FarcasterIDProvider) Close() error {
	return nil
}

// IsExternal returns false because the provider is not an external API.
func (p *FarcasterIDProvider) IsExternal() bool {
	return false
}

// IsSynced returns true if the current state of the provider is synced. It also
// receives an external ID but it is not used by the provider.
func (p *FarcasterIDProvider) IsSynced(_ []byte) bool {
	return p.synced.Load()
}

// Address returns the address of the current token set in the provider.
func (p *FarcasterIDProvider) Address() common.Address {
	return p.address
}

// Type returns the type of the current token set in the provider.
func (p *FarcasterIDProvider) Type() uint64 {
	return providers.CONTRACT_TYPE_FARCASTER_ID_REGISTRY
}

// TypeName returns the type name of the current token set in the provider.
func (p *FarcasterIDProvider) TypeName() string {
	return providers.TokenTypeName(providers.CONTRACT_TYPE_FARCASTER_ID_REGISTRY)
}

// ChainID returns the chain ID of the current token set in the provider.
func (p *FarcasterIDProvider) ChainID() uint64 {
	return p.chainID
}

// Name returns the name of the current token set in the provider. It gets the
// name from the contract. It also receives an external ID but it is not used by
// the provider.
func (p *FarcasterIDProvider) Name(_ []byte) (string, error) {
	return "Farcaster ID Registry", nil
}

// Symbol returns the symbol of the current token set in the provider. It gets
// the symbol from the contract. It also receives an external ID but it is not
// used by the provider.
func (p *FarcasterIDProvider) Symbol(_ []byte) (string, error) {
	return "", nil
}

// Decimals returns the decimals of the current token set in the provider. It
// gets the decimals from the contract. It also receives an external ID but it
// is not used by the provider.
func (p *FarcasterIDProvider) Decimals(_ []byte) (uint64, error) {
	return 0, nil
}

// TotalSupply returns the total supply of the current token set in the provider.
// It gets the total supply from the contract. It also receives an external ID
// but it is not used by the provider.
func (p *FarcasterIDProvider) TotalSupply(_ []byte) (*big.Int, error) {
	return p.contract.IdCounter(nil)
}

// BalanceOf returns the balance of the given address for the current token set
// in the provider. It gets the balance from the contract. It also receives an
// external ID but it is not used by the provider.
func (p *FarcasterIDProvider) BalanceOf(addr common.Address, _ []byte) (*big.Int, error) {
	return big.NewInt(0), nil
}

// BalanceAt returns the balance of the given address for the current token at
// the given block number for the current token set in the provider. It gets
// the balance from the contract. It also receives an external ID but it is not
// used by the provider.
func (p *FarcasterIDProvider) BalanceAt(ctx context.Context, addr common.Address,
	_ []byte, blockNumber uint64,
) (*big.Int, error) {
	return p.client.BalanceAt(ctx, addr, new(big.Int).SetUint64(blockNumber))
}

// BlockTimestamp returns the timestamp of the given block number for the
// current token set in the provider. It gets the timestamp from the client.
func (p *FarcasterIDProvider) BlockTimestamp(ctx context.Context, blockNumber uint64) (string, error) {
	blockHeader, err := p.client.HeaderByNumber(ctx, new(big.Int).SetUint64(blockNumber))
	if err != nil {
		return "", err
	}
	return time.Unix(int64(blockHeader.Time), 0).Format(timeLayout), nil
}

// BlockRootHash returns the root hash of the given block number for the current
// token set in the provider. It gets the root hash from the client.
func (p *FarcasterIDProvider) BlockRootHash(ctx context.Context, blockNumber uint64) ([]byte, error) {
	blockHeader, err := p.client.HeaderByNumber(ctx, new(big.Int).SetInt64(int64(blockNumber)))
	if err != nil {
		return nil, err
	}
	return blockHeader.Root.Bytes(), nil
}

// LatestBlockNumber returns the latest block number of the current token set
// in the provider. It gets the latest block number from the client. It also
// receives an external ID but it is not used by the provider.
func (p *FarcasterIDProvider) LatestBlockNumber(ctx context.Context, _ []byte) (uint64, error) {
	lastBlockHeader, err := p.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return 0, err
	}
	return lastBlockHeader.Number.Uint64(), nil
}

// CreationBlock returns the creation block of the current token set in the
// provider. It gets the creation block from the client. It also receives an
// external ID but it is not used by the provider. It uses the
// creationBlockInRange function to calculate the creation block in the range
// of blocks.
func (p *FarcasterIDProvider) CreationBlock(ctx context.Context, _ []byte) (uint64, error) {
	var err error
	if p.creationBlock == 0 {
		var lastBlock uint64
		lastBlock, err = p.LatestBlockNumber(ctx, nil)
		if err != nil {
			return 0, err
		}
		p.creationBlock, err = creationBlockInRange(p.client, ctx, p.address, 0, lastBlock)
	}
	return p.creationBlock, err
}

// IconURI method is not implemented for Farcaster Key Registry tokens.
func (p *FarcasterIDProvider) IconURI(_ []byte) (string, error) {
	return "", nil
}

// Return the custody address of a given FarcasterID
func (p *FarcasterIDProvider) CustodyOf(fid *big.Int) (common.Address, error) {
	return p.contract.CustodyOf(nil, fid)
}

// Return the ID of a given custody address
func (p *FarcasterIDProvider) IdOf(custody common.Address) (*big.Int, error) {
	return p.contract.IdOf(nil, custody)
}

// Verifies a given FID signature
func (p *FarcasterIDProvider) VerifyFIDSignature(custodyAddress common.Address,
	fid *big.Int,
	digest [32]byte,
	signature []byte,
) (bool, error) {
	return p.contract.VerifyFidSignature(nil, custodyAddress, fid, digest, signature)
}
