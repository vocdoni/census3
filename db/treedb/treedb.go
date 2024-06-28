package treedb

// The treedb package provides a wrapper of key-value database that uses a
// merkle tree under the hood. Every tree is stored in the same database, but
// with a different prefix.

import (
	"fmt"

	"go.vocdoni.io/dvote/db"
	"go.vocdoni.io/dvote/db/prefixeddb"
	"go.vocdoni.io/dvote/log"
	"go.vocdoni.io/dvote/tree"
	"go.vocdoni.io/dvote/tree/arbo"
)

// filterTreeLevels is the number of levels of the tree used to store the
// filter. It is a constant to avoid re-creating the tree with a different
// number of levels. The available number of leaves is 2^filterTreeLevels.
// It also limits the size of the key to filterTreeLevels/8 bytes.
const filterTreeLevels = 64

// ErrNotInitialized is returned when no tree is initialized in a TreeDB
// instance, which means that LoadTree has not been called and the tree is
// not ready to be used.
var ErrNotInitialized = fmt.Errorf("tree not initialized, call LoadTree first")

// TokenFilter is a filter associated with a token.
type TreeDB struct {
	prefix   string
	parentDB db.Database
	tree     *tree.Tree
}

// LoadTree loads a tree from the database identified by the given prefix. If it
// does not exist, it creates a new tree with the given prefix. It also creates
// the index if it does not exist. It returns an error if the tree cannot be
// loaded or created.
func LoadTree(db db.Database, prefix string) (*TreeDB, error) {
	treeDB := prefixeddb.NewPrefixedDatabase(db, []byte(prefix))
	tree, err := tree.New(nil, tree.Options{
		DB:        treeDB,
		MaxLevels: filterTreeLevels,
		HashFunc:  arbo.HashFunctionBlake2b,
	})
	if err != nil {
		return nil, err
	}
	// ensure index is created
	wTx := tree.DB().WriteTx()
	defer wTx.Discard()
	return &TreeDB{
		prefix:   prefix,
		parentDB: db,
		tree:     tree,
	}, wTx.Commit()
}

// Close closes the tree database. If the tree is not nil, it closes the
// underlying database. If the parent database is not nil, it closes it too.
// It returns an error if any of the databases cannot be closed.
func (tdb *TreeDB) Close() error {
	if tdb.tree != nil {
		if err := tdb.tree.DB().Close(); err != nil {
			return err
		}
	}
	if tdb.parentDB != nil {
		return tdb.parentDB.Close()
	}
	return nil
}

// Purge deletes a tree from the database identified by current prefix. It
// iterates over all the keys in the tree and deletes them. If some key cannot
// be deleted, it logs a warning and continues with the next key. It commits the
// transaction at the end.
func (tdb *TreeDB) Purge() error {
	treeDB := prefixeddb.NewPrefixedDatabase(tdb.parentDB, []byte(tdb.prefix))
	wTx := treeDB.WriteTx()
	if err := treeDB.Iterate(nil, func(k, _ []byte) bool {
		if err := wTx.Delete(k); err != nil {
			log.Warnw("error deleting key", "key", k, "err", err)
		}
		return true
	}); err != nil {
		return err
	}
	return wTx.Commit()
}

// Add adds a key to the tree. It no write transaction is provided, it creates
// a new one and commits it at the end. It returns an error if the tree is not
// initialized, if there is an error adding the key-value pair or committing
// the transaction if it was created. If a transaction is provided, it does
// not commit or discard it.
func (tdb *TreeDB) Add(wtx db.WriteTx, key, value []byte) error {
	if tdb.tree == nil {
		return ErrNotInitialized
	}
	commitTx := wtx == nil
	if commitTx {
		wtx = tdb.tree.DB().WriteTx()
		defer wtx.Discard()
	}
	if err := tdb.tree.Add(wtx, key, value); err != nil {
		return err
	}
	if commitTx {
		return wtx.Commit()
	}
	return nil
}

// Del deletes a key from the tree. If no write transaction is provided, it
// creates a new one and commits it at the end. It returns an error if the tree
// is not initialized, if there is an error deleting the key-value pair or
// committing the transaction if it was provided. If a transaction is provided,
// it does not commit or discard it.
func (tdb *TreeDB) Del(wtx db.WriteTx, key []byte) error {
	if tdb.tree == nil {
		return ErrNotInitialized
	}
	commitTx := wtx == nil
	if commitTx {
		wtx = tdb.tree.DB().WriteTx()
		defer wtx.Discard()
	}
	if err := tdb.tree.Del(wtx, key); err != nil {
		return err
	}
	if commitTx {
		return wtx.Commit()
	}
	return nil
}

// AddBatch adds a batch of keys and values to the tree. It is more efficient
// than calling Add for each key-value pair. It returns an error if the length
// of keys and values is different, if the tree is not initialized, if there
// is an error adding a key-value pair or committing the transaction. It uses
// a new write transaction to add all the keys and commits it at the end. If
// something goes wrong, it returns an error and discards the transaction.
func (tdb *TreeDB) AddBatch(keys, values [][]byte) error {
	if tdb.tree == nil {
		return ErrNotInitialized
	}
	if len(keys) != len(values) {
		return fmt.Errorf("keys and values must have the same length")
	}
	wTx := tdb.tree.DB().WriteTx()
	defer wTx.Discard()
	for i := range keys {
		if err := tdb.tree.Add(wTx, keys[i], values[i]); err != nil {
			return err
		}
	}
	return wTx.Commit()
}

// AddKey adds a key to the tree with nil value. It accepts variadic keys. It
// uses a new write transaction to add all the keys and commits it at the end.
// If something goes wrong, it returns an error and discards the transaction.
func (tdb *TreeDB) AddKey(key ...[]byte) error {
	if tdb.tree == nil {
		return ErrNotInitialized
	}
	wtx := tdb.tree.DB().WriteTx()
	defer wtx.Discard()
	for _, k := range key {
		if err := tdb.tree.Add(wtx, k, nil); err != nil {
			return err
		}
	}
	return wtx.Commit()
}

// CheckKeyKey checks if a key is in the tree.
func (tdb *TreeDB) CheckKey(key []byte) (bool, error) {
	if tdb.tree == nil {
		return false, ErrNotInitialized
	}
	_, err := tdb.tree.Get(nil, key)
	if err != nil {
		if err == arbo.ErrKeyNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// CheckAndAddKey checks if a key is in the tree, if not, add it to the tree. It
// is the combination of CheckKey and conditional AddKey.
func (tdb *TreeDB) CheckAndAddKey(key []byte) (bool, error) {
	exists, err := tdb.CheckKey(key)
	if err != nil {
		return false, err
	}
	if exists {
		return true, nil
	}
	return false, tdb.AddKey(key)
}
