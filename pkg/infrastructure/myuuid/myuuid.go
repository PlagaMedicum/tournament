package myuuid

import (
	uuid "github.com/satori/go.uuid"
	"tournament/pkg/interfaces/repositories/postgresql"
)

const (
	idRegex  = "[0-9+a-z]{8}-[0-9+a-z]{4}-[0-9+a-z]{4}-[0-9+a-z]{4}-[0-9+a-z]{12}"
	nullUUID = "00000000-0000-0000-0000-000000000000"
	size = 16
)

type IDFactory struct {}

type IDType struct {}

func (factory IDFactory) NewNullable() postgresql.NullableID {
	return nullable{}
}

func (factory IDFactory) New() postgresql.ID{
	id := uuid.NewV4()
	return &id
}

func (factory IDFactory) NewString() string{
	return uuid.NewV4().String()
}

func (t IDType) Null() string {
	return nullUUID
}

func (t IDType) Regex() string {
	return idRegex
}
