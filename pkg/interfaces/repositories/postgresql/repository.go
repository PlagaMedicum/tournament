package postgresql

import (
	"github.com/jackc/pgx"
)

type handler interface{
	QueryRow(string, ...interface{}) *pgx.Row
	Query(string, ...interface{}) (*pgx.Rows, error)
	Exec(string, ...interface{}) (pgx.CommandTag, error)
}

type NullableID interface{
	IsValid() bool
	String() string
}

type idType interface {
	NewNullable() NullableID
}

type PSQLController struct {
	Handler handler
	IDType  idType
}
