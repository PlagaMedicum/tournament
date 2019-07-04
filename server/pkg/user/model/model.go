package model

import "tournament/pkg/mid"

type User struct {
	mid.ID
	Name    string 			`json:"name"`
	Balance int    			`json:"balance"`
}
