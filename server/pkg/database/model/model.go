package model

import (
	"github.com/jackc/pgx"
	"tournament/pkg/errproc"
)

type DB struct {
	Conn *pgx.Conn
}

func (db DB) Connect(applicationName string) {
	var runtimeParams map[string]string
	runtimeParams = make(map[string]string)
	runtimeParams["application_name"] = applicationName
	connConfig := pgx.ConnConfig{
		User:              "postgres",
		Password:          "postgres",
		Host:              "localhost",
		Port:              5432,
		Database:          "tournament",
		TLSConfig:         nil,
		UseFallbackTLS:    false,
		FallbackTLSConfig: nil,
		RuntimeParams:     runtimeParams,
	}
	conn, err := pgx.Connect(connConfig)
	if err != nil {
		errproc.FprintErr("Unable to establish connection: %v\n", err)
	}
	db.Conn = conn
}
