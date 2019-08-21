package user

// User is a domain for users that can participate in tournaments.
type User struct {
	ID      uint64 `json:"id"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}
