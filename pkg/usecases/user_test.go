package usecases

import (
	"errors"
	"testing"
	"tournament/pkg/domain"
	"tournament/pkg/infrastructure/myuuid"
)

const (
	methodNameInsertUser            = "InsertUser"
	methodNameGetUserByID           = "GetUserByID"
	methodNameGetUsers				= "GetUsers"
	methodNameDeleteUserByID        = "DeleteUserByID"
	methodNameUpdateUserBalanceByID = "UpdateUserBalanceByID"
)

func TestCreateUser(t *testing.T) {
	factory := myuuid.IDFactory{}
	testCases := map[string]testCase{
		"everything ok": {
			requestName: "Josip",
			resultID:    factory.NewString(),
		},
	}

	mo := mockedRepository{}
	c := Controller{
		Repository: &mo,
		IDType:     myuuid.IDType{},
	}

	name := "everything ok"
	tc := testCases[name]

	mo.On(methodNameInsertUser, tc.requestName, defaultBalance).Return(tc.resultID, tc.err)

	gotTC := tc
	gotTC.resultID, gotTC.err = c.CreateUser(tc.requestName)
	handleTestCase(name, tc, gotTC, t)

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}

func TestGetUser(t *testing.T) {
	factory := myuuid.IDFactory{}
	testCases := map[string]testCase{
		"everything ok": {
			userID: factory.NewString(),
			user:   domain.User{ID: factory.NewString()},
		},
	}

	mo := mockedRepository{}
	c := Controller{
		Repository: &mo,
		IDType:     myuuid.IDType{},
	}

	name := "everything ok"
	tc := testCases[name]

	mo.On(methodNameGetUserByID, tc.userID).Return(tc.user, tc.err)

	gotTC := tc
	gotTC.user, gotTC.err = c.GetUser(tc.userID)
	handleTestCase(name, tc, gotTC, t)

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}

func TestDeleteUser(t *testing.T) {
	factory := myuuid.IDFactory{}
	testCases := map[string]testCase{
		"everything ok": {
			userID: factory.NewString(),
		},
	}

	mo := mockedRepository{}
	c := Controller{
		Repository: &mo,
		IDType:     myuuid.IDType{},
	}

	name := "everything ok"
	tc := testCases[name]

	mo.On(methodNameDeleteUserByID, tc.userID).Return(tc.err)

	gotTC := tc
	gotTC.err = c.DeleteUser(tc.userID)
	handleTestCase(name, tc, gotTC, t)

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}

func TestFundUser(t *testing.T) {
	factory := myuuid.IDFactory{}
	testCases := map[string]testCase{
		"everything ok": {
			userID:        factory.NewString(),
			requestPoints: 10,
			user: domain.User{
				ID:      factory.NewString(),
				Balance: 10,
			},
		},
		"get user error": {
			mockingStop:   1,
			userID:        factory.NewString(),
			requestPoints: 10,
			user: domain.User{
				ID:      factory.NewString(),
				Balance: 10,
			},
			err: errors.New("get user error"),
		},
	}

	mo := mockedRepository{}
	c := Controller{
		Repository: &mo,
		IDType:     myuuid.IDType{},
	}

	for _, tc := range testCases {
		if tc.mockingStop == 1 {
			mo.On(methodNameGetUserByID, tc.userID).Return(tc.user, tc.err)
			continue
		} else {
			mo.On(methodNameGetUserByID, tc.userID).Return(tc.user, nil)
		}

		if tc.mockingStop == 2 {
			mo.On(methodNameUpdateUserBalanceByID, tc.user.ID, tc.user.Balance+tc.requestPoints).Return(tc.err)
			continue
		} else {
			mo.On(methodNameUpdateUserBalanceByID, tc.user.ID, tc.user.Balance+tc.requestPoints).Return(nil)
		}
	}

	for name, tc := range testCases {
		gotTC := tc
		gotTC.err = c.FundUser(tc.userID, tc.requestPoints)
		handleTestCase(name, tc, gotTC, t)
	}

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}
