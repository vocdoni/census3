package web3

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	erc20 "github.com/vocdoni/census3/contracts/erc/erc20"
	"github.com/vocdoni/census3/scanner/providers"
	"go.vocdoni.io/dvote/log"
)

type ERC20HolderProvider struct {
	endpoints NetworkEndpoints
	client    *ethclient.Client

	contract      *erc20.ERC20Contract
	address       common.Address
	chainID       uint64
	name          string
	symbol        string
	decimals      uint64
	totalSupply   *big.Int
	creationBlock uint64

	balances      map[common.Address]*big.Int
	balancesMtx   sync.RWMutex
	balancesBlock uint64
}

func (p *ERC20HolderProvider) Init(iconf any) error {
	// parse the config and set the endpoints
	conf, ok := iconf.(Web3ProviderConfig)
	if !ok {
		return errors.New("invalid config type, it must be Web3ProviderConfig")
	}
	p.endpoints = conf.Endpoints
	// reset the internal balances
	p.balances = make(map[common.Address]*big.Int)
	p.balancesMtx = sync.RWMutex{}
	// set the reference if the address and chainID are defined in the config
	if conf.HexAddress != "" && conf.ChainID > 0 {
		return p.SetRef(Web3ProviderRef{
			HexAddress: conf.HexAddress,
			ChainID:    conf.ChainID,
		})
	}
	return nil
}

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
	address := common.HexToAddress(ref.HexAddress)
	if p.contract, err = erc20.NewERC20Contract(address, client); err != nil {
		return errors.Join(ErrInitializingContract, fmt.Errorf("[ERC20] %s: %w", p.address, err))
	}
	// reset the internal attributes
	p.address = address
	p.chainID = ref.ChainID
	p.name = ""
	p.symbol = ""
	p.decimals = 0
	p.totalSupply = nil
	p.creationBlock = 0
	// reset balances
	p.balancesMtx.Lock()
	defer p.balancesMtx.Unlock()
	p.balances = make(map[common.Address]*big.Int)
	return nil
}

func (p *ERC20HolderProvider) SetLastBalances(ctx context.Context, id []byte,
	balances map[common.Address]*big.Int, from uint64,
) error {
	p.balancesMtx.Lock()
	defer p.balancesMtx.Unlock()

	if from < p.balancesBlock {
		return errors.New("from block is lower than the last block analyzed")
	}
	p.balancesBlock = from
	p.balances = balances
	return nil
}

func (p *ERC20HolderProvider) HoldersBalances(ctx context.Context, _ []byte, fromBlock uint64) (
	map[common.Address]*big.Int, uint64, bool, error,
) {
	// calculate the range of blocks to scan, by default take the last block
	// scanned and scan to the latest block
	toBlock, err := p.LatestBlockNumber(ctx, nil)
	if err != nil {
		return nil, fromBlock, false, err
	}
	// if the range is too big, scan only a part of it using the constant
	// BLOCKS_TO_SCAN_AT_ONCE
	if toBlock-fromBlock > BLOCKS_TO_SCAN_AT_ONCE {
		toBlock = fromBlock + MAX_SCAN_BLOCKS_PER_ITERATION
	}
	logCount := 0
	blocksRange := BLOCKS_TO_SCAN_AT_ONCE
	log.Infow("scan iteration",
		"address", p.address,
		"type", providers.TokenTypeStringMap[providers.CONTRACT_TYPE_ERC20],
		"from", fromBlock,
		"to", toBlock)
	// some variables to calculate the progress
	startTime := time.Now()
	initialBlock := fromBlock
	lastBlock := fromBlock
	p.balancesMtx.RLock()
	initialHolders := len(p.balances)
	p.balancesMtx.RUnlock()
	// iterate scanning the logs in the range of blocks until the last block
	// is reached
	for fromBlock < toBlock {
		select {
		case <-ctx.Done():
			log.Warnf("scan graceful canceled by context")
			return p.balances, lastBlock, false, nil
		default:
			if logCount > MAX_SCAN_LOGS_PER_ITERATION {
				return p.balances, lastBlock, false, nil
			}
			// compose the filter to get the logs of the ERC20 Transfer events
			filter := ethereum.FilterQuery{
				Addresses: []common.Address{p.address},
				FromBlock: new(big.Int).SetUint64(fromBlock),
				ToBlock:   new(big.Int).SetUint64(fromBlock + blocksRange),
				Topics: [][]common.Hash{
					{common.HexToHash(LOG_TOPIC_ERC20_TRANSFER)},
				},
			}
			// get the logs and check if there are any errors
			logs, err := p.client.FilterLogs(ctx, filter)
			if err != nil {
				// if the error is about the query returning more than the maximum
				// allowed logs, split the range of blocks in half and try again
				if strings.Contains(err.Error(), "query returned more than") {
					blocksRange /= 2
					log.Warnf("too much results on query, decreasing blocks to %d", blocksRange)
					continue
				}
				return nil, lastBlock, false, errors.Join(ErrScanningTokenLogs, fmt.Errorf("[ERC20] %s: %w", p.address, err))
			}
			// if there are no logs, the range of blocks is empty, so return the
			// balances
			if len(logs) == 0 {
				fromBlock += blocksRange
				continue
			}
			logCount += len(logs)
			// iterate the logs and update the balances
			for _, log := range logs {
				logData, err := p.contract.ERC20ContractFilterer.ParseTransfer(log)
				if err != nil {
					return nil, lastBlock, false, errors.Join(ErrParsingTokenLogs, fmt.Errorf("[ERC20] %s: %w", p.address, err))
				}
				// update balances
				p.balancesMtx.Lock()
				if toBalance, ok := p.balances[logData.To]; ok {
					p.balances[logData.To] = new(big.Int).Add(toBalance, logData.Value)
				} else {
					p.balances[logData.To] = logData.Value
				}
				if fromBalance, ok := p.balances[logData.From]; ok {
					p.balances[logData.From] = new(big.Int).Sub(fromBalance, logData.Value)
				} else {
					p.balances[logData.From] = new(big.Int).Neg(logData.Value)
				}
				p.balancesMtx.Unlock()
				lastBlock = log.BlockNumber
			}
			// update the fromBlock to the last block scanned
			fromBlock += blocksRange
		}
	}
	p.balancesMtx.RLock()
	finalHolders := len(p.balances)
	p.balancesMtx.RUnlock()
	log.Infow("saving blocks",
		"count", finalHolders-initialHolders,
		"logs", logCount,
		"blocks/s", 1000*float32(fromBlock-initialBlock)/float32(time.Since(startTime).Milliseconds()),
		"took", time.Since(startTime).Seconds(),
		"progress", fmt.Sprintf("%d%%", (fromBlock*100)/toBlock))
	return p.balances, lastBlock, true, nil
}

