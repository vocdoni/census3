package web3

import (
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

// Web3EndpointPool struct is a pool of Web3Endpoint that allows to get the next
// available endpoint in a round-robin fashion. It also allows to disable an
// endpoint if it fails. It allows to manage multiple endpoints safely.
type Web3EndpointPool struct {
	nextIndex atomic.Uint32
	available []*Web3Endpoint
	disabled  []*Web3Endpoint
	mtx       sync.Mutex
}

// NewWeb3EndpointPool creates a new Web3EndpointPool with the given endpoints.
func newWeb3EndpointPool(endpoints ...*Web3Endpoint) *Web3EndpointPool {
	if endpoints == nil {
		endpoints = make([]*Web3Endpoint, 0)
	}
	return &Web3EndpointPool{
		available: endpoints,
		disabled:  make([]*Web3Endpoint, 0),
	}
}

// Add adds a new endpoint to the pool, making it available for the next
// requests.
func (w3pp *Web3EndpointPool) add(endpoint *Web3Endpoint) {
	w3pp.mtx.Lock()
	defer w3pp.mtx.Unlock()
	w3pp.available = append(w3pp.available, endpoint)
}

// Next returns the next available endpoint in a round-robin fashion. If there
// are no endpoints, it will return nil. If there are no available endpoints, it
// will reset the disabled endpoints and return the first available endpoint.
func (w3pp *Web3EndpointPool) next() *Web3Endpoint {
	w3pp.mtx.Lock()
	defer w3pp.mtx.Unlock()
	// check if there is any available endpoint
	l := len(w3pp.available)
	if l == 0 {
		// reset the next index and move the disabled endpoints to the available
		w3pp.nextIndex.Store(0)
		w3pp.available = append(w3pp.available, w3pp.disabled...)
		w3pp.disabled = make([]*Web3Endpoint, 0)
		// if continue to have no available endpoints, return nil
		if len(w3pp.available) == 0 {
			return nil
		}
		return w3pp.available[0]
	}
	// get the current next index and endpoint
	currentIndex := w3pp.nextIndex.Load()
	if int(currentIndex) >= l {
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
	if int(nextIndex) >= l {
		nextIndex = 0
	}
	// update the next index and return the current endpoint
	w3pp.nextIndex.Store(nextIndex)
	return currentEndpoint
}

// disable method disables an endpoint, moving it from the available list to the
// the disabled list.
func (w3pp *Web3EndpointPool) disable(uri string) {
	w3pp.mtx.Lock()
	defer w3pp.mtx.Unlock()
	// remove the endpoint from the available list
	for i, e := range w3pp.available {
		if e.URI == uri {
			w3pp.available = append(w3pp.available[:i], w3pp.available[i+1:]...)
			w3pp.disabled = append(w3pp.disabled, e)
		}
	}
}
