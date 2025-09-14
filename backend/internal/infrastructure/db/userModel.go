package db

import (
	"task-manager/internal/domain"

	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	ID       uint
	Username string
	Password string
	Email    string
}

func (um *UserModel) ToDomain() domain.User {
	return domain.User{
		ID:       um.ID,
		Username: um.Username,
		Password: um.Password,
		Email:    um.Email,
	}
}

func FromDomainU(du *domain.User) *UserModel {
	return &UserModel{
		Model:    gorm.Model{ID: du.ID},
		Username: du.Username,
		Password: du.Password,
		Email:    du.Email,
	}
}
