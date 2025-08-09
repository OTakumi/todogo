package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID         uuid.UUID
	Title      string
	Deadline   time.Time
	IsComplete bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (t *Task) Validate() error {
	if t.Title == "" {
		return errors.New("Title is required")
	}

	return nil
}
