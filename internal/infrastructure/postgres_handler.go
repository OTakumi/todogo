package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type PostgreSQLHandler struct {
	DB *sql.DB
}

func NewPostgreSQLHandler(dsn string) (*PostgreSQLHandler, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open postgresql: %w", err)
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

	log.Println("Successfully connected to the database.")
	return &PostgreSQLHandler{DB: db}, err
}
