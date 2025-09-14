package domain

import "time"

// Task is the business entity.
// It represents a unit of work in our Kanban board.
type Task struct {
	ID          uint       // unique identifier
	Title       string     // short name of the task
	Description string     // detailed description of the task
	Status      string     // "to-do", "in-progress", "done"
	Deadline    *time.Time // optional deadline for the task
	UserID      uint       // ID of the user who owns the task
}
