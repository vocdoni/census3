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
	"github.com/vocdoni/census3/state"
)

var web3uri = web3testUri()

func TestNewHolderScanner(t *testing.T) {
	c := qt.New(t)

	testdb := StartTestDB(t)
	defer testdb.Close(t)

	hs, err := NewHoldersScanner(testdb.db, testdb.queries, web3uri)
	c.Assert(err, qt.IsNil)
	c.Assert(hs.lastBlock, qt.Equals, uint64(0))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = testdb.queries.CreateBlock(ctx, queries.CreateBlockParams{
		ID:        1000,
		Timestamp: "test",
		RootHash:  []byte("test"),
	})
	c.Assert(err, qt.IsNil)

	hs, err = NewHoldersScanner(testdb.db, testdb.queries, web3uri)
	c.Assert(err, qt.IsNil)
	c.Assert(hs.lastBlock, qt.Equals, uint64(1000))

	_, err = NewHoldersScanner(nil, nil, web3uri)
	c.Assert(err, qt.IsNotNil)
}

func TestHolderScannerStart(t *testing.T) {
	c := qt.New(t)
	twg := sync.WaitGroup{}

	ctx, cancel := context.WithCancel(context.Background())
	testdb := StartTestDB(t)
	defer testdb.Close(t)

	twg.Add(1)
	hs, err := NewHoldersScanner(testdb.db, testdb.queries, web3uri)
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

	hs, err := NewHoldersScanner(testdb.db, testdb.queries, web3uri)
	c.Assert(err, qt.IsNil)

	res, err := hs.tokenAddresses()
	c.Assert(err, qt.IsNil)
	c.Assert(res, qt.ContentEquals, make(map[common.Address]bool))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = testdb.queries.CreateToken(ctx, testTokenParams("0x1", "test0",
		"test0", MonkeysDecimals, 0, MonkeysTotalSupply.Uint64(),
		uint64(state.CONTRACT_TYPE_ERC20), false))
	c.Assert(err, qt.IsNil)

	res, err = hs.tokenAddresses()
	c.Assert(err, qt.IsNil)
	c.Assert(res[common.HexToAddress("0x1")], qt.IsFalse)

	_, err = testdb.queries.CreateToken(ctx, testTokenParams("0x2", "test2",
		"test3", MonkeysDecimals, 10, MonkeysTotalSupply.Uint64(),
		uint64(state.CONTRACT_TYPE_ERC20), false))
	c.Assert(err, qt.IsNil)

	res, err = hs.tokenAddresses()
	c.Assert(err, qt.IsNil)
	c.Assert(res[common.HexToAddress("0x2")], qt.IsTrue)
}

