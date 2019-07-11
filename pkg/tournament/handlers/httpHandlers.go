package handlers

import (
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"tournament/pkg/errproc"
	"tournament/pkg/mid"
	tournament "tournament/pkg/tournament/model"
	"tournament/pkg/tournament/usecases"
)

func CreateTournament(w http.ResponseWriter, r *http.Request) {
	var t tournament.Tournament
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		errproc.HandleJSONErr(err, w)
		return
	}
	t.Prize = 4000
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
	w.WriteHeader(201)
}

func GetTournament(w http.ResponseWriter, r *http.Request) {
	var id mid.MID
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
	w.WriteHeader(200)
}

func DeleteTournament(w http.ResponseWriter, r *http.Request) {
	var id mid.MID
	id.FromString(r.URL.Path[len("/tournament/"):])
	err := usecases.DeleteTournament(id.Get())
	if err != nil {
		errproc.HandleErr(err, w)
		return
	}
	w.WriteHeader(200)
}

func JoinTournament(w http.ResponseWriter, r *http.Request) {
	var id mid.MID
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
	w.WriteHeader(200)
}

func FinishTournament(w http.ResponseWriter, r *http.Request) {
	var id mid.MID
	id.FromString(r.URL.Path[len("/tournament/") : len("/tournament/")+36])
	err := usecases.FinishTournament(id.Get())
	if err != nil {
		errproc.HandleErr(err, w)
		return
	}
	w.WriteHeader(200)
}
