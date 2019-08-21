package mocks

import (
	"github.com/stretchr/testify/mock"
	"tournament/pkg/domain/user"
)

// MockedUsecases is mock for user.Usecases interface.
type MockedUsecases struct {
	mock.Mock
}

func (m *MockedUsecases) CreateUser(username string) (uint64, error) {
	args := m.Called(username)
	return args.Get(0).(uint64), args.Error(1)
}

func (m *MockedUsecases) GetUser(id uint64) (user.User, error) {
	args := m.Called(id)
	return args.Get(0).(user.User), args.Error(1)
}

func (m *MockedUsecases) DeleteUser(id uint64) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockedUsecases) FundUser(id uint64, points int) error {
	args := m.Called(id, points)
	return args.Error(0)
}
