package postgresql

import (
	"database/sql"
	"flag"
	"fmt"
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
	"os"
	"strconv"
	"tournament/pkg/controllers/repositories/postgresql"
)

const (
	migrationsPath        = "file://databases/postgresql/migrations"
	configurationFilePath = "./databases/postgresql/config.yaml"
)

var (
	cwd_arg = flag.String("cwd", "", "set cwd")
)

// init changes current working directory with provided by
// "-cwd" flag.
func init() {
	flag.Parse()
	if *cwd_arg != "" {
		err := os.Chdir(*cwd_arg)
		if err != nil {
			fmt.Println("Chdir error:", err)
		}
	}
}

type conn struct {
	conn *pgx.Conn
}

func (c conn) QueryRow(sql string, args ...interface{}) postgresql.Row {
	return c.conn.QueryRow(sql, args...)
}

func (c conn) Query(sql string, args ...interface{}) (postgresql.Rows, error) {
	return c.conn.Query(sql, args...)
}

func (c conn) Exec(sql string, args ...interface{}) (interface{}, error) {
	return c.conn.Exec(sql, args...)
}

func (c conn) ErrNoRows() error {
	return pgx.ErrNoRows
}

type DB struct {
	Conn         conn
	User         string `yaml:"User"`
	Password     string `yaml:"Password"`
	Host         string `yaml:"Host"`
	Port         uint64 `yaml:"Port"`
	Database     string `yaml:"Database"`
	SSLMode		 string `yaml:"SSLMode"`
	m            *migrate.Migrate
}

func (db *DB) readConfigFile() {
	file, err := ioutil.ReadFile(configurationFilePath)
	if err != nil {
		log.Printf("Unable to read yaml file: " + err.Error())
	}

	err = yaml.Unmarshal(file, &db)
	if err != nil {
		log.Printf("Unable to unmarshal yaml data: " + err.Error())
	}
}

// connect reads configuration file and initialises a new
// connection with the Postgres DB.
func (db *DB) connect(dbName string) *sql.DB {
	sqldb, err := sql.Open("pgx",
		"user="+db.User+
		" password="+db.Password+
		" host="+db.Host+
		" port="+strconv.FormatUint(db.Port, 10)+
		" database="+dbName+
		" sslmode="+db.SSLMode)
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

	db.Conn.conn = conn

	return sqldb
}

// Connect initialises Postgresql connection with staging DB.
func (db *DB) Connect() *sql.DB {
	db.readConfigFile()

	sqldb := db.connect(db.Database)
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

func (db *DB) MigrateTablesUp() {
	err := db.m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Printf("Unexpected error trying to migrate up: " + err.Error())
		return
	}
}

// MigrateTablesDown migrates DB's tables down.
func (db *DB) MigrateTablesDown() {
	err := db.m.Down()
	if err != nil && err != migrate.ErrNoChange {
		log.Printf("Unexpected error trying to migrate down: " + err.Error())
		return
	}
}

// InitNewPostgresDB initialises new DB connection
// and initializes it's migrations.
func (db *DB) InitNewPostgresDB() {
	sqldb := db.Connect()
	db.createNewMigration(sqldb, migrationsPath)
}
