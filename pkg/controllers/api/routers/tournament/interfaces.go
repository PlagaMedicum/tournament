package tournament

import "net/http"

type httpController interface {
	CreateTournamentHandler(http.ResponseWriter, *http.Request)
	GetTournamentHandler(http.ResponseWriter, *http.Request)
	DeleteTournamentHandler(http.ResponseWriter, *http.Request)
	JoinTournamentHandler(http.ResponseWriter, *http.Request)
	FinishTournamentHandler(http.ResponseWriter, *http.Request)
}
