package tests

import (
	"bytes"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	app "tournament/pkg"
	"tournament/pkg/router"
	tournament "tournament/pkg/tournament/model"
	user "tournament/pkg/user/model"
)

type tester struct {
	caseName			string
	path 				string
	method				string
	methodName 			string
	requestUser			user.User
	requestTournament 	tournament.Tournament
	expectedStatus 		int
}

// TestCreateUser tests creation of user.
func TestCreateUserHandling(t *testing.T) {
	tr1 := tester{
		caseName: "everything ok",
		path: "/user",
		method: http.MethodPost,
		methodName: "CreateUser",
		requestUser: user.User{
			Name: "Daniil Dankovskij",
		},
		expectedStatus: http.StatusCreated,
	}

	obj := mockedObject{}
	obj.On(tr1.methodName, tr1.requestUser.Name).Return(uuid.NewV1(), nil)

	app.DB.InitNewTestPostgresDB()

	h := router.Route()

	r, err := http.NewRequest(
		tr1.method, tr1.path, bytes.NewBuffer([]byte(`{"name": "`+tr1.requestUser.Name+`"}`)))
	if err != nil {
		t.Errorf("Cannot encode request: '%v'", err)
	}

	w := httptest.NewRecorder()

	h.ServeHTTP(w, r)

	respStr := w.Body.String()[1:len(w.Body.String())-2]

	if b, err := regexp.MatchString(respStr, app.UUIDRegex); !b {
		t.Errorf("Response string doesn't correspond UUID regexp: %v\nResponse string: '%s'", err, respStr)
	}

	if w.Code != tr1.expectedStatus {
		t.Errorf("Wrong status code.\nExpected: %v\nGot: %v", http.StatusCreated, w.Code)
	}

	app.DB.MigrateTablesDown()
	obj.AssertExpectations(t)
}
