package http

import (
	"errors"
	"net/http"
	"strconv"
	"testing"
	httpHandler "tournament/pkg/controllers/handlers/http"
	"tournament/pkg/controllers/routers/http/mocks"
	"tournament/pkg/domain"
	"tournament/pkg/infrastructure/myhandler"
)

const (
	methodNameCreateUser = "CreateUser"
	methodNameGetUser    = "GetUser"
	methodNameDeleteUser = "DeleteUser"
	methodNameFundUser   = "FundUser"
)

// TestCreateUserHandler tests creation of user.
func TestCreateUserHandler(t *testing.T) {
	testCases := []testCase{
		{
			caseName: "everything ok",
			resultID: 1,
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
			resultID:  3,
			requestUser: domain.User{
				Name: "Artemij Burah",
			},
			requestBody:    `{"name": "Artemij Burah"}`,
			expectedStatus: http.StatusBadRequest,
		},
	}

	mock := mocks.MockedUsecases{}
	h := myhandler.Handler{}
	Route(&h, httpHandler.Controller{Usecases: &mock})

	for _, tc := range testCases {
		tc.path = httpHandler.UserPath
		tc.method = http.MethodPost

		if !tc.noMock {
			mock.On(methodNameCreateUser, tc.requestUser.Name).Return(tc.resultID, tc.resultErr)
		}

		handleTestCase(t, &h, tc)
	}

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}

// TestGetUserHandler tests getting of user's information.
func TestGetUserHandler(t *testing.T) {
	testCases := []testCase{
		{
			caseName: "everything ok",
			resultUser: domain.User{
				Name: "Anna Angel",
			},
			requestID:      1,
			expectedStatus: http.StatusOK,
		},
		{
			caseName:       "wrong user error",
			resultErr:      errors.New("i'm the bad err"),
			resultUser:     domain.User{},
			requestID:      2,
			expectedStatus: http.StatusBadRequest,
		},
	}

	mock := mocks.MockedUsecases{}
	h := myhandler.Handler{}
	Route(&h, httpHandler.Controller{Usecases: &mock})

	for _, tc := range testCases {
		tc.path = httpHandler.UserPath + "/" + strconv.FormatUint(tc.requestID, 10)
		tc.method = http.MethodGet

		if !tc.noMock {
			mock.On(methodNameGetUser, tc.requestID).Return(tc.resultUser, tc.resultErr)
		}

		handleTestCase(t, &h, tc)
	}

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}

// TestDeleteUserHandler tests deleting of user.
func TestDeleteUserHandler(t *testing.T) {
	testCases := []testCase{
		{
			caseName:       "everything ok",
			requestID:      1,
			expectedStatus: http.StatusOK,
		},
		{
			caseName:       "wrong user error",
			resultErr:      errors.New("i'm the bad err"),
			requestID:      2,
			expectedStatus: http.StatusBadRequest,
		},
	}

	mock := mocks.MockedUsecases{}
	h := myhandler.Handler{}
	Route(&h, httpHandler.Controller{Usecases: &mock})

	for _, tc := range testCases {
		tc.path = httpHandler.UserPath + "/" + strconv.FormatUint(tc.requestID,10)
		tc.method = http.MethodDelete

		if !tc.noMock {
			mock.On(methodNameDeleteUser, tc.requestID).Return(tc.resultErr)
		}

		handleTestCase(t, &h, tc)
	}

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}

// TestTakePointsHandler tests taking points from user.
func TestTakePointsHandler(t *testing.T) {
	testCases := []testCase{
		{
			caseName:       "everything ok",
			requestID:      1,
			requestBody:    `{"points": 1}`,
			expectedStatus: http.StatusOK,
		},
		{
			caseName:       "wrong body",
			noMock:         true,
			requestBody:    `i'm the wrong body"`,
			requestID:      2,
			expectedStatus: http.StatusBadRequest,
		},
		{
			caseName:       "wrong user error",
			resultErr:      errors.New("i'm the bad err"),
			requestID:      3,
			requestBody:    `{"points": 1}`,
			expectedStatus: http.StatusBadRequest,
		},
	}

	mock := mocks.MockedUsecases{}
	h := myhandler.Handler{}
	Route(&h, httpHandler.Controller{Usecases: &mock})

	for _, tc := range testCases {
		tc.path = httpHandler.UserPath + "/" + strconv.FormatUint(tc.requestID,10) + httpHandler.TakingPointsPath
		tc.method = http.MethodPost

		if !tc.noMock {
			mock.On(methodNameFundUser, tc.requestID, -1).Return(tc.resultErr)
		}

		handleTestCase(t, &h, tc)
	}

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}

// TestGivePointsHandler tests giving points to user.
func TestGivePointsHandler(t *testing.T) {
	testCases := []testCase{
		{
			caseName:       "everything ok",
			requestID:      1,
			requestBody:    `{"points": 1}`,
			expectedStatus: http.StatusOK,
		},
		{
			caseName:       "wrong body",
			noMock:         true,
			requestBody:    `i'm the wrong body"`,
			requestID:      2,
			expectedStatus: http.StatusBadRequest,
		},
		{
			caseName:       "wrong user error",
			resultErr:      errors.New("i'm the bad err"),
			requestID:      3,
			requestBody:    `{"points": 1}`,
			expectedStatus: http.StatusBadRequest,
		},
	}

	mock := mocks.MockedUsecases{}
	h := myhandler.Handler{}
	Route(&h, httpHandler.Controller{Usecases: &mock})

	for _, tc := range testCases {
		tc.path = httpHandler.UserPath + "/" + strconv.FormatUint(tc.requestID, 10) + httpHandler.GivingPointsPath
		tc.method = http.MethodPost

		if !tc.noMock {
			mock.On(methodNameFundUser, tc.requestID, 1).Return(tc.resultErr)
		}

		handleTestCase(t, &h, tc)
	}

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}
