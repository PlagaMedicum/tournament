package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
)

type User struct{
	id 		string
	name 	string
	balance int
}

type Tournament struct {
	id 		string
	name  	string
	deposit int
	prize   int
	users   []int
	winner  int
}

var (
	userList   []User
	tournament []Tournament
)

func createUser(response http.ResponseWriter, request *http.Request){
	if request.Method == "POST"{
		response.Header().Set("Content-Type", "application/json")
		var user User
		err := json.NewDecoder(request.Body).Decode(&user); if err != nil {
			panic(err)
		}
		user.id = strconv.Itoa(rand.Intn(1000000))
		user.balance = 700
		userList = append(userList, user)

		response.WriteHeader(201)
		err = json.NewEncoder(response).Encode(user.name); if err != nil {
			panic(err)
		}
	}
}

func processUser(response http.ResponseWriter, request *http.Request){
	switch request.Method {
	case "GET":
		response.Header().Set("Content-Type", "application/json")
		id := request.URL.Path[len("/user/"):]
		for _, item := range userList {
			if item.id == id {
				err := json.NewEncoder(response).Encode(item.name); if err != nil {
					panic(err)
				}
				return
			}
		}
		err := json.NewEncoder(response).Encode(&User{}); if err != nil {
			panic(err)
		}
		response.WriteHeader(200)
	case "DELETE":
		response.WriteHeader(200)
	}
}

func subtractUserPoints(response http.ResponseWriter, request *http.Request){
	if request.Method == "POST"{
		response.WriteHeader(201)
	}
}

func addUserPoints(response http.ResponseWriter, request *http.Request){
	if request.Method == "POST"{
		response.WriteHeader(201)
	}
}

func createTournament(response http.ResponseWriter, request *http.Request){
	if request.Method == "POST"{
		response.WriteHeader(201)
	}
}

func processTournamentInfo(response http.ResponseWriter, request *http.Request){
	switch request.Method {
	case "GET":
		response.WriteHeader(200)
	case "DELETE":
		response.WriteHeader(200)
	}
}

func joinTournament(response http.ResponseWriter, request *http.Request){
	if request.Method == "POST"{
		response.WriteHeader(201)
	}
}

func finishTournament(response http.ResponseWriter, request *http.Request){
	if request.Method == "POST"{
		response.WriteHeader(201)
	}
}

func main(){
	http.HandleFunc("/user", createUser)
	http.HandleFunc("/user/", processUser)
	http.HandleFunc("user/{id}/take", subtractUserPoints)
	http.HandleFunc("user/{id}/fund", addUserPoints)
	http.HandleFunc("/tournament", createTournament)
	http.HandleFunc("/tournament/{id}", processTournamentInfo)
	http.HandleFunc("/tournament/{id}/join", joinTournament)
	http.HandleFunc("/tournament/{id}/finish", finishTournament)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		panic(err)
	}
}
