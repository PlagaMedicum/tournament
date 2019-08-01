package mocks

import (
	"github.com/stretchr/testify/mock"
	"tournament/pkg/domain"
)

type MockedUsecases struct {
	mock.Mock
}

func (m *MockedUsecases) CreateUser(username string) (uint64, error) {
	args := m.Called(username)
	return args.Get(0).(uint64), args.Error(1)
}

func (m *MockedUsecases) GetUser(id uint64) (domain.User, error) {
	args := m.Called(id)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockedUsecases) DeleteUser(id uint64) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockedUsecases) FundUser(id uint64, points int) error {
	args := m.Called(id, points)
	return args.Error(0)
}

func (m *MockedUsecases) CreateTournament(tournamentName string, deposit int) (uint64, error) {
	args := m.Called(tournamentName, deposit)
	return args.Get(0).(uint64), args.Error(1)
}

func (m *MockedUsecases) GetTournament(id uint64) (domain.Tournament, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Tournament), args.Error(1)
}

func (m *MockedUsecases) DeleteTournament(id uint64) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockedUsecases) JoinTournament(id, userID uint64) error {
	args := m.Called(id, userID)
	return args.Error(0)
}

func (m *MockedUsecases) FinishTournament(id uint64) error {
	args := m.Called(id)
	return args.Error(0)
}
