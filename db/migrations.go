package db

import "embed"

// go:embed migrations/*.sql
var Census3Migrations embed.FS
