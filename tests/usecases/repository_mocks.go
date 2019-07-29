package usecases

import (
	"github.com/stretchr/testify/mock"
	"tournament/pkg/domain"
)

type mockedRepository struct {
	mock.Mock
}

func (m *mockedRepository) InsertUser(name string, balance int) (string, error) {
	args := m.Called(name, balance)
	return args.String(0), args.Error(1)
}

func (m *mockedRepository) GetUserByID(uid string) (domain.User, error) {
	args := m.Called(uid)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *mockedRepository) GetUsers() ([]domain.User, error) {
	args := m.Called()
	return args.Get(0).([]domain.User), args.Error(1)
}

func (m *mockedRepository)DeleteUserByID(uid string) error {
	args := m.Called(uid)
	return args.Error(0)
}

func (m *mockedRepository) UpdateUserBalanceByID(uid string, balance int) error {
	args := m.Called(uid, balance)
	return args.Error(0)
}

func (m *mockedRepository) InsertTournament(name string, deposit int, prize int) (string, error) {
	args := m.Called(name, deposit, prize)
	return args.String(0), args.Error(1)
}

func (m *mockedRepository) GetTournamentParticipantList(pid string) ([]string, error) {
	args := m.Called(pid)
	return args.Get(0).([]string), args.Error(1)
}

func (m *mockedRepository) GetTournamentByID(tid string) (domain.Tournament, error) {
	args := m.Called(tid)
	return args.Get(0).(domain.Tournament), args.Error(1)
}

func (m *mockedRepository) GetTournaments() ([]domain.Tournament, error) {
	args := m.Called()
	return args.Get(0).([]domain.Tournament), args.Error(1)
}

func (m *mockedRepository) DeleteTournamentByID(tid string) error {
	args := m.Called(tid)
	return args.Error(0)
}

func (m *mockedRepository) AddUserInTournament(userID, tournamentID string) error {
	args := m.Called(userID, tournamentID)
	return args.Error(0)
}

func (m *mockedRepository) SetWinner(winnerID, tournamentID string) error {
	args := m.Called(winnerID, tournamentID)
	return args.Error(0)
}

