package model

import (
	"database/sql"
	"github.com/go-yaml/yaml"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"strconv"
)

type DB struct {
	Conn         *pgx.Conn
	User         string `yaml:"User"`
	Password     string `yaml:"Password"`
	Host         string `yaml:"Host"`
	Port         uint64 `yaml:"Port"`
	Database     string `yaml:"Database"`
	TestDatabase string `yaml:"TestDatabase"`
	m            *migrate.Migrate
}

const (
	migrationsPath     = "file://database/migrations"
	testMigrationsPath = "file://../database/test_migrations"
)

func (db *DB) readConfigFile() {
	file, err := ioutil.ReadFile("../database/connconf.yaml")
	if err != nil {
		log.Printf("Unable to read yaml file: " + err.Error())
	}

	err = yaml.Unmarshal(file, &db)
	if err != nil {
		log.Printf("Unable to unmarshal yaml data: " + err.Error())
	}
}

// connect reads configuration file and initialises a new
// connection with the postgres database.
func (db *DB) connect(dbName string) *sql.DB {
	sqldb, err := sql.Open("pgx",
		"user="+db.User+
			" password="+db.Password+
			" host="+db.Host+
			" port="+strconv.FormatUint(db.Port, 10)+
			" database="+dbName+
			" sslmode=disable")
	if err != nil {
		log.Printf("Unable to open connection: " + err.Error())
	}

	err = sqldb.Ping()
	if err != nil {
		log.Printf("Postgresql ping: " + err.Error())
	}

	conn, err := stdlib.AcquireConn(sqldb)
	if err != nil {
		log.Printf("Unable to establish connection: " + err.Error())
	}

	db.Conn = conn

	return sqldb
}

//Connect initialises Postgresql connection with stage database.
func (db *DB) Connect() *sql.DB {
	db.readConfigFile()

	sqldb := db.connect(db.Database)
	return sqldb
}

//Connect initialises Postgresql connection with test database.
func (db *DB) ConnectForTests() *sql.DB {
	db.readConfigFile()

	sqldb := db.connect(db.TestDatabase)
	return sqldb
}

func (db *DB) createNewMigration(sqldb *sql.DB, path string) {
	driver, err := postgres.WithInstance(sqldb, &postgres.Config{DatabaseName: db.Database})
	if err != nil {
		log.Printf("Unexpected error trying to create a driver: " + err.Error())
	}

	db.m, err = migrate.NewWithDatabaseInstance(
		path,
		db.Database, driver)
	if err != nil {
		log.Printf("Unexpected error trying to create new migration: " + err.Error())
	}
}

func (db *DB) migrateTablesUp() {
	err := db.m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Printf("Unexpected error trying to migrate up: " + err.Error())
	}
}

// MigrateTablesDown migrates database's tables down.
func (db *DB) MigrateTablesDown() {
	err := db.m.Down()
	if err != nil && err != migrate.ErrNoChange {
		log.Printf("Unexpected error trying to migrate down: " + err.Error())
	}
}

// InitNewPostgresDB initialises new database connection
// and calls it's migrations.
func (db *DB) InitNewPostgresDB() {
	sqldb := db.Connect()
	db.createNewMigration(sqldb, migrationsPath)
	db.migrateTablesUp()
}

// InitNewTestPostgresDB initialises new database connection
// for tests and calls it's migrations.
func (db *DB) InitNewTestPostgresDB() {
	sqldb := db.ConnectForTests()
	db.createNewMigration(sqldb, testMigrationsPath)
	db.migrateTablesUp()
}