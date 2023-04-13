package contractstate

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"go.vocdoni.io/dvote/log"
)

type TokenHolders struct {
	address common.Address
	ctype   ContractType
	holders sync.Map
}

func (h *TokenHolders) Init(address common.Address, ctype ContractType) {
	log.Infof("initializing contract %s of type %d", address, ctype)
	h.address = address
	h.ctype = ctype
	h.holders = sync.Map{}
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
