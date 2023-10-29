package service

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// HolderProvider is the interface that wraps the basic methods to interact with
// a holders provider. It is used by the HoldersScanner to get the balances of
// the token holders. It allows to implement different providers, such as
// external API's, Web3 RPC endpoints, etc.
type HolderProvider interface {
	// Init initializes the provider and its internal structures. Initial
	// attributes values must be defined in the struct that implements this
	// interface before calling this method.
	Init() error
	// SetLastBalances sets the balances of the token holders for the given
	// id and from point in time and store it in a snapshot. It is used to
	// calculate the delta balances in the next call to HoldersBalances from
	// the given from point in time.
	SetLastBalances(ctx context.Context, id []byte, balances map[common.Address]*big.Int, from uint64) error
	// HoldersBalances returns the balances of the token holders for the given
	// id and delta point in time, from the stored last snapshot.
	HoldersBalances(ctx context.Context, id []byte, to uint64) (map[common.Address]*big.Int, error)
	// Close closes the provider and its internal structures.
	Close() error
	// Token realated methods
	Address(ctx context.Context, id []byte) (common.Address, error)
	Name(ctx context.Context, id []byte) (string, error)
	Symbol(ctx context.Context, id []byte) (string, error)
	Decimals(ctx context.Context, id []byte) (uint64, error)
	TotalSupply(ctx context.Context, id []byte) (*big.Int, error)
	BalanceOf(ctx context.Context, id []byte, addr common.Address) (*big.Int, error)
	BlockTimestamp(ctx context.Context, blockNumber uint64) (string, error)
	LatestBlockNumber(ctx context.Context, id []byte) (uint64, error)
	CreationBlock(ctx context.Context, id []byte) (uint64, error)
}
