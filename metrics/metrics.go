package metrics

import (
	"github.com/VictoriaMetrics/metrics"
)

const (
	LastAnalysedBlockByChainPrefix = `census3_last_analysed_block_by_chain_`
	NumberOfCensusesByTypePrefix   = `census3_number_of_censuses_by_type_`
)

var (
	// total number of tokens
	TotalNumberOfTokens = metrics.NewCounter(`census3_total_number_of_tokens`)
	// new tokens by time
	NewTokensByTime = metrics.NewHistogram(`census3_new_tokens_by_time_seconds`)
	// number of analysed transfers transactions
	TotalNumberOfTransfers = metrics.NewCounter(`census3_total_number_of_transfers`)
	// number of analysed transfers transactions by time
	TransfersByTime = metrics.NewHistogram(`census3_transfers_by_time_seconds`)
	// last analysed block by chain
	LastAnalysedBlockByChain = metrics.NewSet()
	// total number of strategies
	TotalNumberOfStrategies = metrics.NewCounter(`census3_total_number_of_strategies`)
	// new strategies by time
	NewStrategiesByTime = metrics.NewHistogram(`census3_new_strategies_by_time_seconds`)
	// total number of censuses
	TotalNumberOfCensuses = metrics.NewCounter(`census3_total_number_of_censuses`)
	// number of censuses by type (anonymous or not)
	NumberOfCensusesByType = metrics.NewSet()
)

func init() {
	metrics.RegisterSet(LastAnalysedBlockByChain)
	metrics.RegisterSet(NumberOfCensusesByType)
}
