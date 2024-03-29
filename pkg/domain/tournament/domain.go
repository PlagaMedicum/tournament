package tournament

// Tournament is a domain of tournaments where users can participate.
type Tournament struct {
	ID           uint64   `json:"id"`
	Name         string   `json:"name"`
	Deposit      int      `json:"deposit"`
	Prize        int      `json:"prize"`
	Participants []uint64 `json:"users"`
	WinnerID     uint64   `json:"winner"`
}
