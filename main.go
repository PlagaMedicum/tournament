package main

import (
	"log"
	"net/http"
	app "tournament/pkg"
	"tournament/pkg/infrastructure/myhandler"
	"tournament/pkg/infrastructure/router"
	"tournament/pkg/usecases"
	user "tournament/pkg/user/usecases"
)

func main() {


	app.DB.InitNewPostgresDB()

	h := myhandler.Handler{}
	u := user.User{}
	router.RouteForUser(&h, &u)
	t := usecases.Tournament{}
	router.RouteForTournament(&h, &t)

	s := http.Server{
		Addr:    ":8080",
		Handler: &h,
	}

	err := s.ListenAndServe()
	if err != nil {
		log.Printf("Unexpected http server error: "+err.Error())
	}
}
