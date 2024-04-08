package web3

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/ethereum/go-ethereum/ethclient"
	"go.vocdoni.io/dvote/log"
)

// NetworkEndpoint struct contains all the required information about a web3
// provider based on its URI. It includes its chain ID, its name (and shortName)
// and the URI.
type NetworkEndpoint struct {
	ChainID   uint64 `json:"chainId"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
	URI       string
	client    *ethclient.Client
	available bool
}

// Client method returns the *ethclient.Client configured for the current
// NetworkEndpoint.
func (e *NetworkEndpoint) Client() *ethclient.Client {
	return e.client
}

// NetworksManager struct contains a map of chainID-[]*NetworkEndpoint, where
// the key is the chainID and the value is a list of NetworkEndpoint. It also
// contains a list of all the NetworkEndpoint metadata. It provides methods to
// add, remove and get endpoints, as well as to get the chainID by short name.
// It allows to support multiple endpoints for the same chainID and switch
// between them looking for the available one.
type NetworksManager struct {
	mtx      sync.RWMutex
	networks map[uint64][]*NetworkEndpoint
	metadata []*NetworkEndpoint
}

// NewNetworksManager method returns a new *NetworksManager instance, initialized
// with the metadata from the external source. It returns an error if the metadata
// cannot be retrieved or decoded.
func NewNetworksManager() (*NetworksManager, error) {
	// get chains information from external source
	res, err := http.Get(shortNameSourceUri)
	if err != nil {
		return nil, fmt.Errorf("error getting chains information from external source: %v", err)
	}
	chainsData := []*NetworkEndpoint{}
	if err := json.NewDecoder(res.Body).Decode(&chainsData); err != nil {
		return nil, fmt.Errorf("error decoding chains information from external source: %v", err)
	}
	return &NetworksManager{networks: make(map[uint64][]*NetworkEndpoint), metadata: chainsData}, nil
}

// AddEndpoint method adds a new web3 provider URI to the *NetworksManager
// instance. It returns an error if the chain metadata is not found or if the
// web3 client cannot be initialized.
func (nm *NetworksManager) AddEndpoint(uri string) error {
	ctx, cancel := context.WithTimeout(context.Background(), checkNetworkEndpointsTimeout)
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
	nm.mtx.Lock()
	defer nm.mtx.Unlock()
	if _, ok := nm.networks[chainID]; !ok {
		nm.networks[chainID] = []*NetworkEndpoint{}
	}
	nm.networks[chainID] = append(nm.networks[chainID], &NetworkEndpoint{
		ChainID:   chainID,
		Name:      name,
		ShortName: shortName,
		URI:       uri,
		client:    client,
		available: true,
	})
	log.Infow("new web3 uri added", "chainID", chainID, "name", name, "shortName", shortName)
	return nil
}

// DelEndpoint method removes a web3 provider URI from the *NetworksManager
// instance. It closes the client and removes the endpoint from the list of
// endpoints for the chainID where it was found.
func (nm *NetworksManager) DelEndoint(uri string) {
	nm.mtx.Lock()
	defer nm.mtx.Unlock()
	// remove the endpoint from the chain manager when the URI is found, closing
	// the client and removing the endpoint from the list of endpoints for the
	// chainID where it was found
	for chainID, endpoints := range nm.networks {
		for i, endpoint := range endpoints {
			if endpoint.URI == uri {
				endpoint.client.Close()
				nm.networks[chainID] = append(endpoints[:i], endpoints[i+1:]...)
			}
		}
	}
}

// GetEndpoint method returns the NetworkEndpoint configured for the chainID
// provided. It returns the first available endpoint and sets its available
// flag to false. If no available endpoint is found, it resets the available
// flag for all and returns the first one.
func (nm *NetworksManager) GetEndpoint(chainID uint64) (*NetworkEndpoint, bool) {
	nm.mtx.RLock()
	defer nm.mtx.RUnlock()
	// get the endpoints for the chainID provided
	endpoints, ok := nm.networks[chainID]
	if !ok {
		return nil, false
	}
	// check if there is an available endpoint and return it if found, after
	// setting its available flag to false
	for i, endpoint := range endpoints {
		if endpoint.available {
			nm.networks[chainID][i].available = false
			return endpoint, true
		}
	}
	// if no available endpoint is found, reset the available flag for all and
	// return the first one
	for i := range endpoints {
		nm.networks[chainID][i].available = true
	}
	return endpoints[0], true
}

// EndpointByChainID method returns the NetworkEndpoint configured for the
// chainID provided.
func (nm *NetworksManager) EndpointsByChainID(chainID uint64) ([]*NetworkEndpoint, bool) {
	nm.mtx.RLock()
	defer nm.mtx.RUnlock()
	endpoints, ok := nm.networks[chainID]
	return endpoints, ok
}

// URIByChainID method returns the URI configured for the chainID provided.
func (nm *NetworksManager) URIsByChainID(chainID uint64) ([]string, bool) {
	endpoints, ok := nm.networks[chainID]
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
func (nm *NetworksManager) ChainIDByShortName(shortName string) (uint64, bool) {
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
func (nps *NetworksManager) ChainAddress(chainID uint64, hexAddress string) (string, bool) {
	for _, data := range nps.metadata {
		if data.ChainID == chainID {
			return fmt.Sprintf("%s:%s", data.ShortName, hexAddress), true
		}
	}
	return "", false
}

// String method returns a string representation of the *NetworksManager list.
func (nm *NetworksManager) String() string {
	shortNames := map[string]bool{}
	for _, endpoint := range nm.networks {
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
func (nm *NetworksManager) CurrentBlockNumbers(ctx context.Context) (map[uint64]uint64, error) {
	blockNumbers := make(map[uint64]uint64)
	for chainID := range nm.networks {
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

// SupportedNetworks method returns a list of all the supported NetworkEndpoint
// metadata. It returns the chainID, name and shortName of unique supported
// chains.
func (nm *NetworksManager) SupportedNetworks() []*NetworkEndpoint {
	nm.mtx.RLock()
	defer nm.mtx.RUnlock()
	var supported []*NetworkEndpoint
	for _, endpoints := range nm.networks {
		supported = append(supported, &NetworkEndpoint{
			ChainID:   endpoints[0].ChainID,
			Name:      endpoints[0].Name,
			ShortName: endpoints[0].ShortName,
		})
	}
	return supported
}

func connect(ctx context.Context, uri string) (client *ethclient.Client, err error) {
	for i := 0; i < DefaultMaxWeb3ClientRetries; i++ {
		if client, err = ethclient.DialContext(ctx, uri); err != nil {
			log.Warnf("error dialing web3 provider, retrying... %v", err)
			continue
		}
		return
	}
	return nil, fmt.Errorf("error dialing web3 provider uri '%s': %w", uri, err)
}
