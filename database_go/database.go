package database_go

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:33066)/go_db?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
