package model

import (
	"database/sql"
	"github.com/go-yaml/yaml"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"io/ioutil"
	"strconv"
	"tournament/pkg/errproc"
)

type DB struct {
	Conn           *pgx.Conn
	User           string            `yaml:"User"`
	Password       string            `yaml:"Password"`
	Host           string            `yaml:"Host"`
	Port           uint64            `yaml:"Port"`
	Database       string            `yaml:"Database"`
}

func (db *DB) Connect() *sql.DB{
	file, err := ioutil.ReadFile("connconf.yaml")
	errproc.FprintErr("Unable to read yaml file: %v\n", err)
	err = yaml.Unmarshal(file, &db)
	errproc.FprintErr("Unable to unmarshal yaml data: %v\n", err)

	db1, err := sql.Open("pgx",
		"user="+db.User+
		" password="+db.Password+
		" host="+db.Host+
		" port="+strconv.FormatUint(db.Port, 10)+
		" database="+db.Database+
		" sslmode=disable")
	errproc.FprintErr("Unable to open connection: %v\n", err)

	err = db1.Ping()
	errproc.FprintErr("Ping error: %v\n", err)

	conn, err := stdlib.AcquireConn(db1)
	errproc.FprintErr("Unable to establish connection: %v\n", err)
	db.Conn = conn
	return db1
}
