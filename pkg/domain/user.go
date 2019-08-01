package domain

type User struct {
	ID      uint64 `json:"id"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}
