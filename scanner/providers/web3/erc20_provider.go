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
	erc20 "github.com/vocdoni/census3/contracts/erc/erc20"
	"github.com/vocdoni/census3/scanner/providers"
	"go.vocdoni.io/dvote/log"
)

type ERC20HolderProvider struct {
	endpoints NetworkEndpoints
	client    *ethclient.Client

	contract         *erc20.ERC20Contract
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

func (p *ERC20HolderProvider) Init(iconf any) error {
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
func (p *ERC20HolderProvider) SetRef(iref any) error {
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
		return errors.Join(ErrConnectingToWeb3Client, fmt.Errorf("[ERC20] %s: %w", ref.HexAddress, err))
	}
	// set the client, parse the address and initialize the contract
	p.client = client
	p.address = common.HexToAddress(ref.HexAddress)
	if p.contract, err = erc20.NewERC20Contract(p.address, client); err != nil {
		return errors.Join(ErrInitializingContract, fmt.Errorf("[ERC20] %s: %w", p.address, err))
	}
	// reset the current creation block, if a creation block is defined in the
	// reference, update it
	p.creationBlock = 0
	if ref.CreationBlock > 0 {
		p.creationBlock = ref.CreationBlock
	}
	// reset the internal attributes
	p.chainID = ref.ChainID
	p.name = ""
	p.symbol = ""
	p.decimals = 0
	p.totalSupply = nil
	p.lastNetworkBlock = 0
	p.synced.Store(false)
	return nil
}

// SetLastBalances method is not implemented for ERC20 tokens, they already
// calculate the partial balances from logs without comparing with the previous
// balances.
func (p *ERC20HolderProvider) SetLastBalances(_ context.Context, _ []byte,
	_ map[common.Address]*big.Int, _ uint64,
) error {
	return nil
}

// SetLastBlockNumber sets the last block number of the token set in the
// provider. It is used to calculate the delta balances in the next call to
// HoldersBalances from the given from point in time. It helps to avoid
// GetBlockNumber calls to the provider.
func (p *ERC20HolderProvider) SetLastBlockNumber(blockNumber uint64) {
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
func (p *ERC20HolderProvider) HoldersBalances(ctx context.Context, _ []byte, fromBlock uint64) (
	map[common.Address]*big.Int, uint64, uint64, bool, *big.Int, error,
) {
	// if the last network block is lower than the last scanned block, and the
	// last scanned block is equal to the creation block, it means that the
	// last network block is outdated, so it returns that it is not synced and
	// an error
	if fromBlock >= p.lastNetworkBlock && fromBlock == p.creationBlock {
		return nil, 0, fromBlock, false, big.NewInt(0),
			errors.New("outdated last network block, it will retry in the next iteration")
	}
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
		"address", p.address,
		"type", p.TypeName(),
		"from", fromBlock,
		"to", toBlock)
	// iterate scanning the logs in the range of blocks until the last block
	// is reached
	startTime := time.Now()
	logs, lastBlock, synced, err := RangeOfLogs(ctx, p.client, p.address, fromBlock, toBlock, LOG_TOPIC_ERC20_TRANSFER)
	if err != nil && !errors.Is(err, ErrTooManyRequests) {
		return nil, 0, fromBlock, false, big.NewInt(0), err
	}
	if errors.Is(err, ErrTooManyRequests) {
		log.Warnf("too many requests, the provider will continue in the next iteration from block %d", lastBlock)
	}
	// encode the number of new transfers
	newTransfers := uint64(len(logs))
	balances := make(map[common.Address]*big.Int)
	// iterate the logs and update the balances
	for _, currentLog := range logs {
		logData, err := p.contract.ERC20ContractFilterer.ParseTransfer(currentLog)
		if err != nil {
			return nil, newTransfers, lastBlock, false, big.NewInt(0),
				errors.Join(ErrParsingTokenLogs, fmt.Errorf("[ERC20] %s: %w", p.address, err))
		}
		// update balances
		if toBalance, ok := balances[logData.To]; ok {
			balances[logData.To] = new(big.Int).Add(toBalance, logData.Value)
		} else {
			balances[logData.To] = logData.Value
		}
		if fromBalance, ok := balances[logData.From]; ok {
			balances[logData.From] = new(big.Int).Sub(fromBalance, logData.Value)
		} else {
			balances[logData.From] = new(big.Int).Neg(logData.Value)
		}
	}
	log.Infow("saving blocks",
		"count", len(balances),
		"logs", len(logs),
		"blocks/s", 1000*float32(lastBlock-fromBlock)/float32(time.Since(startTime).Milliseconds()),
		"took", time.Since(startTime).Seconds(),
		"progress", fmt.Sprintf("%d%%", (fromBlock*100)/toBlock))
	p.synced.Store(synced)
	totalSupply, err := p.TotalSupply(nil)
	if err != nil {
		log.Warn("error getting total supply, it will retry in the next iteration", "error", err)
		return balances, newTransfers, lastBlock, synced, nil, nil
	}
	return balances, newTransfers, lastBlock, synced, totalSupply, nil
}

// Close method is not implemented for ERC20 tokens.
func (p *ERC20HolderProvider) Close() error {
	return nil
}

// IsExternal returns false because the provider is not an external API.
func (p *ERC20HolderProvider) IsExternal() bool {
	return false
}

// IsSynced returns true if the current state of the provider is synced. It also
// receives an external ID but it is not used by the provider.
func (p *ERC20HolderProvider) IsSynced(_ []byte) bool {
	return p.synced.Load()
}

// Address returns the address of the current token set in the provider.
func (p *ERC20HolderProvider) Address(_ []byte) common.Address {
	return p.address
}

// Type returns the type of the current token set in the provider.
func (p *ERC20HolderProvider) Type() uint64 {
	return providers.CONTRACT_TYPE_ERC20
}

// TypeName returns the type name of the current token set in the provider.
func (p *ERC20HolderProvider) TypeName() string {
	return providers.TokenTypeName(providers.CONTRACT_TYPE_ERC20)
}

// ChainID returns the chain ID of the current token set in the provider.
func (p *ERC20HolderProvider) ChainID() uint64 {
	return p.chainID
}

// Name returns the name of the current token set in the provider. It gets the
// name from the contract. It also receives an external ID but it is not used by
// the provider.
func (p *ERC20HolderProvider) Name(_ []byte) (string, error) {
	var err error
	if p.name == "" {
		p.name, err = p.contract.ERC20ContractCaller.Name(nil)
	}
	return p.name, err
}

// Symbol returns the symbol of the current token set in the provider. It gets
// the symbol from the contract. It also receives an external ID but it is not
// used by the provider.
func (p *ERC20HolderProvider) Symbol(_ []byte) (string, error) {
	var err error
	if p.symbol == "" {
		p.symbol, err = p.contract.ERC20ContractCaller.Symbol(nil)
	}
	return p.symbol, err
}

// Decimals returns the decimals of the current token set in the provider. It
// gets the decimals from the contract. It also receives an external ID but it
// is not used by the provider.
func (p *ERC20HolderProvider) Decimals(_ []byte) (uint64, error) {
	if p.decimals == 0 {
		decimals, err := p.contract.ERC20ContractCaller.Decimals(nil)
		if err != nil {
			return 0, err
		}
		p.decimals = uint64(decimals)
	}
	return p.decimals, nil
}

// TotalSupply returns the total supply of the current token set in the provider.
// It gets the total supply from the contract. It also receives an external ID
// but it is not used by the provider.
func (p *ERC20HolderProvider) TotalSupply(_ []byte) (*big.Int, error) {
	return p.contract.ERC20ContractCaller.TotalSupply(nil)
}

// BalanceOf returns the balance of the given address for the current token set
// in the provider. It gets the balance from the contract. It also receives an
// external ID but it is not used by the provider.
func (p *ERC20HolderProvider) BalanceOf(addr common.Address, _ []byte) (*big.Int, error) {
	return p.contract.ERC20ContractCaller.BalanceOf(nil, addr)
}

// BalanceAt returns the balance of the given address for the current token at
// the given block number for the current token set in the provider. It gets
// the balance from the contract. It also receives an external ID but it is not
// used by the provider.
func (p *ERC20HolderProvider) BalanceAt(ctx context.Context, addr common.Address,
	_ []byte, blockNumber uint64,
) (*big.Int, error) {
	return p.client.BalanceAt(ctx, addr, new(big.Int).SetUint64(blockNumber))
}

// BlockTimestamp returns the timestamp of the given block number for the
// current token set in the provider. It gets the timestamp from the client.
func (p *ERC20HolderProvider) BlockTimestamp(ctx context.Context, blockNumber uint64) (string, error) {
	blockHeader, err := p.client.HeaderByNumber(ctx, new(big.Int).SetUint64(blockNumber))
	if err != nil {
		return "", err
	}
	return time.Unix(int64(blockHeader.Time), 0).Format(TimeLayout), nil
}

// BlockRootHash returns the root hash of the given block number for the current
// token set in the provider. It gets the root hash from the client.
func (p *ERC20HolderProvider) BlockRootHash(ctx context.Context, blockNumber uint64) ([]byte, error) {
	blockHeader, err := p.client.HeaderByNumber(ctx, new(big.Int).SetInt64(int64(blockNumber)))
	if err != nil {
		return nil, err
	}
	return blockHeader.Root.Bytes(), nil
}

// LatestBlockNumber returns the latest block number of the current token set
// in the provider. It gets the latest block number from the client. It also
// receives an external ID but it is not used by the provider.
func (p *ERC20HolderProvider) LatestBlockNumber(ctx context.Context, _ []byte) (uint64, error) {
	lastBlockHeader, err := p.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return 0, err
	}
	return lastBlockHeader.Number.Uint64(), nil
}

// CreationBlock returns the creation block of the current token set in the
// provider. It gets the creation block from the client. It also receives an
// external ID but it is not used by the provider. It uses the
// creationBlock function to calculate the creation block of a web3 provider
// contract.
func (p *ERC20HolderProvider) CreationBlock(ctx context.Context, _ []byte) (uint64, error) {
	if p.creationBlock == 0 {
		var err error
		p.creationBlock, err = creationBlock(p.client, ctx, p.address)
		return p.creationBlock, err
	}
	return p.creationBlock, nil
}

// IconURI method is not implemented for ERC20 tokens.
func (p *ERC20HolderProvider) IconURI(_ []byte) (string, error) {
	return "", nil
}

// CensusKeys method returns the holders and balances provided transformed. The
// ERC20 provider does not need to transform the holders and balances, so it
// returns the data as is.
func (p *ERC20HolderProvider) CensusKeys(data map[common.Address]*big.Int) (map[common.Address]*big.Int, error) {
	return data, nil
}
