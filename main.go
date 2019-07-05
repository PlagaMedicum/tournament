package main

import (
	"net/http"
	app "tournament/pkg"
	"tournament/pkg/errproc"
	"tournament/pkg/mhandler"
	tournament "tournament/pkg/tournament/handlers"
	user "tournament/pkg/user/handlers"
)

const (
	uuidRegex = "[0-9+a-z]{8}-[0-9+a-z]{4}-[0-9+a-z]{4}-[0-9+a-z]{4}-[0-9+a-z]{12}"
)

func main() {
	app.Load()

	var h mhandler.Handler
	h.HandleFunc("^/user$", user.CreateUser, "POST")
	h.HandleFunc("^/user/"+uuidRegex+"$", user.GetUser, "GET")
	h.HandleFunc("^/user/"+uuidRegex+"$", user.DeleteUser, "DELETE")
	h.HandleFunc("^/user/"+uuidRegex+"/take$", user.TakePoints, "POST")
	h.HandleFunc("^/user/"+uuidRegex+"/fund$", user.GivePoints, "POST")
	h.HandleFunc("^/tournament$", tournament.CreateTournament, "POST")
	h.HandleFunc("^/tournament/"+uuidRegex+"$", tournament.GetTournament, "GET")
	h.HandleFunc("^/tournament/"+uuidRegex+"$", tournament.DeleteTournament, "DELETE")
	h.HandleFunc("^/tournament/"+uuidRegex+"/join$", tournament.JoinTournament, "POST")
	h.HandleFunc("^/tournament/"+uuidRegex+"/finish$", tournament.FinishTournament, "POST")

	server := &http.Server{
		Addr:    ":8080",
		Handler: &h,
	}
	err := server.ListenAndServe()
	errproc.FprintErr("Unexpected http server error: %v\n", err)
}
