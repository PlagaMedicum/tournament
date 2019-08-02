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
			RequestName: "Josip",
			ResultID:    1,
		},
	}

	mock := mocks.MockedRepository{}
	c := Controller{
		Repository: &mock,
	}

	name := "everything ok"
	tc := testCases[name]

	mock.On(methodNameInsertUser, tc.RequestName, defaultBalance).Return(tc.ResultID, tc.Err)

	gotTC := tc
	gotTC.ResultID, gotTC.Err = c.CreateUser(tc.RequestName)
	usecases.HandleTestCase(name, tc, gotTC, t)

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

	gotTC := tc
	gotTC.User, gotTC.Err = c.GetUser(tc.UserID)
	usecases.HandleTestCase(name, tc, gotTC, t)

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

	gotTC := tc
	gotTC.Err = c.DeleteUser(tc.UserID)
	usecases.HandleTestCase(name, tc, gotTC, t)

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}

func TestFundUser(t *testing.T) {
	testCases := map[string]usecases.TestCase{
		"everything ok": {
			UserID:        1,
			RequestPoints: 10,
			User: userDomain.User{
				ID:      1,
				Balance: 10,
			},
		},
		"get user error": {
			MockingStop:   1,
			UserID:        2,
			RequestPoints: 10,
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
			mock.On(methodNameUpdateUserBalanceByID, tc.User.ID, tc.User.Balance+tc.RequestPoints).Return(tc.Err)
			continue
		} 

		mock.On(methodNameUpdateUserBalanceByID, tc.User.ID, tc.User.Balance+tc.RequestPoints).Return(nil)
	}

	for name, tc := range testCases {
		gotTC := tc
		gotTC.Err = c.FundUser(tc.UserID, tc.RequestPoints)
		usecases.HandleTestCase(name, tc, gotTC, t)
	}

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}
