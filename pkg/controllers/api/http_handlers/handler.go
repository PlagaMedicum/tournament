package http_handlers

import (
	"net/http"
	"strconv"
)

const (
	UserPath             = "/user"
	TakingPointsPath     = "/take"
	GivingPointsPath     = "/give"
	TournamentPath       = "/tournament"
	JoinTournamentPath   = "/join"
	FinishTournamentPath = "/finish"
)

func ScanID(path string, r *http.Request) (uint64, error) {
	var i int
	for i = len(path + "/"); len(r.URL.Path) >= i+1 && (r.URL.Path[i:i+1] != "/"); i++ {}

	id, err := strconv.ParseUint(r.URL.Path[len(path+"/"):i], 10, 64)
	return id, err
}
