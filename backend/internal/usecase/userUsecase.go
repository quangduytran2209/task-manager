package usecase

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"task-manager/internal/domain"
	"task-manager/internal/repository"
)

type UserUsecase struct {
	repo repository.UserReposition
}

func NewUserUsecase(r repository.UserReposition) *UserUsecase{
	return &UserUsecase{repo: r}
}

func (uuc *UserUsecase) SignUp(username, email, password string) error{
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := &domain.User{
		Username: username,
		Password: string(hash),
		Email: email,
	}
	return uuc.repo.CreateUser(user)
}

func (uuc *UserUsecase) Login(email, password string) (*domain.User, error){
	user, err := uuc.repo.FindUserByEmail(email)
	if err != nil {
		return nil, errors.New("user not found")
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, errors.New("invaild password")
	}
	return user, nil
}