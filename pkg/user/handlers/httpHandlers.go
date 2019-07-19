package handlers

import (
	"encoding/json"
	"net/http"
	"tournament/env/errproc"
	"tournament/env/myuuid"
	"tournament/pkg/user/model"
	"tournament/pkg/user/usecases"
)

type UserController struct {
	usecases.UserInterface
}

// CreateUserHandler is the http handler for creating users with
// specified name. Default balance is 700.
// Writes the user id in response body.
func (c UserController) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var u model.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		errproc.HandleJSONErr(err, w)
		return
	}

	id, err := c.CreateUser(u)
	if err != nil {
		errproc.HandleErr(err, w)
		return
	}

	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(id)
	if err != nil {
		errproc.HandleJSONErr(err, w)
		return
	}
}

// GetUserHandler is the http handler for getting
// information about user with specified id.
// Writes user's name and balance in response body.
func (c UserController) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var id myuuid.MyUUID
	id.FromString(r.URL.Path[len("/user/"):])

	u, err := c.GetUser(id.Get())
	if err != nil {
		errproc.HandleErr(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(u)
	if err != nil {
		errproc.HandleJSONErr(err, w)
		return
	}
}

// DeleteUserHandler is the http handler for deleting users by id.
func (c UserController) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var id myuuid.MyUUID
	id.FromString(r.URL.Path[len("/user/"):])

	err := c.DeleteUser(id.Get())
	if err != nil {
		errproc.HandleErr(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// TakePointsHandler is the http handler for taking points
// from user with specified id
func (c UserController) TakePointsHandler(w http.ResponseWriter, r *http.Request) {
	var id myuuid.MyUUID
	id.FromString(r.URL.Path[len("/user/") : len("/user/")+36])

	var st struct{ Points int `json:"points"` }

	err := json.NewDecoder(r.Body).Decode(&st)
	if err != nil {
		errproc.HandleJSONErr(err, w)
		return
	}

	err = c.FundUser(id.Get(), -st.Points)
	if err != nil {
		errproc.HandleErr(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)

}

// GivePointsHandler is the http handler for giving points
// to users with specified id.
func (c UserController) GivePointsHandler(w http.ResponseWriter, r *http.Request) {
	var id myuuid.MyUUID
	id.FromString(r.URL.Path[len("/user/") : len("/user/")+36])

	var st struct{ Points int `json:"points"` }

	err := json.NewDecoder(r.Body).Decode(&st)
	if err != nil {
		errproc.HandleJSONErr(err, w)
		return
	}

	err = c.FundUser(id.Get(), st.Points)
	if err != nil {
		errproc.HandleErr(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}
