package state

import (
	"math/big"
	"sync"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/common"
)

type HoldersCandidates map[common.Address]*big.Int

// TokenHolders struct abstracts the current state of a TokenHolders into the
// Census3 HoldersScanner for a specific token. It contains some information
// about the token such as its address or its type. Also includes some atomic
// variables used to store the state of the token holders safely across
// differents goroutines, such as holders list, analyzed blocks or the last
// block analyzed.
type TokenHolders struct {
	address   common.Address
	ctype     TokenType
	holders   sync.Map
	blocks    sync.Map
	lastBlock atomic.Uint64
	synced    atomic.Bool
	ChainID   uint64
}

// Init function fills the given TokenHolders struct with the address and type
// given, also checks the block number provided as done. It returns the
// TokenHolders struct updated.
func (h *TokenHolders) Init(addr common.Address, ctype TokenType, block, chainID uint64) *TokenHolders {
	h.address = addr
	h.ctype = ctype
	h.holders = sync.Map{}
	h.blocks = sync.Map{}
	h.lastBlock.Store(block)
	h.synced.Store(false)
	h.ChainID = chainID
	return h
}

// Address function returns the given TokenHolders token address.
func (h *TokenHolders) Address() common.Address {
	return h.address
}

// Type function returns the given TokenHolders token type.
func (h *TokenHolders) Type() TokenType {
	return h.ctype
}

// Holders function returns the given TokenHolders current token holders
// addresses and its balances.
func (h *TokenHolders) Holders() HoldersCandidates {
	holders := HoldersCandidates{}
	h.holders.Range(func(rawAddr, rawBalance any) bool {
		address, okAddr := rawAddr.(common.Address)
		balance, okBalance := rawBalance.(*big.Int)
		if !okAddr || !okBalance {
			return true
		}

		holders[address] = balance
		return true
	})
	return holders
}

// Exists function returns if the given TokenHolders list of holders addresss
// includes the provided address.
func (h *TokenHolders) Exists(address common.Address) bool {
	_, exists := h.holders.Load(address)
	return exists
}

// Append function appends the holder address and the balance provided into the
// given TokenHolders list of holders. If the holder already exists, it will
// update its balance.
func (h *TokenHolders) Append(addr common.Address, balance *big.Int) {
	if currentBalance, exists := h.holders.Load(addr); exists {
		h.holders.Store(addr, new(big.Int).Add(currentBalance.(*big.Int), balance))
		return
	}
	h.holders.Store(addr, balance)
}

// Del function marks the holder address provided in the list of current
// TokenHolders as false, which means that it will be removed.
func (h *TokenHolders) Del(address common.Address) {
	h.holders.Store(address, false)
}

// FlushHolders function cleans the current list of token holders from the
// current TokenHolders state.
func (h *TokenHolders) FlushHolders() {
	h.holders = sync.Map{}
}

// BlockDone function checks the block number provided as checked appending it
// to the given TokenHolders list of blocks. If it is greater than the current
// TokenHolders block number, it will be updated.
func (h *TokenHolders) BlockDone(blockNumber uint64) {
	h.blocks.Store(blockNumber, true)
	h.synced.Store(false)
	h.lastBlock.CompareAndSwap(h.lastBlock.Load(), blockNumber)
}

// HasBlock function returns if the provided block number has already checked by
// the given TokenHolders.
func (h *TokenHolders) HasBlock(blockNumber uint64) bool {
	_, exists := h.blocks.Load(blockNumber)
	return exists
}

// LastBlock function returns the number of latest block registered.
func (h *TokenHolders) LastBlock() uint64 {
	return h.lastBlock.Load()
}

// Synced function marks the current TokenHolders struct as synced with the
// latest network status.
func (h *TokenHolders) Synced() {
	h.synced.Store(true)
}

// IsSynced function returns if the current TokenHolders instance is already
// synced with the latest network status.
func (h *TokenHolders) IsSynced() bool {
	return h.synced.Load()
}
