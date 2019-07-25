package myuuid

import uuid "github.com/satori/go.uuid"

type nullable uuid.NullUUID

func (id nullable) IsValid() bool {
	return id.Valid
}

func (id nullable) String() string {
	return id.UUID.String()
}
