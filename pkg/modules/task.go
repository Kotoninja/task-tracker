package modules

import (
	"time"
)

type TaskStatus string

var (
	Todo       TaskStatus = "todo"
	InProgress TaskStatus = "in-progress"
	Done       TaskStatus = "done"
)

type Task struct {
	Id          uint64     `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"createdat"`
	UpdatedAt   time.Time  `json:"updatedat"`
}

func NewTask(id uint64, description string) *Task {
	now := time.Now()
	return &Task{Id: id, Description: description, Status: Todo, CreatedAt: now, UpdatedAt: now}
}
