package filter

// The filter package provides a wrapper of mechanism to store a list of
// identifiers to avoid re-processing them. It splits the list of identifiers
// into buckets, each bucket has a fixed size, and stores the buckets to files.
// This allows to reduce the write operations to the file system.

type TokenFilter struct{}

// LoadFilter loads a filter from a file.
func LoadFilter(basePath, fileName string) (*TokenFilter, error) { return &TokenFilter{}, nil }

// Add adds a key to the filter.
func (tf *TokenFilter) Add(key []byte) {}

// Test checks if a key is in the filter.
func (tf *TokenFilter) Test(key []byte) bool { return false }

// TestAndAdd checks if a key is in the filter, if not, add it to the filter. It
// is the combination of Test and conditional Add.
func (tf *TokenFilter) TestAndAdd(key []byte) bool { return false }

// Commit writes the filter to its file.
func (tf *TokenFilter) Commit() error { return nil }

func (tf *TokenFilter) loadLocalFilters() error { return nil }
func (tf *TokenFilter) add(key ...string)       {}
func (tf *TokenFilter) test(key string) bool    { return false }
