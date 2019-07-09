package pkg

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	database "tournament/pkg/database/model"
	"tournament/pkg/errproc"
)

var DB database.DB

func Load() {
	DB.Connect()
	MigrateTables()
}

func MigrateTables() {
	db, err := sql.Open("postgres", "user="+DB.User+" password="+DB.Password+" sslmode=disable")
	errproc.FprintErr("Unexpected error trying to connect database: %v\n", err)
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	errproc.FprintErr("Unexpected error trying to create a driver: %v\n", err)
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		DB.Database, driver)
	errproc.FprintErr("Unexpected error trying to create new migration: %v\n", err)
	err = m.Up()
	errproc.FprintErr("Unexpected error trying to migrate: %v\n", err)
}
