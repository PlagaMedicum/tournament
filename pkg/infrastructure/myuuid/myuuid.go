package myuuid

import (
	uuid "github.com/satori/go.uuid"
	"tournament/pkg/controllers/repositories/postgresql"
)

const (
	idRegex  = "[0-9+a-z]{8}-[0-9+a-z]{4}-[0-9+a-z]{4}-[0-9+a-z]{4}-[0-9+a-z]{12}"
	nullUUID = "00000000-0000-0000-0000-000000000000"
)

type IDFactory struct {}

type IDType struct {}

// NewNullable returns a pointer to a new nullable ID.
func (factory IDFactory) NewNullable() postgresql.NullableID {
	return &nullable{}
}

// New generates a new UUID and returns a pointer on it.
func (factory IDFactory) New() postgresql.ID{
	id := uuid.NewV4()
	return &id
}

// NewString generates a new UUID and returns it as string.
// Used in tests.
func (factory IDFactory) NewString() string{
	return uuid.NewV4().String()
}

// Null returns a NULL value for this implementation of ID.
func (t IDType) Null() string {
	return nullUUID
}

// Regex returns a regular expression for this implementation of ID.
func (t IDType) Regex() string {
	return idRegex
}
