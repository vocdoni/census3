package web3

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/vocdoni/census3/helpers/web3"
	"github.com/vocdoni/census3/scanner/filter"
	"github.com/vocdoni/census3/scanner/providers"
	"go.vocdoni.io/dvote/db"
	"go.vocdoni.io/dvote/log"
)

type Web3ProviderRef struct {
	HexAddress    string
	ChainID       uint64
	CreationBlock uint64
	Filter        *filter.TokenFilter
}

type Web3ProviderConfig struct {
	Web3ProviderRef
	Endpoints *web3.Web3Pool
	DB        *db.Database
}

// creationBlock function returns the block number of the creation of a contract
// address. It uses the `eth_getCode` method to get the contract code at the
// block number provided. If the method is not supported, it returns 0 and nil.
func creationBlock(client *web3.Client, ctx context.Context, addr common.Address) (uint64, error) {
	// check if the current client supports `eth_getCode` method, if not, return
	// 1 and nil. It is assumed that the contract is created at block 1 to start
	// scanning from the first block.
	getCodeSupport := false
	for i := 0; i < web3.DefaultMaxWeb3ClientRetries; i++ {
		ethClient, err := client.EthClient()
		if err != nil {
			return 0, err
		}
		if getCodeSupport = providers.ClientSupportsGetCode(ctx, ethClient, addr); getCodeSupport {
			break
		}
		time.Sleep(RetryWeb3Cooldown)
	}
	if !getCodeSupport {
		return 1, nil
	}
	// get the latest block number
	var err error
	var lastBlock uint64
	for i := 0; i < web3.DefaultMaxWeb3ClientRetries; i++ {
		lastBlock, err = client.BlockNumber(ctx)
		if err == nil {
			break
		}
		time.Sleep(RetryWeb3Cooldown)
	}
	if err != nil {
		return 0, err
	}
	var creationBlock uint64
	for i := 0; i < web3.DefaultMaxWeb3ClientRetries; i++ {
		creationBlock, err = creationBlockInRange(client, ctx, addr, 0, lastBlock)
		if err == nil {
			break
		}
		time.Sleep(RetryWeb3Cooldown)
	}
	return creationBlock, err
}

// creationBlockInRange function finds the block number of a contract between
// the bounds provided as start and end blocks.
func creationBlockInRange(client *web3.Client, ctx context.Context,
	addr common.Address, start, end uint64,
) (uint64, error) {
	// if both block numbers are equal, return its value as birthblock
	if start == end {
		return start, nil
	}
	// find the middle block between start and end blocks and get the contract
	// code at this block
	midBlock := (start + end) / 2
	codeLen, err := sourceCodeLenAt(client, ctx, addr, midBlock)
	if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("No state available for block %d", midBlock)) &&
		!strings.Contains(err.Error(), "missing trie node") {
		return 0, err
	}
	// if any code is found, keep trying with the lower half of blocks until
	// find the first. if not, keep trying with the upper half
	if codeLen > 2 {
		return creationBlockInRange(client, ctx, addr, start, midBlock)
	} else {
		return creationBlockInRange(client, ctx, addr, midBlock+1, end)
	}
}

// SourceCodeLenAt function returns the length of the current contract bytecode
// at the block number provided.
func sourceCodeLenAt(client *web3.Client, ctx context.Context,
	addr common.Address, atBlockNumber uint64,
) (int, error) {
	blockNumber := new(big.Int).SetUint64(atBlockNumber)
	sourceCode, err := client.CodeAt(ctx, addr, blockNumber)
	return len(sourceCode), err
}

// RangeOfLogs function returns the logs of a token contract between the
// provided block numbers. It returns the logs, the last block scanned and an
// error if any. It filters the logs by the topic hash and for the token
// contract address provided.
func RangeOfLogs(ctx context.Context, client *web3.Client, addr common.Address,
	fromBlock, lastBlock uint64, hexTopics ...string,
) ([]types.Log, uint64, bool, error) {
	// if the range is too big, scan only a part of it using the constant
	// BLOCKS_TO_SCAN_AT_ONCE
	initialLastBlock := lastBlock
	if lastBlock-fromBlock > BLOCKS_TO_SCAN_AT_ONCE && fromBlock+MAX_SCAN_BLOCKS_PER_ITERATION < lastBlock {
		lastBlock = fromBlock + MAX_SCAN_BLOCKS_PER_ITERATION
	}
	if fromBlock > lastBlock {
		fromBlock = lastBlock
	}
	// some variables to calculate the end of the scan and store the logs
	logCount := 0
	finalLogs := []types.Log{}
	blocksRange := BLOCKS_TO_SCAN_AT_ONCE
	topicHashes := make([]common.Hash, len(hexTopics))
	for i, topic := range hexTopics {
		topicHashes[i] = common.HexToHash(topic)
	}
	for fromBlock < lastBlock {
		select {
		case <-ctx.Done():
			log.Warnf("scan graceful canceled by context")
			return finalLogs, fromBlock, false, nil
		default:
			if logCount > MAX_SCAN_LOGS_PER_ITERATION {
				return finalLogs, fromBlock, false, nil
			}
			toBlock := fromBlock + blocksRange - 1
			if toBlock > lastBlock {
				toBlock = lastBlock
			}
			log.Debugw("scanning logs",
				"address", addr.Hex(),
				"fromBlock", fromBlock,
				"toBlock", toBlock)
			// compose the filter to get the logs of the ERC20 Transfer events
			filter := ethereum.FilterQuery{
				Addresses: []common.Address{addr},
				FromBlock: new(big.Int).SetUint64(fromBlock),
				ToBlock:   new(big.Int).SetUint64(toBlock),
				Topics:    [][]common.Hash{topicHashes},
			}
			// get the logs and check if there are any errors
			logs, err := client.FilterLogs(ctx, filter)
			if err != nil {
				// if the error is about the query returning more than the maximum
				// allowed logs, split the range of blocks in half and try again
				if strings.Contains(strings.ToLower(err.Error()), "query returned more than") ||
					strings.Contains(strings.ToLower(err.Error()), "exceeds the range allowed") ||
					strings.Contains(strings.ToLower(err.Error()), "query timeout exceeded") ||
					strings.Contains(strings.ToLower(err.Error()), "execution aborted (timeout") ||
					strings.Contains(strings.ToLower(err.Error()), "size is larger than") {
					blocksRange /= 2
					log.Warnf("too much results on query, decreasing blocks to %d: %v", blocksRange, err)
					time.Sleep(RetryWeb3Cooldown)
					continue
				}
				// if error is about too many requests, return the logs scanned
				// until now and the last block scanned with an specific error
				if strings.Contains(strings.ToLower(err.Error()), "too many requests") {
					return finalLogs, fromBlock, false, errors.Join(ErrTooManyRequests, fmt.Errorf("%s: %w", addr.Hex(), err))
				}
				return finalLogs, fromBlock, false, errors.Join(ErrScanningTokenLogs, fmt.Errorf("%s: %w", addr.Hex(), err))
			}
			// if there are logs, add them to the final list and update the
			// counter
			if len(logs) > 0 {
				finalLogs = append(finalLogs, logs...)
				logCount += len(logs)
			}
			// update the fromBlock to the last block scanned
			fromBlock += blocksRange
		}
	}
	// if the last block scanned is the same as the last block of the range,
	// the scan is completed and the token is synced. If not, the token is not
	// synced and the last block scanned is the last block of the scanned range
	// plus one.
	if fromBlock > lastBlock {
		fromBlock = lastBlock
	}
	return finalLogs, fromBlock, fromBlock >= initialLastBlock, nil
}
