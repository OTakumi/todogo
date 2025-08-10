package main

import (
	"OTakumi/todogo/cmd"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func setupDB(dbDriver string, dsn string) (*sql.DB, error) {
	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		return nil, err
	}

	// 接続が実際に確立されるまで待機
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// 接続確認
	if err := db.Ping(); err != nil {
		db.Close()

		return nil, fmt.Errorf("Failed to ping postgres: %w", err)
	}

	return db, err
}

func main() {
	dbDriver := "postgres"
	dsn := "host=localhost port=5433 user=postgres password=p@ssw0rd dbname=todogo sslmode=disable"

	db, err := setupDB(dbDriver, dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	cmd.Execute()
}
