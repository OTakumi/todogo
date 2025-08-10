package main

import (
	"OTakumi/todogo/cmd"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

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
	// dotenvファイルから環境変数を読み込む
	// 実行環境（本番環境など）に.envファイルがない場合でも、
	// OSの環境変数が設定されていればそちらを優先して利用できるため、
	// err != nil の場合はエラーとせず、ログ出力に留める。
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found")
	}

	// DBパラメータを環境変数から読み込む
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbDriver := "postgres"

	// 必須の環境変数が設定されているか確認
	if dbUser == "" || dbPassword == "" || dbHost == "" || dbPort == "" || dbName == "" {
		log.Fatal("Database environment variables are not set correctly")
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := setupDB(dbDriver, dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	cmd.Execute()
}
