package http_handlers

import (
	"net/http"
	"strconv"
)

// A list of paths used in http endpoints.
const (
	UserPath             = "/user"
	TakingPointsPath     = "/take"
	GivingPointsPath     = "/give"
	TournamentPath       = "/tournament"
	JoinTournamentPath   = "/join"
	FinishTournamentPath = "/finish"
)

// ScanID gets requested url, parses ID from it and returns resulting ID.
func ScanID(path string, r *http.Request) (uint64, error) {
	i := len(path + "/")
	for len(r.URL.Path) >= i+1 && (r.URL.Path[i:i+1] != "/") {
		i++
	}

	id, err := strconv.ParseUint(r.URL.Path[len(path+"/"):i], 10, 64)
	return id, err
}
