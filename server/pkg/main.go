package main

import (
	"net/http"
	app "tournament/pkg/boot"
	"tournament/pkg/errproc"
	"tournament/pkg/mhandler"
	tournament "tournament/pkg/tournament/handlers"
	user "tournament/pkg/user/handlers"
)

const (
	UUIDRegex = "[0-9+a-z]{8}-[0-9+a-z]{4}-[0-9+a-z]{4}-[0-9+a-z]{4}-[0-9+a-z]{12}"
)

func main() {
	app.DB.Connect("tournament-app")
	app.MigrateTables()

	var handler mhandler.Handler
	handler.HandleFunc("^/user$", user.CreateUser, "POST")
	handler.HandleFunc("^/user/" + UUIDRegex + "$", user.GetUser, "GET")
	handler.HandleFunc("^/user/" + UUIDRegex + "$", user.DeleteUser, "DELETE")
	handler.HandleFunc("^/user/" + UUIDRegex + "/take$", user.TakePoints, "POST")
	handler.HandleFunc("^/user/" + UUIDRegex + "/fund$", user.GivePoints, "POST")
	handler.HandleFunc("^/tournament$", tournament.CreateTournament, "POST")
	handler.HandleFunc("^/tournament/" + UUIDRegex + "$", tournament.GetTournament, "GET")
	handler.HandleFunc("^/tournament/" + UUIDRegex + "$", tournament.DeleteTournament, "DELETE")
	handler.HandleFunc("^/tournament/" + UUIDRegex + "/join$", tournament.JoinTournament, "POST")
	handler.HandleFunc("^/tournament/" + UUIDRegex + "/finish$", tournament.FinishTournament, "POST")

	server := &http.Server{
		Addr:    ":8080",
		Handler: &handler,
	}
	err := server.ListenAndServe()
	errproc.FprintErr("Unexpected http server error: %v\n", err)
}