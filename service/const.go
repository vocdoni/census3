package service

import "time"

const (
	snapshotBlocks = 100000 // a snapshot and reset of the tree is performed every snapshotBlocks
	scanSleepTime  = time.Second * 10
)
