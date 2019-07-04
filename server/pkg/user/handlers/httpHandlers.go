package handlers

import (
	"encoding/json"
	"net/http"
	"tournament/pkg/errproc"
	"tournament/pkg/mid"
	"tournament/pkg/user/model"
	"tournament/pkg/user/usecases"
)

func CreateUser(response http.ResponseWriter, request *http.Request) {
	var u model.User
	errproc.HandleJSONErr(json.NewDecoder(request.Body).Decode(&u))
	u.Balance = 700
	usecases.CreateUser(&u)
	response.WriteHeader(201)
	errproc.HandleJSONErr(json.NewEncoder(response).Encode(u.Id))
}

func GetUser(response http.ResponseWriter, request *http.Request) {
	var id mid.ID
	id.FromString(request.URL.Path[len("/u/"):])
	u := usecases.GetUser(id.Get())
	response.WriteHeader(200)
	errproc.HandleJSONErr(json.NewEncoder(response).Encode(u))
}

func DeleteUser(response http.ResponseWriter, request *http.Request) {
	var id mid.ID
	id.FromString(request.URL.Path[len("/user/"):])
	usecases.DeleteUser(id.Get())
	response.WriteHeader(200)
}

func TakePoints(response http.ResponseWriter, request *http.Request) {
	var id mid.ID
	id.FromString(request.URL.Path[len("/user/") : len("/user/") + 36])
	var st struct{ Points int `json:"points"` }
	errproc.HandleJSONErr(json.NewDecoder(request.Body).Decode(&st))
	usecases.FundUser(id.Get(), -st.Points)
	response.WriteHeader(200)

}

func GivePoints(response http.ResponseWriter, request *http.Request) {
	var id mid.ID
	id.FromString(request.URL.Path[len("/user/") : len("/user/") + 36])
	var st struct{ Points int `json:"points"` }
	errproc.HandleJSONErr(json.NewDecoder(request.Body).Decode(&st))
	usecases.FundUser(id.Get(), st.Points)
	response.WriteHeader(200)
}
