package model

import (
	uuid "github.com/satori/go.uuid"
	"tournament/env/myuuid"
)

type Tournament struct {
	myuuid.MyUUID
	Name     string      `json:"name"`
	Deposit  int         `json:"deposit"`
	Prize    int         `json:"prize"`
	Users    []uuid.UUID `json:"users"`
	WinnerID uuid.UUID   `json:"winner"`
}

// GetUsers returns a slice, which contains ids of tournament's participants.
func (t *Tournament) GetUsers() (users []uuid.UUID) {
	users = append(users, t.Users...)

	return
}
