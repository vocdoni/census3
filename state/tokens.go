package state

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type TokenType uint64

var TokenTypeStringMap = map[TokenType]string{
	CONTRACT_TYPE_UNKNOWN:                 "unknown",
	CONTRACT_TYPE_ERC20:                   "erc20",
	CONTRACT_TYPE_ERC721_BURNED:           "erc721burned",
	CONTRACT_TYPE_ERC1155:                 "erc1155",
	CONTRACT_TYPE_ERC777:                  "erc777",
	CONTRACT_TYPE_CUSTOM_NATION3_VENATION: "nation3",
	CONTRACT_TYPE_CUSTOM_ARAGON_WANT:      "want",
	CONTRACT_TYPE_ERC721:                  "erc721",
}

var TokenTypeIntMap = map[string]TokenType{
	"unknown":      CONTRACT_TYPE_UNKNOWN,
	"erc20":        CONTRACT_TYPE_ERC20,
	"erc721":       CONTRACT_TYPE_ERC721,
	"erc1155":      CONTRACT_TYPE_ERC1155,
	"erc777":       CONTRACT_TYPE_ERC777,
	"nation3":      CONTRACT_TYPE_CUSTOM_NATION3_VENATION,
	"want":         CONTRACT_TYPE_CUSTOM_ARAGON_WANT,
	"erc721burned": CONTRACT_TYPE_ERC721_BURNED,
}

func (c TokenType) String() string {
	if s, ok := TokenTypeStringMap[c]; ok {
		return s
	}
	return "unknown"
}

func TokenTypeFromString(s string) TokenType {
	if c, ok := TokenTypeIntMap[s]; ok {
		return c
	}
	return CONTRACT_TYPE_UNKNOWN
}

type TokenData struct {
	Address     common.Address
	Type        TokenType
	Name        string
	Symbol      string
	Decimals    uint64
	TotalSupply *big.Int
}

func (t *TokenData) String() string {
	return fmt.Sprintf(`{"address":%s, "type":%s "name":%s,"symbol":%s,"decimals":%d,"totalSupply":%s}`,
		t.Address, t.Type.String(), t.Name, t.Symbol, t.Decimals, t.TotalSupply.String())
}
