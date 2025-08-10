# Makefile for todogo project
# マイグレーション関連のコマンドを提供

# 環境変数の読み込み
include .env
export

# デフォルトのデータベースURL
# 環境変数から読み込むか、デフォルト値を使用
POSTGRES_HOST ?= localhost
POSTGRES_PORT ?= 5433
POSTGRES_USER ?= $(POSTGRES_USER)
POSTGRES_PASSWORD ?= $(POSTGRES_PASSWORD)
POSTGRES_DB ?= $(POSTGRES_DB)

# データベース接続URL
DATABASE_URL = postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable

# マイグレーションファイルのパス
MIGRATIONS_PATH = ./migrations

# ヘルプメッセージ
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make migrate-up      - Run all pending migrations"
	@echo "  make migrate-down    - Rollback the last migration"
	@echo "  make migrate-reset   - Reset database (down all migrations then up)"
	@echo "  make migrate-status  - Show migration status"
	@echo "  make migrate-create  - Create new migration files (NAME=migration_name)"
	@echo "  make db-up           - Start database container"
	@echo "  make db-down         - Stop database container"
	@echo "  make test            - Run all tests"
	@echo "  make build           - Build the application"

# データベースコンテナの起動
.PHONY: db-up
db-up:
	@echo "Starting database container..."
	docker-compose up -d db
	@echo "Waiting for database to be ready..."
	@until docker-compose exec db pg_isready -U $(POSTGRES_USER) -d $(POSTGRES_DB); do \
		echo "Database is unavailable - sleeping"; \
		sleep 2; \
	done
	@echo "Database is ready!"

# データベースコンテナの停止
.PHONY: db-down
db-down:
	@echo "Stopping database container..."
	docker-compose down

# マイグレーションの実行
.PHONY: migrate-up
migrate-up:
	@echo "Running migrations..."
	migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" up
	@echo "Migrations completed!"

# マイグレーションのロールバック（1つ前に戻る）
.PHONY: migrate-down
migrate-down:
	@echo "Rolling back last migration..."
	migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" down 1

# データベースのリセット（全てのマイグレーションをロールバックしてから再実行）
.PHONY: migrate-reset
migrate-reset:
	@echo "Resetting database..."
	migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" down -all
	migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" up
	@echo "Database reset completed!"

# マイグレーションの状態確認
.PHONY: migrate-status
migrate-status:
	@echo "Migration status:"
	migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" version

# 新しいマイグレーションファイルの作成
# 使用例: make migrate-create NAME=add_user_table
.PHONY: migrate-create
migrate-create:
	@if [ -z "$(NAME)" ]; then \
		echo "Error: NAME is required. Usage: make migrate-create NAME=migration_name"; \
		exit 1; \
	fi
	@echo "Creating new migration: $(NAME)"
	migrate create -ext sql -dir $(MIGRATIONS_PATH) $(NAME)

# テストの実行
.PHONY: test
test:
	@echo "Running tests..."
	go test ./...

# アプリケーションのビルド
.PHONY: build
build:
	@echo "Building application..."
	go build -o todogo .

# 開発環境のセットアップ
.PHONY: setup
setup: db-up migrate-up
	@echo "Development environment setup completed!"

# クリーンアップ
.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -f todogo
	docker-compose down -v
	@echo "Cleanup completed!"
