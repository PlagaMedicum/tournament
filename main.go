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
	dbController := postgresql.PSQLController{
		Handler:   db.Conn,
		IDFactory: myuuid.IDFactory{},
	}
	uc := usecases.Controller{
		Repository: &dbController,
		IDType:     myuuid.IDType{},
	}

	r := httpRouter.Router{IDType: myuuid.IDType{}}
	r.Route(&handler, &httpHandlers.Controller{Usecases: &uc})

	s := http.Server{
		Addr:    ":8080",
		Handler: &handler,
	}

	err := s.ListenAndServe()
	if err != nil {
		log.Printf("Unexpected http server error: " + err.Error())
	}
}
