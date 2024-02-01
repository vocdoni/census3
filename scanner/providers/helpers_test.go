package providers

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	qt "github.com/frankban/quicktest"
)

func TestCalcPartialHolders(t *testing.T) {
	c := qt.New(t)
	currentHolders := map[common.Address]*big.Int{
		common.HexToAddress("0x1"): big.NewInt(100),
		common.HexToAddress("0x2"): big.NewInt(200),
		common.HexToAddress("0x3"): big.NewInt(300),
		common.HexToAddress("0x4"): big.NewInt(400),
		common.HexToAddress("0x5"): big.NewInt(500),
	}
	newHolders := map[common.Address]*big.Int{
		common.HexToAddress("0x1"): big.NewInt(200),
		common.HexToAddress("0x2"): big.NewInt(100),
		common.HexToAddress("0x3"): big.NewInt(300),
		common.HexToAddress("0x6"): big.NewInt(600),
		common.HexToAddress("0x7"): big.NewInt(700),
	}
	expectedPartialHolders := map[common.Address]*big.Int{
		common.HexToAddress("0x1"): big.NewInt(100),
		common.HexToAddress("0x2"): big.NewInt(-100),
		common.HexToAddress("0x4"): big.NewInt(-400),
		common.HexToAddress("0x5"): big.NewInt(-500),
		common.HexToAddress("0x6"): big.NewInt(600),
		common.HexToAddress("0x7"): big.NewInt(700),
	}
	partialHolders := CalcPartialHolders(currentHolders, newHolders)
	c.Assert(len(partialHolders), qt.Equals, len(expectedPartialHolders))
	for addr, balance := range partialHolders {
		expectedBalance, ok := expectedPartialHolders[addr]
		c.Assert(ok, qt.Equals, true)
		c.Assert(balance.String(), qt.Equals, expectedBalance.String())
	}

	newgativeHolders := CalcPartialHolders(currentHolders, nil)
	c.Assert(len(newgativeHolders), qt.Equals, len(currentHolders))
	for addr, balance := range newgativeHolders {
		currentBalance, ok := currentHolders[addr]
		c.Assert(ok, qt.Equals, true)
		c.Assert(balance.String(), qt.Equals, new(big.Int).Neg(currentBalance).String())
	}

	fullNewHolders := CalcPartialHolders(nil, newHolders)
	c.Assert(len(fullNewHolders), qt.Equals, len(newHolders))
	for addr, balance := range fullNewHolders {
		expectedBalance, ok := newHolders[addr]
		c.Assert(ok, qt.Equals, true)
		c.Assert(balance.String(), qt.Equals, expectedBalance.String())
	}

	emptyHolders := CalcPartialHolders(currentHolders, currentHolders)
	c.Assert(len(emptyHolders), qt.Equals, 0)

	doubleHolders := map[common.Address]*big.Int{
		common.HexToAddress("0x1"): big.NewInt(200),
		common.HexToAddress("0x2"): big.NewInt(400),
		common.HexToAddress("0x3"): big.NewInt(600),
		common.HexToAddress("0x4"): big.NewInt(800),
		common.HexToAddress("0x5"): big.NewInt(1000),
	}
	sameHolders := CalcPartialHolders(currentHolders, doubleHolders)
	c.Assert(len(sameHolders), qt.Equals, len(currentHolders))
	for addr, balance := range sameHolders {
		currentBalance, ok := currentHolders[addr]
		c.Assert(ok, qt.Equals, true)
		c.Assert(balance.String(), qt.Equals, currentBalance.String())
	}
}
