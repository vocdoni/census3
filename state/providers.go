package state

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
)

// Web3Provider struct contains all the required information about a web3
// provider based on its URI. It includes its chain ID, its name (and shortName)
// and the URI.
type Web3Provider struct {
	ChainID   uint64 `json:"chainId"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
	URI       string
}

// Web3Providers type envolves a map of uint64-Web3Provider, used to index the
// configured web3 providers by the chainID.
type Web3Providers map[uint64]Web3Provider

// URIByChainID method returns the URI configured for the chainID provided.
func (w3p Web3Providers) URIByChainID(chainID uint64) (string, bool) {
	provider, ok := w3p[chainID]
	if !ok {
		return "", false
	}
	return provider.URI, true
}

// ChainIDByShortName method returns the chainID configured for the network
// short name provided.
func (w3p Web3Providers) ChainIDByShortName(shortName string) (uint64, bool) {
	for _, provider := range w3p {
		if provider.ShortName == shortName {
			return provider.ChainID, true
		}
	}
	return 0, false
}

// ChainAddress method returns a prefixed string of the hex address provided,
// with the short name of the network identified by the chain id provided.
// Read more here: https://eips.ethereum.org/EIPS/eip-3770
func (w3p Web3Providers) ChainAddress(chainID uint64, hexAddress string) (string, bool) {
	provider, ok := w3p[chainID]
	if !ok {
		return "", false
	}
	return fmt.Sprintf("%s:%s", provider.ShortName, hexAddress), true
}

// CurrentBlockNumbers method returns a map of uint64-uint64, where the key is
// the chainID and the value is the current block number of the network.
func (w3p Web3Providers) CurrentBlockNumbers(ctx context.Context) (map[uint64]uint64, error) {
	blockNumbers := make(map[uint64]uint64)
	for _, provider := range w3p {
		if _, ok := blockNumbers[provider.ChainID]; ok {
			continue
		}
		cli, err := ethclient.Dial(provider.URI)
		if err != nil {
			return nil, fmt.Errorf("error dialing web3 provider uri '%s': %w", provider.URI, err)
		}

		blockNumber, err := cli.BlockNumber(ctx)
		if err != nil {
			return nil, fmt.Errorf("error getting the block number from %s network: %w", provider.Name, err)
		}
		blockNumbers[provider.ChainID] = blockNumber
	}
	return blockNumbers, nil
}

// CheckWeb3Providers function initializes a Web3Providers list checking the
// web3 enpoint URI's provided as argument. It checks if the URI's are valid,
// getting its chain ID's and then query to shortNameSourceURI endpoint to get
// the chain name and short name.
func CheckWeb3Providers(providersURIs []string) (Web3Providers, error) {
	if len(providersURIs) == 0 {
		return nil, fmt.Errorf("no URIs provided")
	}
	// get chains information from external source
	res, err := http.Get(shortNameSourceUri)
	if err != nil {
		return nil, fmt.Errorf("error getting chains information from external source: %v", err)
	}
	chainsData := []Web3Provider{}
	if err := json.NewDecoder(res.Body).Decode(&chainsData); err != nil {
		return nil, fmt.Errorf("error decoding chains information from external source: %v", err)
	}
	providers := make(Web3Providers)
	for _, uri := range providersURIs {
		cli, err := ethclient.Dial(uri)
		if err != nil {
			return nil, fmt.Errorf("error dialing web3 provider uri '%s': %w", uri, err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), checkWeb3ProvidersTimeout)
		defer cancel()
		// get the chainID from the web3 endpoint
		chainID, err := cli.ChainID(ctx)
		if err != nil {
			return nil, fmt.Errorf("error getting the chainID from the web3 provider '%s': %w", uri, err)
		}
		// get chain shortName
		ok := false
		for _, info := range chainsData {
			if info.ChainID == chainID.Uint64() {
				ok = true
				info.URI = uri
				providers[info.ChainID] = info
				break
			}
		}
		if !ok {
			return nil, fmt.Errorf("information about %d chain not found", chainID.Uint64())
		}
	}
	return providers, nil
}
