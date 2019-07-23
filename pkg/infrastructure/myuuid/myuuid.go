package myuuid

import (
	uuid "github.com/satori/go.uuid"
	"log"
)

const (
	idRegex = "[0-9+a-z]{8}-[0-9+a-z]{4}-[0-9+a-z]{4}-[0-9+a-z]{4}-[0-9+a-z]{12}"
	nullUUID     = "00000000-0000-0000-0000-000000000000"
)

type MyUUID struct {
	ID uuid.UUID `json:"id"`
}

func (id *MyUUID) String() string {
	return id.ID.String()
}

// FromString initialises the ID with string.
func (id *MyUUID) FromString(s string) {
	err := id.ID.Scan(s)
	if err != nil {
		log.Printf("Error trying to convert string in uuid: "+err.Error())
	}
}

func (id MyUUID) Valid() bool {

	return id.ID.String()
}

func (id MyUUID) Null() string {
	return nullUUID
}

func (id MyUUID) Regex() string {
	return idRegex
}
