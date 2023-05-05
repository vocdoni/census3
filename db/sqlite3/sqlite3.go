package sqlite3

import "embed"

//go:embed migrations/*.sql
var Migrations embed.FS