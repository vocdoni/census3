package evmindexer

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/census3/db/treedb"
	"github.com/vocdoni/census3/helpers/web3"
	"go.vocdoni.io/dvote/log"
)

// BLOCK_BATCH_SIZE is the number of blocks to process in a single batch.
const BLOCK_BATCH_SIZE = 1000

// EVMIndexer struct contains a Web3Pool instance and a TreeDB instance. It encapsulates
// the artifacts and methods to index Ethereum smart contracts from blocks and their transactions into a TreeDB.
type EVMIndexer struct {
	web3                     *web3.Web3Pool
	db                       *treedb.TreeDB
	ctx                      context.Context
	lastProcessedBlockNumber map[uint64]uint64 // chainId -> lastProcessedBlockNumber
	coolDown                 time.Duration
}

// NewEVMIndexer method returns a new EVMIndexer instance, initialized with the provided Web3Pool and TreeDB instances.
func NewEVMIndexer(ctx context.Context, cooldown time.Duration, web3 *web3.Web3Pool, db *treedb.TreeDB) *EVMIndexer {
	return &EVMIndexer{
		web3:     web3,
		db:       db,
		ctx:      ctx,
		coolDown: cooldown,
	}
}

// Start is meant to be used to start the process that will index the EVM smart contracts
// for the supported networks available on the web3 pool.
func (e *EVMIndexer) Start() {
	// create a loop that runs forever and calls the ProcessBlockRange method with the BLOCK_BATCH_SIZE
	// as the range of blocks to process.
	for {
		select {
		case <-e.ctx.Done():
			return
		default:
			// get last block numbers
			blockNumbers, err := e.web3.CurrentBlockNumbers(e.ctx)
			if err != nil {
				log.Warnf("error getting current block numbers: %v", err)
				time.Sleep(e.coolDown)
				continue
			}
			for chainId, blockNumber := range blockNumbers {
				if blockNumber < e.lastProcessedBlockNumber[chainId] {
					client, err := e.web3.Client(chainId)
					if err != nil {
						log.Warnf("error getting client for chain %d: %v", chainId, err)
						continue
					}
					contracts, err := e.ProcessBlockRange(e.lastProcessedBlockNumber[chainId], BLOCK_BATCH_SIZE, client)
					if err != nil {
						log.Warnf("error processing block range: %v", err)
						continue
					}
					if err := e.IndexContracts(chainId, contracts); err != nil {
						log.Warnf(
							"error indexing contracts for blocks between %d and %d: %w",
							e.lastProcessedBlockNumber[chainId],
							BLOCK_BATCH_SIZE,
							err,
						)
					}
					e.lastProcessedBlockNumber[chainId] += BLOCK_BATCH_SIZE
				}
			}
		}
	}
}

// ProcessBlockRange is meant to be used to process a range of blocks by calling the ProcessBlock method for each block in the range.
func (e *EVMIndexer) ProcessBlockRange(from, to uint64, client *web3.Client) (map[uint64][]common.Address, error) {
	contractAddresses := make(map[uint64][]common.Address)
	for blockNumber := from; blockNumber < to; blockNumber++ {
		contractAddresses[blockNumber] = make([]common.Address, 0)
		contracts, err := e.ProcessBlock(blockNumber, client)
		if err != nil {
			return nil, fmt.Errorf("error processing block %d: %w", blockNumber, err)
		}
		contractAddresses[blockNumber] = contracts
	}
	return contractAddresses, nil
}

// ProcessBlock method processes the provided block number. It holds all the logic
// for processing block transactions getting the contract addresses.
func (e *EVMIndexer) ProcessBlock(blockNumber uint64, client *web3.Client) ([]common.Address, error) {
	block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(blockNumber)))
	if err != nil {
		return nil, fmt.Errorf("error getting block %d: %w", blockNumber, err)
	}

	contractAddresses := make([]common.Address, 0)
	for _, tx := range block.Transactions() {
		if tx.To() == nil {
			receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
			if err != nil {
				return nil, fmt.Errorf("error getting transaction receipt for tx %s: %w", tx.Hash().Hex(), err)
			}
			contractAddress := receipt.ContractAddress
			if contractAddress != (common.Address{}) {
				log.Debugf("Block %d: Contract %s created in transaction %s\n", blockNumber, contractAddress.Hex(), tx.Hash().Hex())
				contractAddresses = append(contractAddresses, contractAddress)
			}
		}
	}
	log.Debugf("Finished scanning block %d\n", blockNumber)
	return contractAddresses, nil
}

// IndexContracts method indexes the smart contracts from the provided block number.
func (e *EVMIndexer) IndexContracts(chainId uint64, contracts map[uint64][]common.Address) error {
	for blockNumber, contractAddresses := range contracts {
		for _, contractAddress := range contractAddresses {
			if err := e.IndexContract(chainId, blockNumber, contractAddress); err != nil {
				return fmt.Errorf("error indexing contract %s: %w", contractAddress.Hex(), err)
			}
		}
	}
	return nil
}

// IndexContract method indexes the smart contracts from the provided block number.
func (e *EVMIndexer) IndexContract(chainId, blockNumber uint64, contractAddress common.Address) error {
	// store the chainId, blockNumber and contractAddress in the TreeDB
	fmt.Printf("Indexing contract %s from block %d on chain %d\n", contractAddress.Hex(), blockNumber, chainId)
	return nil
}
