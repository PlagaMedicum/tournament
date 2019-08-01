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
	ErrNoParticipants	  = errors.New("cannot assign winner while there is no participants")
)

type Repository interface {
	InsertUser(string, int) (uint64, error)
	GetUsers() ([]domain.User, error)
	GetUserByID(uint64) (domain.User, error)
	DeleteUserByID(uint64) error
	UpdateUserBalanceByID(uint64, int) error
	InsertTournament(string, int, int) (uint64, error)
	GetTournamentByID(uint64) (domain.Tournament, error)
	DeleteTournamentByID(uint64) error
	AddUserInTournament(uint64, uint64) error
	SetWinner(uint64, uint64) error
}

type Controller struct {
	Repository Repository
}
