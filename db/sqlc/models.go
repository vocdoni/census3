// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package queries

import (
	"database/sql"

	"github.com/vocdoni/census3/db/annotations"
)

type Block struct {
	ID        uint64
	Timestamp string
	RootHash  annotations.Hash
}

type Censuse struct {
	ID         uint64
	StrategyID uint64
	MerkleRoot annotations.Hash
	Uri        sql.NullString
	Size       uint64
	Weight     sql.NullString
	CensusType uint64
	QueueID    string
}

type Holder struct {
	ID annotations.Address
}

type Strategy struct {
	ID        uint64
	Predicate string
	Alias     string
	Uri       string
}

type StrategyToken struct {
	StrategyID uint64
	TokenID    []byte
	MinBalance []byte
	ChainID    uint64
	ExternalID string
}

type Token struct {
	ID            annotations.Address
	Name          string
	Symbol        string
	Decimals      uint64
	TotalSupply   annotations.BigInt
	CreationBlock int64
	TypeID        uint64
	Synced        bool
	Tags          string
	ChainID       uint64
	ChainAddress  string
	ExternalID    string
}

type TokenHolder struct {
	TokenID    annotations.Address
	HolderID   annotations.Address
	Balance    []byte
	BlockID    uint64
	ChainID    uint64
	ExternalID string
}

type TokenType struct {
	ID       uint64
	TypeName string
}
