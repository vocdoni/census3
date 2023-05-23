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

const web3testUri = "https://eth-goerli.api.onfinality.io/public"

func TestNewHolderScanner(t *testing.T) {
	c := qt.New(t)

	testdb := StartTestDB(t)
	defer testdb.Close(t)

	hs, err := NewHoldersScanner(testdb.db, testdb.queries, web3testUri)
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

	hs, err = NewHoldersScanner(testdb.db, testdb.queries, web3testUri)
	c.Assert(err, qt.IsNil)
	c.Assert(hs.lastBlock, qt.Equals, uint64(1000))

	_, err = NewHoldersScanner(nil, nil, web3testUri)
	c.Assert(err, qt.IsNotNil)
}

func TestHolderScannerStart(t *testing.T) {
	c := qt.New(t)
	twg := sync.WaitGroup{}

	ctx, cancel := context.WithCancel(context.Background())
	testdb := StartTestDB(t)
	defer testdb.Close(t)

	twg.Add(1)
	hs, err := NewHoldersScanner(testdb.db, testdb.queries, web3testUri)
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

	hs, err := NewHoldersScanner(testdb.db, testdb.queries, web3testUri)
	c.Assert(err, qt.IsNil)

	res, err := hs.tokenAddresses()
	c.Assert(err, qt.IsNil)
	c.Assert(res, qt.ContentEquals, make(map[common.Address]bool))

	var (
		name          = sql.NullString{}
		symbol        = sql.NullString{}
		decimals      = sql.NullInt64{}
		creationBlock = sql.NullInt32{}
		totalSupply   = big.NewInt(100).Bytes()
	)
	c.Assert(name.Scan("test"), qt.IsNil)
	c.Assert(symbol.Scan("test"), qt.IsNil)
	c.Assert(decimals.Scan(10), qt.IsNil)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = testdb.queries.CreateToken(ctx, queries.CreateTokenParams{
		ID:            common.HexToAddress("0x0").Bytes(),
		Name:          name,
		Symbol:        symbol,
		Decimals:      decimals,
		TotalSupply:   totalSupply,
		CreationBlock: creationBlock,
		TypeID:        int64(state.CONTRACT_TYPE_ERC20),
		Synced:        false,
	})
	c.Assert(err, qt.IsNil)

	res, err = hs.tokenAddresses()
	c.Assert(err, qt.IsNil)
	c.Assert(res, qt.ContentEquals, map[common.Address]bool{
		common.HexToAddress("0x0"): false,
	})

	c.Assert(creationBlock.Scan(10), qt.IsNil)
	_, err = testdb.queries.CreateToken(ctx, queries.CreateTokenParams{
		ID:            common.HexToAddress("0x1").Bytes(),
		Name:          name,
		Symbol:        symbol,
		Decimals:      decimals,
		TotalSupply:   totalSupply,
		CreationBlock: creationBlock,
		TypeID:        int64(state.CONTRACT_TYPE_ERC20),
		Synced:        false,
	})
	c.Assert(err, qt.IsNil)

	res, err = hs.tokenAddresses()
	c.Assert(err, qt.IsNil)
	c.Assert(res, qt.ContentEquals, map[common.Address]bool{
		common.HexToAddress("0x0"): false,
		common.HexToAddress("0x1"): true,
	})
}

func Test_saveHolders(t *testing.T) {
	// check synced
	// check no new holders
	// check init web3
	// check block info
	// check new block creations
	// check new holders
	// check update holders
	// check delete holders
}

func Test_scanHolders(t *testing.T)            {}
func Test_calcTokenCreationBlock(t *testing.T) {}
