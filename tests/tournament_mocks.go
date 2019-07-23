package tests

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/mock"
	"tournament/pkg/domain"
)

type mockedTournament struct {
	mock.Mock
}

func (m *mockedTournament) CreateTournament(t domain.Tournament) (uuid.UUID, error) {
	args := m.Called(t)
	return args.Get(0).(uuid.UUID), args.Error(1)
}

func (m *mockedTournament) GetTournament(id uuid.UUID) (domain.Tournament, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Tournament), args.Error(1)
}

func (m *mockedTournament) DeleteTournament(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *mockedTournament) JoinTournament(id uuid.UUID, userID uuid.UUID) error {
	args := m.Called(id, userID)
	return args.Error(0)
}

func (m *mockedTournament) FinishTournament(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}
