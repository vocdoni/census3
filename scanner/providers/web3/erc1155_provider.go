package web3

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	erc1155 "github.com/vocdoni/census3/contracts/erc/erc1155"
	"github.com/vocdoni/census3/helpers/web3"
	"github.com/vocdoni/census3/scanner/providers"
	"go.vocdoni.io/dvote/log"
)

type ERC1155HolderProvider struct {
	endpoints *web3.Web3Pool
	client    *web3.Client

	contract         *erc1155.ERC1155Contract
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

func (p *ERC1155HolderProvider) Init(_ context.Context, iconf any) error {
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
func (p *ERC1155HolderProvider) SetRef(iref any) error {
	if p.endpoints == nil {
		return errors.New("endpoints not defined")
	}
	ref, ok := iref.(Web3ProviderRef)
	if !ok {
		return errors.New("invalid ref type, it must be Web3ProviderRef")
	}
	var err error
	p.client, err = p.endpoints.Client(ref.ChainID)
	if err != nil {
		return fmt.Errorf("error getting web3 client for the given chainID: %w", err)
	}
	// set the client, parse the address and initialize the contract
	address := common.HexToAddress(ref.HexAddress)
	if p.contract, err = erc1155.NewERC1155Contract(address, p.client); err != nil {
		return errors.Join(ErrInitializingContract, fmt.Errorf("[ERC1155] %s: %w", p.address, err))
	}
	if ref.CreationBlock > 0 {
		p.creationBlock = ref.CreationBlock
	}
	// reset the internal attributes
	p.address = address
	p.chainID = ref.ChainID
	p.name = ""
	p.symbol = ""
	p.decimals = 0
	p.totalSupply = nil
	p.lastNetworkBlock = 0
	p.synced.Store(false)
	return nil
}

// SetLastBalances method is not implemented for ERC1155 tokens, they already
// calculate the partial balances from logs without comparing with the previous
// balances.
func (p *ERC1155HolderProvider) SetLastBalances(_ context.Context, _ []byte,
	_ map[common.Address]*big.Int, _ uint64,
) error {
	return nil
}

// SetLastBlockNumber sets the last block number of the token set in the
// provider. It is used to calculate the delta balances in the next call to
// HoldersBalances from the given from point in time. It helps to avoid
// GetBlockNumber calls to the provider.
func (p *ERC1155HolderProvider) SetLastBlockNumber(blockNumber uint64) {
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
func (p *ERC1155HolderProvider) HoldersBalances(ctx context.Context, _ []byte, fromBlock uint64) (
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
			return nil, 0, fromBlock, false, nil, err
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
	// get single and batch transfer logs
	logs, lastBlock, synced, err := RangeOfLogs(ctx, p.client, p.address,
		fromBlock, toBlock, LOG_TOPIC_ERC1155_TRANSFER_SINGLE, LOG_TOPIC_ERC1155_TRANSFER_BATCH)
	if err != nil && !errors.Is(err, ErrTooManyRequests) {
		return nil, 0, fromBlock, false, big.NewInt(0), err
	}
	if errors.Is(err, ErrTooManyRequests) {
		log.Warnf("too many requests, the provider will continue in the next iteration from block %d", lastBlock)
	}
	// encode the number of new transfers
	newTransfers := uint64(len(logs))
	// decode logs
	balances := make(map[common.Address]*big.Int)
	// iterate the single transfer logs and update the balances
	for _, currentLog := range logs {
		var from, to common.Address
		var value *big.Int
		// try to parse the log as a single transfer
		if singleLogData, err := p.contract.ERC1155ContractFilterer.ParseTransferSingle(currentLog); err == nil {
			to = singleLogData.To
			from = singleLogData.From
			value = singleLogData.Value
		}
		// try to parse the log as a batch transfer
		if batchLogData, err := p.contract.ERC1155ContractFilterer.ParseTransferBatch(currentLog); err == nil {
			to = batchLogData.To
			from = batchLogData.From
			value = new(big.Int)
			for _, v := range batchLogData.Values {
				value.Add(value, v)
			}
		}
		if value == nil {
			return nil, newTransfers, lastBlock, false, nil, fmt.Errorf("[ERC1155] %w: %s", ErrParsingTokenLogs, p.address.Hex())
		}
		log.Infow("erc1155 transfer", "from", from.Hex(), "to", to.Hex(), "value", value.String())
		// update the balances
		if toBalance, ok := balances[to]; ok {
			balances[to] = new(big.Int).Add(toBalance, value)
		} else {
			balances[to] = value
		}
		if fromBalance, ok := balances[from]; ok {
			balances[from] = new(big.Int).Sub(fromBalance, value)
		} else {
			balances[from] = new(big.Int).Neg(value)
		}
	}
	log.Infow("saving blocks",
		"count", len(balances),
		"logs", newTransfers,
		"blocks/s", 1000*float32(lastBlock-fromBlock)/float32(time.Since(startTime).Milliseconds()),
		"took", time.Since(startTime).Seconds(),
		"progress", fmt.Sprintf("%d%%", (fromBlock*100)/toBlock))
	p.synced.Store(synced)
	return balances, newTransfers, lastBlock, synced, nil, nil
}

// Close method is not implemented for ERC1155 tokens.
func (p *ERC1155HolderProvider) Close() error {
	return nil
}

// IsExternal returns false because ERC1155 tokens are not external APIs.
func (p *ERC1155HolderProvider) IsExternal() bool {
	return false
}

// IsSynced returns true if the current state of the provider is synced. It also
// receives an external ID but it is not used by the provider.
func (p *ERC1155HolderProvider) IsSynced(_ []byte) bool {
	return p.synced.Load()
}

// Address returns the address of the current token set in the provider.
func (p *ERC1155HolderProvider) Address(_ []byte) common.Address {
	return p.address
}

// Type returns the type of the current token set in the provider.
func (p *ERC1155HolderProvider) Type() uint64 {
	return providers.CONTRACT_TYPE_ERC1155
}

// TypeName returns the type name of the current token set in the provider.
func (p *ERC1155HolderProvider) TypeName() string {
	return providers.TokenTypeName(providers.CONTRACT_TYPE_ERC1155)
}

// ChainID returns the chain ID of the current token set in the provider.
func (p *ERC1155HolderProvider) ChainID() uint64 {
	return p.chainID
}

// Name returns the name of the current token set in the provider. It also
// receives an external ID but it is not used by the provider. It calls to the
// contract to get the name.
func (p *ERC1155HolderProvider) Name(_ []byte) (string, error) {
	var err error
	if p.name == "" {
		if p.name, err = p.contract.ERC1155ContractCaller.Name(nil); err != nil {
			return "", nil
		}
	}
	return p.name, err
}

// Symbol returns the symbol of the current token set in the provider. It
// also receives an external ID to be used if it is required by the provider.
// It calls to the contract to get the symbol.
func (p *ERC1155HolderProvider) Symbol(_ []byte) (string, error) {
	var err error
	if p.symbol == "" {
		if p.symbol, err = p.contract.ERC1155ContractCaller.Symbol(nil); err != nil {
			return "", nil
		}
	}
	return p.symbol, err
}

// Decimals returns the decimals of the current token set in the provider. It
// also receives an external ID but it is not used by the provider. It calls to
// the contract to get the decimals.
func (p *ERC1155HolderProvider) Decimals(_ []byte) (uint64, error) {
	return 0, nil
}

// TotalSupply method is not implemented for ERC1155 tokens.
func (p *ERC1155HolderProvider) TotalSupply(_ []byte) (*big.Int, error) {
	return nil, nil
}

// BalanceOf returns the balance of the given address for the current token
// set in the provider. It also receives an external ID but it is not used by
// the provider. It calls to the contract to get the balance.
func (p *ERC1155HolderProvider) BalanceOf(addr common.Address, bID []byte) (*big.Int, error) {
	return p.contract.ERC1155ContractCaller.BalanceOf(nil, addr, new(big.Int).SetBytes(bID))
}

// BalanceAt returns the balance of the given address for the current token at
// the given block number for the current token set in the provider. It also
// receives an external ID but it is not used by the provider. It calls to the
// contract to get the balance.
func (p *ERC1155HolderProvider) BalanceAt(ctx context.Context, addr common.Address,
	_ []byte, blockNumber uint64,
) (*big.Int, error) {
	return p.client.BalanceAt(ctx, addr, new(big.Int).SetUint64(blockNumber))
}

// BlockTimestamp returns the timestamp of the given block number for the
// current token set in the provider. It calls to the client to get the block
// header and then it returns the timestamp.
func (p *ERC1155HolderProvider) BlockTimestamp(ctx context.Context, blockNumber uint64) (string, error) {
	blockHeader, err := p.client.HeaderByNumber(ctx, new(big.Int).SetUint64(blockNumber))
	if err != nil {
		return "", err
	}
	return time.Unix(int64(blockHeader.Time), 0).Format(TimeLayout), nil
}

// BlockRootHash returns the root hash of the given block number for the current
// token set in the provider. It calls to the client to get the block header and
// then it returns the root hash.
func (p *ERC1155HolderProvider) BlockRootHash(ctx context.Context, blockNumber uint64) ([]byte, error) {
	blockHeader, err := p.client.HeaderByNumber(ctx, new(big.Int).SetInt64(int64(blockNumber)))
	if err != nil {
		return nil, err
	}
	return blockHeader.Root.Bytes(), nil
}

// LatestBlockNumber returns the latest block number of the current token set
// in the provider. It calls to the client to get the latest block header and
// then it returns the block number. It also receives an external ID but it is
// not used by the provider.
func (p *ERC1155HolderProvider) LatestBlockNumber(ctx context.Context, _ []byte) (uint64, error) {
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
func (p *ERC1155HolderProvider) CreationBlock(ctx context.Context, _ []byte) (uint64, error) {
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

// IconURI method is not implemented for ERC1155 tokens.
func (p *ERC1155HolderProvider) IconURI(_ []byte) (string, error) {
	return "", nil
}

// CensusKeys method returns the holders and balances provided transformed. The
// ERC1155 provider does not need to transform the holders and balances, so it
// returns the data as is.
func (p *ERC1155HolderProvider) CensusKeys(data map[common.Address]*big.Int) (map[common.Address]*big.Int, error) {
	return data, nil
}