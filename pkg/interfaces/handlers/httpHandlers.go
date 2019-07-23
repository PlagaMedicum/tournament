package handlers

import (
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"tournament/pkg/domain"
	"tournament/pkg/infrastructure/errproc"
	"tournament/pkg/infrastructure/myuuid"
	"tournament/pkg/usecases"
)

type Controller struct {
	usecases.RepositoryInteractor
}

// CreateUserHandler is the http handler for creating users with
// specified name. Default balance is 700.
// Writes the user id in response body.
func (c Controller) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var u domain.User

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
func (c Controller) GetUserHandler(w http.ResponseWriter, r *http.Request) {
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
func (c Controller) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
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
func (c Controller) TakePointsHandler(w http.ResponseWriter, r *http.Request) {
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
func (c Controller) GivePointsHandler(w http.ResponseWriter, r *http.Request) {
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

// CreateTournamentHandler is the http handler for creating tournaments
// with specified name and deposit. Default prize is 4000.
// Writes tournament's id in response body.
func (c Controller) CreateTournamentHandler(w http.ResponseWriter, r *http.Request) {
	var t domain.Tournament

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		errproc.HandleJSONErr(err, w)
		return
	}

	id, err := c.CreateTournament(t)
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

// GetTournamentHandler is the http handler for getting information
// about tournaments with specified id.
// Writes tournament's name, deposit, prize, list of participants
// and winner in response body.
func (c Controller) GetTournamentHandler(w http.ResponseWriter, r *http.Request) {
	var id myuuid.MyUUID
	id.FromString(r.URL.Path[len("/tournament/"):])

	t, err := c.GetTournament(id.Get())
	if err != nil {
		errproc.HandleErr(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(t)
	if err != nil {
		errproc.HandleJSONErr(err, w)
		return
	}
}

// DeleteTournamentHandler is the http handler for deleting tournaments by id.
func (c Controller) DeleteTournamentHandler(w http.ResponseWriter, r *http.Request) {
	var id myuuid.MyUUID
	id.FromString(r.URL.Path[len("/tournament/"):])

	err := c.DeleteTournament(id.Get())
	if err != nil {
		errproc.HandleErr(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// JoinTournamentHandler is the http handler for assigning a new participant
// to the tournament using theirs id's.
func (c Controller) JoinTournamentHandler(w http.ResponseWriter, r *http.Request) {
	var id string
	id.FromString(r.URL.Path[len("/tournament/") : len("/tournament/")+36])

	var st struct{ ID string `json:"userId"` }

	err := json.NewDecoder(r.Body).Decode(&st)
	if err != nil {
		errproc.HandleJSONErr(err, w)
		return
	}

	err = c.JoinTournament(id, st.ID)
	if err != nil {
		errproc.HandleErr(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// FinishTournamentHandler is the http handler for finishing the tournament by id.
// Updates winner of the tournament and adding prize to winner's balance.
func (c Controller) FinishTournamentHandler(w http.ResponseWriter, r *http.Request) {
	var id myuuid.MyUUID
	id.FromString(r.URL.Path[len("/tournament/") : len("/tournament/")+36])

	err := c.FinishTournament(id.String())
	if err != nil {
		errproc.HandleErr(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}
