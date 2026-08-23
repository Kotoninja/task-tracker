package modules

import (
	"time"
)

type TaskStatuses string

var (
	Todo       TaskStatuses = "todo"
	InProgress TaskStatuses = "in-progress"
	Done       TaskStatuses = "done"
)

type Task struct {
	Id          uint64       `json:"id"`
	Description string       `json:"description"`
	Status      TaskStatuses `json:"status"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

func NewTask(id uint64, description string) *Task {
	now := time.Now()
	return &Task{Id: id, Description: description, Status: Todo, CreatedAt: now, UpdatedAt: now}
}
