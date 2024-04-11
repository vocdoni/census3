package manager

// package manager provides a manager for providers of different types
// and a way to add and get them by type concurrently safe. It initializes a new
// provider based on the type and the configuration provided every time that a
// provider is requested to avoid data races. It also provides a way to get all
// the provider types and all the providers initialized at once.

import (
	"fmt"
	"sync"

	"github.com/vocdoni/census3/scanner/providers"
	"github.com/vocdoni/census3/scanner/providers/farcaster"
	"github.com/vocdoni/census3/scanner/providers/gitcoin"
	"github.com/vocdoni/census3/scanner/providers/poap"
	"github.com/vocdoni/census3/scanner/providers/web3"
)

type ProviderManager struct {
	confs sync.Map
}

// NewProviderManager creates a new provider manager
func NewProviderManager() *ProviderManager {
	return &ProviderManager{}
}

// AddProvider adds a new provider configuration to the manager assigned to the
// specific type provided
func (m *ProviderManager) AddProvider(providerType uint64, conf any) {
	m.confs.Store(providerType, conf)
}

// GetProvider returns a provider based on the type provided. It initializes the
// provider based on the configuration stored in the manager. It initializes a
// new provider every time to avoid data races. It returns an error if the
// provider type is not found or if the provider cannot be initialized.
func (m *ProviderManager) GetProvider(providerType uint64) (providers.HolderProvider, error) {
	// load the configuration for the provider type
	conf, ok := m.confs.Load(providerType)
	if !ok {
		return nil, fmt.Errorf("provider type %d not found", providerType)
	}
	// initialize the provider based on the type
	var provider providers.HolderProvider
	switch providerType {
	case providers.CONTRACT_TYPE_ERC20:
		provider = &web3.ERC20HolderProvider{}
	case providers.CONTRACT_TYPE_ERC721:
		provider = &web3.ERC721HolderProvider{}
	case providers.CONTRACT_TYPE_ERC777:
		provider = &web3.ERC777HolderProvider{}
	case providers.CONTRACT_TYPE_POAP:
		provider = &poap.POAPHolderProvider{}
	case providers.CONTRACT_TYPE_GITCOIN:
		provider = &gitcoin.GitcoinPassport{}
	case providers.CONTRACT_TYPE_FARCASTER:
		provider = &farcaster.FarcasterProvider{}
	default:
		return nil, fmt.Errorf("provider type %d not found", providerType)
	}
	// initialize the provider with the specific configuration
	if err := provider.Init(conf); err != nil {
		return nil, err
	}
	return provider, nil
}

// GetProviderTypes returns all the provider types stored in the manager as a
// slice of uint64.
func (m *ProviderManager) GetProviderTypes() []uint64 {
	types := []uint64{}
	m.confs.Range(func(t, _ any) bool {
		types = append(types, t.(uint64))
		return true
	})
	return types
}

// Providers returns all the providers stored in the manager associated to their
// types as a map of uint64 to HolderProvider.
func (m *ProviderManager) Providers() map[uint64]providers.HolderProvider {
	providers := make(map[uint64]providers.HolderProvider)
	for _, t := range m.GetProviderTypes() {
		provider, err := m.GetProvider(t)
		if err != nil {
			panic(err)
		}
		providers[t] = provider
	}
	return providers
}
