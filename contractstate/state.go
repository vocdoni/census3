package contractstate

import (
	"fmt"
	"math"
	"math/big"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/arbo"
	"go.vocdoni.io/dvote/db"
	"go.vocdoni.io/dvote/db/metadb"
	"go.vocdoni.io/dvote/log"
	dvoteTree "go.vocdoni.io/dvote/tree"
	"go.vocdoni.io/dvote/types"
)

const (
	amountsTreeName = "amounts"
	blocksTreeName  = "blocks"
)

type ContractState struct {
	tree         *dvoteTree.Tree
	treeDataDir  string
	blocksKV     db.Database
	decimals     *big.Float
	snapshotLock sync.RWMutex
}

type Proof struct {
	Value    types.HexBytes
	Siblings types.HexBytes
}

func (t *ContractState) Init(datadir, contract string, decimals int) error {
	t.treeDataDir = filepath.Join(datadir, contract, "tree")
	if err := t.loadTree(); err != nil {
		return err
	}
	var err error
	t.blocksKV, err = metadb.New(
		db.TypePebble,
		filepath.Join(datadir, contract, "blocks"),
	)
	if err != nil {
		return err
	}
	t.decimals = big.NewFloat(math.Pow(10, float64(decimals)))
	return nil
}

func (t *ContractState) Add(address common.Address, amount *big.Int) error {
	t.snapshotLock.RLock()
	defer t.snapshotLock.RUnlock()
	stAmount, err := t.Get(address)
	if err != nil {
		return err
	}
	stAmount.Add(stAmount, amount)
	return t.store(address, stAmount)
}

func (t *ContractState) Sub(address common.Address, amount *big.Int) error {
	t.snapshotLock.RLock()
	defer t.snapshotLock.RUnlock()
	stAmount, err := t.Get(address)
	if err != nil {
		return err
	}
	stAmount.Sub(stAmount, amount)
	return t.store(address, stAmount)
}

func (t *ContractState) LastRoot() ([]byte, error) {
	t.snapshotLock.RLock()
	defer t.snapshotLock.RUnlock()
	return t.tree.Root(t.tree.DB().ReadTx())
}

func (t *ContractState) Root(blocknum uint64) ([]byte, error) {
	t.snapshotLock.RLock()
	defer t.snapshotLock.RUnlock()
	tx := t.blocksKV.ReadTx()
	defer tx.Discard()
	return tx.Get([]byte(fmt.Sprintf("%d", blocknum)))
}

func (t *ContractState) Snapshot() error {
	t.snapshotLock.Lock()
	defer t.snapshotLock.Unlock()
	root, err := t.tree.Root(t.tree.DB().ReadTx())
	if err != nil {
		return err
	}
	startTime := time.Now()
	log.Infof("performing snapshot at root %x", root)
	dump, err := t.tree.Dump()
	if err != nil {
		return err
	}
	log.Debugf("snapshot took %s", time.Now().Sub(startTime))
	log.Infof("snapshot dump has %d bytes", len(dump))
	log.Debugf("removing tree...")
	if err := t.removeTree(); err != nil {
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
	log.Debugf("snapshot import took %s", time.Now().Sub(startTime))
	return err
}

func (t *ContractState) GenProof(key []byte) (*Proof, error) {
	t.snapshotLock.RLock()
	defer t.snapshotLock.RUnlock()
	tx := t.tree.DB().ReadTx()
	defer tx.Discard()
	value, siblings, err := t.tree.GenProof(tx, key)
	if err != nil {
		return nil, err
	}
	return &Proof{
		Value:    value,
		Siblings: siblings,
	}, nil
}

func (t *ContractState) Export() ([]byte, error) {
	t.snapshotLock.Lock()
	defer t.snapshotLock.Unlock()
	log.Debugf("creating export dump...")
	dump, err := t.tree.Dump()
	if err != nil {
		return nil, err
	}
	log.Debugf("dump bytes %d", len(dump))
	return dump, nil
}

func (t *ContractState) Import(dump []byte) error {
	t.snapshotLock.Lock()
	defer t.snapshotLock.Unlock()
	return t.tree.ImportDump(dump)
}

func (t *ContractState) Save(blocknum uint64) error {
	t.snapshotLock.Lock()
	defer t.snapshotLock.Unlock()
	root, err := t.tree.Root(t.tree.DB().ReadTx())
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
	t.snapshotLock.RLock()
	defer t.snapshotLock.RUnlock()
	tx := t.blocksKV.ReadTx()
	defer tx.Discard()
	_, err := tx.Get([]byte(fmt.Sprintf("%d", blocknum)))
	return err == nil
}

func (t *ContractState) Close() {
	t.snapshotLock.RLock()
	defer t.snapshotLock.RUnlock()
	t.blocksKV.Close()
	t.tree.DB().Close()
}

func (t *ContractState) Get(address common.Address) (*big.Int, error) {
	t.snapshotLock.RLock()
	defer t.snapshotLock.RUnlock()
	stAmountBytes, err := t.tree.Get(t.tree.DB().ReadTx(), address.Bytes())
	if err != nil && err != arbo.ErrKeyNotFound {
		return nil, err
	}
	stAmount := new(big.Int).SetUint64(0)
	if stAmountBytes != nil {
		stAmount.SetBytes(stAmountBytes)
	}
	return stAmount, nil
}

func (t *ContractState) List() map[string]*big.Float {
	t.snapshotLock.RLock()
	defer t.snapshotLock.RUnlock()
	amounts := make(map[string]*big.Float)
	total := big.NewFloat(0)
	zero := big.NewFloat(0)
	null := common.Address{}
	holders := 0
	t.tree.IterateLeaves(nil, func(k, v []byte) bool {
		af := new(big.Float).SetInt(new(big.Int).SetBytes(v))
		af.Quo(af, t.decimals)
		if af.Cmp(zero) > 0 {
			addr := common.BytesToAddress(k)
			if addr != null {
				amounts[addr.Hex()] = af
				total.Add(total, af)
				holders++
			}
		}
		return false
	})
	amounts["holders"] = big.NewFloat(float64(holders))
	amounts["total"] = total
	return amounts
}

func (t *ContractState) loadTree() error {
	treeKV, err := metadb.New(
		db.TypePebble,
		t.treeDataDir,
	)
	if err != nil {
		return err
	}
	t.tree, err = dvoteTree.New(
		nil,
		dvoteTree.Options{
			DB:        treeKV,
			HashFunc:  arbo.HashFunctionPoseidon,
			MaxLevels: 160,
		},
	)
	return err
}

func (t *ContractState) removeTree() error {
	_ = t.tree.DB().Close()
	return os.RemoveAll(t.treeDataDir)
}

func (t *ContractState) store(address common.Address, amount *big.Int) error {
	wTx := t.tree.DB().WriteTx()
	defer wTx.Discard()
	if err := t.tree.Set(wTx, address.Bytes(), amount.Bytes()); err != nil {
		return fmt.Errorf("cannot store amount %s for address %x: %w", amount.String(), address.Bytes(), err)
	}
	return wTx.Commit()
}
