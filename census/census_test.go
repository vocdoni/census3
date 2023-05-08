package census

import (
	"bytes"
	"encoding/json"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"go.vocdoni.io/dvote/api/censusdb"
	"go.vocdoni.io/dvote/censustree"
	"go.vocdoni.io/dvote/data/compressor"
	"go.vocdoni.io/dvote/util"
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

func Test_validateCensus(t *testing.T) {
	cdb, err := NewCensusDB(t.TempDir(), "vocdonidev", util.RandomHex(32))
	if err != nil {
		t.Fatal(err)
	}

	censusDefinition := DefaultCensusDefinition(1, 1, internalAddresses)
	publishedCensus, err := cdb.CreateAndPublish(censusDefinition)
	if err != nil {
		t.Fatal(err)
	}

	importedCensusDefinition := DefaultCensusDefinition(1, 1, nil)
	importedCensusDefinition, err = cdb.newTree(importedCensusDefinition)
	if err != nil {
		t.Fatal(err)
	}

	dump := censusdb.CensusDump{}
	if err := json.Unmarshal(publishedCensus.Dump, &dump); err != nil {
		t.Fatal(err)
	}

	ddata := compressor.NewCompressor().DecompressBytes(dump.Data)
	if err := importedCensusDefinition.tree.ImportDump(ddata); err != nil {
		t.Fatal(err)
	}

	tree := importedCensusDefinition.tree
	root, _ := tree.Root()
	if !bytes.Equal(publishedCensus.RootHash, root) {
		t.Fatal("roots are not equal")
	}
	for addr, val := range internalAddresses {
		key, err := tree.Hash(addr.Bytes())
		if err != nil {
			t.Fatal(err)
		}
		tval, _, err := tree.GenProof(key[:censustree.DefaultMaxKeyLen])
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(tval, tree.BigIntToBytes(val)) {
			t.Fatalf("value not equals for %s", addr.String())
		}
	}
}
