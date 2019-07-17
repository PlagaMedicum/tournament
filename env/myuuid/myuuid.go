package myuuid

import (
	uuid "github.com/satori/go.uuid"
	"log"
)

type MyUUID struct {
	ID uuid.UUID `json:"id"`
}

// Get returns the copy of ID field.
func (id *MyUUID) Get() uuid.UUID {

	return id.ID
}

// FromString initialises the ID with string.
func (id *MyUUID) FromString(s string) {
	err := id.ID.Scan(s)
	if err != nil {
		log.Printf("Error trying to convert string in uuid: "+err.Error())
	}
}
