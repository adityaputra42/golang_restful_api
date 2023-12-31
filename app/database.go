package app

import (
	"database/sql"
	"golang_restful_api/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/golang_database_migration")
	helper.PanicIfError(err)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db

	// migrate create -ext sql -dir db/migrations create_table_first
	// migrate -database "mysql://root:@tcp(localhost:3306)/golang_database_migration" -path db/migrations up
	// migrate -database "mysql://root:@tcp(localhost:3306)/golang_database_migration" -path db/migrations down

	// Fixx Dirty State
	// migrate -database "mysql://root:@tcp(localhost:3306)/golang_database_migration" -path db/migrations force 20230919075726
	// migrate -database "mysql://root:@tcp(localhost:3306)/golang_database_migration" -path db/migrations version

}
