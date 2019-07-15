package mid

import (
	uuid "github.com/satori/go.uuid"
	"log"
)

type MID struct {
	ID uuid.UUID `json:"id"`
}

func (id *MID) Get() uuid.UUID {
	return id.ID
}

func (id *MID) FromString(s string) {
	err := id.ID.Scan(s)
	if err != nil {
		log.Printf("Error trying to convert string in uuid: "+err.Error())
	}
}
