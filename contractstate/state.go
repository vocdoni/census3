package contractstate

import (
	"fmt"
	"math"
	"math/big"
	"path/filepath"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/arbo"
	"go.vocdoni.io/dvote/db"
	"go.vocdoni.io/dvote/db/metadb"
	"go.vocdoni.io/dvote/tree"
)

const (
	amountsTreeName = "amounts"
	blocksTreeName  = "blocks"
)

type ContractState struct {
	storageTree *tree.Tree
	blocksKV    db.Database
	decimals    *big.Float
}

func (t *ContractState) Init(datadir, contract string, decimals int) error {
	treeKV, err := metadb.New(
		db.TypePebble,
		filepath.Join(datadir, contract, "tree"),
	)
	if err != nil {
		return err
	}
	t.storageTree, err = tree.New(
		nil,
		tree.Options{
			DB:        treeKV,
			HashFunc:  arbo.HashFunctionPoseidon,
			MaxLevels: 160,
		},
	)
	if err != nil {
		return err
	}
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

func (t *ContractState) store(address common.Address, amount *big.Int) error {
	wTx := t.storageTree.DB().WriteTx()
	defer wTx.Discard()
	// Warning: key+value creates the index? so we will be always adding new leafs?
	if err := t.storageTree.Set(wTx, address.Bytes(), amount.Bytes()); err != nil {
		return fmt.Errorf("cannot store amount %s for address %x: %w", amount.String(), address.Bytes(), err)
	}
	//log.Debugf("storing address %s with balance %s", address, amount.String())
	return wTx.Commit()

}

func (t *ContractState) Add(address common.Address, amount *big.Int) error {
	stAmount, err := t.Get(address)
	if err != nil {
		return err
	}
	stAmount.Add(stAmount, amount)
	return t.store(address, stAmount)
}

func (t *ContractState) Sub(address common.Address, amount *big.Int) error {
	stAmount, err := t.Get(address)
	if err != nil {
		return err
	}
	stAmount.Sub(stAmount, amount)
	return t.store(address, stAmount)
}

func (t *ContractState) Save(blocknum uint64) error {
	root, err := t.storageTree.Root(t.storageTree.DB().ReadTx())
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
	tx := t.blocksKV.ReadTx()
	defer tx.Discard()
	_, err := tx.Get([]byte(fmt.Sprintf("%d", blocknum)))
	return err == nil
}

func (t *ContractState) Close() {
	t.blocksKV.Close()
	t.storageTree.DB().Close()
}

func (t *ContractState) Get(address common.Address) (*big.Int, error) {
	stAmountBytes, err := t.storageTree.Get(t.storageTree.DB().ReadTx(), address.Bytes())
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
	amounts := make(map[string]*big.Float)
	total := big.NewFloat(0)
	zero := big.NewFloat(0)
	holders := 0
	t.storageTree.IterateLeaves(nil, func(k, v []byte) bool {
		af := new(big.Float).SetInt(new(big.Int).SetBytes(v))
		af.Quo(af, t.decimals)
		if af.Cmp(zero) > 0 {
			amounts[common.BytesToAddress(k).Hex()] = af
			total.Add(total, af)
			holders++

		}
		return false
	})
	amounts["holders"] = big.NewFloat(float64(holders))
	amounts["total"] = total
	return amounts
}
