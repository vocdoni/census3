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
	url           = flag.String("url", "https://eth-goerli.api.onfinality.io/public", "ethereum web3 url")
	tokenContract = flag.String("address", "0xf530280176385af31177d78bbfd5ea3f6d07488a", "token address to scan")
	fromblock     = flag.Uint64("from", 8947005, "from block to scan")
	nblocks       = flag.Uint64("nblock", 100, "number of blocks to scan from the 'from' block")
)

func TestNewContract(t *testing.T)       {}
func TestWeb3Init(t *testing.T)          {}
func TestTokenName(t *testing.T)         {}
func TestTokenSymbol(t *testing.T)       {}
func TestTokenDecimals(t *testing.T)     {}
func TestTokenTotalSupply(t *testing.T)  {}
func TestTokenData(t *testing.T)         {}
func TestTokenBalanceOf(t *testing.T)    {}
func TestBlockTimestamp(t *testing.T)    {}
func TestBlockRootHash(t *testing.T)     {}
func TestLatestBlockNumber(t *testing.T) {}

// go test -v -run TestUpdateTokenHolders -url http://... -block 100
func TestUpdateTokenHolders(t *testing.T) {
	log.Init(log.LogLevelDebug, "stderr", nil)

	c := qt.New(t)

	th := new(TokenHolders)
	th.Init(common.HexToAddress(*tokenContract), CONTRACT_TYPE_ERC20, *fromblock)

	w3 := Web3{}
	ctx, cancel := context.WithTimeout(context.Background(), 3000*time.Second)
	defer cancel()

	err := w3.Init(ctx, *url, th.Address(), th.Type())
	c.Assert(err, qt.IsNil)

	log.Infof("getting new holders from block %d of the token %s \n", *fromblock, th.Address())

	current, end := *fromblock, *fromblock+*nblocks
	c.Assert(err, qt.IsNil)
	for current < end {
		log.Infof("upgrading holders from block %d", current)
		current, err = w3.UpdateTokenHolders(ctx, th)
		if err != nil {
			c.Assert(err, qt.ErrorIs, ErrNoNewBlocks)
		}
		time.Sleep(time.Second)
	}

	for addr, balance := range th.Holders() {
		expectedBalance, ok := MonkeysHolders[addr]
		c.Assert(ok, qt.IsTrue)
		c.Assert(balance.String(), qt.Equals, expectedBalance.String())
	}
}

func Test_transferLogs(t *testing.T)         {}
func Test_calcPartialBalances(t *testing.T)  {}
func Test_commitTokenHolders(t *testing.T)   {}
func TestCreationBlock(t *testing.T)         {}
func Test_creationBlockInRange(t *testing.T) {}
func TestSourceCodeLenAt(t *testing.T)       {}
