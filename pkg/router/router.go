package router

import (
	"net/http"
	"tournament/env/myhandler"
	app "tournament/pkg"
	tournamentHandlers "tournament/pkg/tournament/handlers"
	tournamentUseCases "tournament/pkg/tournament/usecases"
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

// RouteForUser connects endpoints with user's handling functions.
func RouteForUser(h *myhandler.Handler, ui userUseCases.UserInterface) {
	u := &userHandlers.UserController{ UserInterface: ui}

	h.HandleFunc("^"+UserPath+"$", u.CreateUserHandler, http.MethodPost)
	h.HandleFunc("^"+UserPath+"/"+app.UUIDRegex+"$", u.GetUserHandler, http.MethodGet)
	h.HandleFunc("^"+UserPath+"/"+app.UUIDRegex+"$", u.DeleteUserHandler, http.MethodDelete)
	h.HandleFunc("^"+UserPath+"/"+app.UUIDRegex+TakingPointsPath+"$", u.TakePointsHandler, http.MethodPost)
	h.HandleFunc("^"+UserPath+"/"+app.UUIDRegex+GivingPointsPath+"$", u.GivePointsHandler, http.MethodPost)
}

// RouteForUser connects endpoints with tournament's handling functions.
func RouteForTournament(h *myhandler.Handler, ti tournamentUseCases.TournamentInterface) {
	t := &tournamentHandlers.TournamentController{ TournamentInterface: ti}

	h.HandleFunc("^"+TournamentPath+"$", t.CreateTournamentHandler, http.MethodPost)
	h.HandleFunc("^"+TournamentPath+"/"+app.UUIDRegex+"$", t.GetTournamentHandler, http.MethodGet)
	h.HandleFunc("^"+TournamentPath+"/"+app.UUIDRegex+"$", t.DeleteTournamentHandler, http.MethodDelete)
	h.HandleFunc("^"+TournamentPath+"/"+app.UUIDRegex+JoinTournamentPath+"$", t.JoinTournamentHandler, http.MethodPost)
	h.HandleFunc("^"+TournamentPath+"/"+app.UUIDRegex+FinishTournamentPath+"$", t.FinishTournamentHandler, http.MethodPost)
}
