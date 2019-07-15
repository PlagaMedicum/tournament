package model

import (
	uuid "github.com/satori/go.uuid"
	"tournament/env/mid"
)

type Tournament struct {
	mid.MID
	Name     string      `json:"name"`
	Deposit  int         `json:"deposit"`
	Prize    int         `json:"prize"`
	Users    []uuid.UUID `json:"users"`
	WinnerID uuid.UUID   `json:"winner"`
}

func (t *Tournament) GetUsers() (users []uuid.UUID) {
	users = append(users, t.Users...)
	return
}
