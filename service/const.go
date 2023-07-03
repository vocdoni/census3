package service

import "time"

const (
	snapshotBlocks        = 100000            // a snapshot and reset of the tree is performed every snapshotBlocks
	scanSleepTime         = time.Second * 20  // time to sleep between scans
	scanSleepTimeOnceSync = time.Second * 120 // time to sleep between scans, once all the tokens are synced
	// scanIterationDurationPerToken is the time the scanner will spend scanning each token, by iteration.
	scanIterationDurationPerToken = 60 * time.Second
)
