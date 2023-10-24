package service

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/census3/db"
)

// ExternalProvider is the interface that must be implemented by any external
// provider that wants to be used by the census service. An external provider
// is an external service that provides information about the holders of a
// token. The census3 service uses this information to create the census.
type ExternalProvider interface {
	// Init initializes the external provider with the database provided.
	Init(db *db.DB) error
	// GetHolders returns the holders of the token with the IDs provided.
	// It receives an external ID to be used to get holders and their balances.
	// It must return a map with the address of the holder as key and the
	// balance of the token holder as value.
	GetHolders(externalID []byte) (map[common.Address]*big.Int, error)
}
