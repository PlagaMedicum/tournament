package domain

type User struct {
	// TODO: Change to uint?
	ID      string `json:"id"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}
