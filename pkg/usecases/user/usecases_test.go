package user

import (
	"errors"
	"testing"
	userDomain "tournament/pkg/domain/user"
	"tournament/pkg/usecases"
	"tournament/pkg/usecases/user/mocks"
)

const (
	methodNameInsertUser            = "InsertUser"
	methodNameGetUserByID           = "GetUserByID"
	methodNameDeleteUserByID        = "DeleteUserByID"
	methodNameUpdateUserBalanceByID = "UpdateUserBalanceByID"
)

func TestCreateUser(t *testing.T) {
	testCases := map[string]usecases.TestCase{
		"everything ok": {
			ReqName: "Josip",
			ResID:   1,
		},
	}

	mock := mocks.MockedRepository{}
	c := Controller{
		Repository: &mock,
	}

	name := "everything ok"
	tc := testCases[name]

	mock.On(methodNameInsertUser, tc.ReqName, defaultBalance).Return(tc.ResID, tc.Err)

	got := tc
	got.ResID, got.Err = c.CreateUser(tc.ReqName)
	tc.Handle(name, got, t)

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}

func TestGetUser(t *testing.T) {
	testCases := map[string]usecases.TestCase{
		"everything ok": {
			UserID: 1,
			User:   userDomain.User{ID: 1},
		},
	}

	mock := mocks.MockedRepository{}
	c := Controller{
		Repository: &mock,
	}

	name := "everything ok"
	tc := testCases[name]

	mock.On(methodNameGetUserByID, tc.UserID).Return(tc.User, tc.Err)

	got := tc
	got.User, got.Err = c.GetUser(tc.UserID)
	tc.Handle(name, got, t)

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}

func TestDeleteUser(t *testing.T) {
	testCases := map[string]usecases.TestCase{
		"everything ok": {
			UserID: 1,
		},
	}

	mock := mocks.MockedRepository{}
	c := Controller{
		Repository: &mock,
	}

	name := "everything ok"
	tc := testCases[name]

	mock.On(methodNameDeleteUserByID, tc.UserID).Return(tc.Err)

	got := tc
	got.Err = c.DeleteUser(tc.UserID)
	tc.Handle(name, got, t)

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}

func TestFundUser(t *testing.T) {
	testCases := map[string]usecases.TestCase{
		"everything ok": {
			UserID:    1,
			ReqPoints: 10,
			User: userDomain.User{
				ID:      1,
				Balance: 10,
			},
		},
		"get user error": {
			MockingStop: 1,
			UserID:      2,
			ReqPoints:   10,
			User: userDomain.User{
				ID:      2,
				Balance: 10,
			},
			Err: errors.New("get user error"),
		},
	}

	mock := mocks.MockedRepository{}
	c := Controller{
		Repository: &mock,
	}

	for _, tc := range testCases {
		if tc.MockingStop == 1 {
			mock.On(methodNameGetUserByID, tc.UserID).Return(tc.User, tc.Err)
			continue
		}

		mock.On(methodNameGetUserByID, tc.UserID).Return(tc.User, nil)

		if tc.MockingStop == 2 {
			mock.On(methodNameUpdateUserBalanceByID, tc.User.ID, tc.User.Balance+tc.ReqPoints).Return(tc.Err)
			continue
		}

		mock.On(methodNameUpdateUserBalanceByID, tc.User.ID, tc.User.Balance+tc.ReqPoints).Return(nil)
	}

	for name, tc := range testCases {
		got := tc
		got.Err = c.FundUser(tc.UserID, tc.ReqPoints)
		tc.Handle(name, got, t)
	}

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}
