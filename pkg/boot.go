package pkg

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
	database "tournament/pkg/database/model"
)

var DB database.DB

func Init() {
	db := DB.Connect()
	migrateTables(db)
}

func migrateTables(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{DatabaseName: DB.Database})
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
	if err != nil {
		log.Printf("Unexpected error trying to migrate up: "+err.Error())
	}
}
