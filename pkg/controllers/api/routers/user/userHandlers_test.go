package user

import (
	"errors"
	"net/http"
	"strconv"
	"testing"
	httpHandlers "tournament/pkg/controllers/api/http_handlers"
	"tournament/pkg/controllers/api/http_handlers/user"
	"tournament/pkg/controllers/api/routers"
	"tournament/pkg/controllers/api/routers/user/mocks"
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
			ResID:    1,
			ReqUser: userDomain.User{
				Name: "Daniil Dankovskij",
			},
			ReqBody:   `{"name": "Daniil Dankovskij"}`,
			ResStatus: http.StatusCreated,
		},
		{
			CaseName:  "wrong body",
			NoMock:    true,
			ReqBody:   `i'm the wrong body"`,
			ReqUser:   userDomain.User{},
			ResStatus: http.StatusBadRequest,
		},
		{
			CaseName: "wrong user error",
			ResErr:   errors.New("i'm the bad err"),
			ResID:    3,
			ReqUser: userDomain.User{
				Name: "Artemij Burah",
			},
			ReqBody:   `{"name": "Artemij Burah"}`,
			ResStatus: http.StatusBadRequest,
		},
	}

	mock := mocks.MockedUsecases{}
	h := httpHandler.Handler{}
	RouteUser(&h, user.Controller{Usecases: &mock})

	for _, tc := range testCases {
		tc.Path = httpHandlers.UserPath
		tc.Method = http.MethodPost

		if !tc.NoMock {
			mock.On(methodNameCreateUser, tc.ReqUser.Name).Return(tc.ResID, tc.ResErr)
		}

		tc.Handle(t, &h)
	}

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}

// TestGetUserHandler tests getting of user's information.
func TestGetUserHandler(t *testing.T) {
	testCases := []routers.TestCase{
		{
			CaseName: "everything ok",
			ResUser: userDomain.User{
				Name: "Anna Angel",
			},
			ReqID:     1,
			ResStatus: http.StatusOK,
		},
		{
			CaseName:  "wrong user error",
			ResErr:    errors.New("i'm the bad err"),
			ResUser:   userDomain.User{},
			ReqID:     2,
			ResStatus: http.StatusBadRequest,
		},
	}

	mock := mocks.MockedUsecases{}
	h := httpHandler.Handler{}
	RouteUser(&h, user.Controller{Usecases: &mock})

	for _, tc := range testCases {
		tc.Path = httpHandlers.UserPath + "/" + strconv.FormatUint(tc.ReqID, 10)
		tc.Method = http.MethodGet

		if !tc.NoMock {
			mock.On(methodNameGetUser, tc.ReqID).Return(tc.ResUser, tc.ResErr)
		}

		tc.Handle(t, &h)
	}

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}

// TestDeleteUserHandler tests deleting of user.
func TestDeleteUserHandler(t *testing.T) {
	testCases := []routers.TestCase{
		{
			CaseName:  "everything ok",
			ReqID:     1,
			ResStatus: http.StatusOK,
		},
		{
			CaseName:  "wrong user error",
			ResErr:    errors.New("i'm the bad err"),
			ReqID:     2,
			ResStatus: http.StatusBadRequest,
		},
	}

	mock := mocks.MockedUsecases{}
	h := httpHandler.Handler{}
	RouteUser(&h, user.Controller{Usecases: &mock})

	for _, tc := range testCases {
		tc.Path = httpHandlers.UserPath + "/" + strconv.FormatUint(tc.ReqID, 10)
		tc.Method = http.MethodDelete

		if !tc.NoMock {
			mock.On(methodNameDeleteUser, tc.ReqID).Return(tc.ResErr)
		}

		tc.Handle(t, &h)
	}

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}

// TestTakePointsHandler tests taking points from user.
func TestTakePointsHandler(t *testing.T) {
	testCases := []routers.TestCase{
		{
			CaseName:  "everything ok",
			ReqID:     1,
			ReqBody:   `{"points": 1}`,
			ResStatus: http.StatusOK,
		},
		{
			CaseName:  "wrong body",
			NoMock:    true,
			ReqBody:   `i'm the wrong body"`,
			ReqID:     2,
			ResStatus: http.StatusBadRequest,
		},
		{
			CaseName:  "wrong user error",
			ResErr:    errors.New("i'm the bad err"),
			ReqID:     3,
			ReqBody:   `{"points": 1}`,
			ResStatus: http.StatusBadRequest,
		},
	}

	mock := mocks.MockedUsecases{}
	h := httpHandler.Handler{}
	RouteUser(&h, user.Controller{Usecases: &mock})

	for _, tc := range testCases {
		tc.Path = httpHandlers.UserPath + "/" + strconv.FormatUint(tc.ReqID, 10) + httpHandlers.TakingPointsPath
		tc.Method = http.MethodPost

		if !tc.NoMock {
			mock.On(methodNameFundUser, tc.ReqID, -1).Return(tc.ResErr)
		}

		tc.Handle(t, &h)
	}

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}

// TestGivePointsHandler tests giving points to user.
func TestGivePointsHandler(t *testing.T) {
	testCases := []routers.TestCase{
		{
			CaseName:  "everything ok",
			ReqID:     1,
			ReqBody:   `{"points": 1}`,
			ResStatus: http.StatusOK,
		},
		{
			CaseName:  "wrong body",
			NoMock:    true,
			ReqBody:   `i'm the wrong body"`,
			ReqID:     2,
			ResStatus: http.StatusBadRequest,
		},
		{
			CaseName:  "wrong user error",
			ResErr:    errors.New("i'm the bad err"),
			ReqID:     3,
			ReqBody:   `{"points": 1}`,
			ResStatus: http.StatusBadRequest,
		},
	}

	mock := mocks.MockedUsecases{}
	h := httpHandler.Handler{}
	RouteUser(&h, user.Controller{Usecases: &mock})

	for _, tc := range testCases {
		tc.Path = httpHandlers.UserPath + "/" + strconv.FormatUint(tc.ReqID, 10) + httpHandlers.GivingPointsPath
		tc.Method = http.MethodPost

		if !tc.NoMock {
			mock.On(methodNameFundUser, tc.ReqID, 1).Return(tc.ResErr)
		}

		tc.Handle(t, &h)
	}

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}
