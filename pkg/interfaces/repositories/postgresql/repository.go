package postgresql

import (
	"github.com/jackc/pgx"
)

// TODO: Create interfaces for pgx
type handler interface{
	QueryRow(string, ...interface{}) *pgx.Row
	Query(string, ...interface{}) (*pgx.Rows, error)
	Exec(string, ...interface{}) (pgx.CommandTag, error)
}

type ID interface {
	String() string
	UnmarshalText([]byte) error
}

type NullableID interface{
	IsValid() bool
	GetPointer() interface{}
	GetNotNullPointer() ID
	ID
}

type IDFactory interface {
	NewNullable() NullableID
	New() ID
}

type PSQLController struct {
	Handler   handler
	IDFactory IDFactory
}
