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

func NewTask(title string) *Task {
	// UUIDを生成する
	uuid, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}

	return &Task{
		ID:         uuid,
		Title:      title,
		IsComplete: false,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

func (t *Task) Validate() error {
	if t.Title == "" {
		return errors.New("Title is required")
	}

	return nil
}
