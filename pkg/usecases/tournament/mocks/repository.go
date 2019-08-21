package mocks

import (
	"github.com/stretchr/testify/mock"
	"tournament/pkg/domain/tournament"
	"tournament/pkg/domain/user"
)

// MockedRepository is mock for postgresql.tournament.Repository interface.
type MockedRepository struct {
	mock.Mock
}

func (m *MockedRepository) GetUserByID(uid uint64) (user.User, error) {
	args := m.Called(uid)
	return args.Get(0).(user.User), args.Error(1)
}

func (m *MockedRepository) GetUsers() ([]user.User, error) {
	args := m.Called()
	return args.Get(0).([]user.User), args.Error(1)
}

func (m *MockedRepository) UpdateUserBalanceByID(uid uint64, balance int) error {
	args := m.Called(uid, balance)
	return args.Error(0)
}

func (m *MockedRepository) InsertTournament(name string, deposit int, prize int) (uint64, error) {
	args := m.Called(name, deposit, prize)
	return args.Get(0).(uint64), args.Error(1)
}

func (m *MockedRepository) GetTournamentByID(tid uint64) (tournament.Tournament, error) {
	args := m.Called(tid)
	return args.Get(0).(tournament.Tournament), args.Error(1)
}

func (m *MockedRepository) DeleteTournamentByID(tid uint64) error {
	args := m.Called(tid)
	return args.Error(0)
}

func (m *MockedRepository) AddUserInTournament(userID, tournamentID uint64) error {
	args := m.Called(userID, tournamentID)
	return args.Error(0)
}

func (m *MockedRepository) SetWinner(winnerID, tournamentID uint64) error {
	args := m.Called(winnerID, tournamentID)
	return args.Error(0)
}
