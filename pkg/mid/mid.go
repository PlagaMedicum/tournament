package mid

import (
	uuid "github.com/satori/go.uuid"
	"tournament/pkg/errproc"
)

type MID struct {
	ID uuid.UUID `json:"id"`
}

func (id *MID) Get() uuid.UUID {
	return id.ID
}

func (id *MID) FromString(s string) {
	err := id.ID.Scan(s)
	errproc.FprintErr("Error trying to convert string in uuid: %v\n", err)
}
