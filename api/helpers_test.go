package api

import (
	"bytes"
	"encoding/binary"
	"math/big"
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	qt "github.com/frankban/quicktest"
	"go.vocdoni.io/dvote/api/censusdb"
	"go.vocdoni.io/dvote/censustree"
	storagelayer "go.vocdoni.io/dvote/data"
	vocdoniDB "go.vocdoni.io/dvote/db"
	"go.vocdoni.io/dvote/db/metadb"
)

var testHolders = map[common.Address]*big.Int{
	common.HexToAddress("0xe54d702f98E312aBA4318E3c6BDba98ab5e11012"): big.NewInt(16),
	common.HexToAddress("0x38d2BC91B89928f78cBaB3e4b1949e28787eC7a3"): big.NewInt(13),
	common.HexToAddress("0xF752B527E2ABA395D1Ba4C0dE9C147B763dDA1f4"): big.NewInt(12),
	common.HexToAddress("0xe1308a8d0291849bfFb200Be582cB6347FBE90D9"): big.NewInt(9),
	common.HexToAddress("0xdeb8699659bE5d41a0e57E179d6cB42E00B9200C"): big.NewInt(7),
	common.HexToAddress("0xB1F05B11Ba3d892EdD00f2e7689779E2B8841827"): big.NewInt(5),
	common.HexToAddress("0xF3C456FAAa70fea307A073C3DA9572413c77f58B"): big.NewInt(6),
	common.HexToAddress("0x45D3a03E8302de659e7Ea7400C4cfe9CAED8c723"): big.NewInt(6),
	common.HexToAddress("0x313c7f7126486fFefCaa9FEA92D968cbf891b80c"): big.NewInt(3),
	common.HexToAddress("0x1893eD78480267D1854373A99Cee8dE2E08d430F"): big.NewInt(2),
	common.HexToAddress("0xa2E4D94c5923A8dd99c5792A7B0436474c54e1E1"): big.NewInt(2),
	common.HexToAddress("0x2a4636A5a1138e35F7f93e81FA56d3c970BC6777"): big.NewInt(1),
}

func testDBAndStorage(t *testing.T) (*censusdb.CensusDB, storagelayer.Storage) {
	t.Helper()
	c := qt.New(t)
	tempDir := t.TempDir()

	ipfsConfig := storagelayer.IPFSNewConfig(tempDir)
	storage, err := storagelayer.Init(storagelayer.IPFS, ipfsConfig)
	c.Assert(err, qt.IsNil)

	// init the database for the census trees
	censusesDB, err := metadb.New(vocdoniDB.TypePebble, filepath.Join(tempDir, "tempcensusdb"))
	c.Assert(err, qt.IsNil)

	return censusdb.NewCensusDB(censusesDB), storage
}

func TestCreateAndPublish(t *testing.T) {
	c := qt.New(t)
	db, storage := testDBAndStorage(t)
	// create a census
	opts := CensusOptions{
		ID:      1,
		Type:    defaultCensusType,
		Holders: testHolders,
	}
	root, _, dump, err := CreateAndPublishCensus(db, storage, opts)
	c.Assert(err, qt.IsNil)
	// encode id
	bID := make([]byte, 8)
	binary.LittleEndian.PutUint64(bID, opts.ID)
	// import the tree with the id
	c.Assert(db.ImportTree(bID, dump), qt.IsNil)
	// load the tree by id
	ref, err := db.Load(bID, nil)
	c.Assert(err, qt.IsNil)
	// check the root
	importedRoot, err := ref.Tree().Root()
	c.Assert(err, qt.IsNil)
	c.Assert(bytes.Equal(root, importedRoot), qt.IsTrue)
	// check the holders
	for addr, balance := range opts.Holders {
		key, err := ref.Tree().Hash(addr.Bytes())
		c.Assert(err, qt.IsNil)
		key = key[:censustree.DefaultMaxKeyLen]
		value := ref.Tree().BigIntToBytes(balance)
		val, err := ref.Tree().Get(key)
		c.Assert(err, qt.IsNil)
		c.Assert(bytes.Equal(value, val), qt.IsTrue)
	}
	db.UnLoad()
}
