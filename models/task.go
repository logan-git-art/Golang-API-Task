package models

import "time"

type Task struct {
	ID          uint       `json:"id" gorm:"primary_key"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	DueDate     time.Time  `json:"due_date"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type TaskStatus string

const (
	Pending    TaskStatus = "pending"
	Completed  TaskStatus = "completed"
	InProgress TaskStatus = "in_progress"
)
