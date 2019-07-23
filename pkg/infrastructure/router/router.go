package router

import (
	"net/http"
	"tournament/pkg/infrastructure/myhandler"
	tournamentHandlers "tournament/pkg/interfaces/handlers"
	"tournament/pkg/usecases"
	userHandlers "tournament/pkg/user/handlers"
	userUseCases "tournament/pkg/user/usecases"
)

const (
	UserPath             = "/user"
	TakingPointsPath     = "/take"
	GivingPointsPath     = "/give"
	TournamentPath       = "/tournament"
	JoinTournamentPath   = "/join"
	FinishTournamentPath = "/finish"
)

type idInterface interface {
	Regex() string
}

type Router struct {
	id idInterface
}

// RouteForUser connects endpoints with user's handling functions.
func (r Router) RouteForUser(h *myhandler.Handler, ui userUseCases.UserInterface) {
	u := &userHandlers.UserController{ UserInterface: ui}

	h.HandleFunc("^"+UserPath+"$", u.CreateUserHandler, http.MethodPost)
	h.HandleFunc("^"+UserPath+"/"+r.id.Regex()+"$", u.GetUserHandler, http.MethodGet)
	h.HandleFunc("^"+UserPath+"/"+r.id.Regex()+"$", u.DeleteUserHandler, http.MethodDelete)
	h.HandleFunc("^"+UserPath+"/"+r.id.Regex()+TakingPointsPath+"$", u.TakePointsHandler, http.MethodPost)
	h.HandleFunc("^"+UserPath+"/"+r.id.Regex()+GivingPointsPath+"$", u.GivePointsHandler, http.MethodPost)
}

// RouteForTournament connects endpoints with tournament's handling functions.
func (r Router) RouteForTournament(h *myhandler.Handler, ti usecases.RepositoryInteractor) {
	t := &tournamentHandlers.Controller{ RepositoryInteractor: ti}

	h.HandleFunc("^"+TournamentPath+"$", t.CreateTournamentHandler, http.MethodPost)
	h.HandleFunc("^"+TournamentPath+"/"+r.id.Regex()+"$", t.GetTournamentHandler, http.MethodGet)
	h.HandleFunc("^"+TournamentPath+"/"+r.id.Regex()+"$", t.DeleteTournamentHandler, http.MethodDelete)
	h.HandleFunc("^"+TournamentPath+"/"+r.id.Regex()+JoinTournamentPath+"$", t.JoinTournamentHandler, http.MethodPost)
	h.HandleFunc("^"+TournamentPath+"/"+r.id.Regex()+FinishTournamentPath+"$", t.FinishTournamentHandler, http.MethodPost)
}
