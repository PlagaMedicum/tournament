package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"tournament/src/myHandler"
)

type User struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

type Tournament struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Deposit int      `json:"deposit"`
	Prize   int      `json:"prize"`
	Users   []string `json:"users"`
	Winner  string   `json:"winner"`
}

var (
	userList       []User
	tournamentList []Tournament
)

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func createUser(response http.ResponseWriter, request *http.Request) {
	var user User
	handleError(json.NewDecoder(request.Body).Decode(&user))
	user.ID = strconv.Itoa(rand.Intn(1000000))
	user.Balance = 700
	userList = append(userList, user)

	response.WriteHeader(201)
	handleError(json.NewEncoder(response).Encode(user.ID))
}

func getUser(response http.ResponseWriter, request *http.Request) {
	id := request.URL.Path[len("/user/"):]
	for _, user := range userList {
		if user.ID == id {
			response.WriteHeader(200)
			handleError(json.NewEncoder(response).Encode(user))
			return
		}
	}
	response.WriteHeader(200)
	handleError(json.NewEncoder(response).Encode(&User{}))
}

func deleteUser(response http.ResponseWriter, request *http.Request) {
	id := request.URL.Path[len("/user/"):]
	for index := range userList {
		if userList[index].ID == id {
			userList = append(userList[:index], userList[index+1:]...)
			response.WriteHeader(200)
			break
		}
	}
}

func takePoints(response http.ResponseWriter, request *http.Request) {
	id := request.URL.Path[len("/user/") : len("/user/")+6]
	for index := range userList {
		if userList[index].ID == id {
			var uc struct{ Points int `json:"points"` }
			handleError(json.NewDecoder(request.Body).Decode(&uc))
			userList[index].Balance -= uc.Points
			response.WriteHeader(200)
			break
		}
	}
}

func givePoints(response http.ResponseWriter, request *http.Request) {
	id := request.URL.Path[len("/user/") : len("/user/")+6]
	for index := range userList {
		if userList[index].ID == id {
			var uc struct{ Points int `json:"points"` }
			handleError(json.NewDecoder(request.Body).Decode(&uc))
			userList[index].Balance += uc.Points
			response.WriteHeader(200)
			break
		}
	}
}

func createTournament(response http.ResponseWriter, request *http.Request) {
	var tournament Tournament
	handleError(json.NewDecoder(request.Body).Decode(&tournament))
	tournament.ID = strconv.Itoa(rand.Intn(1000000))
	tournament.Prize = 4000
	tournament.Winner = "0"
	tournamentList = append(tournamentList, tournament)

	response.WriteHeader(201)
	handleError(json.NewEncoder(response).Encode(tournament.ID))
}

func getTournament(response http.ResponseWriter, request *http.Request) {
	id := request.URL.Path[len("/tournament/"):]
	for _, tournament := range tournamentList {
		if tournament.ID == id {
			response.WriteHeader(200)
			handleError(json.NewEncoder(response).Encode(tournament))
			return
		}
	}
	response.WriteHeader(200)
	handleError(json.NewEncoder(response).Encode(&Tournament{}))
}

func deleteTournament(response http.ResponseWriter, request *http.Request) {
	id := request.URL.Path[len("/tournament/"):]
	for index := range tournamentList {
		if tournamentList[index].ID == id {
			tournamentList = append(tournamentList[:index], tournamentList[index+1:]...)
			response.WriteHeader(200)
			break
		}
	}
}

func joinTournament(response http.ResponseWriter, request *http.Request) {
	tournamentId := request.URL.Path[len("/tournament/") : len("/tournament/")+6]
	for tindex := range tournamentList {
		if tournamentList[tindex].ID == tournamentId {
			var tc struct { UserID string `json:"userId"` }
			handleError(json.NewDecoder(request.Body).Decode(&tc))
			for uindex := range userList {
				if userList[uindex].ID == tc.UserID {
					if userList[uindex].Balance >= tournamentList[tindex].Deposit {
						userList[uindex].Balance -= tournamentList[tindex].Deposit
						tournamentList[tindex].Users = append(tournamentList[tindex].Users[:], tc.UserID)
						response.WriteHeader(200)
					}
					break
				}
			}
			break
		}
	}
}

func finishTournament(response http.ResponseWriter, request *http.Request) {
	tournamentId := request.URL.Path[len("/tournament/") : len("/tournament/")+6]
	for tindex := range tournamentList {
		if tournamentList[tindex].ID == tournamentId {
			if tournamentList[tindex].Winner == "0" {
				maxBalance := -999999
				for _, userId := range tournamentList[tindex].Users {
					for uindex, user := range userList {
						if user.ID == userId {
							if user.Balance > maxBalance {
								maxBalance = user.Balance
								tournamentList[tindex].Winner = user.ID
								userList[uindex].Balance += tournamentList[tindex].Prize
								response.WriteHeader(200)
							}
							break
						}
					}
				}
			}
			break
		}
	}
}

func main() {
	var myHandler myHandler.MyHandler
	myHandler.HandleFunc("/user", createUser, "POST")
	myHandler.HandleFunc("/user/([0-9]+)", getUser, "GET")
	myHandler.HandleFunc("/user/([0-9]+)", deleteUser, "DELETE")
	myHandler.HandleFunc("/user/([0-9]+)/take", takePoints, "POST")
	myHandler.HandleFunc("/user/([0-9]+)/fund", givePoints, "POST")
	myHandler.HandleFunc("/tournament", createTournament, "POST")
	myHandler.HandleFunc("/tournament/([0-9]+)", getTournament, "GET")
	myHandler.HandleFunc("/tournament/([0-9]+)", deleteTournament, "DELETE")
	myHandler.HandleFunc("/tournament/([0-9]+)/join", joinTournament, "POST")
	myHandler.HandleFunc("/tournament/([0-9]+)/finish", finishTournament, "POST")

	handleError(http.ListenAndServe(":9090", &myHandler))
}
