package http

import (
	"github.com/stretchr/testify/mock"
	"tournament/pkg/domain"
)

type mockedUsecases struct {
	mock.Mock
}

func (m *mockedUsecases) CreateUser(username string) (string, error) {
	args := m.Called(username)
	return args.String(0), args.Error(1)
}

func (m *mockedUsecases) GetUser(id string) (domain.User, error) {
	args := m.Called(id)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *mockedUsecases) DeleteUser(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *mockedUsecases) FundUser(id string, points int) error {
	args := m.Called(id, points)
	return args.Error(0)
}

func (m *mockedUsecases) CreateTournament(tournamentName string, deposit int) (string, error) {
	args := m.Called(tournamentName, deposit)
	return args.String(0), args.Error(1)
}

func (m *mockedUsecases) GetTournament(id string) (domain.Tournament, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Tournament), args.Error(1)
}

func (m *mockedUsecases) DeleteTournament(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *mockedUsecases) JoinTournament(id string, userID string) error {
	args := m.Called(id, userID)
	return args.Error(0)
}

func (m *mockedUsecases) FinishTournament(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
