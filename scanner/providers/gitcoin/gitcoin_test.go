package gitcoin

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	qt "github.com/frankban/quicktest"
)

var (
	expectedOriginalHolders = map[string]string{
		"0x85ff01cff157199527528788ec4ea6336615c989": "12",
		"0x7587cfbd20e5a970209526b4d1f69dbaae8bed37": "9",
		"0x7bec70fa7ef926878858333b0fa581418e2ef0b5": "1",
		"0x2b1a6dd2a80f7e9a2305205572df0f4b38b205a1": "0",
	}
	expectedUpdatedHolders = map[string]string{
		"0x85ff01cff157199527528788ec4ea6336615c989": "10",
		"0x7587cfbd20e5a970209526b4d1f69dbaae8bed37": "9",
		"0x7bec70fa7ef926878858333b0fa581418e2ef0b5": "1",
		"0x2b1a6dd2a80f7e9a2305205572df0f4b38b205a1": "0",
	}
)

func serveStaticFile(original, updated string) (string, string) {
	http.HandleFunc("/original", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, original)
	})
	http.HandleFunc("/updated", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, updated)
	})
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Println("HTTP server error:", err)
		}
	}()
	return "http://localhost:8080/original", "http://localhost:8080/updated"
}

func TestGitcoinPassport(t *testing.T) {
	c := qt.New(t)
	// start the mocked server with the static file
	originalEndpoint, updatedEndpoint := serveStaticFile("./mocked_data.jsonl", "./mocked_data_updated.jsonl")
	// create the provider
	provider := new(GitcoinPassport)
	c.Assert(provider.Init(GitcoinPassportConf{originalEndpoint}), qt.IsNil)
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

	provider.apiEndpoint = updatedEndpoint
	provider.lastUpdate = time.Time{}
	emptyBalances, _, _, _, err = provider.HoldersBalances(context.TODO(), nil, 0)
	c.Assert(err, qt.IsNil)
	c.Assert(len(emptyBalances), qt.Equals, 0)

	time.Sleep(2 * time.Second)
	// check the balances
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
