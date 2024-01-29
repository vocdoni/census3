package providers

// This file contains the data types that can be implemented in the provider
// package. These types identify the type of contract being scanned and the type
// of token being scanned. However, not all contract types may be available.
// The available contract types depend on the providers that the scanner has
// configured. If a new token type is added, it must be added to this file,
// trying to maintain consistency with the existing token types, and avoiding
// using IDs that have already been used. If the IDs change, the database must
// be updated, correcting the IDs of the existing tokens.

const (
	// CONTRACT TYPES
	CONTRACT_TYPE_UNKNOWN uint64 = iota
	CONTRACT_TYPE_ERC20
	CONTRACT_TYPE_ERC721
	CONTRACT_TYPE_ERC777
	CONTRACT_TYPE_POAP
	CONTRACT_TYPE_GITCOIN
	// CONTRACT NAMES
	CONTRACT_NAME_UNKNOWN = "unknown"
	CONTRACT_NAME_ERC20   = "erc20"
	CONTRACT_NAME_ERC721  = "erc721"
	CONTRACT_NAME_ERC777  = "erc777"
	CONTRACT_NAME_POAP    = "poap"
	CONTRACT_NAME_GITCOIN = "gitcoinpassport"
)

var TokenTypeStringMap = map[uint64]string{
	CONTRACT_TYPE_UNKNOWN: CONTRACT_NAME_UNKNOWN,
	CONTRACT_TYPE_ERC20:   CONTRACT_NAME_ERC20,
	CONTRACT_TYPE_ERC721:  CONTRACT_NAME_ERC721,
	CONTRACT_TYPE_ERC777:  CONTRACT_NAME_ERC777,
	CONTRACT_TYPE_POAP:    CONTRACT_NAME_POAP,
	CONTRACT_TYPE_GITCOIN: CONTRACT_NAME_GITCOIN,
}

var TokenTypeIntMap = map[string]uint64{
	CONTRACT_NAME_UNKNOWN: CONTRACT_TYPE_UNKNOWN,
	CONTRACT_NAME_ERC20:   CONTRACT_TYPE_ERC20,
	CONTRACT_NAME_ERC721:  CONTRACT_TYPE_ERC721,
	CONTRACT_NAME_ERC777:  CONTRACT_TYPE_ERC777,
	CONTRACT_NAME_POAP:    CONTRACT_TYPE_POAP,
	CONTRACT_NAME_GITCOIN: CONTRACT_TYPE_GITCOIN,
}
