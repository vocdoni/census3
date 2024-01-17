package web3

import "time"

const (
	DefaultMaxWeb3ClientRetries = 3
)

const (
	shortNameSourceUri           = "https://chainid.network/chains_mini.json"
	checkNetworkEndpointsTimeout = time.Second * 10
	timeLayout                   = "2006-01-02T15:04:05Z07:00"
)

var DefaultNetworkEndpoint = &NetworkEndpoint{
	ChainID:   5,
	Name:      "Goerli",
	ShortName: "gor",
	URIs:      []string{"https://eth-goerli.api.onfinality.io/public"},
}

const (
	// OTHER CONSTANTS
	MAX_SCAN_BLOCKS_PER_ITERATION           = 1000000
	MAX_SCAN_LOGS_PER_ITERATION             = 100000
	MAX_NEW_HOLDER_CANDIDATES_PER_ITERATION = 5000
	BLOCKS_TO_SCAN_AT_ONCE                  = uint64(20000)
	NULL_ADDRESS                            = "0"
)

const (
	// EVM LOG TOPICS
	LOG_TOPIC_VENATION_DEPOSIT        = "4566dfc29f6f11d13a418c26a02bef7c28bae749d4de47e4e6a7cddea6730d59"
	LOG_TOPIC_VENATION_WITHDRAW       = "f279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568"
	LOG_TOPIC_ERC20_TRANSFER          = "ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	LOG_TOPIC_ERC1155_TRANSFER_SINGLE = "c3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62"
	LOG_TOPIC_ERC1155_TRANSFER_BATCH  = "4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb"
	LOG_TOPIC_WANT_DEPOSIT            = "e1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c"
	LOG_TOPIC_WANT_WITHDRAWAL         = "7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65"
	// Add more topics here
)