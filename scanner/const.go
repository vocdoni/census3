package scanner

import "time"

const (
	READ_TIMEOUT = time.Minute
	SCAN_TIMEOUT = 5 * time.Minute
	SAVE_TIMEOUT = 5 * time.Minute
)

const (
	scanSleepTime         = time.Second * 20  // time to sleep between scans
	scanSleepTimeOnceSync = time.Second * 120 // time to sleep between scans, once all the tokens are synced
	blockNumbersCooldown  = 5 * time.Minute   // time to wait to update latest block numbers of every supported networkd
)
