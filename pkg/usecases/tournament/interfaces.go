package tournament

import (
	"tournament/pkg/domain/tournament"
	"tournament/pkg/domain/user"
)

type Repository interface {
	GetUsers() ([]user.User, error)
	GetUserByID(uint64) (user.User, error)
	UpdateUserBalanceByID(uint64, int) error
	InsertTournament(string, int, int) (uint64, error)
	GetTournamentByID(uint64) (tournament.Tournament, error)
	DeleteTournamentByID(uint64) error
	AddUserInTournament(uint64, uint64) error
	SetWinner(uint64, uint64) error
}

type Controller struct {
	Repository
}
