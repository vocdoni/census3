package state

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var (
	MonkeysAddress        = common.HexToAddress("0xF530280176385AF31177D78BbFD5eA3f6D07488A")
	MonkeysCreationBlock  = uint64(8901659)
	MonkeysSymbol         = "MON"
	MonkeysName           = "Monkeys"
	MonkeysDecimals       = int64(18)
	MonkeysTotalSupply, _ = new(big.Int).SetString("82000000000000000000", 10)
	MonkeysHolders        = map[common.Address]*big.Int{
		common.HexToAddress("0xe54d702f98E312aBA4318E3c6BDba98ab5e11012"): new(big.Int).SetUint64(16000000000000000000),
		common.HexToAddress("0x38d2BC91B89928f78cBaB3e4b1949e28787eC7a3"): new(big.Int).SetUint64(13000000000000000000),
		common.HexToAddress("0xF752B527E2ABA395D1Ba4C0dE9C147B763dDA1f4"): new(big.Int).SetUint64(12000000000000000000),
		common.HexToAddress("0xe1308a8d0291849bfFb200Be582cB6347FBE90D9"): new(big.Int).SetUint64(9000000000000000000),
		common.HexToAddress("0xdeb8699659bE5d41a0e57E179d6cB42E00B9200C"): new(big.Int).SetUint64(7000000000000000000),
		common.HexToAddress("0xB1F05B11Ba3d892EdD00f2e7689779E2B8841827"): new(big.Int).SetUint64(5000000000000000000),
		common.HexToAddress("0xF3C456FAAa70fea307A073C3DA9572413c77f58B"): new(big.Int).SetUint64(6000000000000000000),
		common.HexToAddress("0x45D3a03E8302de659e7Ea7400C4cfe9CAED8c723"): new(big.Int).SetUint64(6000000000000000000),
		common.HexToAddress("0x313c7f7126486fFefCaa9FEA92D968cbf891b80c"): new(big.Int).SetUint64(3000000000000000000),
		common.HexToAddress("0x1893eD78480267D1854373A99Cee8dE2E08d430F"): new(big.Int).SetUint64(2000000000000000000),
		common.HexToAddress("0xa2E4D94c5923A8dd99c5792A7B0436474c54e1E1"): new(big.Int).SetUint64(2000000000000000000),
		common.HexToAddress("0x2a4636A5a1138e35F7f93e81FA56d3c970BC6777"): new(big.Int).SetUint64(1000000000000000000),
	}
)
