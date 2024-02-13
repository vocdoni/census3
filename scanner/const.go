package scanner

import "time"

const (
	READ_TIMEOUT = time.Minute
	SCAN_TIMEOUT = 5 * time.Minute
	SAVE_TIMEOUT = time.Minute
)

const (
	snapshotBlocks        = 100000            // a snapshot and reset of the tree is performed every snapshotBlocks
	scanSleepTime         = time.Second * 20  // time to sleep between scans
	scanSleepTimeOnceSync = time.Second * 120 // time to sleep between scans, once all the tokens are synced
	blockNumbersCooldown  = 5 * time.Minute   // time to wait to update latest block numbers of every supported networkd
)
