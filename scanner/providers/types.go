package providers

const (
	// CONTRACT TYPES
	CONTRACT_TYPE_UNKNOWN uint64 = iota
	CONTRACT_TYPE_ERC20
	CONTRACT_TYPE_ERC721
	CONTRACT_TYPE_ERC777
	CONTRACT_TYPE_POAP
)

var TokenTypeStringMap = map[uint64]string{
	CONTRACT_TYPE_UNKNOWN: "unknown",
	CONTRACT_TYPE_ERC20:   "erc20",
	CONTRACT_TYPE_ERC721:  "erc721",
	CONTRACT_TYPE_ERC777:  "erc777",
	CONTRACT_TYPE_POAP:    "poap",
}

var TokenTypeIntMap = map[string]uint64{
	"unknown": CONTRACT_TYPE_UNKNOWN,
	"erc20":   CONTRACT_TYPE_ERC20,
	"erc721":  CONTRACT_TYPE_ERC721,
	"erc777":  CONTRACT_TYPE_ERC777,
	"poap":    CONTRACT_TYPE_POAP,
}
