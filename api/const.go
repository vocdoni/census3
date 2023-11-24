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
	// tokens
	getTokensTimeout   = time.Second * 10
	createTokenTimeout = time.Second * 10
	getTokenTimeout    = time.Second * 15
)

const (
	defaultPageSize               = int32(10)
	defaultCensusType             = models.Census_ARBO_BLAKE2B
	anonymousCensusType           = models.Census_ARBO_POSEIDON
	strategyHoldersCacheThreshold = 500
)
