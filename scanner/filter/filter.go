package filter

// The filter package provides a wrapper of boom.ScalableBloomFilter to store
// the filter to a file and load it from it. The filter is used to store the
// processed transactions to avoid re-processing them, but also rescanning a
// synced token to find missing transactions.

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	cuckoo "github.com/panmari/cuckoofilter"
	"go.vocdoni.io/dvote/log"
)

const (
	FilterSize = 5000000 // 5M items
	MaxSize    = 4000000 // 4M items
)

type filterDump struct {
	Filters [][]byte `json:"filters"`
}

type batchFilter struct {
	filter *cuckoo.Filter
	size   uint
	full   bool
}

// TokenFilter is a wrapper of boom.ScalableBloomFilter to store the filter to
// a file and load it from it. The file that stores the filter is named as
// <address>-<chainID>-<externalID>.filter, where address is the token contract
// address and chainID is the chain ID of the network where the token is
// deployed.
type TokenFilter struct {
	path       string
	filters    []*batchFilter
	address    common.Address
	chainID    uint64
	externalID string
}

// LoadFilter loads the filter from the file, if the file does not exist, create
// a new filter and return it. The filter is stored in the file named as
// <address>-<chainID>-<externalID>.filter in the basePath directory.
func LoadFilter(basePath string, address common.Address, chainID uint64, externalID string) (*TokenFilter, error) {
	tf := &TokenFilter{
		path:       fmt.Sprintf("%s/%s-%d-%s.filter", basePath, address.Hex(), chainID, externalID),
		filters:    []*batchFilter{},
		address:    address,
		chainID:    chainID,
		externalID: externalID,
	}
	// load filters from the local file
	empty, err := tf.loadLocalFilters()
	if err != nil {
		return nil, err
	}
	// append a new filter if there is no filter in the local file
	if empty {
		tf.addFilter()
	}
	return tf, nil
}

// Add adds a key to the filter.
func (tf *TokenFilter) Add(key []byte) {
	for _, f := range tf.filters {
		if f.full {
			continue
		}
		if f.filter.Insert(key) {
			f.size++
			f.full = f.size >= MaxSize
			return
		}
	}
	// add a new filter if all filters are full
	tf.addFilter(key)
}

// Test checks if a key is in the filter.
func (tf *TokenFilter) Test(key []byte) bool {
	for _, f := range tf.filters {
		if f.filter.Lookup(key) {
			return true
		}
	}
	return false
}

// TestAndAdd checks if a key is in the filter, if not, add it to the filter. It
// is the combination of Test and conditional Add.
func (tf *TokenFilter) TestAndAdd(key []byte) bool {
	if tf.Test(key) {
		return true
	}
	tf.Add(key)
	return false
}

// Commit writes the filter to its file.
func (tf *TokenFilter) Commit() error {
	filterDump := filterDump{
		Filters: make([][]byte, 0, len(tf.filters)),
	}
	for _, f := range tf.filters {
		filterDump.Filters = append(filterDump.Filters, f.filter.Encode())
	}
	filterBytes, err := json.Marshal(filterDump)
	if err != nil {
		return err
	}
	return os.WriteFile(tf.path, filterBytes, 0644)
}

func (tf *TokenFilter) loadLocalFilters() (bool, error) {
	filterBytes, err := os.ReadFile(tf.path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	var filterDump filterDump
	if err := json.Unmarshal(filterBytes, &filterDump); err != nil {
		return false, err
	}
	if len(filterDump.Filters) == 0 {
		return false, nil
	}
	for _, bFilter := range filterDump.Filters {
		filter, err := cuckoo.Decode(bFilter)
		if err != nil {
			return false, err
		}
		tf.filters = append(tf.filters, &batchFilter{
			filter: filter,
			size:   filter.Count(),
			full:   filter.Count() >= MaxSize,
		})
	}
	return true, nil
}

func (tf *TokenFilter) addFilter(keys ...[]byte) *batchFilter {
	log.Info("adding new filter")
	f := &batchFilter{
		filter: cuckoo.NewFilter(FilterSize),
		size:   0,
		full:   false,
	}
	for _, key := range keys {
		if f.filter.Insert(key) {
			f.size++
		}
	}

	tf.filters = append(tf.filters, f)
	return f
}
