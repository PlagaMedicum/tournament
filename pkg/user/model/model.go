package model

import "tournament/pkg/mid"

type User struct {
	mid.MID
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}
