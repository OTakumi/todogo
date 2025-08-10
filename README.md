# todogo

A command-line todo management application built with Go, following clean architecture principles.

## Features

- Add, list, update, and delete tasks
- Mark tasks as complete
- Set deadlines for tasks

## Prerequisites

- Go 1.23.0 or higher
- PostgreSQL (or Docker for containerized PostgreSQL)
- golang-migrate CLI tool (for database migrations)
- Make (optional, for using Makefile commands)

## Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/todogo.git
cd todogo
```

2. Install golang-migrate CLI tool:

```bash
# For macOS using Homebrew
brew install golang-migrate

# For Ubuntu/Debian
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz
sudo mv migrate /usr/local/bin/

# For Go users (any platform)
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

3. Install dependencies:

```bash
go mod download
```

4. Set up environment variables:

```bash
cp .env.template .env
```

Edit `.env` file with your database configuration:

```env
POSTGRES_USER=postgres
POSTGRES_PASSWORD=password
POSTGRES_DB=todogo_db
DB_PORT=5433
DB_HOST=localhost
```

5. Set up the database:

Using Docker Compose and Make (recommended):

```bash
# Start database and run migrations in one command
make setup
```

Or manually:

```bash
# Start database container
docker-compose up -d

# Wait for database to be ready, then run migrations
make migrate-up
```

Or use your existing PostgreSQL installation:

```sql
CREATE DATABASE todogo_db;
```

Then run migrations:

```bash
make migrate-up
```

6. Build the application:

```bash
go build -o todogo .
# or
make build
```

## Usage

### Available Commands

#### Add a new task

```bash
todogo add "Complete the project documentation"
```

#### List all tasks

```bash
todogo list
```

#### Update a task

```bash
todogo update <task-id> --title "Updated task title"
```

#### Mark a task as complete

```bash
todogo complete <task-id>
```

#### Delete a task

```bash
todogo delete <task-id>
```

#### Show version information

```bash
todogo version
```

### Command Options

- `--help` - Show help for any command
- `--config` - Specify custom config file location

## Database Management

### Migration Commands

```bash
# Run all pending migrations
make migrate-up

# Rollback the last migration
make migrate-down

# Reset database (rollback all then reapply)
make migrate-reset

# Check migration status
make migrate-status

# Create a new migration
make migrate-create NAME=add_new_feature

# Start database container
make db-up

# Stop database container
make db-down
```

## Development

### Running Tests

```bash
go test ./...
# or
make test
```

### Running with Hot Reload

```bash
go run main.go [command]
```

### Building for Production

```bash
go build -ldflags="-s -w" -o todogo .
# or
make build
```

### Development Environment Setup

```bash
# Complete setup (database + migrations)
make setup

# Clean up everything
make clean
```
