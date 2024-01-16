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
)

type ERC20HolderProvider struct {
	HexAddress string
	ChainID    uint64
	Client     *ethclient.Client

	contract      *erc20.ERC20Contract
	address       common.Address
	name          string
	symbol        string
	decimals      uint64
	totalSupply   *big.Int
	creationBlock uint64

	balances      map[common.Address]*big.Int
	balancesMtx   sync.RWMutex
	balancesBlock uint64
}

func (p *ERC20HolderProvider) Init() error {
	if p.HexAddress == "" || p.ChainID == 0 || p.Client == nil {
		return ErrInvalidProviderAttributes
	}
	p.address = common.HexToAddress(p.HexAddress)
	var err error
	p.contract, err = erc20.NewERC20Contract(p.address, p.Client)
	if err != nil {
		return errors.Join(ErrInitializingContract, fmt.Errorf("[ERC20] %s: %w", p.HexAddress, err))
	}
	p.balancesBlock, err = p.CreationBlock(context.Background(), nil)
	p.balances = make(map[common.Address]*big.Int)
	p.balancesMtx = sync.RWMutex{}
	return err
}

func (p *ERC20HolderProvider) SetLastBalances(ctx context.Context, id []byte, balances map[common.Address]*big.Int, from uint64) error {
	p.balancesMtx.Lock()
	defer p.balancesMtx.Unlock()

	if from < p.balancesBlock {
		return errors.New("from block is lower than the last block analyzed")
	}
	p.balancesBlock = from
	p.balances = balances
	return nil
}

func (p *ERC20HolderProvider) HoldersBalances(ctx context.Context, _ []byte, _ uint64) (map[common.Address]*big.Int, error) {
	// calculate the range of blocks to scan, by default take the last block
	// scanned and scan to the latest block
	fromBlock := p.balancesBlock
	toBlock, err := p.LatestBlockNumber(ctx, nil)
	if err != nil {
		return nil, err
	}
	// if the range is too big, scan only a part of it using the constant
	// BLOCKS_TO_SCAN_AT_ONCE
	if toBlock > fromBlock+BLOCKS_TO_SCAN_AT_ONCE {
		toBlock = fromBlock + BLOCKS_TO_SCAN_AT_ONCE
	}
	// iterate scanning the logs in the range of blocks until the last block
	// is reached
	for fromBlock < toBlock {
		// compose the filter to get the logs of the ERC20 Transfer events
		filter := ethereum.FilterQuery{
			Addresses: []common.Address{p.address},
			FromBlock: new(big.Int).SetUint64(fromBlock),
			ToBlock:   new(big.Int).SetUint64(toBlock),
			Topics: [][]common.Hash{
				{common.HexToHash(LOG_TOPIC_ERC20_TRANSFER)},
			},
		}
		// get the logs and check if there are any errors
		logs, err := p.Client.FilterLogs(ctx, filter)
		if err != nil {
			// if the error is about the query returning more than the maximum
			// allowed logs, split the range of blocks in half and try again
			if strings.Contains(err.Error(), "query returned more than") {
				toBlock = fromBlock + ((toBlock - fromBlock) / 2)
				continue
			}
			return nil, errors.Join(ErrScanningTokenLogs, fmt.Errorf("[ERC20] %s: %w", p.HexAddress, err))
		}
		// iterate the logs and update the balances
		for _, log := range logs {
			logData, err := p.contract.ERC20ContractFilterer.ParseTransfer(log)
			if err != nil {
				return nil, errors.Join(ErrParsingTokenLogs, fmt.Errorf("[ERC20] %s: %w", p.HexAddress, err))
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
		}
	}
	return p.balances, nil
}

func (p *ERC20HolderProvider) Close() error {
	return nil
}

func (p *ERC20HolderProvider) Address() common.Address {
	return p.address
}

func (p *ERC20HolderProvider) Type() TokenType {
	return CONTRACT_TYPE_ERC20
}

func (p *ERC20HolderProvider) NetworkID() uint64 {
	return p.ChainID
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

func (p *ERC20HolderProvider) BalanceOf(_ []byte, addr common.Address) (*big.Int, error) {
	return p.contract.ERC20ContractCaller.BalanceOf(nil, addr)
}

func (p *ERC20HolderProvider) BlockTimestamp(ctx context.Context, blockNumber uint64) (string, error) {
	blockHeader, err := p.Client.HeaderByNumber(ctx, new(big.Int).SetInt64(int64(blockNumber)))
	if err != nil {
		return "", err
	}
	return time.Unix(int64(blockHeader.Time), 0).Format(timeLayout), nil
}

func (p *ERC20HolderProvider) BlockRootHash(ctx context.Context, blockNumber uint64) ([]byte, error) {
	blockHeader, err := p.Client.HeaderByNumber(ctx, new(big.Int).SetInt64(int64(blockNumber)))
	if err != nil {
		return nil, err
	}
	return blockHeader.Root.Bytes(), nil
}

func (p *ERC20HolderProvider) LatestBlockNumber(ctx context.Context, _ []byte) (uint64, error) {
	lastBlockHeader, err := p.Client.HeaderByNumber(ctx, nil)
	if err != nil {
		return 0, err
	}
	return lastBlockHeader.Number.Uint64(), nil
}

func (p *ERC20HolderProvider) CreationBlock(ctx context.Context, _ []byte) (uint64, error) {
	var err error
	if p.creationBlock != 0 {
		var lastBlock uint64
		lastBlock, err = p.LatestBlockNumber(ctx, nil)
		if err != nil {
			return 0, err
		}
		p.creationBlock, err = creationBlockInRange(p.Client, ctx, p.address, 0, lastBlock)
	}
	return p.creationBlock, err
}

func (p *ERC20HolderProvider) IconURI(_ []byte) (string, error) {
	return "", nil
}
