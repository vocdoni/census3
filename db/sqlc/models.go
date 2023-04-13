// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package queries

import (
	"database/sql"

	"github.com/vocdoni/census3/db"
)

type Block struct {
	ID        int64
	Timestamp string
	RootHash  db.Hash
}

type Censusblock struct {
	CensusID int64
	BlockID  int64
}

type Censuse struct {
	ID         int64
	StrategyID int64
	MerkleRoot db.Hash
	Uri        sql.NullString
}

type Holder struct {
	ID db.Address
}

type Strategy struct {
	ID        int64
	Predicate string
}

type Strategytoken struct {
	StrategyID int64
	TokenID    db.Address
	MinBalance db.BigInt
	MethodHash db.MethodHash
}

type Token struct {
	ID            db.BigInt
	Name          sql.NullString
	Symbol        sql.NullString
	Decimals      sql.NullInt32
	TotalSupply   db.Address
	CreationBlock int64
	TypeID        int64
}

type Tokenholder struct {
	TokenID  db.Address
	HolderID db.Address
	Balance  db.BigInt
	BlockID  int64
}

type Tokentype struct {
	ID       int64
	TypeName string
}
