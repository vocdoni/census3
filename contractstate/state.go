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

type ContractState struct {
	// address is the contract address
	address common.Address
	// ctype is the contract type (erc20, erc721, erc1155, erc777, custom...)
	ctype ContractType
	// decimals is the number of decimals if the contract is a token contract
	decimals    *big.Float
	treeDataDir string
	tree        *dvoteTree.Tree
	blocksKV    db.Database

	snapshotLock sync.RWMutex
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
	t.decimals = big.NewFloat(math.Pow(10, float64(decimals)))
	return nil
}

func (t *ContractState) Address() common.Address {
	return t.address
}

func (t *ContractState) Type() ContractType {
	return t.ctype
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

func (t *ContractState) Reset(address common.Address) error {
	t.snapshotLock.Lock()
	defer t.snapshotLock.Unlock()
	return t.store(address, big.NewInt(0))
}

func (t *ContractState) LastRoot() ([]byte, error) {
	t.snapshotLock.RLock()
	defer t.snapshotLock.RUnlock()
	return t.tree.Root(t.tree.DB().ReadTx())
}

func (t *ContractState) BlockRootHash(blocknum uint64) ([]byte, error) {
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
	log.Debugf("snapshot took %s", time.Since(startTime))
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
	log.Debugf("snapshot import took %s", time.Since(startTime))
	return err
}

func (t *ContractState) Remove() error {
	t.snapshotLock.Lock()
	defer t.snapshotLock.Unlock()
	log.Debugf("removing tree...")
	if err := t.removeTree(); err != nil {
		return err
	}
	log.Debugf("create new tree...")
	return t.loadTree()
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

func (t *ContractState) ExportTree() ([]byte, error) {
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

func (t *ContractState) ImportTree(dump []byte) error {
	t.snapshotLock.Lock()
	defer t.snapshotLock.Unlock()
	return t.tree.ImportDump(dump)
}

func (t *ContractState) SaveBlock(blocknum uint64) error {
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

func (t *ContractState) CloseBlocksDB() {
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

func (t *ContractState) Holders() map[common.Address]*big.Int {
	t.snapshotLock.RLock()
	defer t.snapshotLock.RUnlock()
	holders := make(map[common.Address]*big.Int)
	t.tree.IterateLeaves(nil, func(k, v []byte) bool {
		af := new(big.Int).SetBytes(v)
		if af.Cmp(big.NewInt(0)) > 0 {
			holders[common.BytesToAddress(k)] = af
		}
		return false
	})
	return holders
}

// TotalHoldersAndAmount returns the number of holders and the total amount.
func (t *ContractState) TotalHoldersAndAmount() (int, *big.Int) {
	t.snapshotLock.RLock()
	defer t.snapshotLock.RUnlock()
	total := big.NewInt(0)
	holders := 0
	t.tree.IterateLeaves(nil, func(k, v []byte) bool {
		af := new(big.Int).SetBytes(v)
		if af.Cmp(big.NewInt(0)) > 0 {
			holders++
		}
		total.Add(total, af)
		return false
	})
	return holders, total
}

func (t *ContractState) loadTree() error {
	treeKV, err := metadb.New(
		db.TypePebble,
		t.treeDataDir,
	)
	if err != nil {
		return err
	}
	if t.tree, err = dvoteTree.New(
		nil,
		dvoteTree.Options{
			DB:        treeKV,
			HashFunc:  arbo.HashFunctionPoseidon,
			MaxLevels: 160,
		},
	); err != nil {
		return err
	}
	return nil
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
