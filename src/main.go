package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
)

type User struct{
	ID 		string	`json:"id"`
	Name 	string	`json:"name"`
	Balance int		`json:"balance"`
}

type Tournament struct {
	ID 		string	`json:"id"`
	Name  	string	`json:"name"`
	Deposit int		`json:"deposit"`
	Prize   int		`json:"prize"`
	Users   []string`json:"users"`
	Winner  string	`json:"winner"`
}

type UserChanger struct {
	Points int		`json:"points"`
}

type TournamentChanger struct {
	UserID string 	`json:"userId"`
}

var (
	userList   	   []User
	tournamentList []Tournament
)

func processUser(response http.ResponseWriter, request *http.Request){
	if request.URL.Path == "/user" {
		if request.Method == "POST"{
			response.Header().Set("Content-Type", "application/json")
			var user User
			err := json.NewDecoder(request.Body).Decode(&user)
			if err != nil {
				panic(err)
			}
			user.ID = strconv.Itoa(rand.Intn(1000000))
			user.Balance = 700
			userList = append(userList, user)

			response.WriteHeader(201)
			err = json.NewEncoder(response).Encode(user.ID)
			if err != nil {
				panic(err)
			}
		}
	}
	if b,_ := regexp.Match("/user/([0-9]+)", []byte(request.URL.Path)); b {
		if request.Method == "GET"{
			response.Header().Set("Content-Type", "application/json")
			id := request.URL.Path[len("/user/"):]
			for _, user := range userList {
				if user.ID == id {
					response.WriteHeader(200)
					err := json.NewEncoder(response).Encode(user)
					if err != nil {
						panic(err)
					}
					return
				}
			}
			response.WriteHeader(200)
			err := json.NewEncoder(response).Encode(&User{})
			if err != nil {
				panic(err)
			}
		}
		if request.Method == "DELETE"{
			response.Header().Set("Content-Type", "application/json")
			id := request.URL.Path[len("/user/"):]
			for index := range userList {
				if userList[index].ID == id {
					userList = append(userList[:index], userList[index+1:]...)
					response.WriteHeader(200)
					break
				}
			}
		}
	}
	if b,_ := regexp.Match("/user/([0-9]+)/take", []byte(request.URL.Path)); b {
		if request.Method == "POST"{
			response.Header().Set("Content-Type", "application/json")
			id := request.URL.Path[len("/user/"):len("/user/") + 6]
			for index := range userList {
				if userList[index].ID == id {
					var uc UserChanger
					err := json.NewDecoder(request.Body).Decode(&uc)
					if err != nil {
						panic(err)
					}
					userList[index].Balance -= uc.Points
					response.WriteHeader(200)
					break
				}
			}
		}
	}
	if b,_ := regexp.Match("/user/([0-9]+)/fund", []byte(request.URL.Path)); b {
		if request.Method == "POST"{
			response.Header().Set("Content-Type", "application/json")
			id := request.URL.Path[len("/user/"):len("/user/") + 6]
			for index := range userList {
				if userList[index].ID == id {
					var uc UserChanger
					err := json.NewDecoder(request.Body).Decode(&uc)
					if err != nil {
						panic(err)
					}
					userList[index].Balance += uc.Points
					response.WriteHeader(200)
					break
				}
			}
		}
	}
}

func processTournament(response http.ResponseWriter, request *http.Request){
	if request.URL.Path == "/tournament"{
		if request.Method == "POST"{
			response.Header().Set("Content-Type", "application/json")
			var tournament Tournament
			err := json.NewDecoder(request.Body).Decode(&tournament)
			if err != nil {
				panic(err)
			}
			tournament.ID = strconv.Itoa(rand.Intn(1000000))
			tournament.Prize = 4000
			tournament.Winner = "0"
			tournamentList = append(tournamentList, tournament)

			response.WriteHeader(201)
			err = json.NewEncoder(response).Encode(tournament.ID)
			if err != nil {
				panic(err)
			}
		}
	}
	if b,_ := regexp.Match("/tournament/([0-9]+)", []byte(request.URL.Path)); b {
		if request.Method == "GET" {
			response.Header().Set("Content-Type", "application/json")
			id := request.URL.Path[len("/tournament/"):]
			for _, tournament := range tournamentList {
				if tournament.ID == id {
					response.WriteHeader(200)
					err := json.NewEncoder(response).Encode(tournament)
					if err != nil {
						panic(err)
					}
					return
				}
			}
			response.WriteHeader(200)
			err := json.NewEncoder(response).Encode(&Tournament{})
			if err != nil {
				panic(err)
			}
		}
		if request.Method == "DELETE" {
			response.Header().Set("Content-Type", "application/json")
			id := request.URL.Path[len("/tournament/"):]
			for index := range tournamentList {
				if tournamentList[index].ID == id {
					tournamentList = append(tournamentList[:index], tournamentList[index+1:]...)
					response.WriteHeader(200)
					break
				}
			}
		}
	}
	if b,_ := regexp.Match("/tournament/([0-9]+)/join", []byte(request.URL.Path)); b {
		if request.Method == "POST"{
			response.Header().Set("Content-Type", "application/json")
			tournamentId := request.URL.Path[len("/tournament/"):len("/tournament/") + 6]
			for tindex := range tournamentList {
				if tournamentList[tindex].ID == tournamentId {
					var tc TournamentChanger
					err := json.NewDecoder(request.Body).Decode(&tc)
					if err != nil {
						panic(err)
					}
					for uindex := range userList{
						if userList[uindex].ID == tc.UserID {
							if userList[uindex].Balance >= tournamentList[tindex].Deposit{
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
	}
	if b,_ := regexp.Match("/tournament/([0-9]+)/finish", []byte(request.URL.Path)); b {
		if request.Method == "POST"{
			response.Header().Set("Content-Type", "application/json")
			tournamentId := request.URL.Path[len("/tournament/"):len("/tournament/") + 6]
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
	}
}

func main(){
	http.HandleFunc("/user", processUser)
	http.HandleFunc("/user/", processUser)
	http.HandleFunc("/tournament", processTournament)
	http.HandleFunc("/tournament/", processTournament)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		panic(err)
	}
}
