package user

import (
	"encoding/json"
	"net/http"
	"tournament/pkg/controllers/api/http_handlers"
	"tournament/pkg/domain/user"
	errors "tournament/pkg/infrastructure/err"
)

type Controller struct {
	Usecases
}

// CreateUserHandler is the http handler for creating users with
// specified name. Default balance is 700.
// Writes the user id in response body.
func (c Controller) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var u user.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		errors.HandleJSONErr(err, w)
		return
	}

	u.ID, err = c.CreateUser(u.Name)
	if err != nil {
		errors.HandleErr(err, w)
		return
	}

	w.WriteHeader(http.StatusCreated)

	type resp struct{ ID uint64 `json:"id"` }

	err = json.NewEncoder(w).Encode(resp{ID: u.ID})
	if err != nil {
		errors.HandleJSONErr(err, w)
		return
	}
}

// GetUserHandler is the http handler for getting
// information about user with specified id.
// Writes user's name and balance in response body.
func (c Controller) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := http_handlers.ScanID(http_handlers.UserPath, r)
	if err != nil {
		errors.HandleErr(err, w)
		return
	}

	u, err := c.GetUser(id)
	if err != nil {
		errors.HandleErr(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(u)
	if err != nil {
		errors.HandleJSONErr(err, w)
		return
	}
}

// DeleteUserHandler is the http handler for deleting users by id.
func (c Controller) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := http_handlers.ScanID(http_handlers.UserPath, r)
	if err != nil {
		errors.HandleErr(err, w)
		return
	}

	err = c.DeleteUser(id)
	if err != nil {
		errors.HandleErr(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// TakePointsHandler is the http handler for taking points
// from user with specified id
func (c Controller) TakePointsHandler(w http.ResponseWriter, r *http.Request) {
	id, err := http_handlers.ScanID(http_handlers.UserPath, r)
	if err != nil {
		errors.HandleErr(err, w)
		return
	}

	var st struct{ Points int `json:"points"` }

	err = json.NewDecoder(r.Body).Decode(&st)
	if err != nil {
		errors.HandleJSONErr(err, w)
		return
	}

	err = c.FundUser(id, -st.Points)
	if err != nil {
		errors.HandleErr(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)

}

// GivePointsHandler is the http handler for giving points
// to users with specified id.
func (c Controller) GivePointsHandler(w http.ResponseWriter, r *http.Request) {
	id, err := http_handlers.ScanID(http_handlers.UserPath, r)
	if err != nil {
		errors.HandleErr(err, w)
		return
	}

	var st struct{ Points int `json:"points"` }

	err = json.NewDecoder(r.Body).Decode(&st)
	if err != nil {
		errors.HandleJSONErr(err, w)
		return
	}

	err = c.FundUser(id, st.Points)
	if err != nil {
		errors.HandleErr(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}
