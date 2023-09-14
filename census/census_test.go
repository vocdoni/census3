package census

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"math/big"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	qt "github.com/frankban/quicktest"
	"go.vocdoni.io/dvote/api/censusdb"
	"go.vocdoni.io/dvote/censustree"
	"go.vocdoni.io/dvote/data"
	"go.vocdoni.io/dvote/data/compressor"
	"go.vocdoni.io/dvote/db"
	"go.vocdoni.io/dvote/db/metadb"
	"go.vocdoni.io/proto/build/go/models"
)

func NewTestCensusDB(t *testing.T) *CensusDB {
	testTempPath := t.TempDir()
	database, err := metadb.New(db.TypePebble, filepath.Join(testTempPath, "censusdb"))
	qt.Assert(t, err, qt.IsNil)

	storage := new(data.DataMockTest)
	qt.Assert(t, storage.Init(nil), qt.IsNil)
	return &CensusDB{
		treeDB:  database,
		storage: storage,
	}
}

var MonkeysAddresses = map[common.Address]*big.Int{
	common.HexToAddress("0xe54d702f98E312aBA4318E3c6BDba98ab5e11012"): big.NewInt(16),
	common.HexToAddress("0x38d2BC91B89928f78cBaB3e4b1949e28787eC7a3"): big.NewInt(13),
	common.HexToAddress("0xF752B527E2ABA395D1Ba4C0dE9C147B763dDA1f4"): big.NewInt(12),
	common.HexToAddress("0xdeb8699659bE5d41a0e57E179d6cB42E00B9200C"): big.NewInt(11),
	common.HexToAddress("0xe1308a8d0291849bfFb200Be582cB6347FBE90D9"): big.NewInt(9),
	common.HexToAddress("0xB1F05B11Ba3d892EdD00f2e7689779E2B8841827"): big.NewInt(6),
	common.HexToAddress("0xF3C456FAAa70fea307A073C3DA9572413c77f58B"): big.NewInt(6),
	common.HexToAddress("0x45D3a03E8302de659e7Ea7400C4cfe9CAED8c723"): big.NewInt(6),
	common.HexToAddress("0x313c7f7126486fFefCaa9FEA92D968cbf891b80c"): big.NewInt(3),
}

func TestNewCensusDB(t *testing.T) {
	c := qt.New(t)
	_, err := NewCensusDB("/", "")
	c.Assert(err, qt.IsNotNil)
	c.Assert(err, qt.ErrorIs, ErrCreatingCensusDB)

	cdb, err := NewCensusDB(t.TempDir(), "")
	c.Assert(err, qt.IsNil)
	c.Assert(cdb.ipfsConn, qt.IsNil)
	c.Assert(cdb.storage.Stop(), qt.IsNil)

	cdb, err = NewCensusDB(t.TempDir(), "test")
	c.Assert(err, qt.IsNil)
	c.Assert(cdb.ipfsConn, qt.IsNotNil)
	c.Assert(cdb.storage.Stop(), qt.IsNil)
}

func TestCreateAndPublish(t *testing.T) {
	c := qt.New(t)
	cdb := NewTestCensusDB(t)
	defer func() {
		c.Assert(cdb.storage.Stop(), qt.IsNil)
	}()

	censusDefinition := NewCensusDefinition(1, 1, MonkeysAddresses, false)
	publishedCensus, err := cdb.CreateAndPublish(censusDefinition)
	c.Assert(err, qt.IsNil)

	importedCensusDefinition := NewCensusDefinition(1, 1, nil, false)
	importedCensusDefinition, err = cdb.newTree(importedCensusDefinition)
	c.Assert(err, qt.IsNil)

	dump := censusdb.CensusDump{}
	c.Assert(json.Unmarshal(publishedCensus.Dump, &dump), qt.IsNil)
	ddata := compressor.NewCompressor().DecompressBytes(dump.Data)
	c.Assert(importedCensusDefinition.tree.ImportDump(ddata), qt.IsNil)
	root, err := importedCensusDefinition.tree.Root()
	c.Assert(err, qt.IsNil)
	c.Assert(publishedCensus.RootHash, qt.ContentEquals, root)

	for addr, val := range MonkeysAddresses {
		key, err := importedCensusDefinition.tree.Hash(addr.Bytes())
		c.Assert(err, qt.IsNil)
		tval, _, err := importedCensusDefinition.tree.GenProof(key[:censustree.DefaultMaxKeyLen])
		c.Assert(err, qt.IsNil)
		c.Assert(tval, qt.ContentEquals, importedCensusDefinition.tree.BigIntToBytes(val))
	}
}

