package db

import "log/slog"

import "database/sql"

func ConnectDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		slog.Error(err.Error())
	}

	return db, nil
}