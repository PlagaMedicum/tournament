package user

import (
	"errors"
	"net/http"
	"strconv"
	"testing"
	httpHandlers "tournament/pkg/controllers/api/http_handlers"
	"tournament/pkg/controllers/api/http_handlers/user"
	"tournament/pkg/controllers/api/routers"
	"tournament/pkg/controllers/api/routers/mocks"
	userDomain "tournament/pkg/domain/user"
	httpHandler "tournament/pkg/infrastructure/handler"
)

const (
	methodNameCreateUser = "CreateUser"
	methodNameGetUser    = "GetUser"
	methodNameDeleteUser = "DeleteUser"
	methodNameFundUser   = "FundUser"
)

// TestCreateUserHandler tests creation of user.
func TestCreateUserHandler(t *testing.T) {
	testCases := []routers.TestCase{
		{
			CaseName: "everything ok",
			ResultID: 1,
			RequestUser: userDomain.User{
				Name: "Daniil Dankovskij",
			},
			RequestBody:    `{"name": "Daniil Dankovskij"}`,
			ExpectedStatus: http.StatusCreated,
		},
		{
			CaseName:       "wrong body",
			NoMock:         true,
			RequestBody:    `i'm the wrong body"`,
			RequestUser:    userDomain.User{},
			ExpectedStatus: http.StatusBadRequest,
		},
		{
			CaseName:  "wrong user error",
			ResultErr: errors.New("i'm the bad err"),
			ResultID:  3,
			RequestUser: userDomain.User{
				Name: "Artemij Burah",
			},
			RequestBody:    `{"name": "Artemij Burah"}`,
			ExpectedStatus: http.StatusBadRequest,
		},
	}

	mock := mocks.MockedUsecases{}
	h := httpHandler.Handler{}
	RouteUser(&h, user.Controller{Usecases: &mock})

	for _, tc := range testCases {
		tc.Path = httpHandlers.UserPath
		tc.Method = http.MethodPost

		if !tc.NoMock {
			mock.On(methodNameCreateUser, tc.RequestUser.Name).Return(tc.ResultID, tc.ResultErr)
		}

		routers.HandleTestCase(t, &h, tc)
	}

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}

// TestGetUserHandler tests getting of user's information.
func TestGetUserHandler(t *testing.T) {
	testCases := []routers.TestCase{
		{
			CaseName: "everything ok",
			ResultUser: userDomain.User{
				Name: "Anna Angel",
			},
			RequestID:      1,
			ExpectedStatus: http.StatusOK,
		},
		{
			CaseName:       "wrong user error",
			ResultErr:      errors.New("i'm the bad err"),
			ResultUser:     userDomain.User{},
			RequestID:      2,
			ExpectedStatus: http.StatusBadRequest,
		},
	}

	mock := mocks.MockedUsecases{}
	h := httpHandler.Handler{}
	RouteUser(&h, user.Controller{Usecases: &mock})

	for _, tc := range testCases {
		tc.Path = httpHandlers.UserPath + "/" + strconv.FormatUint(tc.RequestID, 10)
		tc.Method = http.MethodGet

		if !tc.NoMock {
			mock.On(methodNameGetUser, tc.RequestID).Return(tc.ResultUser, tc.ResultErr)
		}

		routers.HandleTestCase(t, &h, tc)
	}

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}

// TestDeleteUserHandler tests deleting of user.
func TestDeleteUserHandler(t *testing.T) {
	testCases := []routers.TestCase{
		{
			CaseName:       "everything ok",
			RequestID:      1,
			ExpectedStatus: http.StatusOK,
		},
		{
			CaseName:       "wrong user error",
			ResultErr:      errors.New("i'm the bad err"),
			RequestID:      2,
			ExpectedStatus: http.StatusBadRequest,
		},
	}

	mock := mocks.MockedUsecases{}
	h := httpHandler.Handler{}
	RouteUser(&h, user.Controller{Usecases: &mock})

	for _, tc := range testCases {
		tc.Path = httpHandlers.UserPath + "/" + strconv.FormatUint(tc.RequestID,10)
		tc.Method = http.MethodDelete

		if !tc.NoMock {
			mock.On(methodNameDeleteUser, tc.RequestID).Return(tc.ResultErr)
		}

		routers.HandleTestCase(t, &h, tc)
	}

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}

// TestTakePointsHandler tests taking points from user.
func TestTakePointsHandler(t *testing.T) {
	testCases := []routers.TestCase{
		{
			CaseName:       "everything ok",
			RequestID:      1,
			RequestBody:    `{"points": 1}`,
			ExpectedStatus: http.StatusOK,
		},
		{
			CaseName:       "wrong body",
			NoMock:         true,
			RequestBody:    `i'm the wrong body"`,
			RequestID:      2,
			ExpectedStatus: http.StatusBadRequest,
		},
		{
			CaseName:       "wrong user error",
			ResultErr:      errors.New("i'm the bad err"),
			RequestID:      3,
			RequestBody:    `{"points": 1}`,
			ExpectedStatus: http.StatusBadRequest,
		},
	}

	mock := mocks.MockedUsecases{}
	h := httpHandler.Handler{}
	RouteUser(&h, user.Controller{Usecases: &mock})

	for _, tc := range testCases {
		tc.Path = httpHandlers.UserPath + "/" + strconv.FormatUint(tc.RequestID,10) + httpHandlers.TakingPointsPath
		tc.Method = http.MethodPost

		if !tc.NoMock {
			mock.On(methodNameFundUser, tc.RequestID, -1).Return(tc.ResultErr)
		}

		routers.HandleTestCase(t, &h, tc)
	}

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}

// TestGivePointsHandler tests giving points to user.
func TestGivePointsHandler(t *testing.T) {
	testCases := []routers.TestCase{
		{
			CaseName:       "everything ok",
			RequestID:      1,
			RequestBody:    `{"points": 1}`,
			ExpectedStatus: http.StatusOK,
		},
		{
			CaseName:       "wrong body",
			NoMock:         true,
			RequestBody:    `i'm the wrong body"`,
			RequestID:      2,
			ExpectedStatus: http.StatusBadRequest,
		},
		{
			CaseName:       "wrong user error",
			ResultErr:      errors.New("i'm the bad err"),
			RequestID:      3,
			RequestBody:    `{"points": 1}`,
			ExpectedStatus: http.StatusBadRequest,
		},
	}

	mock := mocks.MockedUsecases{}
	h := httpHandler.Handler{}
	RouteUser(&h, user.Controller{Usecases: &mock})

	for _, tc := range testCases {
		tc.Path = httpHandlers.UserPath + "/" + strconv.FormatUint(tc.RequestID, 10) + httpHandlers.GivingPointsPath
		tc.Method = http.MethodPost

		if !tc.NoMock {
			mock.On(methodNameFundUser, tc.RequestID, 1).Return(tc.ResultErr)
		}

		routers.HandleTestCase(t, &h, tc)
	}

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}
