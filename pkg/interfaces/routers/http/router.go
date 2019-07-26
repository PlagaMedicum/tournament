package http

import (
	"net/http"
	"tournament/pkg/infrastructure/myhandler"
	handlers "tournament/pkg/interfaces/handlers/http"
	"tournament/pkg/usecases"
)

type idInterface interface {
	Regex() string
}

type Router struct {
	IDType idInterface
}

// Route connects endpoints with handling functions.
func (r Router) Route(h *myhandler.Handler, ri usecases.RepositoryInteractor) {
	c := &handlers.Controller{RepositoryInteractor: ri}

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
