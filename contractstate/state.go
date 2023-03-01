package contractstate

import (
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"go.vocdoni.io/dvote/db"
	"go.vocdoni.io/dvote/db/metadb"
	"go.vocdoni.io/dvote/log"
	arbo "go.vocdoni.io/dvote/tree/arbo"
	"go.vocdoni.io/dvote/types"
)

type ContractState struct {
	// address is the contract address
	address common.Address
	// ctype is the contract type (erc20, erc721, erc1155, erc777, custom...)
	ctype       ContractType
	treeDataDir string
	tree        *arbo.Tree
	treeDB      db.Database
	blocksKV    db.Database

	mutex sync.RWMutex
}

type Proof struct {
	Value    types.HexBytes
	Siblings types.HexBytes
}

func (t *ContractState) Init(datadir string, address common.Address, ctype ContractType, decimals int) error {
	log.Infof("initializing contract %s of type %d", address, ctype)
	t.treeDataDir = filepath.Join(datadir, address.Hex(), "tree")
	if err := t.loadTree(); err != nil {
		return err
	}
	var err error
	if t.blocksKV, err = metadb.New(
		db.TypePebble,
		filepath.Join(datadir, address.Hex(), "blocks"),
	); err != nil {
		return err
	}
	t.address = address
	t.ctype = ctype
	return nil
}

func (t *ContractState) Address() common.Address {
	return t.address
}

func (t *ContractState) Type() ContractType {
	return t.ctype
}

func (t *ContractState) Add(address common.Address, amount *big.Int) error {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	tAmount, err := t.Get(address)
	if err != nil {
		return err
	}
	log.Debugf("Adding %s to %s", amount, address.Hex())
	tAmount.Add(tAmount, amount)
	return t.store(address, tAmount)
}

func (t *ContractState) Sub(address common.Address, amount *big.Int) error {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	tAmount, err := t.Get(address)
	if err != nil {
		return err
	}
	log.Debugf("Subtracting %s to %s", amount, address.Hex())
	tAmount.Sub(tAmount, amount)
	return t.store(address, tAmount)
}

func (t *ContractState) LastRoot() ([]byte, error) {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	rTx := t.treeDB.ReadTx()
	defer rTx.Discard()
	return t.tree.RootWithTx(rTx)
}

func (t *ContractState) BlockRootHash(blocknum uint64) ([]byte, error) {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	tx := t.blocksKV.ReadTx()
	defer tx.Discard()
	return tx.Get([]byte(fmt.Sprintf("%d", blocknum)))
}

func (t *ContractState) Snapshot() error {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	rTx := t.treeDB.ReadTx()
	defer rTx.Discard()
	root, err := t.tree.RootWithTx(rTx)
	if err != nil {
		return err
	}
	startTime := time.Now()
	log.Infof("performing snapshot at root %x", root)
	dump, err := t.tree.Dump(nil)
	if err != nil {
		return err
	}
	log.Debugf("snapshot took %s", time.Since(startTime))
	log.Infof("snapshot dump has %d bytes", len(dump))
	log.Debugf("removing tree...")
	t.treeDB.Close()
	if err := os.RemoveAll(t.treeDataDir); err != nil {
		return err
	}
	log.Debugf("create new tree...")
	if err := t.loadTree(); err != nil {
		return err
	}
	startTime = time.Now()
	log.Debugf("importing dump...")
	if err := t.tree.ImportDump(dump); err != nil {
		return err
	}
	log.Debugf("snapshot import took %s", time.Since(startTime))
	return err
}

func (t *ContractState) GenProof(key []byte) (*Proof, error) {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	value, siblings, _, _, err := t.tree.GenProof(key)
	if err != nil {
		return nil, err
	}
	return &Proof{
		Value:    value,
		Siblings: siblings,
	}, nil
}

func (t *ContractState) ExportTree() ([]byte, error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	log.Debugf("creating export dump...")
	dump, err := t.tree.Dump(nil)
	if err != nil {
		return nil, err
	}
	log.Debugf("dump bytes %d", len(dump))
	return dump, nil
}

func (t *ContractState) ImportTree(dump []byte) error {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return t.tree.ImportDump(dump)
}

