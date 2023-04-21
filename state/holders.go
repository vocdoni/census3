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
// variables used to store the state of the token holders safely accross
// differents goroutines, such as holders list, analyzed blocks or the last
// block analyzed.
type TokenHolders struct {
	address   common.Address
	ctype     TokenType
	holders   sync.Map
	blocks    sync.Map
	lastBlock atomic.Uint64
}

// Init function fills the given TokenHolders struct with the address and type
// given, also checks the block number provided as done. It returns the
// TokenHolders struct updated.
func (h *TokenHolders) Init(addr common.Address, ctype TokenType, block uint64) *TokenHolders {
	h.address = addr
	h.ctype = ctype
	h.holders = sync.Map{}
	h.blocks = sync.Map{}
	h.lastBlock.Store(block)
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
// addresses.
func (h *TokenHolders) Holders() []common.Address {
	holders := make([]common.Address, 0)
	h.holders.Range(func(address, _ any) bool {
		holders = append(holders, address.(common.Address))
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

// Append function appends the holder address provided into the given
// TokenHolders list of holders addresss. It register the address to a boolean
// setted to true which means that is a new token holder.
func (h *TokenHolders) Append(candidates ...common.Address) {
	for _, address := range candidates {
		h.holders.Store(address, true)
	}
}

// Del function marks the holder address provided in the list of current
// TokenHolders as false, which means that it will be removed.
func (h *TokenHolders) Del(address common.Address) {
	h.holders.Store(address, false)
}

// HoldersToCreate returns the address of the new token holders. These addresses
// is marked with true in the current TokenHolders state list.
func (h *TokenHolders) HoldersToCreate() []common.Address {
	holdersToCreate := make([]common.Address, 0)
	h.holders.Range(func(address, value any) bool {
		if isNew := value.(bool); isNew {
			holdersToCreate = append(holdersToCreate, address.(common.Address))
		}
		return true
	})
	return holdersToCreate
}

// HoldersToCreate returns the address of token holders to delete. These
// addresses is marked with false in the current TokenHolders state list.
func (h *TokenHolders) HoldersToDelete() []common.Address {
	holdersToDelete := make([]common.Address, 0)
	h.holders.Range(func(address, value any) bool {
		if toDelete := !value.(bool); toDelete {
			holdersToDelete = append(holdersToDelete, address.(common.Address))
		}
		return true
	})
	return holdersToDelete
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
