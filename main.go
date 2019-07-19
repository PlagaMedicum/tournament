package main

import (
	"log"
	"tournament/env/myhandler"
	app "tournament/pkg"
	"tournament/pkg/router"
	user "tournament/pkg/user/usecases"
	tournament "tournament/pkg/tournament/usecases"
)

func main() {
	app.DB.InitNewPostgresDB()

	h := myhandler.Handler{}
	u := user.User{}
	router.RouteForUser(&h, &u)
	t := tournament.Tournament{}
	router.RouteForTournament(&h, &t)

	s:= app.InitServer(&h)

	err := s.ListenAndServe()
	if err != nil {
		log.Printf("Unexpected http server error: "+err.Error())
	}
}
