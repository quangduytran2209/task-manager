package db

import (
	"task-manager/internal/domain"
	"time"

	"gorm.io/gorm"
)

// taskModel is the GORM model for the Task entity.
// It maps to the "tasks" table in the database.
type TaskModel struct {
	gorm.Model
	Title       string     // short name of the task
	Description string     // detailed description of the task
	Status      string     // "to-do", "in-progress", "done"
	Deadline    *time.Time // optional deadline for the task
	UserID      uint       // ID of the user who owns the task
}

// ToDomain converts DB model → domain entity
func (tm *TaskModel) ToDomain() domain.Task {
	return domain.Task{
		ID:		  		tm.ID,
		Title:       	tm.Title,
		Description: 	tm.Description,
		Status:       	tm.Status,
		Deadline:     	tm.Deadline,
		UserID:      	tm.UserID,
	}
}

// FromDomain converts domain entity → DB model
func FromDomainT(dt *domain.Task) *TaskModel {
	return &TaskModel{
		Model: 	  	 gorm.Model{ID: dt.ID},
		Title:    	 dt.Title,
		Description: dt.Description,
		Status:      dt.Status,
		Deadline:    dt.Deadline,
		UserID:      dt.UserID,
	}
}