package web3

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

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
	URIs      []string
}

// GetClient method returns a web3 client for the first URI that can be dialed.
func (n *NetworkEndpoint) GetClient(maxRetries int) (*ethclient.Client, error) {
	for try := 0; try < maxRetries; try++ {
		for _, uri := range n.URIs {
			if cli, err := ethclient.Dial(uri); err == nil {
				return cli, nil
			}
		}
	}
	return nil, fmt.Errorf("error dialing web3 provider uris")
}

// GetChainIDByURI function returns the chainID of the web3 provider URI
// provided. It dials the URI and gets the chainID from the web3 endpoint,
// using the context provided and the GetClient method with the
// DefaultMaxRetries value.
func GetChainIDByURI(ctx context.Context, uri string) (uint64, error) {
	n := &NetworkEndpoint{URIs: []string{uri}}
	cli, err := n.GetClient(DefaultMaxRetries)
	if err != nil {
		return 0, err
	}
	defer cli.Close()
	ctx, cancel := context.WithTimeout(ctx, checkNetworkEndpointsTimeout)
	defer cancel()
	chainID, err := cli.ChainID(ctx)
	if err != nil {
		return 0, err
	}
	n.ChainID = chainID.Uint64()
	return n.ChainID, nil
}

// NetworkEndpoints type envolves a map of uint64-NetworkEndpoint, used to index the
// configured web3 providers by the chainID.
type NetworkEndpoints map[uint64]*NetworkEndpoint

// EndpointByChainID method returns the NetworkEndpoint configured for the
// chainID provided.
func (nps NetworkEndpoints) EndpointByChainID(chainID uint64) (*NetworkEndpoint, bool) {
	provider, ok := nps[chainID]
	return provider, ok
}

// URIByChainID method returns the URI configured for the chainID provided.
func (nps NetworkEndpoints) URIsByChainID(chainID uint64) ([]string, bool) {
	provider, ok := nps[chainID]
	if !ok {
		return nil, false
	}
	return provider.URIs, true
}

// ChainIDByShortName method returns the chainID configured for the networkEndpoint
// short name provided.
func (nps NetworkEndpoints) ChainIDByShortName(shortName string) (uint64, bool) {
	for _, provider := range nps {
		if provider.ShortName == shortName {
			return provider.ChainID, true
		}
	}
	return 0, false
}

// ChainAddress method returns a prefixed string of the hex address provided,
// with the short name of the networkEndpoint identified by the chain id provided.
// Read more here: https://eips.ethereum.org/EIPS/eip-3770
func (nps NetworkEndpoints) ChainAddress(chainID uint64, hexAddress string) (string, bool) {
	provider, ok := nps[chainID]
	if !ok {
		return "", false
	}
	return fmt.Sprintf("%s:%s", provider.ShortName, hexAddress), true
}

// String method returns a string representation of the NetworkEndpoints list.
func (nps NetworkEndpoints) String() string {
	var shortNames []string
	for _, provider := range nps {
		shortNames = append(shortNames, provider.ShortName)
	}
	return fmt.Sprintf("%v", shortNames)
}

// CurrentBlockNumbers method returns a map of uint64-uint64, where the key is
// the chainID and the value is the current block number of the network.
func (nps NetworkEndpoints) CurrentBlockNumbers(ctx context.Context) (map[uint64]uint64, error) {
	blockNumbers := make(map[uint64]uint64)
	for _, endpoint := range nps {
		cli, err := endpoint.GetClient(DefaultMaxRetries)
		if err != nil {
			return blockNumbers, err
		}
		blockNumber, err := cli.BlockNumber(ctx)
		if err != nil {
			return blockNumbers, fmt.Errorf("error getting the block number from %s network: %w", endpoint.Name, err)
		}
		blockNumbers[endpoint.ChainID] = blockNumber
	}
	return blockNumbers, nil
}

// InitNetworkEndpoints function initializes a NetworkEndpoints list checking
// the web3 enpoint URI's provided as argument. It checks if the URI's are
// valid, getting its chain ID's and then query to shortNameSourceURI endpoint
// to get the chain name and short name. If more than one URI is provided for
// the same chainID, the URIs are grouped in the same NetworkEndpoint. If no
// valid URIs are provided, an error is returned.
func InitNetworkEndpoints(providersURIs []string) (NetworkEndpoints, error) {
	if len(providersURIs) == 0 {
		return nil, fmt.Errorf("no URIs provided")
	}
	// get chains information from external source
	res, err := http.Get(shortNameSourceUri)
	if err != nil {
		return nil, fmt.Errorf("error getting chains information from external source: %v", err)
	}
	chainsData := []*NetworkEndpoint{}
	if err := json.NewDecoder(res.Body).Decode(&chainsData); err != nil {
		return nil, fmt.Errorf("error decoding chains information from external source: %v", err)
	}
	providers := make(NetworkEndpoints)
	for _, uri := range providersURIs {
		ctx, cancel := context.WithTimeout(context.Background(), checkNetworkEndpointsTimeout)
		defer cancel()

		cli, err := ethclient.DialContext(ctx, uri)
		if err != nil {
			log.Errorf("error dialing web3 provider uri '%s': %v", uri, err)
			continue
		}
		// get the chainID from the web3 endpoint
		bChainID, err := cli.ChainID(ctx)
		if err != nil {
			log.Errorf("error getting the chainID from the web3 provider '%s': %v", uri, err)
			continue
		}
		chainID := bChainID.Uint64()
		if provider, ok := providers[chainID]; ok {
			provider.URIs = append(provider.URIs, uri)
			providers[chainID] = provider
			continue
		}
		// get chain shortName
		for _, info := range chainsData {
			if info.ChainID == chainID {
				info.URIs = []string{uri}
				providers[info.ChainID] = info
				break
			}
		}
	}
	if len(providers) == 0 {
		return nil, fmt.Errorf("no valid URIs provided")
	}
	return providers, nil
}

// TestNetworkEndpoint function returns a NetworkEndpoint for testing purposes.
// It checks if the WEB3_URI environment variable is set, and if so, it uses
// its value to initialize a NetworkEndpoint. If not, it returns the
// DefaultNetworkEndpoint.
func TestNetworkEndpoint() (*NetworkEndpoint, error) {
	if uri := os.Getenv("WEB3_URI"); uri != "" {
		endpoints, err := InitNetworkEndpoints([]string{uri})
		if err != nil {
			return nil, err
		}
		var chainID uint64
		if chainID, err = GetChainIDByURI(context.Background(), uri); err != nil {
			return nil, err
		}
		return endpoints[chainID], nil
	}
	return DefaultNetworkEndpoint, nil
}
