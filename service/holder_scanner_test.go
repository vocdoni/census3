package service

import (
	"context"
	"database/sql"
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	qt "github.com/frankban/quicktest"
	queries "github.com/vocdoni/census3/db/sqlc"
	"github.com/vocdoni/census3/service/web3"
	"github.com/vocdoni/census3/state"
)

var (
	web3endpoint, _ = web3.TestNetworkEndpoint()
	web3Endpoints   = map[uint64]*web3.NetworkEndpoint{
		web3endpoint.ChainID: web3endpoint,
	}
)

func TestNewHolderScanner(t *testing.T) {
	c := qt.New(t)

	testdb := StartTestDB(t)
	defer testdb.Close(t)

	hs, err := NewHoldersScanner(testdb.db, web3Endpoints, nil, 20)
	c.Assert(err, qt.IsNil)
	c.Assert(hs.lastBlock, qt.Equals, uint64(0))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = testdb.db.QueriesRW.CreateBlock(ctx, queries.CreateBlockParams{
		ID:        1000,
		Timestamp: "test",
		RootHash:  []byte("test"),
	})
	c.Assert(err, qt.IsNil)

	hs, err = NewHoldersScanner(testdb.db, web3Endpoints, nil, 20)
	c.Assert(err, qt.IsNil)
	c.Assert(hs.lastBlock, qt.Equals, uint64(1000))

	_, err = NewHoldersScanner(nil, web3Endpoints, nil, 20)
	c.Assert(err, qt.IsNotNil)
}

func TestHolderScannerStart(t *testing.T) {
	c := qt.New(t)
	twg := sync.WaitGroup{}

	ctx, cancel := context.WithCancel(context.Background())
	testdb := StartTestDB(t)
	defer testdb.Close(t)

	twg.Add(1)
	hs, err := NewHoldersScanner(testdb.db, web3Endpoints, nil, 20)
	c.Assert(err, qt.IsNil)
	go func() {
		hs.Start(ctx)
		twg.Done()
	}()

	cancel()
	twg.Wait()
}

