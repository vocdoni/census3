package state

import (
	"context"
	"flag"
	"fmt"
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

var expectedHolders = map[common.Address]string{
	common.HexToAddress("0xe54d702f98E312aBA4318E3c6BDba98ab5e11012"): "16000000000000000000",
	common.HexToAddress("0x38d2BC91B89928f78cBaB3e4b1949e28787eC7a3"): "13000000000000000000",
	common.HexToAddress("0xF752B527E2ABA395D1Ba4C0dE9C147B763dDA1f4"): "12000000000000000000",
	common.HexToAddress("0xdeb8699659bE5d41a0e57E179d6cB42E00B9200C"): "9000000000000000000",
	common.HexToAddress("0xe1308a8d0291849bfFb200Be582cB6347FBE90D9"): "9000000000000000000",
	common.HexToAddress("0xB1F05B11Ba3d892EdD00f2e7689779E2B8841827"): "6000000000000000000",
	common.HexToAddress("0xF3C456FAAa70fea307A073C3DA9572413c77f58B"): "6000000000000000000",
	common.HexToAddress("0x45D3a03E8302de659e7Ea7400C4cfe9CAED8c723"): "6000000000000000000",
	common.HexToAddress("0x313c7f7126486fFefCaa9FEA92D968cbf891b80c"): "3000000000000000000",
	common.HexToAddress("0x1893eD78480267D1854373A99Cee8dE2E08d430F"): "2000000000000000000",
}

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
		fmt.Println(addr.String())
		fmt.Println(balance.String())
		expectedBalance, ok := expectedHolders[addr]
		c.Assert(ok, qt.IsTrue)
		c.Assert(balance.String(), qt.Equals, expectedBalance)
	}
}

func Test_transferLogs(t *testing.T)         {}
func Test_calcPartialBalances(t *testing.T)  {}
func Test_commitTokenHolders(t *testing.T)   {}
func TestCreationBlock(t *testing.T)         {}
func Test_creationBlockInRange(t *testing.T) {}
func TestSourceCodeLenAt(t *testing.T)       {}
