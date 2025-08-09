package model

import (
	"errors"
	"time"
)

type Task struct {
	ID         string
	Title      string
	Deadline   time.Time
	IsComplete bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewTask(id string, title string) *Task {
	return &Task{
		ID:         id,
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
