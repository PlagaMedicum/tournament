package mocks

import (
	"github.com/stretchr/testify/mock"
	"tournament/pkg/domain/user"
)

type MockedRepository struct {
	mock.Mock
}

func (m *MockedRepository) InsertUser(name string, balance int) (uint64, error) {
	args := m.Called(name, balance)
	return args.Get(0).(uint64), args.Error(1)
}

func (m *MockedRepository) GetUserByID(uid uint64) (user.User, error) {
	args := m.Called(uid)
	return args.Get(0).(user.User), args.Error(1)
}

func (m *MockedRepository) GetUsers() ([]user.User, error) {
	args := m.Called()
	return args.Get(0).([]user.User), args.Error(1)
}

func (m *MockedRepository)DeleteUserByID(uid uint64) error {
	args := m.Called(uid)
	return args.Error(0)
}

func (m *MockedRepository) UpdateUserBalanceByID(uid uint64, balance int) error {
	args := m.Called(uid, balance)
	return args.Error(0)
}
