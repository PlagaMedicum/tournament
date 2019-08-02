package tournament

import (
	"net/http"
	"tournament/pkg/controllers/api/http_handlers"
	"tournament/pkg/controllers/api/routers"
)

// RouteTournament connects tournament endpoints with handling functions.
func RouteTournament(h routers.Handler, c httpController) {
	h.HandleFunc("^"+http_handlers.TournamentPath+"$", c.CreateTournamentHandler, http.MethodPost)
	h.HandleFunc("^"+http_handlers.TournamentPath+routers.Regex+"$", c.GetTournamentHandler, http.MethodGet)
	h.HandleFunc("^"+http_handlers.TournamentPath+routers.Regex+"$", c.DeleteTournamentHandler, http.MethodDelete)
	h.HandleFunc("^"+http_handlers.TournamentPath+routers.Regex+http_handlers.JoinTournamentPath+"$", c.JoinTournamentHandler, http.MethodPost)
	h.HandleFunc("^"+http_handlers.TournamentPath+routers.Regex+http_handlers.FinishTournamentPath+"$", c.FinishTournamentHandler, http.MethodPost)
}
