package myuuid

import (
	uuid "github.com/satori/go.uuid"
	"tournament/pkg/interfaces/repositories/postgresql"
)

const (
	idRegex  = "[0-9+a-z]{8}-[0-9+a-z]{4}-[0-9+a-z]{4}-[0-9+a-z]{4}-[0-9+a-z]{12}"
	nullUUID = "00000000-0000-0000-0000-000000000000"
)

//type MyUUID uuid.UUID

type IDFabric struct {}

type IDType struct {}

func (fabric IDFabric) NewNullable() postgresql.NullableID {
	return nullable{}
}

func (fabric IDFabric) New() string{
	return uuid.NewV1().String()
}

func (t IDType) Null() string {
	return nullUUID
}

func (t IDType) Regex() string {
	return idRegex
}
