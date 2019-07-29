package usecases

import (
	"errors"
	"testing"
	"tournament/pkg/domain"
	"tournament/pkg/infrastructure/myuuid"
	"tournament/pkg/usecases"
)

const (
	methodNameInsertUser            = "InsertUser"
	methodNameGetUserByID           = "GetUserByID"
	methodNameDeleteUserByID        = "DeleteUserByID"
	methodNameUpdateUserBalanceByID = "UpdateUserBalanceByID"
)

func TestCreateUser(t *testing.T) {
	factory := myuuid.IDFactory{}
	testCases := []testCase{
		{
			caseName:    "everything ok",
			requestName: "Josip",
			resultID:    factory.NewString(),
			resultErr:   nil,
		},
	}

	mo := mockedRepository{}
	c := usecases.Controller{
		Repository: &mo,
		IDType:     myuuid.IDType{},
	}

	for _, tc := range testCases {
		if tc.mockingStop == 1 {
			mo.On(methodNameInsertUser, tc.requestName, 700).Return(tc.resultID, tc.resultErr)
			continue
		} else {
			mo.On(methodNameInsertUser, tc.requestName, 700).Return(tc.resultID, nil)
		}
	}

	for _, tc := range testCases {
		gotTC := tc
		gotTC.resultID, gotTC.resultErr = c.CreateUser(tc.requestName)
		handleTestCase(tc, gotTC, t)
	}

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}

func TestGetUser(t *testing.T) {
	factory := myuuid.IDFactory{}
	testCases := []testCase{
		{
			caseName:      "everything ok",
			requestUserID: factory.NewString(),
			resultUser:    domain.User{ID: factory.NewString()},
			resultErr:     nil,
		},
	}

	mo := mockedRepository{}
	c := usecases.Controller{
		Repository: &mo,
		IDType:     myuuid.IDType{},
	}

	for _, tc := range testCases {
		if tc.mockingStop == 1 {
			mo.On(methodNameGetUserByID, tc.requestUserID).Return(tc.resultUser, tc.resultErr)
			continue
		} else {
			mo.On(methodNameGetUserByID, tc.requestUserID).Return(tc.resultUser, nil)
		}
	}

	for _, tc := range testCases {
		gotTC := tc
		gotTC.resultUser, gotTC.resultErr = c.GetUser(tc.requestUserID)
		handleTestCase(tc, gotTC, t)
	}

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}

func TestDeleteUser(t *testing.T) {
	factory := myuuid.IDFactory{}
	testCases := []testCase{
		{
			caseName:      "everything ok",
			requestUserID: factory.NewString(),
			resultErr:     nil,
		},
	}

	mo := mockedRepository{}
	c := usecases.Controller{
		Repository: &mo,
		IDType:     myuuid.IDType{},
	}

	for _, tc := range testCases {
		if tc.mockingStop == 1 {
			mo.On(methodNameDeleteUserByID, tc.requestUserID).Return(tc.resultErr)
			continue
		} else {
			mo.On(methodNameDeleteUserByID, tc.requestUserID).Return(nil)
		}
	}

	for _, tc := range testCases {
		gotTC := tc
		gotTC.resultErr = c.DeleteUser(tc.requestUserID)
		handleTestCase(tc, gotTC, t)
	}

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}

func TestFundUser(t *testing.T) {
	factory := myuuid.IDFactory{}
	testCases := []testCase{
		{
			caseName:      "everything ok",
			requestUserID: factory.NewString(),
			requestPoints: 10,
			resultUser: domain.User{
				ID:      factory.NewString(),
				Balance: 10,
			},
			resultErr: nil,
		},
		{
			caseName:      "get user error",
			mockingStop:   1,
			requestUserID: factory.NewString(),
			requestPoints: 10,
			resultUser: domain.User{
				ID:      factory.NewString(),
				Balance: 10,
			},
			resultErr: errors.New("get user error"),
		},
	}

	mo := mockedRepository{}
	c := usecases.Controller{
		Repository: &mo,
		IDType:     myuuid.IDType{},
	}

	for _, tc := range testCases {
		if tc.mockingStop == 1 {
			mo.On(methodNameGetUserByID, tc.requestUserID).Return(tc.resultUser, tc.resultErr)
			continue
		} else {
			mo.On(methodNameGetUserByID, tc.requestUserID).Return(tc.resultUser, nil)
		}

		if tc.mockingStop == 2 {
			mo.On(methodNameUpdateUserBalanceByID, tc.resultUser.ID, tc.resultUser.Balance+tc.requestPoints).Return(tc.resultErr)
			continue
		} else {
			mo.On(methodNameUpdateUserBalanceByID, tc.resultUser.ID, tc.resultUser.Balance+tc.requestPoints).Return(nil)
		}
	}

	for _, tc := range testCases {
		gotTC := tc
		gotTC.resultErr = c.FundUser(tc.requestUserID, tc.requestPoints)
		handleTestCase(tc, gotTC, t)
	}

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}
