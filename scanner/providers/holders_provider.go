package providers

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// BlocksDelta struct defines the delta of blocks processed by any
// HolderProvider. It includes the total number of logs processed, the new logs
// processed, the logs already processed, the last block processed, and if the
// provider is synced. It also includes the current total supply of the token
// set in the provider.
type BlocksDelta struct {
	LogsCount                 uint64
	NewLogsCount              uint64
	AlreadyProcessedLogsCount uint64
	Block                     uint64
	Synced                    bool
	TotalSupply               *big.Int
	NewLogs                   [][]byte
}

// Filter interface defines the basic methods to interact with a filter to
// store the processed transfers identifiers and avoid to process them again,
// for example, if a token is rescanned. It allows to implement different
// filters, such as in-memory, disk, merkle tree, etc.
type Filter interface {
	CheckKey(key []byte) (bool, error)
	CheckAndAddKey(key []byte) (bool, error)
}

// HolderProvider is the interface that wraps the basic methods to interact with
// a holders provider. It is used by the HoldersScanner to get the balances of
// the token holders. It allows to implement different providers, such as
// external API's, Web3 RPC endpoints, etc.
type HolderProvider interface {
	// Init initializes the provider and its internal structures. Initial
	// attributes values must be defined in the struct that implements this
	// interface before calling this method.
	Init(ctx context.Context, conf any) error
	// SetRef sets the reference to the provider. It is used to define the
	// required token information to interact with the provider.
	SetRef(ref any) error
	// SetLastBalances sets the balances of the token holders for the given
	// id and from point in time and store it in a snapshot. It is used to
	// calculate the delta balances in the next call to HoldersBalances from
	// the given from point in time.
	SetLastBalances(ctx context.Context, id []byte, balances map[common.Address]*big.Int, from uint64) error
	// SetLastBlockNumber sets the last block number of the token set in the
	// provider. It is used to calculate the delta balances in the next call
	// to HoldersBalances from the given from point in time. It helps to avoid
	// GetBlockNumber calls to the provider.
	SetLastBlockNumber(blockNumber uint64)
	// HoldersBalances returns the balances of the token holders for the given
	// id and delta point in time, from the stored last snapshot. It also
	// returns the total supply of tokens as a *big.Int.
	HoldersBalances(ctx context.Context, id []byte, to uint64) (map[common.Address]*big.Int, *BlocksDelta, error)
	// Close closes the provider and its internal structures.
	Close() error
	// IsExternal returns true if the provider is an external API.
	IsExternal() bool
	// IsSynced returns true if the current state of the provider is synced. It
	// also receives an external ID to be used if it is required by the provider.
	IsSynced(id []byte) bool
	// Address returns the address of the current token set in the provider. It
	// also receives an external ID to be used if it is required by the
	// provider.
	Address(id []byte) common.Address
	// Type returns the type of the current token set in the provider
	Type() uint64
	// TypeName returns the type name of the current token set in the provider
	TypeName() string
	// ChainID returns the chain ID of the current token set in the provider
	ChainID() uint64
	// Name returns the name of the current token set in the provider. It also
	// receives an external ID to be used if it is required by the provider.
	Name(id []byte) (string, error)
	// Symbol returns the symbol of the current token set in the provider. It
	// also receives an external ID to be used if it is required by the provider.
	Symbol(id []byte) (string, error)
	// Decimals returns the decimals of the current token set in the provider.
	// It also receives an external ID to be used if it is required by the
	// provider.
	Decimals(id []byte) (uint64, error)
	// TotalSupply returns the total supply of the current token set in the
	// provider. It also receives an external ID to be used if it is required
	// by the provider.
	TotalSupply(id []byte) (*big.Int, error)
	// BalanceOf returns the balance of the given address for the current token
	// set in the provider. It also receives an external ID to be used if it is
	// required by the provider.
	BalanceOf(addr common.Address, id []byte) (*big.Int, error)
	// BalanceAt returns the balance of the given address for the current token
	// at the given block number for the current token set in the provider. It
	// also receives an external ID to be used if it is required by the provider.
	BalanceAt(ctx context.Context, addr common.Address, id []byte, blockNumber uint64) (*big.Int, error)
	// BlockTimestamp returns the timestamp of the given block number for the
	// current token set in the provider
	BlockTimestamp(ctx context.Context, blockNumber uint64) (string, error)
	// BlockRootHash returns the root hash of the given block number for the
	// current token set in the provider
	BlockRootHash(ctx context.Context, blockNumber uint64) ([]byte, error)
	// LatestBlockNumber returns the latest block number for the network of the
	// current token set in the provider
	LatestBlockNumber(ctx context.Context, id []byte) (uint64, error)
	// CreationBlock returns the creation block number for the contract of the
	// current token set in the provider
	CreationBlock(ctx context.Context, id []byte) (uint64, error)
	// IconURI returns the icon URI of the icon asset of the current token set
	// in the provider.
	IconURI(id []byte) (string, error)
	// CensusKeys method returns the holders and balances provided transformed.
	// The transformation strategy is defined by the provider. The returned
	// map will be used to build the census.
	CensusKeys(holders map[common.Address]*big.Int) (map[common.Address]*big.Int, error)
}
