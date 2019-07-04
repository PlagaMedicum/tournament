package model

import (
	uuid "github.com/satori/go.uuid"
	"tournament/pkg/mid"
)

type Tournament struct {
	mid.ID
	Name     string        	`json:"name"`
	Deposit  int           	`json:"deposit"`
	Prize    int           	`json:"prize"`
	Users    []uuid.UUID   	`json:"users"`
	WinnerId uuid.UUID 		`json:"winner"`
}

func (t *Tournament) GetUsers() (users []uuid.UUID) {
	for _, user := range t.Users {
		users = append(users, user)
	}
	return
}
