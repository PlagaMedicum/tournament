package tests

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"testing"
	"tournament/env/myhandler"
	"tournament/pkg/router"
	user "tournament/pkg/user/model"
)

// TestCreateUserHandler tests creation of user.
func TestCreateUserHandler(t *testing.T) {
	trList := []tester{
		{
			caseName:  "everything ok",
			resultErr: nil,
			resultID:  uuid.NewV1(),
			requestUser: user.User{
				Name: "Daniil Dankovskij",
			},
			requestBody: []byte(`{"name": "Daniil Dankovskij"}`),
			expectedStatus: http.StatusCreated,
		},
		{
			caseName: "wrong body",
			requestBody: []byte(`i'm the wrong body"`),
			requestUser: user.User{},
			expectedStatus: http.StatusBadRequest,
		},
		{
			caseName: "wrong user error",
			resultErr: errors.New("i'm the bad err"),
			resultID: uuid.NewV1(),
			requestUser: user.User{
				Name: "Artemij Burah",
			},
			requestBody: []byte(`{"name": "Artemij Burah"}`),
			expectedStatus: http.StatusBadRequest,
		},
	}

	mo := mockedUser{}
	h := myhandler.Handler{}
	router.RouteForUser(&h, &mo)

	for _, tr := range trList {
		tr.path = router.UserPath
		tr.method = http.MethodPost

		if tr.requestUser != (user.User{}) {
			mo.On("CreateUser", tr.requestUser).Return(tr.resultID, tr.resultErr)
		}

		handleTester(t, &h, tr)
	}

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}

// TestGetUserHandler tests getting of user's information.
func TestGetUserHandler(t *testing.T) {
	trList := []tester{
		{
			caseName:  "everything ok",
			resultErr: nil,
			resultUser: user.User{
				Name: "Anna Angel",
			},
			requestID: uuid.NewV1(),
			expectedStatus: http.StatusOK,
		},
		{
			caseName: "wrong user error",
			resultErr: errors.New("i'm the bad err"),
			resultUser: user.User{},
			requestID: uuid.NewV1(),
			expectedStatus: http.StatusBadRequest,
		},
	}

	mo := mockedUser{}
	h := myhandler.Handler{}
	router.RouteForUser(&h, &mo)

	for _, tr := range trList {
		tr.path = router.UserPath + "/" + tr.requestID.String()
		tr.method = http.MethodGet

		if tr.requestID != uuid.Nil{
			mo.On("GetUser", tr.requestID).Return(tr.resultUser, tr.resultErr)
		}

		handleTester(t, &h, tr)
	}

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}

// TestDeleteUserHandler tests deleting of user.
func TestDeleteUserHandler(t *testing.T) {
	trList := []tester{
		{
			caseName: "everything ok",
			resultErr: nil,
			requestID: uuid.NewV1(),
			expectedStatus: http.StatusOK,
		},
		{
			caseName: "wrong user error",
			resultErr: errors.New("i'm the bad err"),
			requestID: uuid.NewV1(),
			expectedStatus: http.StatusBadRequest,
		},
	}

	mo := mockedUser{}
	h := myhandler.Handler{}
	router.RouteForUser(&h, &mo)

	for _, tr := range trList {
		tr.path = router.UserPath + "/" + tr.requestID.String()
		tr.method = http.MethodDelete

		if tr.requestID != uuid.Nil{
			mo.On("DeleteUser", tr.requestID).Return(tr.resultErr)
		}

		handleTester(t, &h, tr)
	}

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}

// TestTakePointsHandler tests taking points from user.
func TestTakePointsHandler(t *testing.T) {
	trList := []tester{
		{
			caseName: "everything ok",
			resultErr: nil,
			requestID: uuid.NewV1(),
			requestBody: []byte(`{"points": 1}`),
			expectedStatus: http.StatusOK,
		},
		{
			caseName: "wrong body",
			requestBody: []byte(`i'm the wrong body"`),
			requestID: uuid.Nil,
			expectedStatus: http.StatusBadRequest,
		},
		{
			caseName: "wrong user error",
			resultErr: errors.New("i'm the bad err"),
			requestID: uuid.NewV1(),
			requestBody: []byte(`{"points": 1}`),
			expectedStatus: http.StatusBadRequest,
		},
	}

	mo := mockedUser{}
	h := myhandler.Handler{}
	router.RouteForUser(&h, &mo)

	for _, tr := range trList {
		tr.path = router.UserPath + "/" + tr.requestID.String() + router.TakingPointsPath
		tr.method = http.MethodPost

		if tr.requestID != uuid.Nil{
			mo.On("FundUser", tr.requestID, -1).Return(tr.resultErr)
		}

		handleTester(t, &h, tr)
	}

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}

// TestGivePointsHandler tests giving points to user.
func TestGivePointsHandler(t *testing.T) {
	trList := []tester{
		{
			caseName: "everything ok",
			resultErr: nil,
			requestID: uuid.NewV1(),
			requestBody: []byte(`{"points": 1}`),
			expectedStatus: http.StatusOK,
		},
		{
			caseName: "wrong body",
			requestBody: []byte(`i'm the wrong body"`),
			requestID: uuid.Nil,
			expectedStatus: http.StatusBadRequest,
		},
		{
			caseName: "wrong user error",
			resultErr: errors.New("i'm the bad err"),
			requestID: uuid.NewV1(),
			requestBody: []byte(`{"points": 1}`),
			expectedStatus: http.StatusBadRequest,
		},
	}

	mo := mockedUser{}
	h := myhandler.Handler{}
	router.RouteForUser(&h, &mo)

	for _, tr := range trList {
		tr.path = router.UserPath + "/" + tr.requestID.String() + router.GivingPointsPath
		tr.method = http.MethodPost

		if tr.requestID != uuid.Nil{
			mo.On("FundUser", tr.requestID, 1).Return(tr.resultErr)
		}

		handleTester(t, &h, tr)
	}

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}
