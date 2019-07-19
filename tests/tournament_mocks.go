package tests

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/mock"
	tournament "tournament/pkg/tournament/model"
)

type mockedTournament struct {
	mock.Mock
}

func (m *mockedTournament) CreateTournament(t *tournament.Tournament) error {
	args := m.Called(t)
	t.ID = uuid.NewV1()
	return args.Error(1)
}

func (m *mockedTournament) GetTournament(id uuid.UUID) (tournament.Tournament, error) {
	args := m.Called(id)
	return args.Get(0).(tournament.Tournament), args.Error(1)
}

func (m *mockedTournament) DeleteTournament(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(1)
}

func (m *mockedTournament) JoinTournament(id uuid.UUID, userID uuid.UUID) error {
	args := m.Called(id, userID)
	return args.Error(1)
}

func (m *mockedTournament) FinishTournament(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(1)
}
