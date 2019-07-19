package tests

import (
	"bytes"
	"errors"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"net/http/httptest"
	"testing"
	"tournament/env/myhandler"
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
	requestBody			[]byte
	expectedStatus 		int
}

func handleTester(t *testing.T, h myhandler.Handler, tr tester){
	r, err := http.NewRequest(
		tr.method, tr.path, bytes.NewBuffer(tr.requestBody))
	if err != nil {
		t.Errorf("Cannot encode request: '%v'", err)
	}

	w := httptest.NewRecorder()

	h.ServeHTTP(w, r)

	if w.Code != tr.expectedStatus {
		t.Errorf("Test case: %s. Wrong status code.\nExpected: %v\nGot: %v\nResponse body is: %s",
			tr.caseName, tr.expectedStatus, w.Code, w.Body)
	}
}

// TestCreateUser tests creation of user.
func TestCreateUserHandling(t *testing.T) {
	trList := []tester{
		{
			caseName:   "everything ok",
			path:       router.UserPath,
			method:     http.MethodPost,
			methodName: "CreateUser",
			requestUser: user.User{
				Name: "Daniil Dankovskij",
			},
			requestBody: []byte(`{"name": "Daniil Dankovskij"}`),
			expectedStatus: http.StatusCreated,
		},
		{
			caseName:   "wrong body",
			path:       router.UserPath,
			method:     http.MethodPost,
			methodName: "CreateUser",
			requestBody: []byte(`{"name": "Vladislav Olgimskij"`),
			expectedStatus: http.StatusBadRequest,
		},
		{
			caseName:   "wrong user error",
			path:       router.UserPath,
			method:     http.MethodPost,
			methodName: "CreateUser",
			requestUser: user.User{
				Name: "Artemij Burah",
			},
			requestBody: []byte(`{"name": "Artemij Burah"}`),
			expectedStatus: http.StatusBadRequest,
		},
	}

	mo := mockedUser{}
	mo.On(trList[0].methodName, trList[0].requestUser).Return(uuid.NewV1(), nil)
	mo.On(trList[2].methodName, trList[2].requestUser).Return(uuid.NewV1(), errors.New("wrong user error"))

	h := myhandler.Handler{}
	router.RouteForUser(&h, &mo)

	for _, tr := range trList {
		handleTester(t, h, tr)
	}

	mo.AssertExpectations(t)
}
