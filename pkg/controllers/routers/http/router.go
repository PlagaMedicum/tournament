package http

import (
	"net/http"
	handlers "tournament/pkg/controllers/handlers/http"
)

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

type idInterface interface {
	Regex() string
}

type Router struct {
	IDType idInterface
}

// Route connects endpoints with handling functions.
func (r Router) Route(h handler, c httpController) {
	h.HandleFunc("^"+handlers.UserPath+"$", c.CreateUserHandler, http.MethodPost)
	h.HandleFunc("^"+handlers.UserPath+"/"+r.IDType.Regex()+"$", c.GetUserHandler, http.MethodGet)
	h.HandleFunc("^"+handlers.UserPath+"/"+r.IDType.Regex()+"$", c.DeleteUserHandler, http.MethodDelete)
	h.HandleFunc("^"+handlers.UserPath+"/"+r.IDType.Regex()+handlers.TakingPointsPath+"$", c.TakePointsHandler, http.MethodPost)
	h.HandleFunc("^"+handlers.UserPath+"/"+r.IDType.Regex()+handlers.GivingPointsPath+"$", c.GivePointsHandler, http.MethodPost)
	h.HandleFunc("^"+handlers.TournamentPath+"$", c.CreateTournamentHandler, http.MethodPost)
	h.HandleFunc("^"+handlers.TournamentPath+"/"+r.IDType.Regex()+"$", c.GetTournamentHandler, http.MethodGet)
	h.HandleFunc("^"+handlers.TournamentPath+"/"+r.IDType.Regex()+"$", c.DeleteTournamentHandler, http.MethodDelete)
	h.HandleFunc("^"+handlers.TournamentPath+"/"+r.IDType.Regex()+handlers.JoinTournamentPath+"$", c.JoinTournamentHandler, http.MethodPost)
	h.HandleFunc("^"+handlers.TournamentPath+"/"+r.IDType.Regex()+handlers.FinishTournamentPath+"$", c.FinishTournamentHandler, http.MethodPost)
}
