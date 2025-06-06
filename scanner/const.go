package scanner

import "time"

const (
	READ_TIMEOUT = time.Minute
	SCAN_TIMEOUT = time.Minute * 5
	SAVE_TIMEOUT = time.Minute * 5
)

const (
	scanSleepTime         = time.Second * 20  // time to sleep between scans
	scanSleepTimeOnceSync = time.Second * 120 // time to sleep between scans, once all the tokens are synced
	blockNumbersCooldown  = time.Minute * 5   // time to wait to update latest block numbers of every supported networkd
)
