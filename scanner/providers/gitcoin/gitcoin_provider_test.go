package gitcoin

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"

	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/census3/scanner/providers"
	"github.com/vocdoni/census3/scanner/providers/gitcoin/db"
)

var (
	expectedOriginalHolders = map[string]string{
		"0x85ff01cff157199527528788ec4ea6336615c989": "2244",
		"0x7587cfbd20e5a970209526b4d1f69dbaae8bed37": "2458",
		"0x7bec70fa7ef926878858333b0fa581418e2ef0b5": "2289",
	}
	expectedUpdatedHolders = map[string]string{
		"0x85ff01cff157199527528788ec4ea6336615c989": "-200",
	}
)

func TestGitcoinPassport(t *testing.T) {
	c := qt.New(t)

	tempDBDir := t.TempDir()
	defer func() {
		c.Assert(os.RemoveAll(tempDBDir), qt.IsNil)
	}()
	testDB, err := db.Init(tempDBDir, "gitcoinpassport.sql")
	c.Assert(err, qt.IsNil)

	// start the mocked server with the static file
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	endpoints := providers.ServeTestStaticFiles(ctx, 5500, map[string]string{
		"/original": "./mocked_data.jsonl",
		"/updated":  "./mocked_data_updated.jsonl",
	})
	// create the provider
	provider := new(GitcoinPassport)
	c.Assert(provider.Init(ctx, GitcoinPassportConf{endpoints["/original"], time.Second, testDB}), qt.IsNil)
	// start the first download
	emptyBalances, _, err := provider.HoldersBalances(context.TODO(), nil, 0)
	c.Assert(err, qt.IsNil)
	c.Assert(len(emptyBalances), qt.Equals, 0)
	// wait for the download to finish
	time.Sleep(2 * time.Second)
	// check the balances
	holders, _, err := provider.HoldersBalances(context.TODO(), nil, 0)
	c.Assert(err, qt.IsNil)
	c.Assert(len(holders), qt.Equals, len(expectedOriginalHolders))
	for addr, balance := range holders {
		strAddr := strings.ToLower(addr.Hex())
		expectedBalance, exists := expectedOriginalHolders[strAddr]
		c.Assert(exists, qt.Equals, true, qt.Commentf(strAddr))
		c.Assert(balance.String(), qt.Equals, expectedBalance)
	}
	c.Assert(provider.SetLastBalances(context.TODO(), nil, holders, 0), qt.IsNil)
	// start the second download expecting to use the cached data
	sameBalances, _, err := provider.HoldersBalances(context.TODO(), nil, 0)
	c.Assert(err, qt.IsNil)
	// empty results because the data the same
	c.Assert(len(sameBalances), qt.Equals, 0)
	// start a new one over the new endpoint
	testDB, err = db.Init(tempDBDir, "gitcoinpassport.sql")
	c.Assert(err, qt.IsNil)
	newProvider := new(GitcoinPassport)
	c.Assert(newProvider.Init(ctx, GitcoinPassportConf{endpoints["/updated"], time.Second * 2, testDB}), qt.IsNil)
	// new endpoint with one change
	time.Sleep(time.Second * 5)
	c.Assert(newProvider.SetLastBalances(context.TODO(), nil, holders, 0), qt.IsNil)
	holders, _, err = newProvider.HoldersBalances(context.TODO(), nil, 1)
	c.Assert(err, qt.IsNil)
	c.Assert(len(holders), qt.Equals, len(expectedUpdatedHolders))
	for addr, balance := range holders {
		strAddr := strings.ToLower(addr.Hex())
		expectedBalance, exists := expectedUpdatedHolders[strAddr]
		c.Assert(exists, qt.Equals, true, qt.Commentf(strAddr))
		c.Assert(balance.String(), qt.Equals, expectedBalance)
	}
}
