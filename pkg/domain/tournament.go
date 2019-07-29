package domain

type Tournament struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Deposit      int      `json:"deposit"`
	Prize        int      `json:"prize"`
	Participants []string `json:"users"`
	WinnerID     string   `json:"winner"`
}
