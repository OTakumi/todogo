package model

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID        uuid.UUID
	Title     string
	Deadline  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
