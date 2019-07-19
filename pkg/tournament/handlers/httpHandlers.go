package handlers

import (
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"tournament/env/errproc"
	"tournament/env/myuuid"
	"tournament/pkg/tournament/model"
	"tournament/pkg/tournament/usecases"
)

type TournamentController struct {
	usecases.TournamentInterface
}

// CreateTournamentHandler is the http handler for creating tournaments
// with specified name and deposit. Default prize is 4000.
// Writes tournament's id in response body.
func (c TournamentController) CreateTournamentHandler(w http.ResponseWriter, r *http.Request) {
	var t model.Tournament

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		errproc.HandleJSONErr(err, w)
		return
	}

	err = c.CreateTournament(&t)
	if err != nil {
		errproc.HandleErr(err, w)
		return
	}

	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(t.ID)
	if err != nil {
		errproc.HandleJSONErr(err, w)
		return
	}
}

// GetTournamentHandler is the http handler for getting information
// about tournaments with specified id.
// Writes tournament's name, deposit, prize, list of participants
// and winner in response body.
func (c TournamentController) GetTournamentHandler(w http.ResponseWriter, r *http.Request) {
	var id myuuid.MyUUID
	id.FromString(r.URL.Path[len("/tournament/"):])

	t, err := c.GetTournament(id.Get())
	if err != nil {
		errproc.HandleErr(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(t)
	if err != nil {
		errproc.HandleJSONErr(err, w)
		return
	}
}

// DeleteTournamentHandler is the http handler for deleting tournaments by id.
func (c TournamentController) DeleteTournamentHandler(w http.ResponseWriter, r *http.Request) {
	var id myuuid.MyUUID
	id.FromString(r.URL.Path[len("/tournament/"):])

	err := c.DeleteTournament(id.Get())
	if err != nil {
		errproc.HandleErr(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// JoinTournamentHandler is the http handler for assigning a new participant
// to the tournament using theirs id's.
func (c TournamentController) JoinTournamentHandler(w http.ResponseWriter, r *http.Request) {
	var id myuuid.MyUUID
	id.FromString(r.URL.Path[len("/tournament/") : len("/tournament/")+36])

	var st struct{ ID uuid.UUID `json:"userId"` }

	err := json.NewDecoder(r.Body).Decode(&st)
	if err != nil {
		errproc.HandleJSONErr(err, w)
		return
	}

	err = c.JoinTournament(id.Get(), st.ID)
	if err != nil {
		errproc.HandleErr(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// FinishTournamentHandler is the http handler for finishing the tournament by id.
// Updates winner of the tournament and adding prize to winner's balance.
func (c TournamentController) FinishTournamentHandler(w http.ResponseWriter, r *http.Request) {
	var id myuuid.MyUUID
	id.FromString(r.URL.Path[len("/tournament/") : len("/tournament/")+36])

	err := c.FinishTournament(id.Get())
	if err != nil {
		errproc.HandleErr(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}
