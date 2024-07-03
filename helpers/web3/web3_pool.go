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
	"math/big"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"go.vocdoni.io/dvote/log"
)

const (
	// DefaultMaxWeb3ClientRetries is the default number of retries to connect to
	// a web3 provider.
	DefaultMaxWeb3ClientRetries = 5
	// shortNameSourceUri is the URI to get the chain metadata from an external
	shortNameSourceUri = "https://chainid.network/chains_mini.json"
	// checkWeb3EndpointsTimeout is the timeout to check the web3 endpoints.
	checkWeb3EndpointsTimeout = time.Second * 10
	// foundTxErrMessage is the error message when a transaction is found but it
	// is not supported.
	foundTxErrMessage = "transaction type not supported"
	// chainAddressFormat is the format to represent a chain address string.
	chainAddressFormat = "%s:%s"
)

var (
	// notFoundTxRgx is a regular expression to match the error message when a
	// transaction is not found.
	notFoundTxRgx = regexp.MustCompile(`not\s[be\s|]*found`)
	// chainAddressRgx is a regular expression to match the chain address format.
	chainAddressRgx = regexp.MustCompile(`^(.+):(.+)$`)
)

// Web3Pool struct contains a map of chainID-[]*Web3Endpoint, where
// the key is the chainID and the value is a list of Web3Endpoint. It also
// contains a list of all the Web3Endpoint metadata. It provides methods to
// add, remove and get endpoints, as well as to get the chainID by short name.
// It allows to support multiple endpoints for the same chainID and switch
// between them looking for the available one.
type Web3Pool struct {
	endpoints map[uint64]*Web3Iterator
	metadata  []*Web3Endpoint
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
		endpoints: make(map[uint64]*Web3Iterator),
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
	// check if the endpoint is an archive node or not
	isArchive, err := isArchiveNode(ctx, client)
	if err != nil {
		log.Warnw("error checking if the web3 provider is an archive node", "chainID", chainID, "error", err)
	}
	// add the endpoint to the pool
	endpoint := &Web3Endpoint{
		ChainID:   chainID,
		Name:      name,
		ShortName: shortName,
		URI:       uri,
		client:    client,
		IsArchive: isArchive,
	}
	if _, ok := nm.endpoints[chainID]; !ok {
		nm.endpoints[chainID] = NewWeb3Iterator(endpoint)
	} else {
		nm.endpoints[chainID].Add(endpoint)
	}
	return nil
}

// DelEndpoint method removes a web3 provider URI from the *Web3Pool
// instance. It closes the client and removes the endpoint from the list of
// endpoints for the chainID where it was found.
func (nm *Web3Pool) DelEndoint(uri string) {
	for _, endpoints := range nm.endpoints {
		endpoints.Disable(uri)
	}
}

// Endpoint method returns the Web3Endpoint configured for the chainID
// provided. It returns the first available endpoint. If no available endpoint
// is found, returns an error.
func (nm *Web3Pool) Endpoint(chainID uint64) (*Web3Endpoint, error) {
	if endpoints, ok := nm.endpoints[chainID]; ok {
		return endpoints.Next()
	}
	return nil, fmt.Errorf("no endpoint found for chainID %d", chainID)
}

// DisableEndpoint method sets the available flag to false for the URI provided
// in the chainID provided.
func (nm *Web3Pool) DisableEndpoint(chainID uint64, uri string) {
	if endpoints, ok := nm.endpoints[chainID]; ok {
		endpoints.Disable(uri)
	}
}

// NumberOfEndpoints method returns the total number (or just the available ones)
// of endpoints for the chainID provided.
func (nm *Web3Pool) NumberOfEndpoints(chainID uint64, onlyAvailable bool) int {
	if endpoints, ok := nm.endpoints[chainID]; ok {
		n := endpoints.Available()
		if !onlyAvailable {
			n += endpoints.Disabled()
		}
		return n
	}
	return 0
}

