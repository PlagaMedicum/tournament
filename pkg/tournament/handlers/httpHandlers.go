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

// CreateTournament is the http handler for creating tournaments
// with specified name and deposit. Default prize is 4000.
// Writes tournament's id in response body.
func CreateTournament(w http.ResponseWriter, r *http.Request) {
	var t model.Tournament
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		errproc.HandleJSONErr(err, w)

		return
	}

	err = usecases.CreateTournament(t)
	if err != nil {
		errproc.HandleErr(err, w)

		return
	}

	err = json.NewEncoder(w).Encode(t.ID)
	if err != nil {
		errproc.HandleJSONErr(err, w)

		return
	}
	w.WriteHeader(http.StatusCreated)
}

// GetTournament is the http handler for getting information
// about tournaments with specified id.
// Writes tournament's name, deposit, prize, list of participants
// and winner in response body.
func GetTournament(w http.ResponseWriter, r *http.Request) {
	var id myuuid.MyUUID
	id.FromString(r.URL.Path[len("/tournament/"):])

	t, err := usecases.GetTournament(id.Get())
	if err != nil {
		errproc.HandleErr(err, w)

		return
	}

	err = json.NewEncoder(w).Encode(t)
	if err != nil {
		errproc.HandleJSONErr(err, w)

		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteTournament is the http handler for deleting tournaments by id.
func DeleteTournament(w http.ResponseWriter, r *http.Request) {
	var id myuuid.MyUUID
	id.FromString(r.URL.Path[len("/tournament/"):])

	err := usecases.DeleteTournament(id.Get())
	if err != nil {
		errproc.HandleErr(err, w)

		return
	}

	w.WriteHeader(http.StatusOK)
}

// JoinTournament is the http handler for adding a new participant
// in the tournament using theirs id's.
func JoinTournament(w http.ResponseWriter, r *http.Request) {
	var id myuuid.MyUUID
	id.FromString(r.URL.Path[len("/tournament/") : len("/tournament/")+36])
	var st struct{ ID uuid.UUID `json:"userId"` }

	err := json.NewDecoder(r.Body).Decode(&st)
	if err != nil {
		errproc.HandleJSONErr(err, w)

		return
	}

	err = usecases.JoinTournament(id.Get(), st.ID)
	if err != nil {
		errproc.HandleErr(err, w)

		return
	}

	w.WriteHeader(http.StatusOK)
}

// FinishTournament is the http handler for finishing the tournament by id.
// Updates winner of the tournament and adding prize to winner's balance.
func FinishTournament(w http.ResponseWriter, r *http.Request) {
	var id myuuid.MyUUID
	id.FromString(r.URL.Path[len("/tournament/") : len("/tournament/")+36])

	err := usecases.FinishTournament(id.Get())
	if err != nil {
		errproc.HandleErr(err, w)

		return
	}

	w.WriteHeader(http.StatusOK)
}
