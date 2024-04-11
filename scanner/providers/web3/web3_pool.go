package web3

// This package contains the Web3Pool struct, which is a pool of Web3Endpoint
// instances. It allows to add, remove and get endpoints, as well as to get the
// chainID by short name. It also provides an implementation of the
// bind.ContractBackend interface for a web3 pool with an specific chainID.
// It allows to interact with the blockchain using the methods provided by the
// interface balancing the load between the available endpoints in the pool for
// every chainID.
// The pool balances the load between the available endpoints in the pool for
// every chainID, allowing to use the endpoints concurrently and switch between
// them flagging them as available if they fail to keep the pool healthy. If
// every endpoint fails for a chainID, the pool resets the available flag for
// all the endpoints and starts again.

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

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

// Web3Pool struct contains a map of chainID-[]*Web3Endpoint, where
// the key is the chainID and the value is a list of Web3Endpoint. It also
// contains a list of all the Web3Endpoint metadata. It provides methods to
// add, remove and get endpoints, as well as to get the chainID by short name.
// It allows to support multiple endpoints for the same chainID and switch
// between them looking for the available one.
type Web3Pool struct {
	nextAvailable sync.Map // chainID-int
	unavailable   sync.Map // chainID-[]int
	endpointsMtx  sync.RWMutex
	endpoints     map[uint64][]*Web3Endpoint
	metadata      []*Web3Endpoint
}

// NewWeb3Pool method returns a new *Web3Pool instance, initialized
// with the metadata from the external source. It returns an error if the metadata
// cannot be retrieved or decoded.
func NewWeb3Pool() (*Web3Pool, error) {
	// get chains information from external source
	res, err := http.Get(shortNameSourceUri)
	if err != nil {
		return nil, fmt.Errorf("error getting chains information from external source: %v", err)
	}
	chainsData := []*Web3Endpoint{}
	if err := json.NewDecoder(res.Body).Decode(&chainsData); err != nil {
		return nil, fmt.Errorf("error decoding chains information from external source: %v", err)
	}
	return &Web3Pool{
		endpoints: make(map[uint64][]*Web3Endpoint),
		metadata:  chainsData,
	}, nil
}

// AddEndpoint method adds a new web3 provider URI to the *Web3Pool
// instance. It returns an error if the chain metadata is not found or if the
// web3 client cannot be initialized.
func (nm *Web3Pool) AddEndpoint(uri string) error {
	ctx, cancel := context.WithTimeout(context.Background(), checkWeb3EndpointsTimeout)
	defer cancel()
	// init the web3 client
	client, err := connect(ctx, uri)
	if err != nil {
		return fmt.Errorf("error dialing web3 provider uri '%s': %w", uri, err)
	}
	// get the chainID from the web3 endpoint
	bChainID, err := client.ChainID(ctx)
	if err != nil {
		return fmt.Errorf("error getting the chainID from the web3 provider '%s': %w", uri, err)
	}
	chainID := bChainID.Uint64()
	// get chain name and the shortName
	var name, shortName string
	for _, info := range nm.metadata {
		if info.ChainID == chainID {
			name = info.Name
			shortName = info.ShortName
			break
		}
	}
	// check if the chain metadata was found, if not, return an error
	if name == "" || shortName == "" {
		return fmt.Errorf("no chain metadata found for chainID %d", chainID)
	}
	// add the endpoint to the chain manager
	nm.endpointsMtx.Lock()
	defer nm.endpointsMtx.Unlock()
	endpoint := &Web3Endpoint{
		ChainID:   chainID,
		Name:      name,
		ShortName: shortName,
		URI:       uri,
		client:    client,
	}
	if _, ok := nm.endpoints[chainID]; !ok {
		nm.endpoints[chainID] = []*Web3Endpoint{}
	}
	nm.endpoints[chainID] = append(nm.endpoints[chainID], endpoint)
	// set the next available endpoint to the last one added if there is no next
	// available endpoint for the chainID
	if _, ok := nm.nextAvailable.Load(chainID); !ok {
		nm.nextAvailable.Store(chainID, len(nm.endpoints[chainID])-1)
	}
	return nil
}

// DelEndpoint method removes a web3 provider URI from the *Web3Pool
// instance. It closes the client and removes the endpoint from the list of
// endpoints for the chainID where it was found.
func (nm *Web3Pool) DelEndoint(uri string) {
	nm.endpointsMtx.Lock()
	defer nm.endpointsMtx.Unlock()
	// remove the endpoint from the chain manager when the URI is found, closing
	// the client and removing the endpoint from the list of endpoints for the
	// chainID where it was found
	for chainID, endpoints := range nm.endpoints {
		for i, endpoint := range endpoints {
			if endpoint.URI == uri {
				endpoint.client.Close()
				nm.endpoints[chainID] = append(endpoints[:i], endpoints[i+1:]...)
				// if the endpoint removed was the next available, update it
				if next, ok := nm.nextAvailable.Load(chainID); ok && next.(int) == i {
					// if the endpoint is not the last in the poll, set the next
					// available to the previous one, otherwise, remove it
					if i > 0 {
						nm.nextAvailable.Store(chainID, i-1)
					} else {
						nm.nextAvailable.Delete(chainID)
					}
				}
			}
		}
	}
}

