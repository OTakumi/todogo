# todogo

A command-line todo management application built with Go, following clean architecture principles.

## Features

- Add, list, update, and delete tasks
- Mark tasks as complete
- Set deadlines for tasks

## Prerequisites

- Go 1.22.2 or higher
- PostgreSQL (or Docker for containerized PostgreSQL)
- Make (optional, for using Makefile commands)

## Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/todogo.git
cd todogo
```

2. Install dependencies:

```bash
go mod download
```

3. Set up environment variables:

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

4. Set up the database:

Using Docker Compose (recommended):

```bash
docker-compose up -d
```

Or use your existing PostgreSQL installation and create a database:

```sql
CREATE DATABASE todogo_db;
```

5. Build the application:

```bash
go build -o todogo .
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

## Development

### Running Tests

```bash
go test ./...
```

### Running with Hot Reload

```bash
go run main.go [command]
```

### Building for Production

```bash
go build -ldflags="-s -w" -o todogo .
```
