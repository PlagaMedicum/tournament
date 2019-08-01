package usecases

import (
	"errors"
	"testing"
	"tournament/pkg/domain"
	"tournament/pkg/usecases/mocks"
)

const (
	methodNameInsertUser            = "InsertUser"
	methodNameGetUserByID           = "GetUserByID"
	methodNameGetUsers				= "GetUsers"
	methodNameDeleteUserByID        = "DeleteUserByID"
	methodNameUpdateUserBalanceByID = "UpdateUserBalanceByID"
)

func TestCreateUser(t *testing.T) {
	testCases := map[string]testCase{
		"everything ok": {
			requestName: "Josip",
			resultID:    1,
		},
	}

	mock := mocks.MockedRepository{}
	c := Controller{
		Repository: &mock,
	}

	name := "everything ok"
	tc := testCases[name]

	mock.On(methodNameInsertUser, tc.requestName, defaultBalance).Return(tc.resultID, tc.err)

	gotTC := tc
	gotTC.resultID, gotTC.err = c.CreateUser(tc.requestName)
	handleTestCase(name, tc, gotTC, t)

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}

func TestGetUser(t *testing.T) {
	testCases := map[string]testCase{
		"everything ok": {
			userID: 1,
			user:   domain.User{ID: 1},
		},
	}

	mock := mocks.MockedRepository{}
	c := Controller{
		Repository: &mock,
	}

	name := "everything ok"
	tc := testCases[name]

	mock.On(methodNameGetUserByID, tc.userID).Return(tc.user, tc.err)

	gotTC := tc
	gotTC.user, gotTC.err = c.GetUser(tc.userID)
	handleTestCase(name, tc, gotTC, t)

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}

func TestDeleteUser(t *testing.T) {
	testCases := map[string]testCase{
		"everything ok": {
			userID: 1,
		},
	}

	mock := mocks.MockedRepository{}
	c := Controller{
		Repository: &mock,
	}

	name := "everything ok"
	tc := testCases[name]

	mock.On(methodNameDeleteUserByID, tc.userID).Return(tc.err)

	gotTC := tc
	gotTC.err = c.DeleteUser(tc.userID)
	handleTestCase(name, tc, gotTC, t)

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}

func TestFundUser(t *testing.T) {
	testCases := map[string]testCase{
		"everything ok": {
			userID:        1,
			requestPoints: 10,
			user: domain.User{
				ID:      1,
				Balance: 10,
			},
		},
		"get user error": {
			mockingStop:   1,
			userID:        2,
			requestPoints: 10,
			user: domain.User{
				ID:      2,
				Balance: 10,
			},
			err: errors.New("get user error"),
		},
	}

	mock := mocks.MockedRepository{}
	c := Controller{
		Repository: &mock,
	}

	for _, tc := range testCases {
		if tc.mockingStop == 1 {
			mock.On(methodNameGetUserByID, tc.userID).Return(tc.user, tc.err)
			continue
		} 

		mock.On(methodNameGetUserByID, tc.userID).Return(tc.user, nil)

		if tc.mockingStop == 2 {
			mock.On(methodNameUpdateUserBalanceByID, tc.user.ID, tc.user.Balance+tc.requestPoints).Return(tc.err)
			continue
		} 

		mock.On(methodNameUpdateUserBalanceByID, tc.user.ID, tc.user.Balance+tc.requestPoints).Return(nil)
	}

	for name, tc := range testCases {
		gotTC := tc
		gotTC.err = c.FundUser(tc.userID, tc.requestPoints)
		handleTestCase(name, tc, gotTC, t)
	}

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}
