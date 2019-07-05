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

func CreateTournament(response http.ResponseWriter, request *http.Request) {
	var t tournament.Tournament
	errproc.HandleJSONErr(json.NewDecoder(request.Body).Decode(&t))
	t.Prize = 4000
	usecases.CreateTournament(&t)
	response.WriteHeader(201)
	errproc.HandleJSONErr(json.NewEncoder(response).Encode(t.ID))
}

func GetTournament(response http.ResponseWriter, request *http.Request) {
	var id mid.MID
	id.FromString(request.URL.Path[len("/tournament/"):])
	t := usecases.GetTournament(id.Get())
	response.WriteHeader(200)
	errproc.HandleJSONErr(json.NewEncoder(response).Encode(t))
}

func DeleteTournament(response http.ResponseWriter, request *http.Request) {
	var id mid.MID
	id.FromString(request.URL.Path[len("/tournament/"):])
	usecases.DeleteTournament(id.Get())
	response.WriteHeader(200)
}

func JoinTournament(response http.ResponseWriter, request *http.Request) {
	var id mid.MID
	id.FromString(request.URL.Path[len("/tournament/") : len("/tournament/")+36])
	var st struct{ ID uuid.UUID `json:"userId"` }
	errproc.HandleJSONErr(json.NewDecoder(request.Body).Decode(&st))
	usecases.JoinTournament(id.Get(), st.ID)
	response.WriteHeader(200)
}

func FinishTournament(response http.ResponseWriter, request *http.Request) {
	var id mid.MID
	id.FromString(request.URL.Path[len("/tournament/") : len("/tournament/")+36])
	code := usecases.FinishTournament(id.Get())
	response.WriteHeader(code)
}
