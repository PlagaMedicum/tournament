package domain

type idInterface interface {
	String() string
	FromString(string)
	Valid() bool
	Null() string
}

type User struct {
	ID 		idInterface
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

type Tournament struct {
	ID 		 idInterface
	Name     string      `json:"name"`
	Deposit  int         `json:"deposit"`
	Prize    int         `json:"prize"`
	Users    []string `json:"users"`
	WinnerID string   `json:"winner"`
}

// GetParticipants returns a slice, which contains ids of tournament's participants.
func (t *Tournament) GetParticipants() (p []string) {
	p = append(p, t.Users...)

	return
}
