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

type ID interface {
	String() string
	UnmarshalText([]byte) error
}

type IDFactory interface {
	NewNullable() NullableID
	New() ID
}

type PSQLController struct {
	Handler   handler
	IDFactory IDFactory
}
