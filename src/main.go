package main

import (
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"tournament/src/database"
	"tournament/src/errproc"
	"tournament/src/mhandler"
)

var (
	db database.DB
)

func createUser(response http.ResponseWriter, request *http.Request) {
	var user database.User
	errproc.HandleJSONErr(json.NewDecoder(request.Body).Decode(&user))
	user.Balance = 700
	db.CreateUser(&user)
	response.WriteHeader(201)
	errproc.HandleJSONErr(json.NewEncoder(response).Encode(user.Id))
}

func getUser(response http.ResponseWriter, request *http.Request) {
	var id database.ID
	id.FromString(request.URL.Path[len("/user/"):])
	user := db.GetUser(id.Get())
	response.WriteHeader(200)
	errproc.HandleJSONErr(json.NewEncoder(response).Encode(user))
}

func deleteUser(response http.ResponseWriter, request *http.Request) {
	var id database.ID
	id.FromString(request.URL.Path[len("/user/"):])
	db.DeleteUser(id.Get())
	response.WriteHeader(200)
}

func takePoints(response http.ResponseWriter, request *http.Request) {
	var id database.ID
	id.FromString(request.URL.Path[len("/user/") : len("/user/") + 36])
	var st struct{ Points int `json:"points"` }
	errproc.HandleJSONErr(json.NewDecoder(request.Body).Decode(&st))
	db.FundUser(id.Get(), -st.Points)
	response.WriteHeader(200)

}

func givePoints(response http.ResponseWriter, request *http.Request) {
	var id database.ID
	id.FromString(request.URL.Path[len("/user/") : len("/user/") + 36])
	var st struct{ Points int `json:"points"` }
	errproc.HandleJSONErr(json.NewDecoder(request.Body).Decode(&st))
	db.FundUser(id.Get(), st.Points)
	response.WriteHeader(200)
}

func createTournament(response http.ResponseWriter, request *http.Request) {
	var tournament database.Tournament
	errproc.HandleJSONErr(json.NewDecoder(request.Body).Decode(&tournament))
	tournament.Prize = 4000
	tournament.Winner.FromString("00000000-0000-0000-0000-000000000000")

	db.CreateTournament(&tournament)

	response.WriteHeader(201)
	errproc.HandleJSONErr(json.NewEncoder(response).Encode(tournament.Id))
}

func getTournament(response http.ResponseWriter, request *http.Request) {
	var id database.ID
	id.FromString(request.URL.Path[len("/tournament/"):])
	tournament := db.GetTournament(id)
	response.WriteHeader(200)
	errproc.HandleJSONErr(json.NewEncoder(response).Encode(tournament))
}

func deleteTournament(response http.ResponseWriter, request *http.Request) {
	var id database.ID
	id.FromString(request.URL.Path[len("/tournament/"):])
	db.DeleteTournament(id.Get())
	response.WriteHeader(200)
}

func joinTournament(response http.ResponseWriter, request *http.Request) {
	var id database.ID
	id.FromString(request.URL.Path[len("/tournament/") : len("/tournament/") + 36])
	var st struct{ ID uuid.UUID `json:"userId"` }
	errproc.HandleJSONErr(json.NewDecoder(request.Body).Decode(&st))
	db.JoinTournament(id.Get(), st.ID)
	response.WriteHeader(200)
}

func finishTournament(response http.ResponseWriter, request *http.Request) {
	var id database.ID
	id.FromString(request.URL.Path[len("/tournament/") : len("/tournament/") + 36])
	db.FinishTournament(id.Get())
	response.WriteHeader(200)
}

func main() {
	db.Connect("tournament-app")
	db.InitTables()

	var handler mhandler.Handler
	handler.HandleFunc("/user", createUser, "POST")
	handler.HandleFunc("/user/([0-9]+)", getUser, "GET")
	handler.HandleFunc("/user/([0-9]+)", deleteUser, "DELETE")
	handler.HandleFunc("/user/([0-9]+)/take", takePoints, "POST")
	handler.HandleFunc("/user/([0-9]+)/fund", givePoints, "POST")
	handler.HandleFunc("/tournament", createTournament, "POST")
	handler.HandleFunc("/tournament/([0-9]+)", getTournament, "GET")
	handler.HandleFunc("/tournament/([0-9]+)", deleteTournament, "DELETE")
	handler.HandleFunc("/tournament/([0-9]+)/join", joinTournament, "POST")
	handler.HandleFunc("/tournament/([0-9]+)/finish", finishTournament, "POST")

	server := &http.Server{
		Addr:    ":9090",
		Handler: &handler,
	}

	err := server.ListenAndServe()
	errproc.FprintErr("Unexpected http server error: %v\n", err)
	defer db.Close()
}