package repository

import (
	"github.com/setiawannuha/go_book/internal/models"
	"github.com/stretchr/testify/mock"
)

type AuthRepositoryMock struct {
	mock.Mock
}

func (m *AuthRepositoryMock) GetByEmail(email string) (*models.User, error) {
	args := m.Mock.Called()
	return args.Get(0).(*models.User), args.Error(1)
}
