package scanner

// The filter package provides a wrapper of boom.ScalableBloomFilter to store
// the filter to a file and load it from it. The filter is used to store the
// processed transactions to avoid re-processing them, but also rescanning a
// synced token to find missing transactions.

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	boom "github.com/tylertreat/BoomFilters"
)

// TokenFilter is a wrapper of boom.ScalableBloomFilter to store the filter to
// a file and load it from it. The file that stores the filter is named as
// <address>-<chainID>.filter, where address is the token contract address and
// chainID is the chain ID of the network where the token is deployed.
type TokenFilter struct {
	filter  *boom.ScalableBloomFilter
	address common.Address
	chainID uint64
	path    string
}

// LoadFilter loads the filter from the file, if the file does not exist, create
// a new filter and return it. The filter is stored in the file named as
// <address>-<chainID>.filter in the basePath directory.
func LoadFilter(basePath string, address common.Address, chainID uint64) (*TokenFilter, error) {
	// compose the filter path: path/<address>-<chainID>.filter
	// by default, create a empty filter
	tf := &TokenFilter{
		filter:  boom.NewDefaultScalableBloomFilter(0.01),
		address: address,
		chainID: chainID,
		path:    fmt.Sprintf("%s/%s-%d.filter", basePath, address.Hex(), chainID),
	}
	// read the filter from the file, if it not exists, create a new one
	bFilter, err := os.ReadFile(tf.path)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
		return tf, nil
	}
	// decode the filter from the file content
	if err := tf.filter.GobDecode(bFilter); err != nil {
		return nil, err
	}
	return tf, nil
}

// Add adds a key to the filter.
func (tf *TokenFilter) Add(key []byte) {
	tf.filter.Add(key)
}

// Test checks if a key is in the filter.
func (tf *TokenFilter) Test(key []byte) bool {
	return tf.filter.Test(key)
}

// TestAndAdd checks if a key is in the filter, if not, add it to the filter. It
// is the combination of Test and conditional Add.
func (tf *TokenFilter) TestAndAdd(key []byte) bool {
	return tf.filter.TestAndAdd(key)
}

// Commit writes the filter to its file.
func (tf *TokenFilter) Commit() error {
	// encode the filter
	bFilter, err := tf.filter.GobEncode()
	if err != nil {
		return err
	}
	// write the filter to the file
	if err := os.WriteFile(tf.path, bFilter, 0o644); err != nil {
		return err
	}
	return nil
}
