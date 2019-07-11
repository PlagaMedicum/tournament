package pkg

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	database "tournament/pkg/database/model"
	"tournament/pkg/errproc"
)

var DB database.DB

func Init() {
	db := DB.Connect()
	migrateTables(db)
}

func migrateTables(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{DatabaseName: DB.Database})
	errproc.FprintErr("Unexpected error trying to create a driver: %v\n", err)

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		DB.Database, driver)
	errproc.FprintErr("Unexpected error trying to create new migration: %v\n", err)

	err = m.Up()
	errproc.FprintErr("Unexpected error trying to migrate: %v\n", err)
}
