package http

import (
	"net/http"
	"strconv"
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
	CreateUser(string) (uint64, error)
	GetUser(uint64) (domain.User, error)
	DeleteUser(uint64) error
	FundUser(uint64, int) error
	CreateTournament(string, int) (uint64, error)
	GetTournament(uint64) (domain.Tournament, error)
	DeleteTournament(uint64) error
	JoinTournament(uint64, uint64) error
	FinishTournament(uint64) error
}

type Controller struct {
	Usecases
}

func scanID(path string, r *http.Request) (uint64, error) {
	var i int
	s := ""
	for i = len(path + "/"); (s != "/") || len(r.URL.Path) < i; i++ {
		s = r.URL.Path[i:i+1]
	}

	id, err := strconv.ParseUint(r.URL.Path[len(path+"/"):i-1], 10, 64)
	return id, err
}
