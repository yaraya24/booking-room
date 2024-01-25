package db

import (
	"context"

	"github.com/jmoiron/sqlx"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

type Database struct {
	DB *sqlx.DB
}

func SetupDB(dbFile string) (*Database, error) {
	db, err := sqlx.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}
	return &Database{DB: db}, nil
}

func (d Database) Read(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return d.DB.SelectContext(ctx, dest, query, args...)
}

func (d Database) Write(ctx context.Context, query string, args ...interface{}) (int64, error) {
	result, err := d.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return int64(0), err
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		return int64(0), err
	}
	return affectedRows, nil
}
