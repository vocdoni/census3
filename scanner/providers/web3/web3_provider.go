package web3

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/vocdoni/census3/scanner/providers"
	"go.vocdoni.io/dvote/log"
)

type Web3ProviderRef struct {
	HexAddress string
	ChainID    uint64
}

type Web3ProviderConfig struct {
	Web3ProviderRef
	Endpoints NetworkEndpoints
}

// TokenTypeFromString function returns the token type ID from a string value.
// If the string is not recognized, it returns CONTRACT_TYPE_UNKNOWN.
func TokenTypeFromString(s string) uint64 {
	if c, ok := providers.TokenTypeIntMap[s]; ok {
		return c
	}
	return providers.CONTRACT_TYPE_UNKNOWN
}

// creationBlockInRange function finds the block number of a contract between
// the bounds provided as start and end blocks.
func creationBlockInRange(client *ethclient.Client, ctx context.Context, addr common.Address, start, end uint64) (uint64, error) {
	// if both block numbers are equal, return its value as birthblock
	if start == end {
		return start, nil
	}
	// find the middle block between start and end blocks and get the contract
	// code at this block
	midBlock := (start + end) / 2
	codeLen, err := sourceCodeLenAt(client, ctx, addr, midBlock)
	if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("No state available for block %d", midBlock)) {
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
func sourceCodeLenAt(client *ethclient.Client, ctx context.Context, addr common.Address, atBlockNumber uint64) (int, error) {
	blockNumber := new(big.Int).SetUint64(atBlockNumber)
	sourceCode, err := client.CodeAt(ctx, addr, blockNumber)
	return len(sourceCode), err
}

// rangeOfLogs function returns the logs of a token contract between the
// provided block numbers. It returns the logs, the last block scanned and an
// error if any. It filters the logs by the topic hash and for the token
// contract address provided.
func rangeOfLogs(ctx context.Context, client *ethclient.Client, addr common.Address,
	fromBlock, lastBlock uint64, hexTopics ...string,
) ([]types.Log, uint64, bool, error) {
	// if the range is too big, scan only a part of it using the constant
	// BLOCKS_TO_SCAN_AT_ONCE
	initialLastBlock := lastBlock
	if lastBlock-fromBlock > BLOCKS_TO_SCAN_AT_ONCE {
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
			log.Debugw("scanning logs",
				"address", addr.Hex(),
				"fromBlock", fromBlock,
				"toBlock", fromBlock+blocksRange-1)
			// compose the filter to get the logs of the ERC20 Transfer events
			filter := ethereum.FilterQuery{
				Addresses: []common.Address{addr},
				FromBlock: new(big.Int).SetUint64(fromBlock),
				ToBlock:   new(big.Int).SetUint64(fromBlock + blocksRange - 1),
				Topics:    [][]common.Hash{topicHashes},
			}
			// get the logs and check if there are any errors
			logs, err := client.FilterLogs(ctx, filter)
			if err != nil {
				// if the error is about the query returning more than the maximum
				// allowed logs, split the range of blocks in half and try again
				if strings.Contains(err.Error(), "query returned more than") {
					blocksRange /= 2
					log.Warnf("too much results on query, decreasing blocks to %d", blocksRange)
					continue
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
	synced := fromBlock >= initialLastBlock
	lastSyncedBlock := initialLastBlock
	if !synced {
		lastSyncedBlock = uint64(0)
		for _, l := range finalLogs {
			if l.BlockNumber > lastSyncedBlock {
				lastSyncedBlock = l.BlockNumber + 1
			}
		}
	}
	return finalLogs, lastSyncedBlock, synced, nil
}
