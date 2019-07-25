package usecases

import (
	"errors"
	"tournament/pkg/domain"
)

var (
	ErrNoUserWithID       = errors.New("no user with target id found")
	ErrNotEnoughPoints    = errors.New("user have not enough points to join the tournament")
	ErrNoTournamentWithID = errors.New("no tournament with target id found")
	ErrParticipantExists  = errors.New("this user already joined the tournament")
	ErrFinishedTournament = errors.New("this tournament already finished")
)

type Repository interface {
	InsertUser(string, int) (string, error)
	GetUsers() ([]domain.User, error)
	GetUserByID(string) (domain.User, error)
	DeleteUserByID(string) error
	UpdateUserBalanceByID(int, string) error
	InsertTournament(string, int, int) (string, error)
	GetTournaments() ([]domain.Tournament, error)
	GetTournamentByID(string) (domain.Tournament, error)
	GetTournamentParticipants(string) ([]string, error)
	DeleteTournamentByID(string) error
	AddUserInTournament(string, string) error
	SetWinner(string, string) error
}

type RepositoryInteractor interface {
	CreateUser(string) (string, error)
	GetUser(string) (domain.User, error)
	DeleteUser(string) error
	FundUser(string, int) error
	CreateTournament(string, int) (string, error)
	GetTournament(string) (domain.Tournament, error)
	DeleteTournament(string) error
	JoinTournament(string, string) error
	FinishTournament(string) error
}

type idType interface {
	Null() string
}

type Controller struct {
	Repository Repository
	IDType idType
}
