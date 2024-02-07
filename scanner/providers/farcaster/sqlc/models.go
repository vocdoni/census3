// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package queries

import (
	"database/sql"

	"github.com/vocdoni/census3/db/annotations"
)

type User struct {
	Fid             uint64
	Username        string
	Signer          annotations.Bytes
	Custodyaddress  []byte
	Appkeys         sql.NullString
	Recoveryaddress []byte
}
