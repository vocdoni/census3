package tokenstate

import (
	"fmt"
	"math/big"

	state "gitlab.com/vocdoni/go-dvote/statedb/gravitonstate"
)

const amountstree = "amounts"
const blockstree = "blocks"

type TokenState struct {
	storage *state.GravitonState
}

func (t *TokenState) Init(datadir, name string) error {
	var st state.GravitonState
	if err := st.Init(datadir+"/"+name, "disk"); err != nil {
		return err
	}
	if err := st.AddTree(amountstree); err != nil {
		return err
	}
	if err := st.AddTree(blockstree); err != nil {
		return err
	}
	t.storage = &st
	return nil
}

func (t *TokenState) store(address string, amount *big.Float) error {
	return t.storage.Tree(amountstree).Add([]byte(address), []byte(amount.String()))
}

func (t *TokenState) Add(address string, amount *big.Float) error {
	stAmount, err := t.Get(address)
	if err != nil {
		return err
	}
	stAmount.Add(stAmount, amount)
	return t.store(address, stAmount)
}

func (t *TokenState) Sub(address string, amount *big.Float) error {
	stAmount, err := t.Get(address)
	if err != nil {
		return err
	}
	stAmount.Sub(stAmount, amount)
	return t.store(address, stAmount)
}

func (t *TokenState) Save(blocknum uint64) error {
	t.storage.Tree(blockstree).Add([]byte(fmt.Sprintf("%d", blocknum)), t.storage.Tree(amountstree).Hash())
	_, err := t.storage.Commit()
	return err
}

func (t *TokenState) HasBlock(blocknum uint64) bool {
	return t.storage.Tree(blockstree).Get([]byte(fmt.Sprintf("%d", blocknum))) != nil
}

func (t *TokenState) Close() {
	t.storage.Close()
}

func (t *TokenState) Get(address string) (*big.Float, error) {
	stAmountBytes := t.storage.Tree(amountstree).Get([]byte(address))
	stAmount := big.NewFloat(0)
	if stAmountBytes != nil {
		if _, ok := stAmount.SetString(string(stAmountBytes)); !ok {
			return nil, fmt.Errorf("cannot read amount from state tree")
		}
	}
	return stAmount, nil
}

func (t *TokenState) List(blocknumber uint64) map[string]*big.Float {
	amounts := make(map[string]*big.Float)
	total := big.NewFloat(0)
	balance := big.NewFloat(0)
	zero := big.NewFloat(0)
	t.storage.Tree(amountstree).Iterate(nil, func(k, v []byte) bool {
		a := big.Float{}
		a.SetString(string(v))
		amounts[string(k)] = &a
		if a.Cmp(zero) > 0 {
			total.Add(total, &a)
		}
		balance.Add(balance, &a)
		return false
	})
	amounts["total"] = total
	amounts["balance"] = balance
	return amounts
}
