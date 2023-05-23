package census

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	qt "github.com/frankban/quicktest"
	"go.vocdoni.io/dvote/api/censusdb"
	"go.vocdoni.io/dvote/censustree"
	"go.vocdoni.io/dvote/data/compressor"
	"go.vocdoni.io/proto/build/go/models"
)

var internalAddresses = map[common.Address]*big.Int{
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
	err = cdb.storage.Stop()
	c.Assert(err, qt.IsNil)

	cdb, err = NewCensusDB(t.TempDir(), "test")
	c.Assert(err, qt.IsNil)
	c.Assert(cdb.ipfsConn, qt.IsNotNil)
}

func TestCreateAndPublish(t *testing.T) {
	c := qt.New(t)
	cdb, err := NewCensusDB(t.TempDir(), "")
	c.Assert(err, qt.IsNil)

	censusDefinition := DefaultCensusDefinition(1, 1, internalAddresses)
	publishedCensus, err := cdb.CreateAndPublish(censusDefinition)
	c.Assert(err, qt.IsNil)

	importedCensusDefinition := DefaultCensusDefinition(1, 1, nil)
	importedCensusDefinition, err = cdb.newTree(importedCensusDefinition)
	c.Assert(err, qt.IsNil)

	dump := censusdb.CensusDump{}
	err = json.Unmarshal(publishedCensus.Dump, &dump)
	c.Assert(err, qt.IsNil)
	ddata := compressor.NewCompressor().DecompressBytes(dump.Data)
	err = importedCensusDefinition.tree.ImportDump(ddata)
	c.Assert(err, qt.IsNil)
	tree := importedCensusDefinition.tree
	root, err := tree.Root()
	c.Assert(err, qt.IsNil)
	c.Assert(publishedCensus.RootHash, qt.ContentEquals, root)

	for addr, val := range internalAddresses {
		key, err := tree.Hash(addr.Bytes())
		c.Assert(err, qt.IsNil)
		tval, _, err := tree.GenProof(key[:censustree.DefaultMaxKeyLen])
		c.Assert(err, qt.IsNil)
		c.Assert(tval, qt.ContentEquals, tree.BigIntToBytes(val))
	}
}

func Test_newTree(t *testing.T) {
	c := qt.New(t)
	cdb, err := NewCensusDB(t.TempDir(), "")
	c.Assert(err, qt.IsNil)

	_, err = cdb.newTree(&CensusDefinition{
		ID:        1,
		MaxLevels: defaultMaxLevels,
		Type:      models.Census_UNKNOWN,
	})
	c.Assert(err, qt.IsNotNil)

	_, err = cdb.newTree(DefaultCensusDefinition(0, 0, map[common.Address]*big.Int{}))
	c.Assert(err, qt.IsNil)
}

func Test_save(t *testing.T) {
	c := qt.New(t)
	cdb, err := NewCensusDB(t.TempDir(), "")
	c.Assert(err, qt.IsNil)

	def := DefaultCensusDefinition(0, 0, map[common.Address]*big.Int{})
	rtx := cdb.treeDB.ReadTx()
	_, err = rtx.Get([]byte(censusDBKey(def.ID)))
	c.Assert(err, qt.IsNotNil)

	bdef := bytes.Buffer{}
	encoder := gob.NewEncoder(&bdef)
	err = encoder.Encode(def)
	c.Assert(err, qt.IsNil)

	err = cdb.save(def)
	c.Assert(err, qt.IsNil)
	res, err := rtx.Get([]byte(censusDBKey(def.ID)))
	c.Assert(err, qt.IsNil)
	c.Assert(res, qt.ContentEquals, bdef.Bytes())
}

func Test_publish(t *testing.T) {
	c := qt.New(t)
	cdb, err := NewCensusDB(t.TempDir(), "")
	c.Assert(err, qt.IsNil)

	def, err := cdb.newTree(DefaultCensusDefinition(0, 0, internalAddresses))
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	dump, err := cdb.storage.Retrieve(ctx, ref.URI, 0)
	c.Assert(err, qt.IsNil)
	c.Assert(dump, qt.ContentEquals, ref.Dump)
}

func Test_delete(t *testing.T) {
	c := qt.New(t)
	cdb, err := NewCensusDB(t.TempDir(), "")
	c.Assert(err, qt.IsNil)

	def := DefaultCensusDefinition(0, 0, map[common.Address]*big.Int{})
	err = cdb.save(def)
	c.Assert(err, qt.IsNil)

	rtx := cdb.treeDB.ReadTx()
	_, err = rtx.Get([]byte(censusDBKey(def.ID)))
	c.Assert(err, qt.IsNil)

	err = cdb.delete(def)
	c.Assert(err, qt.IsNil)

	_, err = rtx.Get([]byte(censusDBKey(def.ID)))
	c.Assert(err, qt.IsNotNil)
}
