// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package queries

import (
	"database/sql"

	"github.com/vocdoni/census3/db/annotations"
)

type Block struct {
	ID        int64
	Timestamp string
	RootHash  annotations.Hash
}

type CensusBlock struct {
	CensusID int64
	BlockID  int64
}

type Censuse struct {
	ID         int64
	StrategyID int64
	MerkleRoot annotations.Hash
	Uri        sql.NullString
	Size       sql.NullInt32
	Weight     sql.NullString
	CensusType int64
	QueueID    string
}

type Holder struct {
	ID annotations.Address
}

type Metadatum struct {
	Chainid int64
}

type Strategy struct {
	ID        int64
	Predicate string
}

type StrategyToken struct {
	StrategyID int64
	TokenID    []byte
	MinBalance []byte
	MethodHash []byte
}

type Token struct {
	ID            annotations.Address
	Name          sql.NullString
	Symbol        sql.NullString
	Decimals      sql.NullInt64
	TotalSupply   annotations.BigInt
	CreationBlock sql.NullInt32
	TypeID        int64
	Synced        bool
	Tag           sql.NullString
}

type TokenHolder struct {
	TokenID  []byte
	HolderID []byte
	Balance  []byte
	BlockID  int64
}

type TokenType struct {
	ID       int64
	TypeName string
}
