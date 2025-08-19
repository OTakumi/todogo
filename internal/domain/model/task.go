package model

import (
	"errors"
	"time"
)

type Task struct {
	ID         string
	Title      string
	Deadline   *time.Time // NULLを許可するためポインタ型
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

	// Deadlineが設定されている場合、現在時刻より未来でなければならない
	if t.Deadline != nil && t.Deadline.Before(time.Now()) {
		return errors.New("Deadline must be in the future")
	}

	return nil
}
