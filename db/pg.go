package db

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

var _db *bun.DB

func Init(dsn string, debug bool) error {
	// Open a PostgreSQL database.
	pgdb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	// Create a Bun db on top of it.
	_db = bun.NewDB(pgdb, pgdialect.New(), bun.WithDiscardUnknownColumns())
	_db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	return nil
}

func PG() *bun.DB {
	return _db
}

func Begin() (bun.Tx, error) {
	return _db.Begin()
}
