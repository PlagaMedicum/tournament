package pkg

import (
	"database/sql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	database "tournament/pkg/database/model"
	"tournament/pkg/errproc"
)

var DB database.DB

func Load() {
	DB.Connect()
	MigrateTables()
}

func MigrateTables() {
	dab, err := sql.Open("postgres", "user="+DB.User+" password="+DB.Password+" sslmode=disable")
	errproc.FprintErr("Unexpected error trying to connect database: %v\n", err)
	driver, err := postgres.WithInstance(dab, &postgres.Config{})
	errproc.FprintErr("Unexpected error trying to create a driver: %v\n", err)
	m, err := migrate.NewWithDatabaseInstance(
		"file:///go/src/tournament/migrations",
		DB.Database, driver)
	errproc.FprintErr("Unexpected error trying to create new migration: %v\n", err)
	err = m.Steps(1)
	errproc.FprintErr("Unexpected error trying to migrate: %v\n", err)
}
