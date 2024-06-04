package scanner

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	boom "github.com/tylertreat/BoomFilters"
)

type TokenFilter struct {
	filter  *boom.ScalableBloomFilter
	address common.Address
	chainID uint64
	path    string
}

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

func (tf *TokenFilter) Add(key []byte) {
	tf.filter.Add(key)
}

func (tf *TokenFilter) Test(key []byte) bool {
	return tf.filter.Test(key)
}

func (tf *TokenFilter) TestAndAdd(key []byte) bool {
	return tf.filter.TestAndAdd(key)
}

func (tf *TokenFilter) Commit() error {
	// encode the filter
	bFilter, err := tf.filter.GobEncode()
	if err != nil {
		return err
	}
	// write the filter to the file
	if err := os.WriteFile(tf.path, bFilter, 0644); err != nil {
		return err
	}
	return nil
}
