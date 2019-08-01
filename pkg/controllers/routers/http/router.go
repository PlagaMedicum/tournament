package http

import (
	"net/http"
	handlers "tournament/pkg/controllers/handlers/http"
)

const regex = "/[0-9]+"

type handler interface {
	Handle(string, http.Handler, string)
	HandleFunc(string, func(http.ResponseWriter, *http.Request), string)
	http.Handler
}

type httpController interface {
	CreateUserHandler(http.ResponseWriter, *http.Request)
	GetUserHandler(http.ResponseWriter, *http.Request)
	DeleteUserHandler(http.ResponseWriter, *http.Request)
	TakePointsHandler(http.ResponseWriter, *http.Request)
	GivePointsHandler(http.ResponseWriter, *http.Request)
	CreateTournamentHandler(http.ResponseWriter, *http.Request)
	GetTournamentHandler(http.ResponseWriter, *http.Request)
	DeleteTournamentHandler(http.ResponseWriter, *http.Request)
	JoinTournamentHandler(http.ResponseWriter, *http.Request)
	FinishTournamentHandler(http.ResponseWriter, *http.Request)
}

// Route connects endpoints with handling functions.
func Route(h handler, c httpController) {
	h.HandleFunc("^"+handlers.UserPath+"$", c.CreateUserHandler, http.MethodPost)
	h.HandleFunc("^"+handlers.UserPath+regex+"$", c.GetUserHandler, http.MethodGet)
	h.HandleFunc("^"+handlers.UserPath+regex+"$", c.DeleteUserHandler, http.MethodDelete)
	h.HandleFunc("^"+handlers.UserPath+regex+handlers.TakingPointsPath+"$", c.TakePointsHandler, http.MethodPost)
	h.HandleFunc("^"+handlers.UserPath+regex+handlers.GivingPointsPath+"$", c.GivePointsHandler, http.MethodPost)
	h.HandleFunc("^"+handlers.TournamentPath+"$", c.CreateTournamentHandler, http.MethodPost)
	h.HandleFunc("^"+handlers.TournamentPath+regex+"$", c.GetTournamentHandler, http.MethodGet)
	h.HandleFunc("^"+handlers.TournamentPath+regex+"$", c.DeleteTournamentHandler, http.MethodDelete)
	h.HandleFunc("^"+handlers.TournamentPath+regex+handlers.JoinTournamentPath+"$", c.JoinTournamentHandler, http.MethodPost)
	h.HandleFunc("^"+handlers.TournamentPath+regex+handlers.FinishTournamentPath+"$", c.FinishTournamentHandler, http.MethodPost)
}
