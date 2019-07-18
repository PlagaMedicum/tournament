package handlers

import (
	"encoding/json"
	"net/http"
	"tournament/env/errproc"
	"tournament/env/myuuid"
	"tournament/pkg/user/model"
	"tournament/pkg/user/usecases"
)

// CreateUser is the http handler for creating users with
// specified name. Default balance is 700.
// Writes the user id in response body.
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

// GetUser is the http handler for getting
// information about user with specified id.
// Writes user's name and balance in response body.
func GetUser(w http.ResponseWriter, r *http.Request) {
	var id myuuid.MyUUID
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

// DeleteUser is the http handler for deleting users by id.
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var id myuuid.MyUUID
	id.FromString(r.URL.Path[len("/user/"):])

	err := usecases.DeleteUser(id.Get())
	if err != nil {
		errproc.HandleErr(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// TakePoints is the http handler for taking points
// from user with specified id
func TakePoints(w http.ResponseWriter, r *http.Request) {
	var id myuuid.MyUUID
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

// GivePoints is the http handler for giving points
// to users with specified id.
func GivePoints(w http.ResponseWriter, r *http.Request) {
	var id myuuid.MyUUID
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
