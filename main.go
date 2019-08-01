package main

import (
	"log"
	"net/http"
	httpHandlers "tournament/pkg/controllers/handlers/http"
	"tournament/pkg/controllers/repositories/postgresql"
	httpRouter "tournament/pkg/controllers/routers/http"
	"tournament/pkg/infrastructure/myhandler"
	"tournament/pkg/infrastructure/myuuid"
	"tournament/pkg/infrastructure/postgresqlDB"
	"tournament/pkg/usecases"
)

func main() {
	db := postgresqlDB.DB{}
	db.InitNewPostgresDB()

	handler := myhandler.Handler{}

	httpRouter.Route(&handler, &httpHandlers.Controller{
		Usecases: &usecases.Controller{
			Repository: &postgresql.PSQLController{
				Database: db.Conn,
			},
		},
	})

	s := http.Server{
		Addr:    ":8080",
		Handler: &handler,
	}

	err := s.ListenAndServe()
	if err != nil {
		log.Printf("Unexpected http server error: " + err.Error())
	}
}
