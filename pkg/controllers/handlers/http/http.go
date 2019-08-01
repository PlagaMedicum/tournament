package http

import (
	"tournament/pkg/domain"
)

const (
	UserPath             = "/user"
	TakingPointsPath     = "/take"
	GivingPointsPath     = "/give"
	TournamentPath       = "/tournament"
	JoinTournamentPath   = "/join"
	FinishTournamentPath = "/finish"
)

type Usecases interface {
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

type Controller struct {
	Usecases
}
