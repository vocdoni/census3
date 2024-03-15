package api

import (
	"time"

	"go.vocdoni.io/proto/build/go/models"
)

const (
	// censuses
	getCensusTimeout              = time.Second * 10
	createAndPublishCensusTimeout = time.Minute * 10
	publishCensusTimeout          = time.Second * 60
	enqueueCensusCreationTimeout  = time.Second * 10
	getStrategyCensusesTimeout    = time.Second * 10
	// strategies
	createDummyStrategyTimeout   = time.Second * 10
	importStrategyTimeout        = time.Second * 10
	enqueueStrategyImportTimeout = time.Second * 10
	getStrategiesTimeout         = time.Second * 10
	getStrategyTimeout           = time.Second * 10
	getTokensStrategyTimeout     = time.Second * 10
	getStrategyHoldersTimeout    = time.Second * 20
	// tokens
	getTokensTimeout       = time.Second * 20
	createTokenTimeout     = time.Second * 10
	getTokenTimeout        = time.Second * 15
	deleteTokenTimeout     = time.Minute * 5
	tokenHoldersCSVTimeout = time.Minute * 5
)

const (
	defaultPageSize               = int32(1000)
	defaultCensusType             = models.Census_ARBO_BLAKE2B
	anonymousCensusType           = models.Census_ARBO_POSEIDON
	strategyHoldersCacheThreshold = 500
	apiCacheKeySize               = 16
	apiCacheSize                  = 128
)
