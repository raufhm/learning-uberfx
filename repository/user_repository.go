package repository

import (
	"database/sql"
	"github.com/raufhm/learning-uberfx/domain"
)

type UserRepository interface {
	GetUserByID(id string) (*domain.User, error)
	CreateUser(user *domain.User) error
}

type UserRepositoryImpl struct {
	DBConnection *sql.DB
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repo *UserRepositoryImpl) GetUserByID(id string) (*domain.User, error) {
	return nil, nil
}

func (repo *UserRepositoryImpl) CreateUser(user *domain.User) error {
	return nil
}
