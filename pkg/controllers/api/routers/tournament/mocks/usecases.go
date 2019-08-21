package mocks

import (
	"github.com/stretchr/testify/mock"
	"tournament/pkg/domain/tournament"
)

// MockedUsecases is mock for tournament.Usecases interface.
type MockedUsecases struct {
	mock.Mock
}

func (m *MockedUsecases) CreateTournament(tournamentName string, deposit int) (uint64, error) {
	args := m.Called(tournamentName, deposit)
	return args.Get(0).(uint64), args.Error(1)
}

func (m *MockedUsecases) GetTournament(id uint64) (tournament.Tournament, error) {
	args := m.Called(id)
	return args.Get(0).(tournament.Tournament), args.Error(1)
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
