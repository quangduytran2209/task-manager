package usecase

import (
	"errors"
	"task-manager/internal/domain"
	"task-manager/internal/repository"
)

// TaskUsecase is the interface that defines the methods for task use case operations.
type TaskUsecase struct {
	repo repository.TaskRepository
}



// NewTaskUsecase creates a new instance of TaskUsecase.
func NewTaskUsecase(r repository.TaskRepository) *TaskUsecase {
	return &TaskUsecase{repo: r}
}

func (tuc *TaskUsecase) ListTasksByUser(userID uint) (any, any) {
	panic("unimplemented")
}

// CreateTask creates a new task after validating the input.
func (tuc *TaskUsecase) CreateTask(task *domain.Task) error {
	if task.Title == "" {
		return errors.New("task title is required")
	}
	if task.Status == "" {
		task.Status = "to-do"
	}
	return tuc.repo.CreateTask(task)
}

// // FindTaskByID retrieves a task by its ID.
// func (tuc *TaskUsecase) FindTaskByID(id uint) (*domain.Task, error) {
// 	return tuc.repo.FindTaskByID(id)
// }

// // UpdateTask updates an existing task after validating the input.
// func (tuc *TaskUsecase) UpdateTask(task *domain.Task) error {
// 	if task.ID == 0 {
// 		return errors.New("task ID is required for update")
// 	}
// 	return tuc.repo.UpdateTask(task)
// }

// // DeleteTask deletes a task by its ID.
// func (tuc *TaskUsecase) DeleteTask(id uint) error {
// 	if id == 0 {
// 		return errors.New("task ID is required for deletion")
// 	}
// 	return tuc.repo.DeleteTask(id)
// }

// ListByUser retrieves all tasks for a specific user.
func (tuc *TaskUsecase) ListByUser(userID uint) ([]*domain.Task, error) {
	return tuc.repo.ListByUser(userID)
}
