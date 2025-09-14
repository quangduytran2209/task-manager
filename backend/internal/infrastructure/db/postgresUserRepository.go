package db

import (
	"task-manager/internal/domain"
	"task-manager/internal/repository"

	"gorm.io/gorm"
)

type postgresUserRepository struct {
	DB *gorm.DB
}

func NewPostgresUserRepository(db *gorm.DB) repository.UserReposition {
    return &postgresUserRepository{DB: db}
}

func (repo *postgresUserRepository) CreateUser(user *domain.User) error {
	return repo.DB.Create(user).Error
}

func (repo *postgresUserRepository) FindUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := repo.DB.Where("email = ?", email).First(user).Error
	if err != nil{
		return nil, err
	}
	return &user, nil
}