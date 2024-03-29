package web3

import "fmt"

var (
	ErrClientNotSet           = fmt.Errorf("client not set")
	ErrConnectingToWeb3Client = fmt.Errorf("client not set")
	ErrInitializingContract   = fmt.Errorf("error initializing token contract")
	ErrScanningTokenLogs      = fmt.Errorf("error scanning token logs")
	ErrTooManyRequests        = fmt.Errorf("web3 endpoint returns too many requests")
	ErrParsingTokenLogs       = fmt.Errorf("error parsing token logs")
	ErrGettingTotalSupply     = fmt.Errorf("error getting total supply")
)
