package myuuid

import (
	uuid "github.com/satori/go.uuid"
	"tournament/pkg/interfaces/repositories/postgresql"
)

type nullable struct {
	uuid.NullUUID
}

// IsValid returns true if ID is correct and not NULL.
func (id *nullable) IsValid() bool {
	return id.Valid
}

func (id *nullable) String() string {
	return id.UUID.String()
}

func (id *nullable) UnmarshalText(text[]byte) error {
	err := id.UUID.UnmarshalText(text)
	if err != nil {
		return err
	}
	if string(text) != nullUUID {
		id.NullUUID.Valid = true
	}
	return nil
}

func (id *nullable) GetPointer() interface{}{
	return &id.NullUUID
}

func (id *nullable) GetNotNullPointer() postgresql.ID {
	return &id.UUID
}