func Test_newTree(t *testing.T) {
	c := qt.New(t)
	cdb := NewTestCensusDB(t)
	defer func() {
		c.Assert(cdb.storage.Stop(), qt.IsNil)
	}()

	_, err := cdb.newTree(&CensusDefinition{
		ID:        1,
		MaxLevels: defaultMaxLevels,
		Type:      models.Census_UNKNOWN,
	})
	c.Assert(err, qt.IsNotNil)

	_, err = cdb.newTree(NewCensusDefinition(0, 0, map[common.Address]*big.Int{}, false))
	c.Assert(err, qt.IsNil)
}

func Test_save(t *testing.T) {
	c := qt.New(t)
	cdb := NewTestCensusDB(t)
	defer func() {
		c.Assert(cdb.storage.Stop(), qt.IsNil)
	}()

	def := NewCensusDefinition(0, 0, map[common.Address]*big.Int{}, false)
	_, err := cdb.treeDB.Get([]byte(censusDBKey(def.ID)))
	c.Assert(err, qt.IsNotNil)

	bdef := bytes.Buffer{}
	encoder := gob.NewEncoder(&bdef)
	c.Assert(encoder.Encode(def), qt.IsNil)

	c.Assert(cdb.save(def), qt.IsNil)
	res, err := cdb.treeDB.Get([]byte(censusDBKey(def.ID)))
	c.Assert(err, qt.IsNil)
	c.Assert(res, qt.ContentEquals, bdef.Bytes())
}

func Test_publish(t *testing.T) {
	c := qt.New(t)
	cdb := NewTestCensusDB(t)
	defer func() {
		c.Assert(cdb.storage.Stop(), qt.IsNil)
	}()

	def, err := cdb.newTree(NewCensusDefinition(0, 0, MonkeysAddresses, false))
	c.Assert(err, qt.IsNil)

	keys, values := [][]byte{}, [][]byte{}
	for addr, balance := range def.Holders {
		value := def.tree.BigIntToBytes(balance)
		key, err := def.tree.Hash(addr.Bytes())
		c.Assert(err, qt.IsNil)

		keys = append(keys, key[:censustree.DefaultMaxKeyLen])
		values = append(values, value)
	}

	_, err = def.tree.AddBatch(keys, values)
	c.Assert(err, qt.IsNil)
	ref, err := cdb.publish(def)
	c.Assert(err, qt.IsNil)
	cid, ok := strings.CutPrefix(ref.URI, cdb.storage.URIprefix())
	c.Assert(ok, qt.IsTrue)
	dump, err := cdb.storage.Retrieve(context.TODO(), cid, 0)
	c.Assert(err, qt.IsNil)
	c.Assert(dump, qt.ContentEquals, ref.Dump)
}

func Test_delete(t *testing.T) {
	c := qt.New(t)
	cdb := NewTestCensusDB(t)
	defer func() {
		c.Assert(cdb.storage.Stop(), qt.IsNil)
	}()

	def := NewCensusDefinition(0, 0, map[common.Address]*big.Int{}, false)
	c.Assert(cdb.save(def), qt.IsNil)

	_, err := cdb.treeDB.Get([]byte(censusDBKey(def.ID)))
	c.Assert(err, qt.IsNil)
	c.Assert(cdb.delete(def), qt.IsNil)

	_, err = cdb.treeDB.Get([]byte(censusDBKey(def.ID)))
	c.Assert(err, qt.IsNotNil)
}
