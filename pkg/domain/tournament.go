package domain

type Tournament struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Deposit      int      `json:"deposit"`
	Prize        int      `json:"prize"`
	Participants []string `json:"users"`
	WinnerID     string   `json:"winner"`
}

// GetParticipants returns a slice, which contains ids of tournament's participants.
func (t *Tournament) GetParticipants() (p []string) {
	p = append(p, t.Participants...)
	return
}
