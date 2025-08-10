package main

import (
	"OTakumi/todogo/cmd"
	"OTakumi/todogo/internal/infrastructure"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

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

	// 必須の環境変数が設定されているか確認
	if dbUser == "" || dbPassword == "" || dbHost == "" || dbPort == "" || dbName == "" {
		log.Fatal("Database environment variables are not set correctly")
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	dbHandler, err := infrastructure.NewPostgreSQLHandler(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbHandler.DB.Close()

	cmd.Execute()
}
