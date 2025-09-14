package repository

import "task-manager/internal/domain"

type TaskRepository interface {
	CreateTask(task *domain.Task) error
	FindTaskByID(id uint) (*domain.Task, error)
	UpdateTask(task *domain.Task) error
	DeleteTask(id uint) error
	ListByUser(userID uint) ([]*domain.Task, error)
}
