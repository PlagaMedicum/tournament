package http

import (
	"encoding/json"
	"net/http"
	"tournament/pkg/domain"
	"tournament/pkg/infrastructure/errHandler"
)

// CreateUserHandler is the http handler for creating users with
// specified name. Default balance is 700.
// Writes the user id in response body.
func (c Controller) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var u domain.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		errHandler.HandleJSONErr(err, w)
		return
	}

	u.ID, err = c.CreateUser(u.Name)
	if err != nil {
		errHandler.HandleErr(err, w)
		return
	}

	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(u.ID)
	if err != nil {
		errHandler.HandleJSONErr(err, w)
		return
	}
}

// GetUserHandler is the http handler for getting
// information about user with specified id.
// Writes user's name and balance in response body.
func (c Controller) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len(UserPath+"/"):]

	u, err := c.GetUser(id)
	if err != nil {
		errHandler.HandleErr(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(u)
	if err != nil {
		errHandler.HandleJSONErr(err, w)
		return
	}
}

// DeleteUserHandler is the http handler for deleting users by id.
func (c Controller) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len(UserPath+"/"):]

	err := c.DeleteUser(id)
	if err != nil {
		errHandler.HandleErr(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// TakePointsHandler is the http handler for taking points
// from user with specified id
func (c Controller) TakePointsHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len(UserPath+"/") : len(UserPath+"/")+36]

	var st struct{ Points int `json:"points"` }

	err := json.NewDecoder(r.Body).Decode(&st)
	if err != nil {
		errHandler.HandleJSONErr(err, w)
		return
	}

	err = c.FundUser(id, -st.Points)
	if err != nil {
		errHandler.HandleErr(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)

}

// GivePointsHandler is the http handler for giving points
// to users with specified id.
func (c Controller) GivePointsHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len(UserPath+"/") : len(UserPath+"/")+36]

	var st struct{ Points int `json:"points"` }

	err := json.NewDecoder(r.Body).Decode(&st)
	if err != nil {
		errHandler.HandleJSONErr(err, w)
		return
	}

	err = c.FundUser(id, st.Points)
	if err != nil {
		errHandler.HandleErr(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}
