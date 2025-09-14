package usecase_test

import (
	"fmt"
	"task-manager/internal/domain"
	"task-manager/internal/usecase"
	"testing"
)

type InMemoryRepo struct {
	data map[uint]*domain.Task
	next uint
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{
		data: map[uint]*domain.Task{},
		next: 1,
	}
}

func (r *InMemoryRepo) CreateTask(task *domain.Task) error {
	task.ID = r.next
	r.next++
	r.data[task.ID] = task
	return nil
}

func (r *InMemoryRepo) ListByUser(userID uint) ([]*domain.Task, error) {
	var tasks []*domain.Task
	for _, task := range r.data {
		if task.UserID == userID {
			tasks = append(tasks, task)
		}
	}
	return tasks, nil
}

// FindTaskByID implements the repository.TaskRepository interface.
func (r *InMemoryRepo) FindTaskByID(id uint) (*domain.Task, error) {
	return r.FindByID(id)
}

func (r *InMemoryRepo) FindByID(id uint) (*domain.Task, error) {
	panic("unimplemented")
}

func (r *InMemoryRepo) UpdateTask(task *domain.Task) error {
	if _, ok := r.data[task.ID]; !ok {
		return fmt.Errorf("task not found")
	}
	r.data[task.ID] = task
	return nil
}

func (r *InMemoryRepo) DeleteTask(id uint) error {
	if _, ok := r.data[id]; !ok {
		return fmt.Errorf("task not found")
	}
	delete(r.data, id)
	return nil
}

func TestCreateTask_Requires(t *testing.T) {
	repo := NewInMemoryRepo()
	uc := usecase.NewTaskUsecase(repo)

	err := uc.CreateTask(&domain.Task{Title: ""})

	if err == nil {
		t.Fatal("expected error for missing title")
	}
}
