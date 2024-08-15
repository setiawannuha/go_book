package repository

import (
	"github.com/setiawannuha/go_book/internal/models"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) CreateUser(body *models.User) (string, error) {
	args := m.Mock.Called()
	return args.Get(0).(string), args.Error(1)
}
