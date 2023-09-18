package state

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	qt "github.com/frankban/quicktest"
)

var web3URI = web3testUri()

func TestWeb3Init(t *testing.T) {
	c := qt.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c.Assert(new(Web3).Init(ctx, "", MonkeysAddress, CONTRACT_TYPE_ERC20), qt.IsNotNil)
	c.Assert(new(Web3).Init(ctx, web3URI, MonkeysAddress, CONTRACT_TYPE_ERC20), qt.IsNil)
}

func TestNewContract(t *testing.T) {
	c := qt.New(t)
	w := new(Web3)
	_, err := w.NewContract()
	c.Assert(err, qt.IsNotNil)

	w.contractType = CONTRACT_TYPE_UNKNOWN
	_, err = w.NewContract()
	c.Assert(err, qt.IsNotNil)

	w.contractType = CONTRACT_TYPE_ERC20
	_, err = w.NewContract()
	c.Assert(err, qt.IsNil)
}

func TestTokenName(t *testing.T) {
	c := qt.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	w := new(Web3)
	c.Assert(w.Init(ctx, "https://google.com", MonkeysAddress, CONTRACT_TYPE_ERC20), qt.IsNil)
	_, err := w.TokenName()
	c.Assert(err, qt.IsNotNil)

	c.Assert(w.Init(ctx, web3URI, common.HexToAddress("0x0"), CONTRACT_TYPE_ERC20), qt.IsNil)
	_, err = w.TokenName()
	c.Assert(err, qt.IsNotNil)

	c.Assert(w.Init(ctx, web3URI, MonkeysAddress, CONTRACT_TYPE_ERC20), qt.IsNil)
	name, err := w.TokenName()
	c.Assert(err, qt.IsNil)
	c.Assert(name, qt.Equals, MonkeysName)
}

func TestTokenSymbol(t *testing.T) {
	c := qt.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	w := new(Web3)
	c.Assert(w.Init(ctx, "https://google.com", MonkeysAddress, CONTRACT_TYPE_ERC20), qt.IsNil)
	_, err := w.TokenSymbol()
	c.Assert(err, qt.IsNotNil)

	c.Assert(w.Init(ctx, web3URI, common.HexToAddress("0x0"), CONTRACT_TYPE_ERC20), qt.IsNil)
	_, err = w.TokenSymbol()
	c.Assert(err, qt.IsNotNil)

	c.Assert(w.Init(ctx, web3URI, MonkeysAddress, CONTRACT_TYPE_ERC20), qt.IsNil)
	symbol, err := w.TokenSymbol()
	c.Assert(err, qt.IsNil)
	c.Assert(symbol, qt.Equals, MonkeysSymbol)
}

func TestTokenDecimals(t *testing.T) {
	c := qt.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	w := new(Web3)
	c.Assert(w.Init(ctx, "https://google.com", MonkeysAddress, CONTRACT_TYPE_ERC20), qt.IsNil)
	_, err := w.TokenDecimals()
	c.Assert(err, qt.IsNotNil)

	c.Assert(w.Init(ctx, web3URI, common.HexToAddress("0x0"), CONTRACT_TYPE_ERC20), qt.IsNil)
	_, err = w.TokenDecimals()
	c.Assert(err, qt.IsNotNil)

	c.Assert(w.Init(ctx, web3URI, MonkeysAddress, CONTRACT_TYPE_ERC20), qt.IsNil)
	decimals, err := w.TokenDecimals()
	c.Assert(err, qt.IsNil)
	c.Assert(int64(decimals), qt.Equals, MonkeysDecimals)
}

func TestTokenTotalSupply(t *testing.T) {
	c := qt.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	w := new(Web3)
	c.Assert(w.Init(ctx, "https://google.com", MonkeysAddress, CONTRACT_TYPE_ERC20), qt.IsNil)
	_, err := w.TokenTotalSupply()
	c.Assert(err, qt.IsNotNil)

	c.Assert(w.Init(ctx, web3URI, common.HexToAddress("0x0"), CONTRACT_TYPE_ERC20), qt.IsNil)
	_, err = w.TokenTotalSupply()
	c.Assert(err, qt.IsNotNil)

	c.Assert(w.Init(ctx, web3URI, MonkeysAddress, CONTRACT_TYPE_ERC20), qt.IsNil)
	totalSupply, err := w.TokenTotalSupply()
	c.Assert(err, qt.IsNil)
	c.Assert(totalSupply.String(), qt.Equals, MonkeysTotalSupply.String())
}

