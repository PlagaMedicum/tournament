package http

import (
	"encoding/json"
	"net/http"
	"tournament/pkg/domain"
	"tournament/pkg/infrastructure/errHandler"
)

// CreateTournamentHandler is the http handler for creating tournaments
// with specified name and deposit. Default prize is 4000.
// Writes tournament's id in response body.
func (c Controller) CreateTournamentHandler(w http.ResponseWriter, r *http.Request) {
	var t domain.Tournament

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		errHandler.HandleJSONErr(err, w)
		return
	}

	id, err := c.CreateTournament(t.Name, t.Deposit)
	if err != nil {
		errHandler.HandleErr(err, w)
		return
	}

	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(id)
	if err != nil {
		errHandler.HandleJSONErr(err, w)
		return
	}
}

// GetTournamentHandler is the http handler for getting information
// about tournaments with specified id.
// Writes tournament's name, deposit, prize, list of participants
// and winner in response body.
func (c Controller) GetTournamentHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len(TournamentPath+"/"):]

	t, err := c.GetTournament(id)
	if err != nil {
		errHandler.HandleErr(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(t)
	if err != nil {
		errHandler.HandleJSONErr(err, w)
		return
	}
}

// DeleteTournamentHandler is the http handler for deleting tournaments by id.
func (c Controller) DeleteTournamentHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len(TournamentPath+"/"):]

	err := c.DeleteTournament(id)
	if err != nil {
		errHandler.HandleErr(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// JoinTournamentHandler is the http handler for assigning a new participant
// to the tournament using theirs id's.
func (c Controller) JoinTournamentHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len(TournamentPath+"/") : len(TournamentPath+"/")+36]

	var st struct{ ID string `json:"userId"` }

	err := json.NewDecoder(r.Body).Decode(&st)
	if err != nil {
		errHandler.HandleJSONErr(err, w)
		return
	}

	err = c.JoinTournament(id, st.ID)
	if err != nil {
		errHandler.HandleErr(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// FinishTournamentHandler is the http handler for finishing the tournament by id.
// Updates winner of the tournament and adding prize to winner's balance.
func (c Controller) FinishTournamentHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len(TournamentPath+"/") : len(TournamentPath+"/")+36]

	err := c.FinishTournament(id)
	if err != nil {
		errHandler.HandleErr(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}
