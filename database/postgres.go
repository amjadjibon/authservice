package database

import (
	"database/sql"
	"time"
)

func GetDB(dsn string) *sql.DB {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(10 * time.Minute)
	db.SetConnMaxLifetime(30 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(20)

	return db
}