func TestTokenData(t *testing.T) {
	c := qt.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	w := new(Web3)
	c.Assert(w.Init(ctx, "https://google.com", MonkeysAddress, CONTRACT_TYPE_ERC20), qt.IsNil)
	_, err := w.TokenData()
	c.Assert(err, qt.IsNotNil)

	c.Assert(w.Init(ctx, web3URI, common.HexToAddress("0x0"), CONTRACT_TYPE_ERC20), qt.IsNil)
	_, err = w.TokenData()
	c.Assert(err, qt.IsNotNil)

	c.Assert(w.Init(ctx, web3URI, MonkeysAddress, CONTRACT_TYPE_ERC20), qt.IsNil)
	data, err := w.TokenData()
	c.Assert(err, qt.IsNil)
	c.Assert(data.Address, qt.Equals, MonkeysAddress)
	c.Assert(data.Type, qt.Equals, CONTRACT_TYPE_ERC20)
	c.Assert(data.Name, qt.Equals, MonkeysName)
	c.Assert(data.Symbol, qt.Equals, MonkeysSymbol)
	c.Assert(int64(data.Decimals), qt.Equals, MonkeysDecimals)
	c.Assert(data.TotalSupply.String(), qt.Equals, MonkeysTotalSupply.String())
}

func TestTokenBalanceOf(t *testing.T) {
	c := qt.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	w := new(Web3)
	c.Assert(w.Init(ctx, web3URI, MonkeysAddress, CONTRACT_TYPE_ERC20), qt.IsNil)

	holderAddress := common.HexToAddress("0xB1F05B11Ba3d892EdD00f2e7689779E2B8841827")
	balance, err := w.TokenBalanceOf(holderAddress)
	c.Assert(err, qt.IsNil)
	c.Assert(balance.String(), qt.Equals, MonkeysHolders[holderAddress].String())

	balance, err = w.TokenBalanceOf(common.HexToAddress("0x0"))
	c.Assert(err, qt.IsNil)
	c.Assert(balance.String(), qt.Equals, "0")
}

func TestBlockTimestamp(t *testing.T) {
	c := qt.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	w := new(Web3)
	c.Assert(w.Init(ctx, "https://google.com", MonkeysAddress, CONTRACT_TYPE_ERC20), qt.IsNil)
	_, err := w.BlockTimestamp(ctx, uint(MonkeysCreationBlock))
	c.Assert(err, qt.IsNotNil)

	c.Assert(w.Init(ctx, web3URI, MonkeysAddress, CONTRACT_TYPE_ERC20), qt.IsNil)

	expected, err := time.Parse(timeLayout, "2023-04-27T19:21:24+02:00")
	c.Assert(err, qt.IsNil)
	timestamp, err := w.BlockTimestamp(ctx, uint(MonkeysCreationBlock))
	c.Assert(err, qt.IsNil)
	result, err := time.Parse(timeLayout, timestamp)
	c.Assert(err, qt.IsNil)
	c.Assert(expected.Equal(result), qt.IsTrue)
}

func TestBlockRootHash(t *testing.T) {
	c := qt.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	w := new(Web3)
	c.Assert(w.Init(ctx, "https://google.com", MonkeysAddress, CONTRACT_TYPE_ERC20), qt.IsNil)
	_, err := w.BlockRootHash(ctx, uint(MonkeysCreationBlock))
	c.Assert(err, qt.IsNotNil)

	c.Assert(w.Init(ctx, web3URI, MonkeysAddress, CONTRACT_TYPE_ERC20), qt.IsNil)

	expected := common.HexToHash("0x541528e4030f29a87e81ea034e485ede7ea3086784212a3a4863a7de32415de0")
	bhash, err := w.BlockRootHash(ctx, uint(MonkeysCreationBlock))
	c.Assert(err, qt.IsNil)
	c.Assert(common.BytesToHash(bhash), qt.ContentEquals, expected)
}

func TestLatestBlockNumber(t *testing.T) {
	c := qt.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	w := new(Web3)
	c.Assert(w.Init(ctx, "https://google.com", MonkeysAddress, CONTRACT_TYPE_ERC20), qt.IsNil)
	_, err := w.LatestBlockNumber(ctx)
	c.Assert(err, qt.IsNotNil)

	c.Assert(w.Init(ctx, web3URI, MonkeysAddress, CONTRACT_TYPE_ERC20), qt.IsNil)
	blockNumber, err := w.LatestBlockNumber(ctx)
	c.Assert(err, qt.IsNil)
	c.Assert(uint64(blockNumber) > MonkeysCreationBlock, qt.IsTrue)
}

