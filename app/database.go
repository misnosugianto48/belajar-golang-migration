package app

import (
	"database/sql"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/belajar-golang-migration")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(10 * time.Minute)
	db.SetConnMaxIdleTime(60 * time.Minute)

	return db
}
