package db

import (
	"task-manager/internal/domain"
	"task-manager/internal/repository"

	"gorm.io/gorm"
)

type postgresTaskRepository struct {
	DB *gorm.DB
}

// NewPostgresTaskRepository returns a repo that satisfies usecase.TaskRepository
func NewPostgresTaskRepository(db *gorm.DB) repository.TaskRepository {
	return &postgresTaskRepository{DB: db}
}

// CreateTask creates a new task in the database.
func (repo *postgresTaskRepository) CreateTask(task *domain.Task) error {
	fd := FromDomainT(task)
	if err := repo.DB.Create(fd).Error; err != nil {
		return err
	}

	task.ID = fd.ID
	return nil
}

// FindTaskByID retrieves a task by its ID.
func (repo *postgresTaskRepository) ListByUser(userID uint) ([]*domain.Task, error) {
	var models []TaskModel
	if err := repo.DB.Where("user_id = ?", userID).Find(&models).Error; err != nil {
		return nil, err
	}

	// map db models to domain entities
	var out []*domain.Task
	for _, r := range models {
		d := r.ToDomain()
		out = append(out, &d)
	}
	return out, nil
}

// DeleteTask implements repository.TaskRepository.
func (repo *postgresTaskRepository) DeleteTask(id uint) error {
	panic("unimplemented")
}

// FindTaskByID implements repository.TaskRepository.
func (repo *postgresTaskRepository) FindTaskByID(id uint) (*domain.Task, error) {
	panic("unimplemented")
}

// UpdateTask implements repository.TaskRepository.
func (repo *postgresTaskRepository) UpdateTask(task *domain.Task) error {
	panic("unimplemented")
}
