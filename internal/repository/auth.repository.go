package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/setiawannuha/go_book/internal/models"
)

type AuthRepositoryInterface interface {
	GetByEmail(email string) (*models.User, error)
}

type AuthRepository struct {
	*sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db}
}

func (r *AuthRepository) GetByEmail(email string) (*models.User, error) {
	result := models.User{}
	query := `SELECT * FROM tbl_users WHERE email=$1`
	err := r.Get(&result, query, email)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