// GetEndpoint method returns the Web3Endpoint configured for the chainID
// provided. It returns the first available endpoint. If no available endpoint
// is found, it resets the available flag for all, resets the next available to
// the first one and returns it.
func (nm *Web3Pool) GetEndpoint(chainID uint64) (*Web3Endpoint, bool) {
	// Cases:
	//  - is there any available endpoint for the chainID?
	//    - yes, continue
	//    - no, reset the available flag for all the endpoints, return the first
	//      one and set the second one as the next available (if there is one)
	//  - do the next available endpoint exists?
	//    - yes, continue
	//    - no, return the first one and set the second one as the next
	//		available (if there is one)
	//  - update the next available endpoint to the next one
	//  - is the current endpoint available?
	//    - yes, return it
	//    - no, start again
	nm.endpointsMtx.RLock()
	defer nm.endpointsMtx.RUnlock()
	// check if there is any available endpoint for the chainID
	unavailable, ok := nm.unavailable.Load(chainID)
	if ok && len(unavailable.([]int)) == len(nm.endpoints[chainID]) {
		// if all the endpoints are unavailable, reset the available flag for
		// all the endpoints, set the second one as the next available (if
		// there) and return the first one
		nm.unavailable.Delete(chainID)
		if len(nm.endpoints[chainID]) > 1 {
			nm.nextAvailable.Store(chainID, 1)
		} else {
			nm.nextAvailable.Store(chainID, 0)
		}
		return nm.endpoints[chainID][0], true
	}
	// get the next available endpoint for the chainID
	currentEndpointIdx, ok := nm.nextAvailable.Load(chainID)
	if !ok {
		// if there is no next available endpoint, set the second one as the next
		// available (if there is one) and return the first one, if there is no
		// endpoint, return false
		if len(nm.endpoints[chainID]) == 0 {
			return nil, false
		}
		if len(nm.endpoints[chainID]) > 1 {
			nm.nextAvailable.Store(chainID, 1)
		} else {
			nm.nextAvailable.Store(chainID, 0)
		}
		return nm.endpoints[chainID][0], true
	}
	// update the next available endpoint to the next one
	nextAvailable := currentEndpointIdx.(int) + 1
	if nextAvailable >= len(nm.endpoints[chainID]) {
		nextAvailable = 0
	}
	nm.nextAvailable.Store(chainID, nextAvailable)
	// check if the current endpoint is available
	if unavailable != nil {
		for _, unavailableIdx := range unavailable.([]int) {
			if unavailableIdx == currentEndpointIdx.(int) {
				return nm.GetEndpoint(chainID)
			}
		}
	}
	// return the current endpoint
	return nm.endpoints[chainID][currentEndpointIdx.(int)], true
}

// DisableEndpoint method sets the available flag to false for the URI provided
// in the chainID provided.
func (nm *Web3Pool) DisableEndpoint(chainID uint64, uri string) {
	nm.endpointsMtx.RLock()
	defer nm.endpointsMtx.RUnlock()
	// get the chain endpoints
	chainEndpoints, ok := nm.endpoints[chainID]
	if !ok {
		return
	}
	// check if the endpoint is already unavailable
	if unavailable, ok := nm.unavailable.Load(chainID); ok {
		for _, unavailableIdx := range unavailable.([]int) {
			if chainEndpoints[unavailableIdx].URI == uri {
				return
			}
		}
	}
	// set the endpoint as unavailable
	for i, endpoint := range chainEndpoints {
		if endpoint.URI == uri {
			if unavailable, ok := nm.unavailable.Load(chainID); !ok {
				nm.unavailable.Store(chainID, []int{i})
			} else {
				nm.unavailable.Store(chainID, append(unavailable.([]int), i))
			}
			if next, ok := nm.nextAvailable.Load(chainID); ok && next.(int) == i {
				// if the endpoint will be the next available and it is the last
				// one, set the next available to the previous one, if there is
				// no previous one, remove it
				if i > 0 {
					nm.nextAvailable.Store(chainID, i-1)
				} else {
					nm.nextAvailable.Delete(chainID)
				}
				continue
			}
		}
	}
}

