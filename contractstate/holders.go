package contractstate

import (
	"sync"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/common"
)

type TokenHolders struct {
	address     common.Address
	ctype       ContractType
	holders     sync.Map
	blocks      sync.Map
	blockNumber atomic.Uint64
}

func (h *TokenHolders) Init(address common.Address, ctype ContractType, blockNumber uint64) *TokenHolders {
	h.address = address
	h.ctype = ctype
	h.holders = sync.Map{}
	h.blocks = sync.Map{}
	h.blockNumber.Store(blockNumber)
	return h
}

func (h *TokenHolders) Address() common.Address {
	return h.address
}

func (h *TokenHolders) Type() ContractType {
	return h.ctype
}

func (h *TokenHolders) Holders() []common.Address {
	holders := make([]common.Address, 0)
	h.holders.Range(func(address, _ any) bool {
		holders = append(holders, address.(common.Address))
		return true
	})
	return holders
}

func (h *TokenHolders) Exists(address common.Address) bool {
	_, exists := h.holders.Load(address)
	return exists
}

func (h *TokenHolders) Append(candidates ...common.Address) {
	for _, address := range candidates {
		h.holders.Store(address, nil)
	}
}

func (h *TokenHolders) Del(address common.Address) {
	h.holders.Delete(address)
}

func (h *TokenHolders) BlockDone(blockNumber uint64) {
	h.blocks.Store(blockNumber, true)
	h.blockNumber.Store(blockNumber)
}

func (h *TokenHolders) HasBlock(blockNumber uint64) bool {
	_, exists := h.blocks.Load(blockNumber)
	return exists
}

func (h *TokenHolders) Block() uint64 {
	return h.blockNumber.Load()
}
