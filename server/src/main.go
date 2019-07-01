package main

import (
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"tournament/src/database"
	"tournament/src/errproc"
	"tournament/src/mhandler"
)

const (
	UUIDRegex = "[0-9+a-z]{8}-[0-9+a-z]{4}-[0-9+a-z]{4}-[0-9+a-z]{4}-[0-9+a-z]{12}"
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
	db.CreateTournament(&tournament)
	response.WriteHeader(201)
	errproc.HandleJSONErr(json.NewEncoder(response).Encode(tournament.Id))
}

func getTournament(response http.ResponseWriter, request *http.Request) {
	var id database.ID
	id.FromString(request.URL.Path[len("/tournament/"):])
	tournament := db.GetTournament(id.Get())
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
	code := db.FinishTournament(id.Get())
	response.WriteHeader(code)
}

func main() {
	db.Connect("tournament-app")
	db.InitTables()

	var handler mhandler.Handler
	handler.HandleFunc("^/user$", createUser, "POST")
	handler.HandleFunc("^/user/" +UUIDRegex+ "$", getUser, "GET")
	handler.HandleFunc("^/user/" +UUIDRegex+ "$", deleteUser, "DELETE")
	handler.HandleFunc("^/user/" +UUIDRegex+ "/take$", takePoints, "POST")
	handler.HandleFunc("^/user/" +UUIDRegex+ "/fund$", givePoints, "POST")
	handler.HandleFunc("^/tournament$", createTournament, "POST")
	handler.HandleFunc("^/tournament/" +UUIDRegex+ "$", getTournament, "GET")
	handler.HandleFunc("^/tournament/" +UUIDRegex+ "$", deleteTournament, "DELETE")
	handler.HandleFunc("^/tournament/" +UUIDRegex+ "/join$", joinTournament, "POST")
	handler.HandleFunc("^/tournament/" +UUIDRegex+ "/finish$", finishTournament, "POST")

	server := &http.Server{
		Addr:    ":8080",
		Handler: &handler,
	}

	err := server.ListenAndServe()
	errproc.FprintErr("Unexpected http server error: %v\n", err)
	defer db.Close()
}