func Test_saveHolders(t *testing.T) {
	c := qt.New(t)

	testdb := StartTestDB(t)
	defer testdb.Close(t)

	hs, err := NewHoldersScanner(testdb.db, testdb.queries, "http://google.com")
	c.Assert(err, qt.IsNil)

	th := new(state.TokenHolders).Init(MonkeysAddress, state.CONTRACT_TYPE_ERC20, MonkeysCreationBlock)
	// no registered token
	c.Assert(hs.saveHolders(th), qt.ErrorIs, ErrTokenNotExists)
	_, err = testdb.queries.CreateToken(context.Background(), testTokenParams(
		MonkeysAddress.String(), MonkeysName, MonkeysSymbol, MonkeysDecimals,
		MonkeysCreationBlock, MonkeysTotalSupply.Uint64(),
		uint64(state.CONTRACT_TYPE_ERC20), false))
	c.Assert(err, qt.IsNil)
	// check no new holders
	c.Assert(hs.saveHolders(th), qt.IsNil)
	// mock holder
	holderAddr := common.HexToAddress("0x1")
	holderBalance := new(big.Int).SetUint64(12)
	th.Append(holderAddr, holderBalance)
	th.BlockDone(MonkeysCreationBlock)
	// check wrong web3
	c.Assert(hs.saveHolders(th), qt.IsNotNil)
	// check new block created
	_, err = testdb.queries.BlockByID(context.Background(), int64(MonkeysCreationBlock))
	c.Assert(err, qt.ErrorIs, sql.ErrNoRows)
	// check good web3
	hs.web3 = web3uri
	c.Assert(hs.saveHolders(th), qt.IsNil)
	// check new holders
	res, err := testdb.queries.TokenHolderByTokenIDAndHolderID(context.Background(),
		queries.TokenHolderByTokenIDAndHolderIDParams{
			TokenID:  MonkeysAddress.Bytes(),
			HolderID: holderAddr.Bytes(),
		})
	c.Assert(err, qt.IsNil)
	c.Assert([]byte(res.Balance), qt.ContentEquals, holderBalance.Bytes())
	// check update holders
	th.Append(holderAddr, holderBalance)
	c.Assert(hs.saveHolders(th), qt.IsNil)
	res, err = testdb.queries.TokenHolderByTokenIDAndHolderID(context.Background(),
		queries.TokenHolderByTokenIDAndHolderIDParams{
			TokenID:  MonkeysAddress.Bytes(),
			HolderID: holderAddr.Bytes(),
		})
	c.Assert(err, qt.IsNil)
	resBalance := new(big.Int).SetBytes(res.Balance)
	c.Assert(resBalance.String(), qt.Equals, "24")
	// check delete holders
	th.Append(holderAddr, big.NewInt(-24))
	c.Assert(hs.saveHolders(th), qt.IsNil)
	_, err = testdb.queries.TokenHolderByTokenIDAndHolderID(context.Background(),
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

	hs, err := NewHoldersScanner(testdb.db, testdb.queries, web3uri)
	c.Assert(err, qt.IsNil)

	// token does not exists
	ctx1, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = hs.scanHolders(ctx1, MonkeysAddress)
	c.Assert(err, qt.IsNotNil)

	_, err = testdb.queries.CreateToken(context.Background(), testTokenParams(
		MonkeysAddress.String(), MonkeysName, MonkeysSymbol, MonkeysDecimals, MonkeysCreationBlock, 10,
		uint64(state.CONTRACT_TYPE_ERC20), false))
	c.Assert(err, qt.IsNil)
	// token exists and the scanner gets the holders
	ctx2, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	c.Assert(hs.scanHolders(ctx2, MonkeysAddress), qt.IsNil)

	res, err := testdb.queries.TokenHoldersByTokenID(context.Background(), MonkeysAddress.Bytes())
	c.Assert(err, qt.IsNil)
	for _, holder := range res {
		balance, ok := MonkeysHolders[common.BytesToAddress(holder.ID)]
		c.Assert(ok, qt.IsTrue)
		c.Assert(new(big.Int).SetBytes(holder.Balance).String(), qt.ContentEquals, balance.String())
	}
}

func Test_calcTokenCreationBlock(t *testing.T) {
	c := qt.New(t)

	testdb := StartTestDB(t)
	defer testdb.Close(t)

	hs, err := NewHoldersScanner(testdb.db, testdb.queries, web3uri)
	c.Assert(err, qt.IsNil)
	c.Assert(hs.calcTokenCreationBlock(context.Background(), MonkeysAddress), qt.IsNotNil)

	_, err = testdb.queries.CreateToken(context.Background(), testTokenParams(
		MonkeysAddress.String(), MonkeysName, MonkeysSymbol, MonkeysDecimals,
		MonkeysCreationBlock, MonkeysTotalSupply.Uint64(),
		uint64(state.CONTRACT_TYPE_ERC20), false))
	c.Assert(err, qt.IsNil)

	c.Assert(hs.calcTokenCreationBlock(context.Background(), MonkeysAddress), qt.IsNil)
	token, err := testdb.queries.TokenByID(context.Background(), MonkeysAddress.Bytes())
	c.Assert(err, qt.IsNil)
	c.Assert(uint64(token.CreationBlock.Int32), qt.Equals, MonkeysCreationBlock)
}
