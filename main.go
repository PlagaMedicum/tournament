package main

import (
	"log"
	"net/http"
	"tournament/pkg/infrastructure/myhandler"
	"tournament/pkg/infrastructure/myuuid"
	"tournament/pkg/infrastructure/postgresqlDB"
	"tournament/pkg/interfaces/repositories/postgresql"
	httpRouter "tournament/pkg/interfaces/routers/http"
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
	controller := usecases.Controller{
		Repository: &dbController,
		IDType: myuuid.IDType{},
	}

	r := httpRouter.Router{IDType: myuuid.IDType{}}
	r.Route(&handler, &controller)

	s := http.Server{
		Addr:    ":8080",
		Handler: &handler,
	}

	err := s.ListenAndServe()
	if err != nil {
		log.Printf("Unexpected http server error: "+err.Error())
	}
}