func (t *ContractState) SaveBlock(blocknum uint64) error {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	rTx := t.treeDB.ReadTx()
	defer rTx.Discard()
	root, err := t.tree.RootWithTx(rTx)
	if err != nil {
		return err
	}
	wtx := t.blocksKV.WriteTx()
	defer wtx.Discard()
	if err := wtx.Set([]byte(fmt.Sprintf("%d", blocknum)), root); err != nil {
		return err
	}
	return wtx.Commit()
}

func (t *ContractState) HasBlock(blocknum uint64) bool {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	tx := t.blocksKV.ReadTx()
	defer tx.Discard()
	_, err := tx.Get([]byte(fmt.Sprintf("%d", blocknum)))
	return err == nil
}

func (t *ContractState) Get(address common.Address) (*big.Int, error) {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	rTx := t.treeDB.ReadTx()
	defer rTx.Discard()
	_, amountBytes, err := t.tree.GetWithTx(rTx, address.Bytes())
	if err != nil && err != arbo.ErrKeyNotFound {
		return nil, err
	}
	if err == arbo.ErrKeyNotFound {
		return new(big.Int).SetUint64(0), nil
	}
	return arbo.BytesToBigInt(amountBytes), nil
}

func (t *ContractState) Holders() map[common.Address]*big.Int {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	holders := make(map[common.Address]*big.Int)
	rTx := t.treeDB.ReadTx()
	defer rTx.Discard()
	if err := t.tree.IterateWithTx(rTx, nil, func(k, v []byte) {
		if v[0] != arbo.PrefixValueLeaf {
			return // skip non-leaf nodes
		}
		leafK, leafV := arbo.ReadLeafValue(v)
		af := arbo.BytesToBigInt(leafV)
		if af.Cmp(big.NewInt(0)) > 0 {
			holders[common.BytesToAddress(leafK)] = af
		}
		log.Debugf("Got %s with amount %s", common.BytesToAddress(leafK).String(), af.String())
	}); err != nil {
		log.Errorf("error iterating tree: %v", err)
		return nil
	}
	return holders
}

// TotalHoldersAndAmount returns the number of holders and the total amount.
func (t *ContractState) TotalHoldersAndAmount() (int, *big.Int) {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	total := big.NewInt(0)
	holders := 0
	rTx := t.treeDB.ReadTx()
	defer rTx.Discard()
	if err := t.tree.IterateWithTx(rTx, nil, func(k, v []byte) {
		if v[0] != arbo.PrefixValueLeaf {
			return // skip non-leaf nodes
		}
		_, leafV := arbo.ReadLeafValue(v)
		af := arbo.BytesToBigInt(leafV)
		if af.Cmp(big.NewInt(0)) > 0 {
			holders++
		}
		total.Add(total, af)
	}); err != nil {
		log.Errorf("error iterating tree: %v", err)
		return 0, nil
	}
	return holders, total
}

func (t *ContractState) loadTree() error {
	var err error
	if t.treeDB, err = metadb.New(
		db.TypePebble,
		t.treeDataDir,
	); err != nil {
		return err
	}
	if t.tree, err = arbo.NewTree(
		arbo.Config{
			Database:     t.treeDB,
			HashFunction: arbo.HashFunctionPoseidon,
			MaxLevels:    160,
		},
	); err != nil {
		return err
	}
	return nil
}

func (t *ContractState) store(address common.Address, amount *big.Int) error {
	log.Debugf("Storing %s for %s", amount, address.Hex())
	rTx := t.treeDB.ReadTx()
	defer rTx.Discard()
	if _, _, err := t.tree.GetWithTx(rTx, address.Bytes()); err != nil {
		if err != arbo.ErrKeyNotFound {
			return err
		}
		wTx := t.treeDB.WriteTx()
		defer wTx.Discard()
		if err := t.tree.AddWithTx(wTx, address.Bytes(), arbo.BigIntToBytes(len(amount.Bytes()), amount)); err != nil {
			return err
		}
		return wTx.Commit()
	}
	wTx := t.treeDB.WriteTx()
	defer wTx.Discard()
	if err := t.tree.Update(address.Bytes(), arbo.BigIntToBytes(len(amount.Bytes()), amount)); err != nil {
		return err
	}
	return wTx.Commit()
}
