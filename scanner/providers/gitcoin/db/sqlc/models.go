// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package queries

import (
	"time"

	"github.com/vocdoni/census3/db/annotations"
)

type Metadatum struct {
	Attr  string
	Value string
}

type Score struct {
	Address annotations.Address
	Score   annotations.BigInt
	Date    time.Time
}

type Stamp struct {
	Address annotations.Address
	Name    string
	Score   annotations.BigInt
}

type TotalSupply struct {
	Name        string
	TotalSupply string
}
