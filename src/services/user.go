package services

import (
	"cine-resenha-go/src/entities"
	"cine-resenha-go/src/repositories"
	"cine-resenha-go/src/utils/context"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (service *UserService) CreateUser(email, password string) error {
	hashedPswd, err := HashPassword(password)
	if err != nil {
		return err
	}

	user := entities.User{
		Email:     email,
		Password:  hashedPswd,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	contextServer := context.CreateContextServerWithTimeout()

	err = service.repo.Create(contextServer, user)
	if err != nil {
		return err
	}

	return nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), err
}
