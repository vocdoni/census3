package gitcoin

import (
	"context"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/census3/scanner/providers"
)

var (
	expectedOriginalHolders = map[string]string{
		"0x85ff01cff157199527528788ec4ea6336615c989": "12",
		"0x7587cfbd20e5a970209526b4d1f69dbaae8bed37": "9",
		"0x7bec70fa7ef926878858333b0fa581418e2ef0b5": "1",
		"0x2b1a6dd2a80f7e9a2305205572df0f4b38b205a1": "0",
	}
	expectedUpdatedHolders = map[string]string{
		"0x85ff01cff157199527528788ec4ea6336615c989": "-2",
	}
)

func TestGitcoinPassport(t *testing.T) {
	c := qt.New(t)
	// start the mocked server with the static file
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	endpoints := providers.ServeTestStaticFiles(ctx, 5500, map[string]string{
		"/original": "./mocked_data.jsonl",
		"/updated":  "./mocked_data_updated.jsonl",
	})
	// create the provider
	provider := new(GitcoinPassport)
	c.Assert(provider.Init(GitcoinPassportConf{endpoints["/original"], time.Second * 2}), qt.IsNil)
	// start the first download
	emptyBalances, _, _, _, err := provider.HoldersBalances(context.TODO(), nil, 0)
	c.Assert(err, qt.IsNil)
	c.Assert(len(emptyBalances), qt.Equals, 0)
	// wait for the download to finish
	time.Sleep(2 * time.Second)
	// check the balances
	holders, _, _, _, err := provider.HoldersBalances(context.TODO(), nil, 0)
	c.Assert(err, qt.IsNil)
	c.Assert(len(holders), qt.Equals, len(expectedOriginalHolders))
	for addr, balance := range holders {
		strAddr := strings.ToLower(addr.Hex())
		expectedBalance, exists := expectedOriginalHolders[strAddr]
		c.Assert(exists, qt.Equals, true, qt.Commentf(strAddr))
		c.Assert(balance.String(), qt.Equals, expectedBalance)
	}
	// start the second download expecting to use the cached data
	sameBalances, _, _, _, err := provider.HoldersBalances(context.TODO(), nil, 0)
	c.Assert(err, qt.IsNil)
	// empty results because the data the same
	c.Assert(len(sameBalances), qt.Equals, 0)

	provider.apiEndpoint = endpoints["/updated"]
	provider.lastUpdate.Store(time.Time{})
	emptyBalances, _, _, _, err = provider.HoldersBalances(context.TODO(), nil, 0)
	c.Assert(err, qt.IsNil)
	c.Assert(len(emptyBalances), qt.Equals, 0)

	time.Sleep(2 * time.Second)
	// check the balances
	currentHolders := make(map[common.Address]*big.Int)
	for addr, balance := range expectedOriginalHolders {
		currentHolders[common.HexToAddress(addr)], _ = new(big.Int).SetString(balance, 10)
	}
	c.Assert(provider.SetLastBalances(context.TODO(), nil, currentHolders, 0), qt.IsNil)
	holders, _, _, _, err = provider.HoldersBalances(context.TODO(), nil, 0)
	c.Assert(err, qt.IsNil)
	c.Assert(len(holders), qt.Equals, len(expectedUpdatedHolders))
	for addr, balance := range holders {
		strAddr := strings.ToLower(addr.Hex())
		expectedBalance, exists := expectedUpdatedHolders[strAddr]
		c.Assert(exists, qt.Equals, true, qt.Commentf(strAddr))
		c.Assert(balance.String(), qt.Equals, expectedBalance)
	}
}
