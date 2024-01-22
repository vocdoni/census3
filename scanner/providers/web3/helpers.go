package web3

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/vocdoni/census3/scanner/providers"
)

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
