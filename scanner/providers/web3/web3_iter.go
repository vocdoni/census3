package web3

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/ethclient"
)

// Web3Endpoint struct contains all the required information about a web3
// provider based on its URI. It includes its chain ID, its name (and shortName)
// and the URI.
type Web3Endpoint struct {
	ChainID   uint64 `json:"chainId"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
	URI       string
	client    *ethclient.Client
}

// Web3Iterator struct is a pool of Web3Endpoint that allows to get the next
// available endpoint in a round-robin fashion. It also allows to disable an
// endpoint if it fails. It allows to manage multiple endpoints safely.
type Web3Iterator struct {
	nextIndex atomic.Uint32
	available []*Web3Endpoint
	disabled  []*Web3Endpoint
	mtx       sync.Mutex
}

// NewWeb3Iterator creates a new Web3Iterator with the given endpoints.
func NewWeb3Iterator(endpoints ...*Web3Endpoint) *Web3Iterator {
	if endpoints == nil {
		endpoints = make([]*Web3Endpoint, 0)
	}
	return &Web3Iterator{
		available: endpoints,
		disabled:  make([]*Web3Endpoint, 0),
	}
}

// Add adds a new endpoint to the pool, making it available for the next
// requests.
func (w3pp *Web3Iterator) Add(endpoint *Web3Endpoint) {
	w3pp.mtx.Lock()
	defer w3pp.mtx.Unlock()
	w3pp.available = append(w3pp.available, endpoint)
}

// Next returns the next available endpoint in a round-robin fashion. If
// there are no endpoints, it will return an error. If there are no available
// endpoints, it will reset the disabled endpoints and return the first
// available endpoint.
func (w3pp *Web3Iterator) Next() (*Web3Endpoint, error) {
	w3pp.mtx.Lock()
	defer w3pp.mtx.Unlock()
	l := uint32(len(w3pp.available))
	if l == 0 {
		return nil, fmt.Errorf("no available endpoints")
	}
	// get the current next index and endpoint
	currentIndex := w3pp.nextIndex.Load()
	if currentIndex >= l {
		// if the current index is out of bounds, reset it to the first one
		currentIndex = 0
	}
	currentEndpoint := w3pp.available[currentIndex]
	if currentEndpoint == nil {
		// if the current endpoint is nil, reset the index and get the first one
		currentIndex = 0
		currentEndpoint = w3pp.available[0]
	}
	// calculate the following next endpoint index based on the current one
	nextIndex := currentIndex + 1
	if nextIndex >= l {
		nextIndex = 0
	}
	// update the next index and return the current endpoint
	w3pp.nextIndex.Store(nextIndex)
	return currentEndpoint, nil
}

// Disable method disables an endpoint, moving it from the available list to the
// the disabled list.
func (w3pp *Web3Iterator) Disable(uri string) {
	w3pp.mtx.Lock()
	defer w3pp.mtx.Unlock()
	// remove the endpoint from the available list
	for i, e := range w3pp.available {
		if e.URI == uri {
			w3pp.available = append(w3pp.available[:i], w3pp.available[i+1:]...)
			w3pp.disabled = append(w3pp.disabled, e)
		}
	}
	// if there are no available endpoints, reset all the disabled ones to
	// available ones and reset the next index to the first one
	if l := len(w3pp.available); l == 0 {
		// reset the next index and move the disabled endpoints to the available
		w3pp.nextIndex.Store(0)
		w3pp.available = append(w3pp.available, w3pp.disabled...)
		w3pp.disabled = make([]*Web3Endpoint, 0)
	}
}
