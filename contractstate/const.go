package contractstate

type ContractType int

const (
	// CONTRACT TYPES
	CONTRACT_TYPE_UNKNOWN ContractType = iota
	CONTRACT_TYPE_ERC20
	CONTRACT_TYPE_ERC721
	CONTRACT_TYPE_ERC1155
	CONTRACT_TYPE_ERC777
	CONTRACT_TYPE_CUSTOM_NATION3_VENATION
	CONTRACT_TYPE_CUSTOM_ARAGON_WANT
)

var ContractTypeStringMap = map[ContractType]string{
	CONTRACT_TYPE_UNKNOWN:                 "unknown",
	CONTRACT_TYPE_ERC20:                   "erc20",
	CONTRACT_TYPE_ERC721:                  "erc721",
	CONTRACT_TYPE_ERC1155:                 "erc1155",
	CONTRACT_TYPE_ERC777:                  "erc777",
	CONTRACT_TYPE_CUSTOM_NATION3_VENATION: "nation3",
	CONTRACT_TYPE_CUSTOM_ARAGON_WANT:      "want",
}

var ContractTypeIntMap = map[string]ContractType{
	"unknown": CONTRACT_TYPE_UNKNOWN,
	"erc20":   CONTRACT_TYPE_ERC20,
	"erc721":  CONTRACT_TYPE_ERC721,
	"erc1155": CONTRACT_TYPE_ERC1155,
	"erc777":  CONTRACT_TYPE_ERC777,
	"nation3": CONTRACT_TYPE_CUSTOM_NATION3_VENATION,
	"want":    CONTRACT_TYPE_CUSTOM_ARAGON_WANT,
}

func (c ContractType) String() string {
	if s, ok := ContractTypeStringMap[c]; ok {
		return s
	}
	return "unknown"
}

func ContractTypeFromString(s string) ContractType {
	if c, ok := ContractTypeIntMap[s]; ok {
		return c
	}
	return CONTRACT_TYPE_UNKNOWN
}

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

	// OTHER CONSTANTS
	MAX_SCAN_BLOCKS_PER_ITERATION = 1000000
	MAX_SCAN_LOGS_PER_ITERATION   = 80000
	BLOCKS_TO_SCAN_AT_ONCE        = uint64(5000)
	NULL_ADDRESS                  = "0"
)
