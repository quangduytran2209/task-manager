package repository

import "task-manager/internal/domain"

type UserReposition interface {
	CreateUser(user *domain.User) error
	FindUserByEmail(email string) (*domain.User, error)
}