func (p *ERC20HolderProvider) Close() error {
	return nil
}

func (p *ERC20HolderProvider) IsExternal() bool {
	return false
}

func (p *ERC20HolderProvider) Address() common.Address {
	return p.address
}

func (p *ERC20HolderProvider) Type() uint64 {
	return providers.CONTRACT_TYPE_ERC20
}

func (p *ERC20HolderProvider) ChainID() uint64 {
	return p.chainID
}

func (p *ERC20HolderProvider) Name(_ []byte) (string, error) {
	var err error
	if p.name == "" {
		p.name, err = p.contract.ERC20ContractCaller.Name(nil)
	}
	return p.name, err
}

func (p *ERC20HolderProvider) Symbol(_ []byte) (string, error) {
	var err error
	if p.symbol == "" {
		p.symbol, err = p.contract.ERC20ContractCaller.Symbol(nil)
	}
	return p.symbol, err
}

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

func (p *ERC20HolderProvider) TotalSupply(_ []byte) (*big.Int, error) {
	var err error
	if p.totalSupply == nil {
		p.totalSupply, err = p.contract.ERC20ContractCaller.TotalSupply(nil)
	}
	return p.totalSupply, err
}

func (p *ERC20HolderProvider) BalanceOf(addr common.Address, _ []byte) (*big.Int, error) {
	return p.contract.ERC20ContractCaller.BalanceOf(nil, addr)
}

func (p *ERC20HolderProvider) BalanceAt(ctx context.Context, addr common.Address,
	_ []byte, blockNumber uint64,
) (*big.Int, error) {
	return p.client.BalanceAt(ctx, addr, new(big.Int).SetUint64(blockNumber))
}

func (p *ERC20HolderProvider) BlockTimestamp(ctx context.Context, blockNumber uint64) (string, error) {
	blockHeader, err := p.client.HeaderByNumber(ctx, new(big.Int).SetUint64(blockNumber))
	if err != nil {
		return "", err
	}
	return time.Unix(int64(blockHeader.Time), 0).Format(timeLayout), nil
}

func (p *ERC20HolderProvider) BlockRootHash(ctx context.Context, blockNumber uint64) ([]byte, error) {
	blockHeader, err := p.client.HeaderByNumber(ctx, new(big.Int).SetInt64(int64(blockNumber)))
	if err != nil {
		return nil, err
	}
	return blockHeader.Root.Bytes(), nil
}

func (p *ERC20HolderProvider) LatestBlockNumber(ctx context.Context, _ []byte) (uint64, error) {
	lastBlockHeader, err := p.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return 0, err
	}
	return lastBlockHeader.Number.Uint64(), nil
}

func (p *ERC20HolderProvider) CreationBlock(ctx context.Context, _ []byte) (uint64, error) {
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

func (p *ERC20HolderProvider) IconURI(_ []byte) (string, error) {
	return "", nil
}
