package service

import "time"

const (
	/*
	   The key value stores the relation:
	   c_<contractAddres> = #block

	*/
	contractPrefix = "c_"   // KV prefix for identify contracts
	snapshotBlocks = 100000 // a snapshot and reset of the tree is performed every snapshotBlocks
	scanSleepTime  = time.Second * 10
)
