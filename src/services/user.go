package services

import (
	"cine-resenha-go/src/dtos"
	"cine-resenha-go/src/repositories"
	"cine-resenha-go/src/utils/context"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (service *UserService) CreateUser(email, password string) error {
	user := dtos.User{
		Email:    email,
		Password: password,
	}

	contextServer := context.CreateContextServerWithTimeout()

	err := service.repo.Create(contextServer, user)
	if err != nil {
		return err
	}

	return nil
}
