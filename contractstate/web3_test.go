package contractstate

import (
	"context"
	"flag"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	qt "github.com/frankban/quicktest"
	"go.vocdoni.io/dvote/log"
)

var url = flag.String("url", "", "ethereum web3 url")

// go test -v -run TestUpdateTokenHolders -url http://...
func TestUpdateTokenHolders(t *testing.T) {
	log.Init(log.LogLevelDebug, "stderr")

	c := qt.New(t)

	th := new(TokenHolders)
	th.Init(common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"), CONTRACT_TYPE_ERC20)

	w3 := Web3{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := w3.Init(ctx, *url, th.Address(), th.Type())
	c.Assert(err, qt.IsNil)

	td, err := w3.GetTokenData()
	c.Assert(err, qt.IsNil)
	log.Infof("getting new holders on the last 100 blocks of the token %s (%s)\n", td.Name, th.Address().String())

	fromBlockNumber, err := w3.client.BlockNumber(ctx)
	c.Assert(err, qt.IsNil)
	_, err = w3.UpdateTokenHolders(ctx, th, fromBlockNumber-100)
	c.Assert(err, qt.IsNil)

	log.Infof("test finished, new tokens found on the last 100 blocks: %d", len(th.Holders()))
}
