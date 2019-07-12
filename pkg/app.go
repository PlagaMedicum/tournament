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
	database "tournament/pkg/database/model"
	"tournament/pkg/mhandler"
	tournament "tournament/pkg/tournament/handlers"
	user "tournament/pkg/user/handlers"
)

const (
	UUIDRegex = "[0-9+a-z]{8}-[0-9+a-z]{4}-[0-9+a-z]{4}-[0-9+a-z]{4}-[0-9+a-z]{12}"
)

var DB database.DB
var Handler mhandler.Handler

func migrateTables(sqldb *sql.DB) {
	driver, err := postgres.WithInstance(sqldb, &postgres.Config{DatabaseName: DB.Database})
	if err != nil {
		log.Printf("Unexpected error trying to create a driver: "+err.Error())
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		DB.Database, driver)
	if err != nil {
		log.Printf("Unexpected error trying to create new migration: "+err.Error())
	}
	err = m.Up()
	if err != nil || err != migrate.ErrNoChange {
		log.Printf("Unexpected error trying to migrate up: "+err.Error())
	}
}

func initNewHandler() (h mhandler.Handler) {
	h.HandleFunc("^/user$", user.CreateUser, "POST")
	h.HandleFunc("^/user/"+UUIDRegex+"$", user.GetUser, "GET")
	h.HandleFunc("^/user/"+UUIDRegex+"$", user.DeleteUser, "DELETE")
	h.HandleFunc("^/user/"+UUIDRegex+"/take$", user.TakePoints, "POST")
	h.HandleFunc("^/user/"+UUIDRegex+"/fund$", user.GivePoints, "POST")
	h.HandleFunc("^/tournament$", tournament.CreateTournament, "POST")
	h.HandleFunc("^/tournament/"+UUIDRegex+"$", tournament.GetTournament, "GET")
	h.HandleFunc("^/tournament/"+UUIDRegex+"$", tournament.DeleteTournament, "DELETE")
	h.HandleFunc("^/tournament/"+UUIDRegex+"/join$", tournament.JoinTournament, "POST")
	h.HandleFunc("^/tournament/"+UUIDRegex+"/finish$", tournament.FinishTournament, "POST")
	return
}

func initServer(h mhandler.Handler) *http.Server {
	s := &http.Server{
		Addr:    ":8080",
		Handler: &h,
	}
	return s
}

func Init() {
	sqldb := DB.Connect()
	migrateTables(sqldb)
	Handler = initNewHandler()
	s := initServer(Handler)
	err := s.ListenAndServe()
	if err != nil {
		log.Printf("Unexpected http server error: "+err.Error())
	}
}

