package tests

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/mock"
	tournament "tournament/pkg/tournament/model"
	user "tournament/pkg/user/model"
)

type mockedObject struct {
	mock.Mock
}

func (m *mockedObject) CreateUser(user *user.User) error {
	args := m.Called(user)
	return args.Error(1)
}

func (m *mockedObject) GetUser(id uuid.UUID) (user.User, error) {
	args := m.Called(id)
	return args.Get(0).(user.User), args.Error(1)
}

func (m *mockedObject) DeleteUser(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(1)
}

func (m *mockedObject) FundUser(id uuid.UUID, points int) error {
	args := m.Called(id)
	return args.Error(1)
}

func (m *mockedObject) CreateTournament(t tournament.Tournament) error {
	args := m.Called(t)
	return args.Error(1)
}

func (m *mockedObject) GetTournament(id uuid.UUID) (tournament.Tournament, error) {
	args := m.Called(id)
	return args.Get(0).(tournament.Tournament), args.Error(1)
}

func (m *mockedObject) DeleteTournament(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(1)
}

func (m *mockedObject) JoinTournament(id uuid.UUID, userID uuid.UUID) error {
	args := m.Called(id, userID)
	return args.Error(1)
}

func (m *mockedObject) FinishTournament(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(1)
}