// GetClient method returns a new *Client instance for the chainID provided.
// It returns an error if the endpoint is not found.
func (nm *Web3Pool) GetClient(chainID uint64) (*Client, error) {
	if _, ok := nm.GetEndpoint(chainID); !ok {
		return nil, fmt.Errorf("error getting endpoint for chainID %d", chainID)
	}
	return &Client{w3p: nm, chainID: chainID}, nil
}

// EndpointByChainID method returns the Web3Endpoint configured for the
// chainID provided.
func (nm *Web3Pool) EndpointsByChainID(chainID uint64) ([]*Web3Endpoint, bool) {
	nm.endpointsMtx.RLock()
	defer nm.endpointsMtx.RUnlock()
	endpoints, ok := nm.endpoints[chainID]
	return endpoints, ok
}

// URIByChainID method returns the URI configured for the chainID provided.
func (nm *Web3Pool) URIsByChainID(chainID uint64) ([]string, bool) {
	nm.endpointsMtx.RLock()
	defer nm.endpointsMtx.RUnlock()
	endpoints, ok := nm.endpoints[chainID]
	if !ok {
		return nil, false
	}
	var uris []string
	for _, endpoint := range endpoints {
		uris = append(uris, endpoint.URI)
	}
	return uris, true
}

// ChainIDByShortName method returns the chainID configured for the networkEndpoint
// short name provided.
func (nm *Web3Pool) ChainIDByShortName(shortName string) (uint64, bool) {
	for _, endpoint := range nm.metadata {
		if endpoint.ShortName == shortName {
			return endpoint.ChainID, true
		}
	}
	return 0, false
}

// ChainAddress method returns a prefixed string of the hex address provided,
// with the short name of the networkEndpoint identified by the chain id provided.
// Read more here: https://eips.ethereum.org/EIPS/eip-3770
func (nps *Web3Pool) ChainAddress(chainID uint64, hexAddress string) (string, bool) {
	for _, data := range nps.metadata {
		if data.ChainID == chainID {
			return fmt.Sprintf("%s:%s", data.ShortName, hexAddress), true
		}
	}
	return "", false
}

// String method returns a string representation of the *Web3Pool list.
func (nm *Web3Pool) String() string {
	nm.endpointsMtx.RLock()
	defer nm.endpointsMtx.RUnlock()
	shortNames := map[string]bool{}
	for _, endpoint := range nm.endpoints {
		for _, ep := range endpoint {
			shortNames[ep.ShortName] = true
		}
	}
	var shortNamesSlice []string
	for shortName := range shortNames {
		shortNamesSlice = append(shortNamesSlice, shortName)
	}
	return fmt.Sprintf("%v", shortNamesSlice)
}

// CurrentBlockNumbers method returns a map of uint64-uint64, where the key is
// the chainID and the value is the current block number of the network.
func (nm *Web3Pool) CurrentBlockNumbers(ctx context.Context) (map[uint64]uint64, error) {
	nm.endpointsMtx.RLock()
	defer nm.endpointsMtx.RUnlock()
	blockNumbers := make(map[uint64]uint64)
	for chainID := range nm.endpoints {
		cli, ok := nm.GetEndpoint(chainID)
		if !ok {
			return nil, fmt.Errorf("error getting endpoint for chainID %d", chainID)
		}
		blockNumber, err := cli.client.BlockNumber(ctx)
		if err != nil {
			return nil, fmt.Errorf("error getting block number for chainID %d: %w", chainID, err)
		}
		blockNumbers[chainID] = blockNumber
	}
	return blockNumbers, nil
}

// SupportedNetworks method returns a list of all the supported Web3Endpoint
// metadata. It returns the chainID, name and shortName of unique supported
// chains.
func (nm *Web3Pool) SupportedNetworks() []*Web3Endpoint {
	nm.endpointsMtx.RLock()
	defer nm.endpointsMtx.RUnlock()
	var supported []*Web3Endpoint
	for _, endpoints := range nm.endpoints {
		supported = append(supported, &Web3Endpoint{
			ChainID:   endpoints[0].ChainID,
			Name:      endpoints[0].Name,
			ShortName: endpoints[0].ShortName,
		})
	}
	return supported
}

// connect method returns a new *ethclient.Client instance for the URI provided.
// It retries to connect to the web3 provider if it fails, up to the
// DefaultMaxWeb3ClientRetries times.
func connect(ctx context.Context, uri string) (client *ethclient.Client, err error) {
	for i := 0; i < DefaultMaxWeb3ClientRetries; i++ {
		if client, err = ethclient.DialContext(ctx, uri); err != nil {
			continue
		}
		return
	}
	return nil, fmt.Errorf("error dialing web3 provider uri '%s': %w", uri, err)
}
