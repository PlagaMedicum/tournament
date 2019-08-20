package tournament

import (
	"encoding/json"
	"net/http"
	"tournament/pkg/controllers/api/http_handlers"
	"tournament/pkg/domain/tournament"
	errors "tournament/pkg/infrastructure/err"
)

type Controller struct {
	Usecases
}

// CreateTournamentHandler is the http http_handlers for creating tournaments
// with specified name and deposit. Default prize is 4000.
// Writes tournament's id in response body.
func (c Controller) CreateTournamentHandler(w http.ResponseWriter, r *http.Request) {
	var t tournament.Tournament

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		errors.HandleJSONErr(err, w)
		return
	}

	t.ID, err = c.CreateTournament(t.Name, t.Deposit)
	if err != nil {
		errors.HandleErr(err, w)
		return
	}

	w.WriteHeader(http.StatusCreated)

	type resp struct{ ID uint64 `json:"id"` }

	err = json.NewEncoder(w).Encode(resp{ID: t.ID})
	if err != nil {
		errors.HandleJSONErr(err, w)
		return
	}
}

// GetTournamentHandler is the http http_handlers for getting information
// about tournaments with specified id.
// Writes tournament's name, deposit, prize, list of participants
// and winner in response body.
func (c Controller) GetTournamentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := http_handlers.ScanID(http_handlers.TournamentPath, r)
	if err != nil {
		errors.HandleErr(err, w)
		return
	}

	t, err := c.GetTournament(id)
	if err != nil {
		errors.HandleErr(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(t)
	if err != nil {
		errors.HandleJSONErr(err, w)
		return
	}
}

// DeleteTournamentHandler is the http http_handlers for deleting tournaments by id.
func (c Controller) DeleteTournamentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := http_handlers.ScanID(http_handlers.TournamentPath, r)
	if err != nil {
		errors.HandleErr(err, w)
		return
	}

	err = c.DeleteTournament(id)
	if err != nil {
		errors.HandleErr(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// JoinTournamentHandler is the http http_handlers for assigning a new participant
// to the tournament using theirs id's.
func (c Controller) JoinTournamentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := http_handlers.ScanID(http_handlers.TournamentPath, r)
	if err != nil {
		errors.HandleErr(err, w)
		return
	}

	var st struct{ ID uint64 `json:"userId"` }

	err = json.NewDecoder(r.Body).Decode(&st)
	if err != nil {
		errors.HandleJSONErr(err, w)
		return
	}

	err = c.JoinTournament(id, st.ID)
	if err != nil {
		errors.HandleErr(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// FinishTournamentHandler is the http http_handlers for finishing the tournament by id.
// Updates winner of the tournament and adding prize to winner's balance.
func (c Controller) FinishTournamentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := http_handlers.ScanID(http_handlers.TournamentPath, r)
	if err != nil {
		errors.HandleErr(err, w)
		return
	}

	err = c.FinishTournament(id)
	if err != nil {
		errors.HandleErr(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}
