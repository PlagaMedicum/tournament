package http

import (
	"github.com/stretchr/testify/mock"
	"tournament/pkg/domain"
)

type mockedRepositoryInteractor struct {
	mock.Mock
}

func (m *mockedRepositoryInteractor) CreateUser(username string) (string, error) {
	args := m.Called(username)
	return args.String(0), args.Error(1)
}

func (m *mockedRepositoryInteractor) GetUser(id string) (domain.User, error) {
	args := m.Called(id)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *mockedRepositoryInteractor) DeleteUser(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *mockedRepositoryInteractor) FundUser(id string, points int) error {
	args := m.Called(id, points)
	return args.Error(0)
}

func (m *mockedRepositoryInteractor) CreateTournament(tournamentName string, deposit int) (string, error) {
	args := m.Called(tournamentName, deposit)
	return args.String(0), args.Error(1)
}

func (m *mockedRepositoryInteractor) GetTournament(id string) (domain.Tournament, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Tournament), args.Error(1)
}

func (m *mockedRepositoryInteractor) DeleteTournament(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *mockedRepositoryInteractor) JoinTournament(id string, userID string) error {
	args := m.Called(id, userID)
	return args.Error(0)
}

func (m *mockedRepositoryInteractor) FinishTournament(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
