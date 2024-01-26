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
	erc721 "github.com/vocdoni/census3/contracts/erc/erc721"
	"github.com/vocdoni/census3/internal"
	"github.com/vocdoni/census3/scanner/providers"
	"go.vocdoni.io/dvote/log"
)

type ERC721HolderProvider struct {
	endpoints NetworkEndpoints
	client    *ethclient.Client

	contract      *erc721.ERC721Contract
	address       common.Address
	chainID       uint64
	name          string
	symbol        string
	decimals      uint64
	totalSupply   *big.Int
	creationBlock uint64
	synced        atomic.Bool
}

func (p *ERC721HolderProvider) Init(iconf any) error {
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

func (p *ERC721HolderProvider) SetRef(iref any) error {
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
		return errors.Join(ErrConnectingToWeb3Client, fmt.Errorf("[ERC721] %s: %w", ref.HexAddress, err))
	}
	// set the client, parse the address and initialize the contract
	p.client = client
	address := common.HexToAddress(ref.HexAddress)
	if p.contract, err = erc721.NewERC721Contract(address, client); err != nil {
		return errors.Join(ErrInitializingContract, fmt.Errorf("[ERC721] %s: %w", p.address, err))
	}
	// reset the internal attributes
	p.address = address
	p.chainID = ref.ChainID
	p.name = ""
	p.symbol = ""
	p.decimals = 0
	p.totalSupply = nil
	p.creationBlock = 0
	p.synced.Store(false)
	return nil
}

func (p *ERC721HolderProvider) SetLastBalances(_ context.Context, _ []byte,
	_ map[common.Address]*big.Int, _ uint64,
) error {
	return nil
}

func (p *ERC721HolderProvider) HoldersBalances(ctx context.Context, _ []byte, fromBlock uint64) (
	map[common.Address]*big.Int, uint64, uint64, bool, error,
) {
	// calculate the range of blocks to scan, by default take the last block
	// scanned and scan to the latest block
	toBlock, err := p.LatestBlockNumber(ctx, nil)
	if err != nil {
		return nil, 0, fromBlock, false, err
	}
	log.Infow("scan iteration",
		"address", p.address,
		"type", providers.TokenTypeStringMap[providers.CONTRACT_TYPE_ERC721],
		"from", fromBlock,
		"to", toBlock)
	// iterate scanning the logs in the range of blocks until the last block
	// is reached
	startTime := time.Now()
	logs, lastBlock, synced, err := rangeOfLogs(ctx, p.client, p.address, fromBlock, toBlock, LOG_TOPIC_ERC20_TRANSFER)
	if err != nil {
		return nil, 0, fromBlock, false, err
	}
	// encode the number of new transfers
	newTransfers := uint64(len(logs))
	balances := make(map[common.Address]*big.Int)
	// iterate the logs and update the balances
	for _, currentLog := range logs {
		logData, err := p.contract.ERC721ContractFilterer.ParseTransfer(currentLog)
		if err != nil {
			return nil, newTransfers, lastBlock, false, fmt.Errorf("[ERC721] %w: %s: %w", ErrParsingTokenLogs, p.address.Hex(), err)
		}
		// update balances
		if toBalance, ok := balances[logData.To]; ok {
			balances[logData.To] = new(big.Int).Add(toBalance, big.NewInt(1))
		} else {
			balances[logData.To] = big.NewInt(1)
		}
		if fromBalance, ok := balances[logData.From]; ok {
			balances[logData.From] = new(big.Int).Sub(fromBalance, big.NewInt(1))
		} else {
			balances[logData.From] = big.NewInt(-1)
		}
	}
	log.Infow("saving blocks",
		"count", len(balances),
		"logs", len(logs),
		"blocks/s", 1000*float32(lastBlock-fromBlock)/float32(time.Since(startTime).Milliseconds()),
		"took", time.Since(startTime).Seconds(),
		"progress", fmt.Sprintf("%d%%", (fromBlock*100)/toBlock))
	p.synced.Store(synced)
	return balances, newTransfers, lastBlock, synced, nil
}

func (p *ERC721HolderProvider) Close() error {
	return nil
}

func (p *ERC721HolderProvider) IsExternal() bool {
	return false
}

func (p *ERC721HolderProvider) IsSynced(_ []byte) bool {
	return p.synced.Load()
}

func (p *ERC721HolderProvider) Address() common.Address {
	return p.address
}

func (p *ERC721HolderProvider) Type() uint64 {
	return providers.CONTRACT_TYPE_ERC721
}

func (p *ERC721HolderProvider) ChainID() uint64 {
	return p.chainID
}

func (p *ERC721HolderProvider) Name(_ []byte) (string, error) {
	var err error
	if p.name == "" {
		p.name, err = p.contract.ERC721ContractCaller.Name(nil)
	}
	return p.name, err
}

func (p *ERC721HolderProvider) Symbol(_ []byte) (string, error) {
	var err error
	if p.symbol == "" {
		p.symbol, err = p.contract.ERC721ContractCaller.Symbol(nil)
	}
	return p.symbol, err
}

func (p *ERC721HolderProvider) Decimals(_ []byte) (uint64, error) {
	return 0, nil
}

func (p *ERC721HolderProvider) TotalSupply(_ []byte) (*big.Int, error) {
	return nil, nil
}

func (p *ERC721HolderProvider) BalanceOf(addr common.Address, _ []byte) (*big.Int, error) {
	return p.contract.ERC721ContractCaller.BalanceOf(nil, addr)
}

func (p *ERC721HolderProvider) BalanceAt(ctx context.Context, addr common.Address,
	_ []byte, blockNumber uint64,
) (*big.Int, error) {
	return p.client.BalanceAt(ctx, addr, new(big.Int).SetUint64(blockNumber))
}

func (p *ERC721HolderProvider) BlockTimestamp(ctx context.Context, blockNumber uint64) (string, error) {
	internal.GetBlockByNumberCounter.Add(1)
	blockHeader, err := p.client.HeaderByNumber(ctx, new(big.Int).SetUint64(blockNumber))
	if err != nil {
		return "", err
	}
	return time.Unix(int64(blockHeader.Time), 0).Format(timeLayout), nil
}

func (p *ERC721HolderProvider) BlockRootHash(ctx context.Context, blockNumber uint64) ([]byte, error) {
	internal.GetBlockByNumberCounter.Add(1)
	blockHeader, err := p.client.HeaderByNumber(ctx, new(big.Int).SetInt64(int64(blockNumber)))
	if err != nil {
		return nil, err
	}
	return blockHeader.Root.Bytes(), nil
}

func (p *ERC721HolderProvider) LatestBlockNumber(ctx context.Context, _ []byte) (uint64, error) {
	internal.GetBlockByNumberCounter.Add(1)
	lastBlockHeader, err := p.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return 0, err
	}
	return lastBlockHeader.Number.Uint64(), nil
}

func (p *ERC721HolderProvider) CreationBlock(ctx context.Context, _ []byte) (uint64, error) {
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

func (p *ERC721HolderProvider) IconURI(_ []byte) (string, error) {
	return "", nil
}
