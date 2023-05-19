package state

import (
	"context"
	"flag"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	qt "github.com/frankban/quicktest"
	"go.vocdoni.io/dvote/log"
)

var (
	url       = flag.String("url", "", "ethereum web3 url")
	fromblock = flag.Uint64("blocks", 17060829, "number of blocks from the latest")
)

// go test -v -run TestUpdateTokenHolders -url http://... -block 100
func TestUpdateTokenHolders(t *testing.T) {
	log.Init(log.LogLevelDebug, "stderr", nil)

	c := qt.New(t)

	th := new(TokenHolders)
	th.Init(common.HexToAddress("0x2868dD9aBF1A88D5be7025858A55180D59bb1689"), CONTRACT_TYPE_ERC721, *fromblock)

	w3 := Web3{}
	ctx, cancel := context.WithTimeout(context.Background(), 3000*time.Second)
	defer cancel()

	err := w3.Init(ctx, *url, th.Address(), th.Type())
	c.Assert(err, qt.IsNil)

	log.Infof("getting new holders from block %d of the token %s \n", *fromblock, th.Address())

	currentBlock, err := w3.client.BlockNumber(ctx)
	lastCheckedBlock := *fromblock
	c.Assert(err, qt.IsNil)
	for lastCheckedBlock < currentBlock {
		log.Infof("upgrading holders from block %d", lastCheckedBlock)
		_, err = w3.UpdateTokenHolders(ctx, th)
		c.Assert(err, qt.IsNil)
		time.Sleep(time.Second)
	}
	log.Infof("test finished, new tokens found from block %d: %d", *fromblock, len(th.Holders()))
}
