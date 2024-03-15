package web3

import "time"

const (
	DefaultMaxWeb3ClientRetries = 3
)

const (
	shortNameSourceUri           = "https://chainid.network/chains_mini.json"
	checkNetworkEndpointsTimeout = time.Second * 10
	TimeLayout                   = "2006-01-02T15:04:05Z07:00"
)

var DefaultNetworkEndpoint = &NetworkEndpoint{
	ChainID:   11155111,
	Name:      "Sepolia",
	ShortName: "sep",
	URIs:      []string{"https://rpc2.sepolia.org"},
}

const (
	// OTHER CONSTANTS
	MAX_SCAN_BLOCKS_PER_ITERATION           = 1000000
	MAX_SCAN_LOGS_PER_ITERATION             = 100000
	MAX_NEW_HOLDER_CANDIDATES_PER_ITERATION = 5000
	BLOCKS_TO_SCAN_AT_ONCE                  = uint64(2000)
	NULL_ADDRESS                            = "0"
)

const (
	// EVM LOG TOPICS
	LOG_TOPIC_ERC20_TRANSFER          = "ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	LOG_TOPIC_ERC1155_TRANSFER_SINGLE = "c3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62"
	LOG_TOPIC_ERC1155_TRANSFER_BATCH  = "4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb"
	LOG_TOPIC_FARCASTER_REGISTER      = "f2e19a901b0748d8b08e428d0468896a039ac751ec4fec49b44b7b9c28097e45"
	LOG_TOPIC_FARCASTER_ADDKEY        = "7d285df41058466977811345cd453c0c52e8d841ffaabc74fc050f277ad4de02"
	LOG_TOPIC_FARCASTER_REMOVEKEY     = "09e77066e0155f46785be12f6938a6b2e4be4381e59058129ce15f355cb96958"
	// Add more topics here
)
