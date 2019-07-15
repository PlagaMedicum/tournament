package model

import (
	"database/sql"
	"github.com/go-yaml/yaml"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"io/ioutil"
	"log"
	"strconv"
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
	file, err := ioutil.ReadFile("./database/connconf.yaml")
	if err != nil {
		log.Printf("Unable to read yaml file: "+err.Error())
	}
	err = yaml.Unmarshal(file, &db)
	if err != nil {
		log.Printf("Unable to unmarshal yaml data: "+err.Error())
	}
	sqldb, err := sql.Open("pgx",
		"user="+db.User+
		" password="+db.Password+
		" host="+db.Host+
		" port="+strconv.FormatUint(db.Port, 10)+
		" database="+db.Database+
		" sslmode=disable")
	if err != nil {
		log.Printf("Unable to open connection: "+err.Error())
	}
	err = sqldb.Ping()
	if err != nil {
		log.Printf("Postgresql ping: "+err.Error())
	}
	conn, err := stdlib.AcquireConn(sqldb)
	if err != nil {
		log.Printf("Unable to establish connection: "+err.Error())
	}
	db.Conn = conn
	return sqldb
}
