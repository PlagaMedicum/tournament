package mid

import (
	uuid "github.com/satori/go.uuid"
	"tournament/pkg/errproc"
)

type ID struct {
	Id 		uuid.UUID 		`json:"id"`
}

func (id *ID) Get() uuid.UUID {
	return id.Id
}

func (id *ID) FromString(s string) {
	err := id.Id.Scan(s)
	errproc.FprintErr("Error trying to convert string in uuid: %v\n", err)
}
