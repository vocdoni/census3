package api

import "time"

const (
	getCensusTimeout              = time.Second * 10
	createAndPublishCensusTimeout = time.Minute * 10
	enqueueCensusCreationTimeout  = time.Second * 10
	getStrategyCensusesTimeout    = time.Second * 10
)