// Client method returns a new *Client instance for the chainID provided.
// It returns an error if the endpoint is not found.
func (nm *Web3Pool) Client(chainID uint64) (*Client, error) {
	if _, err := nm.Endpoint(chainID); err != nil {
		return nil, fmt.Errorf("error getting endpoint for chainID %d: %w", chainID, err)
	}
	return &Client{w3p: nm, chainID: chainID}, nil
}

// ChainAddress method returns a prefixed string of the hex address provided,
// with the short name of the networkEndpoint identified by the chain id provided.
// Read more here: https://eips.ethereum.org/EIPS/eip-3770
func (nps *Web3Pool) ChainAddress(chainID uint64, hexAddress string) (string, bool) {
	for _, data := range nps.metadata {
		if data.ChainID == chainID {
			return fmt.Sprintf(chainAddressFormat, data.ShortName, hexAddress), true
		}
	}
	return "", false
}

// AddressFrom method extracts the short name and hex address from a chain
// address string, and returns the hex address and the corresponding
// Web3Endpoint.
func (nps *Web3Pool) AddressFrom(chainAddress string) (string, *Web3Endpoint) {
	if !chainAddressRgx.MatchString(chainAddress) {
		return "", nil
	}
	parts := chainAddressRgx.FindStringSubmatch(chainAddress)
	shortName := parts[1]
	hexAddress := parts[2]
	for _, data := range nps.metadata {
		if data.ShortName == shortName {
			return hexAddress, data
		}
	}
	return "", nil
}

// String method returns a string representation of the *Web3Pool list.
func (nm *Web3Pool) String() string {
	shortNames := map[string]bool{}
	for chainID := range nm.endpoints {
		for _, data := range nm.metadata {
			if data.ChainID == chainID {
				shortNames[data.ShortName] = true
			}
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
	blockNumbers := make(map[uint64]uint64)
	for chainID := range nm.endpoints {
		cli, err := nm.Endpoint(chainID)
		if err != nil {
			return nil, fmt.Errorf("error getting endpoint for chainID %d: %w", chainID, err)
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
	var supported []*Web3Endpoint
	for chainID := range nm.endpoints {
		for _, data := range nm.metadata {
			if data.ChainID == chainID {
				supported = append(supported, &Web3Endpoint{
					ChainID:   chainID,
					Name:      data.Name,
					ShortName: data.ShortName,
				})
				break
			}
		}
	}
	return supported
}

// NetworkInfoByShortName method returns the Web3Endpoint metadata for the
// shortName provided. It returns nil if the shortName is not found.
func (nm *Web3Pool) NetworkInfoByShortName(shortName string) *Web3Endpoint {
	for _, data := range nm.metadata {
		if data.ShortName == shortName {
			return data
		}
	}
	return nil
}

// NetworkInfoByChainID method returns the Web3Endpoint metadata for the
// chainID provided. It returns nil if the chainID is not found.
func (nm *Web3Pool) NetworkInfoByChainID(chainID uint64) *Web3Endpoint {
	for _, data := range nm.metadata {
		if data.ChainID == chainID {
			return data
		}
	}
	return nil
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

// isArchiveNode method returns true if the web3 client is an archive node. To
// determine if the client is an archive node, checks the transactions of the
// block 1 of the chain. If client finds transactions, it is an archive node. If
// it does not find transactions, it is not an archive node. If an error occurs,
// it returns false and the error.
func isArchiveNode(ctx context.Context, client *ethclient.Client) (bool, error) {
	block, err := client.BlockByNumber(ctx, big.NewInt(1))
	if err != nil {
		if strings.Contains(err.Error(), foundTxErrMessage) {
			return true, nil
		}
		return false, fmt.Errorf("error getting block 1: %w", err)
	}

	if _, err := client.TransactionCount(ctx, block.Hash()); err != nil {
		if notFoundTxRgx.MatchString(err.Error()) {
			return false, nil
		}
		if strings.Contains(err.Error(), foundTxErrMessage) {
			return true, nil
		}
		return false, fmt.Errorf("error getting transaction in block 1: %w", err)
	}
	return true, nil
}
