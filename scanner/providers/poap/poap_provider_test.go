package poap

import (
	"context"
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/census3/scanner/providers"
)

var (
	expectedOriginalHolders = map[string]string{
		"0x0000000000000000000000000000000000000001": "1",
		"0x0000000000000000000000000000000000000002": "1",
		"0x0000000000000000000000000000000000000003": "1",
		"0x0000000000000000000000000000000000000004": "1",
		"0x0000000000000000000000000000000000000005": "1",
		"0x0000000000000000000000000000000000000006": "1",
		"0x0000000000000000000000000000000000000007": "1",
		"0x0000000000000000000000000000000000000008": "1",
		"0x0000000000000000000000000000000000000009": "1",
		"0x0000000000000000000000000000000000000010": "1",
	}
	expectedUpdatedHolders = map[string]string{
		"0x0000000000000000000000000000000000000001": "-1",
		"0x0000000000000000000000000000000000000002": "-1",
	}
)

func TestPOAP(t *testing.T) {
	c := qt.New(t)
	// start the mocked server with the static file
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	endpoints := providers.ServeTestStaticFiles(ctx, 5501, map[string]string{
		"/original": "./mocked_data.json",
		"/updated":  "./mocked_data_updated.json",
	})

	provider := new(POAPHolderProvider)
	c.Assert(provider.Init(POAPConfig{endpoints["/original"], "no-token"}), qt.IsNil)
	holders, _, _, _, err := provider.HoldersBalances(context.TODO(), nil, 0)
	c.Assert(err, qt.IsNil)
	c.Assert(len(holders), qt.Equals, len(expectedOriginalHolders))
	for addr, balance := range holders {
		expectedBalance, exists := expectedOriginalHolders[addr.Hex()]
		c.Assert(exists, qt.Equals, true)
		c.Assert(balance.String(), qt.Equals, expectedBalance)
	}
	sameBalances, _, _, _, err := provider.HoldersBalances(context.TODO(), nil, 0)
	c.Assert(err, qt.IsNil)
	// empty results because the data the same
	c.Assert(len(sameBalances), qt.Equals, 0)

	provider.apiEndpoint = endpoints["/updated"]
	holders, _, _, _, err = provider.HoldersBalances(context.TODO(), nil, 0)
	c.Assert(err, qt.IsNil)
	c.Assert(len(holders), qt.Equals, len(expectedUpdatedHolders))
	for addr, balance := range holders {
		expectedBalance, exists := expectedUpdatedHolders[addr.Hex()]
		c.Assert(exists, qt.Equals, true)
		c.Assert(balance.String(), qt.Equals, expectedBalance)
	}
}
