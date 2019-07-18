package router

import (
	"net/http"
	"tournament/env/myhandler"
	app "tournament/pkg"
	tournament "tournament/pkg/tournament/handlers"
	user "tournament/pkg/user/handlers"
)

// Route connects endpoints with handling functions.
func Route() *myhandler.Handler {
	var h myhandler.Handler
	h.HandleFunc("^/user$", user.CreateUser, http.MethodPost)
	h.HandleFunc("^/user/"+app.UUIDRegex+"$", user.GetUser, http.MethodGet)
	h.HandleFunc("^/user/"+app.UUIDRegex+"$", user.DeleteUser, http.MethodDelete)
	h.HandleFunc("^/user/"+app.UUIDRegex+"/take$", user.TakePoints, http.MethodPost)
	h.HandleFunc("^/user/"+app.UUIDRegex+"/fund$", user.GivePoints, http.MethodPost)
	h.HandleFunc("^/tournament$", tournament.CreateTournament, http.MethodPost)
	h.HandleFunc("^/tournament/"+app.UUIDRegex+"$", tournament.GetTournament, http.MethodGet)
	h.HandleFunc("^/tournament/"+app.UUIDRegex+"$", tournament.DeleteTournament, http.MethodDelete)
	h.HandleFunc("^/tournament/"+app.UUIDRegex+"/join$", tournament.JoinTournament, http.MethodPost)
	h.HandleFunc("^/tournament/"+app.UUIDRegex+"/finish$", tournament.FinishTournament, http.MethodPost)

	return &h
}