func Test_tokenAddresses(t *testing.T) {
	c := qt.New(t)

	testdb := StartTestDB(t)
	defer testdb.Close(t)

	hs, err := NewHoldersScanner(testdb.db, web3Endpoints, nil, 20)
	c.Assert(err, qt.IsNil)

	res, err := hs.tokenAddresses()
	c.Assert(err, qt.IsNil)
	c.Assert(res, qt.HasLen, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = testdb.db.QueriesRW.CreateToken(ctx, testTokenParams("0x1", "test0",
		"test0", 0, MonkeysDecimals, uint64(state.CONTRACT_TYPE_ERC20),
		MonkeysTotalSupply.Int64(), false, 5, ""))
	c.Assert(err, qt.IsNil)

	res, err = hs.tokenAddresses()
	c.Assert(err, qt.IsNil)
	c.Assert(res[0].ready, qt.IsFalse)
	c.Assert(res[0].addr.String(), qt.Equals, common.HexToAddress("0x1").String())

	_, err = testdb.db.QueriesRW.CreateToken(ctx, testTokenParams("0x2", "test2",
		"test3", 10, MonkeysDecimals, uint64(state.CONTRACT_TYPE_ERC20),
		MonkeysTotalSupply.Int64(), false, 5, ""))
	c.Assert(err, qt.IsNil)

	res, err = hs.tokenAddresses()
	c.Assert(err, qt.IsNil)
	c.Assert(res[1].ready, qt.IsTrue)
	c.Assert(res[1].addr.String(), qt.Equals, common.HexToAddress("0x2").String())
}

func Test_saveHolders(t *testing.T) {
	c := qt.New(t)

	testdb := StartTestDB(t)
	defer testdb.Close(t)

	hs, err := NewHoldersScanner(testdb.db, web3Endpoints, nil, 20)
	c.Assert(err, qt.IsNil)

	th := new(state.TokenHolders).Init(MonkeysAddress, state.CONTRACT_TYPE_ERC20, MonkeysCreationBlock, 5, "")
	// no registered token
	c.Assert(hs.saveHolders(th), qt.ErrorIs, ErrTokenNotExists)
	_, err = testdb.db.QueriesRW.CreateToken(context.Background(), testTokenParams(
		MonkeysAddress.String(), MonkeysName, MonkeysSymbol, MonkeysCreationBlock,
		MonkeysDecimals, uint64(state.CONTRACT_TYPE_ERC20), MonkeysTotalSupply.Int64(), false, 5, ""))
	c.Assert(err, qt.IsNil)
	// check no new holders
	c.Assert(hs.saveHolders(th), qt.IsNil)
	// mock holder
	holderAddr := common.HexToAddress("0x1")
	holderBalance := new(big.Int).SetInt64(12)
	th.Append(holderAddr, holderBalance)
	th.BlockDone(MonkeysCreationBlock)
	// check web3
	c.Assert(hs.saveHolders(th), qt.IsNil)
	// check new holders
	res, err := testdb.db.QueriesRO.TokenHolderByTokenIDAndHolderID(context.Background(),
		queries.TokenHolderByTokenIDAndHolderIDParams{
			TokenID:  MonkeysAddress.Bytes(),
			HolderID: holderAddr.Bytes(),
			ChainID:  th.ChainID,
		})
	c.Assert(err, qt.IsNil)
	c.Assert([]byte(res.Balance), qt.ContentEquals, holderBalance.Bytes())
	// check update holders
	th.Append(holderAddr, holderBalance)
	c.Assert(hs.saveHolders(th), qt.IsNil)
	res, err = testdb.db.QueriesRO.TokenHolderByTokenIDAndHolderID(context.Background(),
		queries.TokenHolderByTokenIDAndHolderIDParams{
			TokenID:  MonkeysAddress.Bytes(),
			HolderID: holderAddr.Bytes(),
			ChainID:  th.ChainID,
		})
	c.Assert(err, qt.IsNil)
	resBalance, ok := new(big.Int).SetString(res.Balance, 10)
	c.Assert(ok, qt.IsTrue)
	c.Assert(resBalance.String(), qt.Equals, "12")
	// check delete holders
	th.Append(holderAddr, big.NewInt(-24))
	c.Assert(hs.saveHolders(th), qt.IsNil)
	_, err = testdb.db.QueriesRO.TokenHolderByTokenIDAndHolderID(context.Background(),
		queries.TokenHolderByTokenIDAndHolderIDParams{
			TokenID:  MonkeysAddress.Bytes(),
			HolderID: holderAddr.Bytes(),
		})
	c.Assert(err, qt.ErrorIs, sql.ErrNoRows)
}

func Test_scanHolders(t *testing.T) {
	c := qt.New(t)

	testdb := StartTestDB(t)
	defer testdb.Close(t)

	hs, err := NewHoldersScanner(testdb.db, web3Endpoints, nil, 20)
	c.Assert(err, qt.IsNil)

	// token does not exists
	ctx1, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = hs.scanHolders(ctx1, MonkeysAddress, 5, []byte{})
	c.Assert(err, qt.IsNotNil)

	_, err = testdb.db.QueriesRW.CreateToken(context.Background(), testTokenParams(
		MonkeysAddress.String(), MonkeysName, MonkeysSymbol, MonkeysCreationBlock,
		MonkeysDecimals, uint64(state.CONTRACT_TYPE_ERC20), 10, false, 5, ""))
	c.Assert(err, qt.IsNil)
	// token exists and the scanner gets the holders
	ctx2, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = hs.scanHolders(ctx2, MonkeysAddress, 5, []byte{})
	c.Assert(err, qt.IsNil)

	res, err := testdb.db.QueriesRW.TokenHoldersByTokenID(context.Background(), MonkeysAddress.Bytes())
	c.Assert(err, qt.IsNil)
	for _, holder := range res {
		balance, ok := MonkeysHolders[common.BytesToAddress(holder.ID)]
		c.Assert(ok, qt.IsTrue)
		currentBalance, ok := new(big.Int).SetString(holder.Balance, 10)
		c.Assert(ok, qt.IsTrue)
		c.Assert(currentBalance, qt.ContentEquals, balance.String())
	}
}

func Test_calcTokenCreationBlock(t *testing.T) {
	c := qt.New(t)

	testdb := StartTestDB(t)
	defer testdb.Close(t)

	hs, err := NewHoldersScanner(testdb.db, web3Endpoints, nil, 20)
	c.Assert(err, qt.IsNil)
	c.Assert(hs.calcTokenCreationBlock(context.Background(), MonkeysAddress, 5), qt.IsNotNil)

	_, err = testdb.db.QueriesRW.CreateToken(context.Background(), testTokenParams(
		MonkeysAddress.String(), MonkeysName, MonkeysSymbol, MonkeysCreationBlock,
		MonkeysDecimals, uint64(state.CONTRACT_TYPE_ERC20), MonkeysTotalSupply.Int64(), false, 5, ""))
	c.Assert(err, qt.IsNil)

	c.Assert(hs.calcTokenCreationBlock(context.Background(), MonkeysAddress, 5), qt.IsNil)
	token, err := testdb.db.QueriesRW.TokenByID(context.Background(), MonkeysAddress.Bytes())
	c.Assert(err, qt.IsNil)
	c.Assert(uint64(token.CreationBlock), qt.Equals, MonkeysCreationBlock)
}
