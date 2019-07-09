package model

import (
	"github.com/go-yaml/yaml"
	"github.com/jackc/pgx"
	"io/ioutil"
	"tournament/pkg/errproc"
)

type DB struct {
	Conn           *pgx.Conn
	User           string            `yaml:"User"`
	Password       string            `yaml:"Password"`
	Host           string            `yaml:"Host"`
	Port           uint16            `yaml:"Port"`
	Database       string            `yaml:"Database"`
	UseFallbackTLS bool              `yaml:"UseFallbackTLS"`
	RuntimeParams  map[string]string `yaml:"RuntimeParams,omitempty"`
}

func (db *DB) Connect() {
	file, err := ioutil.ReadFile("connconf.yaml")
	errproc.FprintErr("Unable to read yaml file: %v\n", err)
	err = yaml.Unmarshal(file, &db)
	errproc.FprintErr("Unable to unmarshal yaml data: %v\n", err)
	connConfig := pgx.ConnConfig{
		User:              db.User,
		Password:          db.Password,
		Host:              db.Host,
		Port:              db.Port,
		Database:          db.Database,
		TLSConfig:         nil,
		UseFallbackTLS:    db.UseFallbackTLS,
		FallbackTLSConfig: nil,
		RuntimeParams:     db.RuntimeParams,
	}
	conn, err := pgx.Connect(connConfig)
	if err != nil {
		errproc.FprintErr("Unable to establish connection: %v\n", err)
	}
	db.Conn = conn
}
