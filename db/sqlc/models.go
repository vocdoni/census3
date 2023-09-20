// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

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
	Alias     string
	Predicate string
}

type StrategyToken struct {
	StrategyID uint64
	TokenID    []byte
	ChainID    uint64
	MinBalance []byte
}

type Token struct {
	ID            annotations.Address
	Name          sql.NullString
	Symbol        sql.NullString
	Decimals      uint64
	TotalSupply   annotations.BigInt
	CreationBlock sql.NullInt64
	TypeID        uint64
	Synced        bool
	Tags          sql.NullString
	ChainID       uint64
}

type TokenHolder struct {
	TokenID  []byte
	HolderID []byte
	Balance  []byte
	BlockID  uint64
	ChainID  uint64
}

type TokenType struct {
	ID       uint64
	TypeName string
}
