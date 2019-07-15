package handlers

import (
	"encoding/json"
	"net/http"
	"tournament/env/errproc"
	"tournament/env/mid"
	"tournament/pkg/user/model"
	"tournament/pkg/user/usecases"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var u model.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		errproc.HandleJSONErr(err, w)
		return
	}
	u.Balance = 700
	err = usecases.CreateUser(&u)
	if err != nil {
		errproc.HandleErr(err, w)
		return
	}
	err = json.NewEncoder(w).Encode(u.ID)
	if err != nil {
		errproc.HandleJSONErr(err, w)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var id mid.MID
	id.FromString(r.URL.Path[len("/user/"):])
	u, err := usecases.GetUser(id.Get())
	if err != nil {
		errproc.HandleErr(err, w)
		return
	}
	err = json.NewEncoder(w).Encode(u)
	if err != nil {
		errproc.HandleJSONErr(err, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var id mid.MID
	id.FromString(r.URL.Path[len("/user/"):])
	err := usecases.DeleteUser(id.Get())
	if err != nil {
		errproc.HandleErr(err, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func TakePoints(w http.ResponseWriter, r *http.Request) {
	var id mid.MID
	id.FromString(r.URL.Path[len("/user/") : len("/user/")+36])
	var st struct{ Points int `json:"points"` }
	err := json.NewDecoder(r.Body).Decode(&st)
	if err != nil {
		errproc.HandleJSONErr(err, w)
		return
	}
	err = usecases.FundUser(id.Get(), -st.Points)
	if err != nil {
		errproc.HandleErr(err, w)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func GivePoints(w http.ResponseWriter, r *http.Request) {
	var id mid.MID
	id.FromString(r.URL.Path[len("/user/") : len("/user/")+36])
	var st struct{ Points int `json:"points"` }
	err := json.NewDecoder(r.Body).Decode(&st)
	if err != nil {
		errproc.HandleJSONErr(err, w)
		return
	}
	err = usecases.FundUser(id.Get(), st.Points)
	if err != nil {
		errproc.HandleErr(err, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}
