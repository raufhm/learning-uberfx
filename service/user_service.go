package service

import (
	"github.com/raufhm/learning-uberfx/domain"
	"github.com/raufhm/learning-uberfx/repository"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepo,
	}
}

func (svc *UserService) GetUserByID(id string) (*domain.User, error) {
	return svc.UserRepository.GetUserByID(id)
}

func (svc *UserService) CreateUser(user *domain.User) error {
	// Implement the logic to create a user using the UserRepository
	// ...
	return svc.UserRepository.CreateUser(user)
}
