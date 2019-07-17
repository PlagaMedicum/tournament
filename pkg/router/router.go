package router

import (
	"tournament/env/myhandler"
	app "tournament/pkg"
	tournament "tournament/pkg/tournament/handlers"
	user "tournament/pkg/user/handlers"
)

// Route connects endpoints with handling functions.
func Route() (h myhandler.Handler) {
	h.HandleFunc("^/user$", user.CreateUser, "POST")
	h.HandleFunc("^/user/"+app.UUIDRegex+"$", user.GetUser, "GET")
	h.HandleFunc("^/user/"+app.UUIDRegex+"$", user.DeleteUser, "DELETE")
	h.HandleFunc("^/user/"+app.UUIDRegex+"/take$", user.TakePoints, "POST")
	h.HandleFunc("^/user/"+app.UUIDRegex+"/fund$", user.GivePoints, "POST")
	h.HandleFunc("^/tournament$", tournament.CreateTournament, "POST")
	h.HandleFunc("^/tournament/"+app.UUIDRegex+"$", tournament.GetTournament, "GET")
	h.HandleFunc("^/tournament/"+app.UUIDRegex+"$", tournament.DeleteTournament, "DELETE")
	h.HandleFunc("^/tournament/"+app.UUIDRegex+"/join$", tournament.JoinTournament, "POST")
	h.HandleFunc("^/tournament/"+app.UUIDRegex+"/finish$", tournament.FinishTournament, "POST")

	return
}
