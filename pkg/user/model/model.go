package model

import "tournament/env/mid"

type User struct {
	mid.MID
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}
