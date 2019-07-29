package usecases

import (
	"errors"
	"testing"
	"tournament/pkg/domain"
	"tournament/pkg/infrastructure/myuuid"
	"tournament/pkg/usecases"
)

const (
	methodNameInsertTournament     = "InsertTournament"
	methodNameGetTournamentByID    = "GetTournamentByID"
	methodNameDeleteTournamentByID = "DeleteTournamentByID"
	methodNameAddUserInTournament  = "AddUserInTournament"
)

func TestCreateTournament(t *testing.T) {
	factory := myuuid.IDFactory{}
	testCases := []testCase{
		{
			caseName:       "everything ok",
			requestName:    "Tournament",
			requestDeposit: 500,
			resultID:       factory.NewString(),
			resultErr:      nil,
		},
	}

	mo := mockedRepository{}
	c := usecases.Controller{
		Repository: &mo,
		IDType:     myuuid.IDType{},
	}

	for _, tc := range testCases {
		if tc.mockingStop == 1 {
			mo.On(methodNameInsertTournament, tc.requestName, tc.requestDeposit, 4000).Return(tc.resultID, tc.resultErr)
			continue
		} else {
			mo.On(methodNameInsertTournament, tc.requestName, tc.requestDeposit, 4000).Return(tc.resultID, nil)
		}
	}

	for _, tc := range testCases {
		gotTC := tc
		gotTC.resultID, gotTC.resultErr = c.CreateTournament(tc.requestName, tc.requestDeposit)
		handleTestCase(tc, gotTC, t)
	}

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}

func TestGetTournament(t *testing.T) {
	factory := myuuid.IDFactory{}
	testCases := []testCase{
		{
			caseName:            "everything ok",
			requestTournamentID: factory.NewString(),
			resultTournament:    domain.Tournament{ID: factory.NewString()},
			resultErr:           nil,
		},
	}

	mo := mockedRepository{}
	c := usecases.Controller{
		Repository: &mo,
		IDType:     myuuid.IDType{},
	}

	for _, tc := range testCases {
		if tc.mockingStop == 1 {
			mo.On(methodNameGetTournamentByID, tc.requestTournamentID).Return(tc.resultTournament, tc.resultErr)
			continue
		} else {
			mo.On(methodNameGetTournamentByID, tc.requestTournamentID).Return(tc.resultTournament, nil)
		}
	}

	for _, tc := range testCases {
		gotTC := tc
		gotTC.resultTournament, gotTC.resultErr = c.GetTournament(tc.requestTournamentID)
		handleTestCase(tc, gotTC, t)
	}

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}

func TestDeleteTournament(t *testing.T) {
	factory := myuuid.IDFactory{}
	testCases := []testCase{
		{
			caseName:            "everything ok",
			requestTournamentID: factory.NewString(),
			resultErr:           nil,
		},
	}

	mo := mockedRepository{}
	c := usecases.Controller{
		Repository: &mo,
		IDType:     myuuid.IDType{},
	}

	for _, tc := range testCases {
		if tc.mockingStop == 1 {
			mo.On(methodNameDeleteTournamentByID, tc.requestTournamentID).Return(tc.resultErr)
			continue
		} else {
			mo.On(methodNameDeleteTournamentByID, tc.requestTournamentID).Return(nil)
		}
	}

	for _, tc := range testCases {
		gotTC := tc
		gotTC.resultErr = c.DeleteTournament(tc.requestTournamentID)
		handleTestCase(tc, gotTC, t)
	}

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}

func TestJoinTournament(t *testing.T) {
	factory := myuuid.IDFactory{}
	testCases := []testCase{
		{
			caseName:            "everything ok",
			requestTournamentID: factory.NewString(),
			requestUserID:       factory.NewString(),
			resultTournament:    domain.Tournament{
				ID: factory.NewString(),
				Deposit: 100,
			},
			resultUser:          domain.User{
				ID: factory.NewString(),
				Balance: 120,
			},
			resultErr:           nil,
		},
		{
			caseName:            "error getting tournament",
			mockingStop:         1,
			requestTournamentID: factory.NewString(),
			requestUserID:       factory.NewString(),
			resultTournament:    domain.Tournament{},
			resultErr:           errors.New("error"),
		},
		{
			caseName:            "error getting user",
			mockingStop:         2,
			requestTournamentID: factory.NewString(),
			requestUserID:       factory.NewString(),
			resultTournament:    domain.Tournament{},
			resultUser:          domain.User{},
			resultErr:           errors.New("error"),
		},
	}

	mo := mockedRepository{}
	c := usecases.Controller{
		Repository: &mo,
		IDType:     myuuid.IDType{},
	}

	for _, tc := range testCases {
		if tc.mockingStop == 1 {
			mo.On(methodNameGetTournamentByID, tc.requestTournamentID).Return(tc.resultTournament, tc.resultErr)
			continue
		} else {
			mo.On(methodNameGetTournamentByID, tc.requestTournamentID).Return(tc.resultTournament, nil)
		}

		if tc.mockingStop == 2 {
			mo.On(methodNameGetUserByID, tc.requestUserID).Return(tc.resultUser, tc.resultErr)
			continue
		} else {
			mo.On(methodNameGetUserByID, tc.requestUserID).Return(tc.resultUser, nil)
		}

		if tc.mockingStop == 3 {
			mo.On(methodNameAddUserInTournament, tc.resultUser.ID, tc.resultTournament.ID).Return(tc.resultErr)
			continue
		} else {
			mo.On(methodNameAddUserInTournament, tc.resultUser.ID, tc.resultTournament.ID).Return(nil)
		}

		if tc.mockingStop == 4 {
			mo.On(methodNameUpdateUserBalanceByID, tc.resultUser.ID, tc.resultUser.Balance - tc.resultTournament.Deposit).Return(tc.resultErr)
			continue
		} else {
			mo.On(methodNameUpdateUserBalanceByID, tc.resultUser.ID, tc.resultUser.Balance - tc.resultTournament.Deposit).Return(nil)

		}
	}

	for _, tc := range testCases {
		gotTC := tc
		gotTC.resultErr = c.JoinTournament(tc.requestTournamentID, tc.requestUserID)
		handleTestCase(tc, gotTC, t)
	}

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}

func TestFinishTournament(t *testing.T) {

}
