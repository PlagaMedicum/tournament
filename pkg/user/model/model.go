package model

import "tournament/env/myuuid"

type User struct {
	myuuid.MyUUID
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}
