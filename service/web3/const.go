package web3

import "time"

const (
	DefaultMaxRetries = 3
)

const (
	shortNameSourceUri           = "https://chainid.network/chains_mini.json"
	checkNetworkEndpointsTimeout = time.Second * 10
)

var DefaultNetworkEndpoint = &NetworkEndpoint{
	ChainID:   5,
	Name:      "Goerli",
	ShortName: "gor",
	URIs:      []string{"https://eth-goerli.api.onfinality.io/public"},
}
