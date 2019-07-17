package pkg

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	database "tournament/database/model"
	"tournament/env/myhandler"
)

const (
	UUIDRegex = "[0-9+a-z]{8}-[0-9+a-z]{4}-[0-9+a-z]{4}-[0-9+a-z]{4}-[0-9+a-z]{12}" // regular expression for uuid.
)

var DB database.DB
var Handler myhandler.Handler

func migrateTables(sqldb *sql.DB) {
	driver, err := postgres.WithInstance(sqldb, &postgres.Config{DatabaseName: DB.Database})
	if err != nil {
		log.Printf("Unexpected error trying to create a driver: "+err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migrations",
		DB.Database, driver)
	if err != nil {
		log.Printf("Unexpected error trying to create new migration: "+err.Error())
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Printf("Unexpected error trying to migrate up: "+err.Error())
	}
}

func initServer(h myhandler.Handler) *http.Server {
	s := &http.Server{
		Addr:    ":8080",
		Handler: &h,
	}

	return s
}

// Init initialises database, routing endpoints and starting up the server
func Init(h myhandler.Handler) {
	sqldb := DB.Connect()
	migrateTables(sqldb)
	Handler = h
	s := initServer(Handler)

	err := s.ListenAndServe()
	if err != nil {
		log.Printf("Unexpected http server error: "+err.Error())
	}
}
