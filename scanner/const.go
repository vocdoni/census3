package scanner

import "time"

const (
	// READ_TIMEOUT is the timeout to get sorted tokens to scan from the database
	READ_TIMEOUT = time.Minute
	// SAVE_TIMEOUT is the timeout to save the scanned tokens to the database
	SAVE_TIMEOUT = 5 * time.Minute
	// PREPARE_TIMEOUT is the timeout to prepare the tokens to scan (calculate
	// the birth block number, etc.)
	PREPARE_TIMEOUT = 5 * time.Minute
	// UPDATE_TIMEOUT is the timeout to update the tokens from their holders
	// providers
	UPDATE_TIMEOUT = 15 * time.Minute
)

const (
	coolDown              = 15 * time.Second  // time to wait between updates
	scanSleepTime         = time.Second * 20  // time to sleep between scans
	scanSleepTimeOnceSync = time.Second * 120 // time to sleep between scans, once all the tokens are synced
	blockNumbersCooldown  = 5 * time.Minute   // time to wait to update latest block numbers of every supported networkd
)