func TestUpdateTokenHolders(t *testing.T) {
	c := qt.New(t)

	th := new(TokenHolders)
	th = th.Init(MonkeysAddress, CONTRACT_TYPE_ERC20, MonkeysCreationBlock, 5)

	w3 := Web3{}
	ctx, cancel := context.WithTimeout(context.Background(), 3000*time.Second)
	defer cancel()

	c.Assert(w3.Init(ctx, web3URI, th.Address(), th.Type()), qt.IsNil)

	current, end := MonkeysCreationBlock, MonkeysCreationBlock+1000
	var err error
	for current < end {
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

func Test_transferLogs(t *testing.T) {
	c := qt.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	w := new(Web3)
	c.Assert(w.Init(ctx, web3URI, MonkeysAddress, CONTRACT_TYPE_ERC20), qt.IsNil)

	logs, err := w.transferLogs(MonkeysCreationBlock, MonkeysCreationBlock+500)
	c.Assert(err, qt.IsNil)
	c.Assert(logs, qt.HasLen, 10)
}

func Test_calcPartialBalances(t *testing.T) {
	c := qt.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	w := new(Web3)
	c.Assert(w.Init(ctx, web3URI, MonkeysAddress, CONTRACT_TYPE_ERC20), qt.IsNil)

	logs, err := w.transferLogs(MonkeysCreationBlock, MonkeysCreationBlock+500)
	c.Assert(err, qt.IsNil)

	hc := HoldersCandidates{}
	for _, log := range logs {
		hc, err = w.calcPartialBalances(hc, log)
		c.Assert(err, qt.IsNil)
	}
	for addr, balance := range MonkeysHolders {
		res, ok := hc[addr]
		c.Assert(ok, qt.IsTrue)
		c.Assert(res.String(), qt.Equals, balance.String())
	}
}

func Test_commitTokenHolders(t *testing.T) {
	c := qt.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	w := new(Web3)
	c.Assert(w.Init(ctx, web3URI, MonkeysAddress, CONTRACT_TYPE_ERC20), qt.IsNil)

	hc := HoldersCandidates(MonkeysHolders)
	th := new(TokenHolders).Init(MonkeysAddress, CONTRACT_TYPE_ERC20, MonkeysCreationBlock, 5)
	c.Assert(w.commitTokenHolders(th, hc, MonkeysCreationBlock+1000), qt.IsNil)

	c.Assert(th.LastBlock(), qt.Equals, MonkeysCreationBlock)
	for addr, balance := range hc {
		c.Assert(th.Exists(addr), qt.IsTrue)
		val, ok := th.holders.Load(addr)
		c.Assert(ok, qt.IsTrue)
		c.Assert(val.(*big.Int).String(), qt.Equals, balance.String())
	}
}

func TestCreationBlock(t *testing.T) {
	c := qt.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	w := new(Web3)
	c.Assert(w.Init(ctx, web3URI, common.HexToAddress(""), CONTRACT_TYPE_ERC20), qt.IsNil)

	// for an invalid contract address, returns the latest block number, the
	// test uses a range of block numbers to cover also the case where any block
	// is mined during test execution
	creationBlock, err := w.ContractCreationBlock(ctx)
	c.Assert(err, qt.IsNil)
	latestBlock, err := w.LatestBlockNumber(ctx)
	c.Assert(err, qt.IsNil)
	c.Assert(creationBlock > latestBlock-5 && creationBlock < latestBlock+5, qt.IsTrue)

	c.Assert(w.Init(ctx, web3URI, MonkeysAddress, CONTRACT_TYPE_ERC20), qt.IsNil)
	creationBlock, err = w.ContractCreationBlock(ctx)
	c.Assert(err, qt.IsNil)
	c.Assert(creationBlock, qt.Equals, MonkeysCreationBlock)
}

func Test_creationBlockInRange(t *testing.T) {
	c := qt.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	w := new(Web3)
	c.Assert(w.Init(ctx, web3URI, MonkeysAddress, CONTRACT_TYPE_ERC20), qt.IsNil)

	blockNumber, err := w.creationBlockInRange(ctx, 0, 10)
	c.Assert(err, qt.IsNil)
	c.Assert(blockNumber, qt.Equals, uint64(10))

	blockNumber, err = w.creationBlockInRange(ctx, 0, 1)
	c.Assert(err, qt.IsNil)
	c.Assert(blockNumber, qt.Equals, uint64(1))

	blockNumber, err = w.creationBlockInRange(ctx, 0, 9000000)
	c.Assert(err, qt.IsNil)
	c.Assert(blockNumber, qt.Equals, MonkeysCreationBlock)
}

func TestSourceCodeLenAt(t *testing.T) {
	c := qt.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	w := new(Web3)
	c.Assert(w.Init(ctx, web3URI, MonkeysAddress, CONTRACT_TYPE_ERC20), qt.IsNil)

	codeLen, err := w.SourceCodeLenAt(ctx, MonkeysCreationBlock)
	c.Assert(err, qt.IsNil)
	c.Assert(codeLen > 2, qt.IsTrue)

	codeLen, err = w.SourceCodeLenAt(ctx, MonkeysCreationBlock-1)
	c.Assert(err, qt.IsNil)
	c.Assert(codeLen > 2, qt.IsFalse)
}
