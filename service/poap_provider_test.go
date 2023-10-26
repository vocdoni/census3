package service

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	qt "github.com/frankban/quicktest"
)

func TestPOAPHolderProvider_calcPartials(t *testing.T) {
	c := qt.New(t)
	// create a new POAPHolderProvider
	p := &POAPHolderProvider{}
	p.snapshots = make(map[string]*POAPSnapshot)
	// calculate the partial balances with the mocked current and new snapshots
	eventID := "1234"
	currentSnapshot := map[common.Address]*big.Int{
		common.HexToAddress("0x1"): big.NewInt(1),
		common.HexToAddress("0x2"): big.NewInt(2),
		common.HexToAddress("0x3"): big.NewInt(3),
	}
	initialSnapshot := p.calcPartials(eventID, currentSnapshot)
	c.Assert(len(initialSnapshot), qt.Equals, len(currentSnapshot))
	for addr, balance := range currentSnapshot {
		resultingBalance, exist := initialSnapshot[addr]
		c.Assert(exist, qt.Equals, true)
		c.Assert(resultingBalance.Cmp(balance), qt.Equals, 0, qt.Commentf("address %s", addr.Hex()))
	}
	// create a new snapshot with the mocked changes and set the current
	// snapshot as last balances of the event
	newSnapshot := map[common.Address]*big.Int{
		common.HexToAddress("0x1"): big.NewInt(1), // keep 0x1 unchanged
		// delete 0x2
		common.HexToAddress("0x3"): big.NewInt(2), // update 0x3
		common.HexToAddress("0x4"): big.NewInt(1), // add 0x4
	}
	expected := map[common.Address]*big.Int{
		common.HexToAddress("0x1"): big.NewInt(0),
		common.HexToAddress("0x2"): big.NewInt(-2),
		common.HexToAddress("0x3"): big.NewInt(-1),
		common.HexToAddress("0x4"): big.NewInt(1),
	}
	// check that the calcPartials method returns the expected results
	c.Assert(p.SetLastBalances(context.TODO(), []byte(eventID), currentSnapshot, 0), qt.IsNil)
	partialBalances := p.calcPartials(eventID, newSnapshot)
	c.Assert(len(partialBalances), qt.Equals, len(expected))
	for addr, balance := range expected {
		resultingBalance, exist := partialBalances[addr]
		c.Assert(exist, qt.Equals, true)
		c.Assert(resultingBalance.Cmp(balance), qt.Equals, 0, qt.Commentf("address %s", addr.Hex()))
	}
	// combine the results of calcPartials with the current snapshot
	computedNewSnapshot := make(map[common.Address]*big.Int)
	for addr, partialBalance := range partialBalances {
		balance := new(big.Int).Set(partialBalance)
		if currentBalance, exist := currentSnapshot[addr]; exist {
			balance = new(big.Int).Add(currentBalance, partialBalance)
		}
		if balance.Cmp(big.NewInt(0)) != 0 {
			computedNewSnapshot[addr] = balance
		}
	}
	// check that the computed new snapshot is the same as the mocked new
	// snapshot
	c.Assert(len(computedNewSnapshot), qt.Equals, len(newSnapshot))
	for addr, balance := range newSnapshot {
		resultingBalance, exist := computedNewSnapshot[addr]
		c.Assert(exist, qt.Equals, true)
		c.Assert(resultingBalance.Cmp(balance), qt.Equals, 0, qt.Commentf("address %s", addr.Hex()))
	}
}
