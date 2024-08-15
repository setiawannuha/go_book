package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/setiawannuha/go_book/internal/models"
)

type UserRepositoryInterface interface {
	CreateUser(body *models.User) (string, error)
}

type UserRepository struct {
	*sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) CreateUser(body *models.User) (string, error) {
	query := `INSERT INTO tbl_users (email, password) VALUES (:email, :password)`
	_, err := r.NamedExec(query, body)
	if err != nil {
		return "", err
	}
	return "data created success", nil
}
