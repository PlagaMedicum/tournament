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
	user "tournament/pkg/user/model"
)

type tester struct {
	caseName			string
	path 				string
	method				string
	methodName 			string
	resultErr			error
	resultID			uuid.UUID
	requestUser			user.User
	requestBody			[]byte
	expectedStatus 		int
}

func handleTester(t *testing.T, mo *mockedUser, tr tester) {
	h := myhandler.Handler{}
	router.RouteForUser(&h, mo)

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
			resultErr:	nil,
			resultID:	uuid.NewV1(),
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
			resultErr:	errors.New("i'm the bad err"),
			resultID:	uuid.NewV1(),
			requestUser: user.User{
				Name: "Artemij Burah",
			},
			requestBody: []byte(`{"name": "Artemij Burah"}`),
			expectedStatus: http.StatusBadRequest,
		},
	}

	mo := mockedUser{}

	for _, tr := range trList {
		mo.On(tr.methodName, tr.requestUser).Return(tr.resultId, tr.resultErr)
		handleTester(t, &mo, tr)
	}

	mo.AssertExpectations(t)
}

// TestCreateUser tests getting of user information.
func TestGetUserHandling(t *testing.T) {
	sampleID := uuid.NewV1()
	trList := []tester{
		{
			caseName:   "everything ok",
			path:       router.UserPath+sampleID,
			method:     http.MethodPost,
			methodName: "GetUser",
			resultErr:	nil,
			resultUser: ,
			expectedStatus: http.StatusCreated,
		},
		{
			caseName:   "wrong user error",
			path:       router.UserPath,
			method:     http.MethodPost,
			methodName: "GetUser",
			resultErr:	errors.New("i'm the bad err"),
			resultUser: ,
			expectedStatus: http.StatusBadRequest,
		},
	}

	mo := mockedUser{}

	for _, tr := range trList {
		mo.On(tr.methodName, tr.requestUser).Return(tr., tr.resultErr)		
		handleTester(t, &mo, tr)
	}

	mo.AssertExpectations(t)
}
