package web3

import "fmt"

var (
	ErrInvalidProviderAttributes = fmt.Errorf("invalid provider attributes")
	ErrInitializingContract      = fmt.Errorf("error initializing token contract")
	ErrScanningTokenLogs         = fmt.Errorf("error scanning token logs")
	ErrParsingTokenLogs          = fmt.Errorf("error parsing token logs")
)
