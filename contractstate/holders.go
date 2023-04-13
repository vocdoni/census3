package contractstate

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"go.vocdoni.io/dvote/log"
)

type ContractHolders struct {
	address common.Address
	ctype   ContractType
	holders []common.Address
	mutex   sync.RWMutex
}

func (h *ContractHolders) Init(address common.Address, ctype ContractType) {
	log.Infof("initializing contract %s of type %d", address, ctype)
	h.address = address
	h.ctype = ctype
	h.holders = make([]common.Address, 0)
	h.mutex = sync.RWMutex{}
}

func (h *ContractHolders) Address() common.Address {
	return h.address
}

func (h *ContractHolders) Type() ContractType {
	return h.ctype
}

func (h *ContractHolders) Holders() []common.Address {
	h.mutex.RLock()
	defer h.mutex.RUnlock()
	holders := make([]common.Address, len(h.holders))
	copy(holders, h.holders)
	return holders
}

func (h *ContractHolders) Exists(address common.Address) bool {
	h.mutex.RLock()
	defer h.mutex.RUnlock()
	for _, currentHolder := range h.holders {
		if currentHolder.String() == address.String() {
			return true
		}
	}
	return false
}

func (h *ContractHolders) Append(candidates ...common.Address) {
	for _, address := range candidates {
		h.mutex.RLock()
		for _, currentHolder := range h.holders {
			if currentHolder.String() == address.String() {
				h.mutex.Unlock()
				return
			}
		}
		h.mutex.Unlock()

		h.mutex.Lock()
		h.holders = append(h.holders, address)
		h.mutex.Unlock()
	}
}

func (h *ContractHolders) Del(address common.Address) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	for idx, currentHolder := range h.holders {
		if currentHolder.String() == address.String() {
			h.holders = append(h.holders[:idx], h.holders[idx+1:]...)
			return
		}
	}
	return
}
