package http

import (
	"errors"
	"net/http"
	"testing"
	httpHandler "tournament/pkg/controllers/handlers/http"
	"tournament/pkg/domain"
	"tournament/pkg/infrastructure/myhandler"
	"tournament/pkg/infrastructure/myuuid"
)

const (
	methodNameCreateUser = "CreateUser"
	methodNameGetUser    = "GetUser"
	methodNameDeleteUser = "DeleteUser"
	methodNameFundUser   = "FundUser"
)

// TestCreateUserHandler tests creation of user.
func TestCreateUserHandler(t *testing.T) {
	idFactory := myuuid.IDFactory{}
	testCases := []testCase{
		{
			caseName: "everything ok",
			resultID: idFactory.NewString(),
			requestUser: domain.User{
				Name: "Daniil Dankovskij",
			},
			requestBody:    `{"name": "Daniil Dankovskij"}`,
			expectedStatus: http.StatusCreated,
		},
		{
			caseName:       "wrong body",
			noMock:         true,
			requestBody:    `i'm the wrong body"`,
			requestUser:    domain.User{},
			expectedStatus: http.StatusBadRequest,
		},
		{
			caseName:  "wrong user error",
			resultErr: errors.New("i'm the bad err"),
			resultID:  idFactory.NewString(),
			requestUser: domain.User{
				Name: "Artemij Burah",
			},
			requestBody:    `{"name": "Artemij Burah"}`,
			expectedStatus: http.StatusBadRequest,
		},
	}

	idType := myuuid.IDType{}
	mo := mockedUsecases{}
	h := myhandler.Handler{}
	r := Router{IDType: idType}
	r.Route(&h, httpHandler.Controller{Usecases: &mo})

	for _, tc := range testCases {
		tc.path = httpHandler.UserPath
		tc.method = http.MethodPost

		if !tc.noMock {
			mo.On(methodNameCreateUser, tc.requestUser.Name).Return(tc.resultID, tc.resultErr)
		}

		handleTestCase(t, &h, tc)
	}

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}

// TestGetUserHandler tests getting of user's information.
func TestGetUserHandler(t *testing.T) {
	idFactory := myuuid.IDFactory{}
	testCases := []testCase{
		{
			caseName: "everything ok",
			resultUser: domain.User{
				Name: "Anna Angel",
			},
			requestID:      idFactory.NewString(),
			expectedStatus: http.StatusOK,
		},
		{
			caseName:       "wrong user error",
			resultErr:      errors.New("i'm the bad err"),
			resultUser:     domain.User{},
			requestID:      idFactory.NewString(),
			expectedStatus: http.StatusBadRequest,
		},
	}

	idType := myuuid.IDType{}
	mo := mockedUsecases{}
	h := myhandler.Handler{}
	r := Router{IDType: idType}
	r.Route(&h, httpHandler.Controller{Usecases: &mo})

	for _, tc := range testCases {
		tc.path = httpHandler.UserPath + "/" + tc.requestID
		tc.method = http.MethodGet

		if !tc.noMock {
			mo.On(methodNameGetUser, tc.requestID).Return(tc.resultUser, tc.resultErr)
		}

		handleTestCase(t, &h, tc)
	}

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}

// TestDeleteUserHandler tests deleting of user.
func TestDeleteUserHandler(t *testing.T) {
	idFactory := myuuid.IDFactory{}
	testCases := []testCase{
		{
			caseName:       "everything ok",
			requestID:      idFactory.NewString(),
			expectedStatus: http.StatusOK,
		},
		{
			caseName:       "wrong user error",
			resultErr:      errors.New("i'm the bad err"),
			requestID:      idFactory.NewString(),
			expectedStatus: http.StatusBadRequest,
		},
	}

	idType := myuuid.IDType{}
	mo := mockedUsecases{}
	h := myhandler.Handler{}
	r := Router{IDType: idType}
	r.Route(&h, httpHandler.Controller{Usecases: &mo})

	for _, tc := range testCases {
		tc.path = httpHandler.UserPath + "/" + tc.requestID
		tc.method = http.MethodDelete

		if !tc.noMock {
			mo.On(methodNameDeleteUser, tc.requestID).Return(tc.resultErr)
		}

		handleTestCase(t, &h, tc)
	}

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}

// TestTakePointsHandler tests taking points from user.
func TestTakePointsHandler(t *testing.T) {
	idFactory := myuuid.IDFactory{}
	testCases := []testCase{
		{
			caseName:       "everything ok",
			requestID:      idFactory.NewString(),
			requestBody:    `{"points": 1}`,
			expectedStatus: http.StatusOK,
		},
		{
			caseName:       "wrong body",
			noMock:         true,
			requestBody:    `i'm the wrong body"`,
			requestID:      idFactory.NewString(),
			expectedStatus: http.StatusBadRequest,
		},
		{
			caseName:       "wrong user error",
			resultErr:      errors.New("i'm the bad err"),
			requestID:      idFactory.NewString(),
			requestBody:    `{"points": 1}`,
			expectedStatus: http.StatusBadRequest,
		},
	}

	idType := myuuid.IDType{}
	mo := mockedUsecases{}
	h := myhandler.Handler{}
	r := Router{IDType: idType}
	r.Route(&h, httpHandler.Controller{Usecases: &mo})

	for _, tc := range testCases {
		tc.path = httpHandler.UserPath + "/" + tc.requestID + httpHandler.TakingPointsPath
		tc.method = http.MethodPost

		if !tc.noMock {
			mo.On(methodNameFundUser, tc.requestID, -1).Return(tc.resultErr)
		}

		handleTestCase(t, &h, tc)
	}

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}

// TestGivePointsHandler tests giving points to user.
func TestGivePointsHandler(t *testing.T) {
	idFactory := myuuid.IDFactory{}
	testCases := []testCase{
		{
			caseName:       "everything ok",
			requestID:      idFactory.NewString(),
			requestBody:    `{"points": 1}`,
			expectedStatus: http.StatusOK,
		},
		{
			caseName:       "wrong body",
			noMock:         true,
			requestBody:    `i'm the wrong body"`,
			requestID:      idFactory.NewString(),
			expectedStatus: http.StatusBadRequest,
		},
		{
			caseName:       "wrong user error",
			resultErr:      errors.New("i'm the bad err"),
			requestID:      idFactory.NewString(),
			requestBody:    `{"points": 1}`,
			expectedStatus: http.StatusBadRequest,
		},
	}

	idType := myuuid.IDType{}
	mo := mockedUsecases{}
	h := myhandler.Handler{}
	r := Router{IDType: idType}
	r.Route(&h, httpHandler.Controller{Usecases: &mo})

	for _, tc := range testCases {
		tc.path = httpHandler.UserPath + "/" + tc.requestID + httpHandler.GivingPointsPath
		tc.method = http.MethodPost

		if !tc.noMock {
			mo.On(methodNameFundUser, tc.requestID, 1).Return(tc.resultErr)
		}

		handleTestCase(t, &h, tc)
	}

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}
