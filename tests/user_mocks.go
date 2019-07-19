package tests

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/mock"
	user "tournament/pkg/user/model"
)

type mockedUser struct {
	mock.Mock
}

func (m *mockedUser) CreateUser(u user.User) (uuid.UUID, error) {
	args := m.Called(u)
	return args.Get(0).(uuid.UUID), args.Error(1)
}

func (m *mockedUser) GetUsers() ([]user.User, error) {
	args := m.Called()
	return args.Get(0).([]user.User), args.Error(1)
}

func (m *mockedUser) GetUser(id uuid.UUID) (user.User, error) {
	args := m.Called(id)
	return args.Get(0).(user.User), args.Error(1)
}

func (m *mockedUser) DeleteUser(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(1)
}

func (m *mockedUser) FundUser(id uuid.UUID, points int) error {
	args := m.Called(id)
	return args.Error(1)
}
