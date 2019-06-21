package main

import (
	"net/http"
)

type User struct{
	id uint
	name string
	balance int
}

type Tournament struct {
	id uint
	name string
	deposit int
	prize int
	users []int
	winner int
}

func createUser(response http.ResponseWriter, request *http.Request){
	if request.Method == "POST"{
		response.WriteHeader(201)
	}
}

func processUser(response http.ResponseWriter, request *http.Request){
	switch request.Method {
	case "GET":
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
	http.HandleFunc("/user/{id}", processUser)
	http.HandleFunc("user/{id}/take", subtractUserPoints)
	http.HandleFunc("user/{id}/fund", addUserPoints)
	http.HandleFunc("/tournament", createTournament)
	http.HandleFunc("/tournament/{id}", processTournamentInfo)
	http.HandleFunc("/tournament/{id}/join", joinTournament)
	http.HandleFunc("/tournament/{id}/finish", finishTournament)

	if err := http.ListenAndServe(":9090", nil); err != nil{
		panic(err)
	}
}
