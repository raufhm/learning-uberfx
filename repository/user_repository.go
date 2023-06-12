package repository

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/raufhm/learning-uberfx/domain"
)

type UserRepository interface {
	GetUserByID(id string) (*domain.User, error)
	CreateUser(user *domain.User) error
}

type UserRepositoryImpl struct {
	DBConnection *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		DBConnection: db,
	}
}

func (repo *UserRepositoryImpl) GetUserByID(id string) (*domain.User, error) {
	query := `SELECT row_to_json(u) FROM users u WHERE u.id = $1`
	row := repo.DBConnection.QueryRow(query, id)

	var userJSON []byte
	err := row.Scan(&userJSON)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to retrieve user: %w", err)
	}

	var user domain.User
	err = json.Unmarshal(userJSON, &user)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal user JSON: %w", err)
	}

	return &user, nil
}

func (repo *UserRepositoryImpl) CreateUser(user *domain.User) error {
	var result string
	query := `INSERT INTO users (first_name, last_name, email, mobile) VALUES ($1, $2, $3) RETURNING id`
	err := repo.DBConnection.QueryRow(query, user.FirstName, user.LastName, user.Email, user.Mobile).Scan(&result)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	if result == "" {
		return errors.New("unable to insert users")
	}

	var res domain.User
	err = json.Unmarshal([]byte(result), &res)
	if err != nil {
		return err
	}

	return nil
}
