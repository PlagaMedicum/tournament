package main

import (
	"log"
	app "tournament/pkg"
	"tournament/pkg/router"
)

func main() {
	app.DB.InitNewPostgresDB()

	h := router.Route()
	s:= app.InitServer(h)

	err := s.ListenAndServe()
	if err != nil {
		log.Printf("Unexpected http server error: "+err.Error())
	}
}
