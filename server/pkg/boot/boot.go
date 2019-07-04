package app

import (
	"database/sql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	database "tournament/pkg/database/model"
	"tournament/pkg/errproc"
)

var (
	DB database.DB
)

func MigrateTables() {
	dab, err := sql.Open("postgres", "user=postgres password=postgres sslmode=disable")
	errproc.FprintErr("Unexpected error trying to connect database: %v\n", err)
	driver, err := postgres.WithInstance(dab, &postgres.Config{})
	errproc.FprintErr("Unexpected error trying to create a driver: %v\n", err)
	m, err := migrate.NewWithDatabaseInstance(
		"file:///server/migrations",
		"tournament", driver)
	errproc.FprintErr("Unexpected error trying to create new migration: %v\n", err)
	err = m.Steps(2)
	errproc.FprintErr("Unexpected error trying to migrate: %v\n", err)

}
