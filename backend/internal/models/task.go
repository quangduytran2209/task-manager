package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"` //"to-do", "in-progress", "done"
	Deadline    time.Time `json:"deadline"`
	UserID      uint      `json:"userid"`
}
