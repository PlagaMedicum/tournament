package usecases

import (
	"errors"
	"testing"
	"tournament/pkg/domain"
	"tournament/pkg/infrastructure/myuuid"
)

const (
	methodNameInsertTournament     = "InsertTournament"
	methodNameGetTournamentByID    = "GetTournamentByID"
	methodNameDeleteTournamentByID = "DeleteTournamentByID"
	methodNameAddUserInTournament  = "AddUserInTournament"
	methodNameSetWinner            = "SetWinner"
)

func TestCreateTournament(t *testing.T) {
	factory := myuuid.IDFactory{}
	testCases := map[string]testCase{
		"everything ok": {
			requestName:    "Tournament",
			requestDeposit: 500,
			resultID:       factory.NewString(),
		},
	}

	mo := mockedRepository{}
	c := Controller{
		Repository: &mo,
		IDType:     myuuid.IDType{},
	}

	name := "everything ok"
	tc := testCases[name]

	mo.On(methodNameInsertTournament, tc.requestName, tc.requestDeposit, defaultPrize).Return(tc.resultID, tc.err)

	gotTC := tc
	gotTC.resultID, gotTC.err = c.CreateTournament(tc.requestName, tc.requestDeposit)
	handleTestCase(name, tc, gotTC, t)

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}

func TestGetTournament(t *testing.T) {
	factory := myuuid.IDFactory{}
	testCases := map[string]testCase{
		"everything ok": {
			tournamentID: factory.NewString(),
			tournament:   domain.Tournament{ID: factory.NewString()},
		},
	}

	mo := mockedRepository{}
	c := Controller{
		Repository: &mo,
		IDType:     myuuid.IDType{},
	}

	name := "everything ok"
	tc := testCases[name]

	mo.On(methodNameGetTournamentByID, tc.tournamentID).Return(tc.tournament, tc.err)

	gotTC := tc
	gotTC.tournament, gotTC.err = c.GetTournament(tc.tournamentID)
	handleTestCase(name, tc, gotTC, t)

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}

func TestDeleteTournament(t *testing.T) {
	factory := myuuid.IDFactory{}
	testCases := map[string]testCase{
		"everything ok": {
			tournamentID: factory.NewString(),
		},
	}

	mo := mockedRepository{}
	c := Controller{
		Repository: &mo,
		IDType:     myuuid.IDType{},
	}

	name := "everything ok"
	tc := testCases[name]

	mo.On(methodNameDeleteTournamentByID, tc.tournamentID).Return(tc.err)

	gotTC := tc
	gotTC.err = c.DeleteTournament(tc.tournamentID)
	handleTestCase(name, tc, gotTC, t)

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}

func TestJoinTournament(t *testing.T) {
	factory := myuuid.IDFactory{}
	dublicateUserID := factory.NewString()
	testCases := map[string]testCase{
		"everything ok": {
			tournamentID: factory.NewString(),
			userID:       factory.NewString(),
			tournament: domain.Tournament{
				ID:      factory.NewString(),
				Deposit: 100,
			},
			user: domain.User{
				ID:      factory.NewString(),
				Balance: 120,
			},
		},
		"error getting tournament": {
			mockingStop:  1,
			tournamentID: factory.NewString(),
			userID:       factory.NewString(),
			tournament:   domain.Tournament{},
			err:          errors.New("error"),
		},
		"error getting user": {
			mockingStop:  2,
			tournamentID: factory.NewString(),
			userID:       factory.NewString(),
			tournament:   domain.Tournament{},
			user:         domain.User{},
			err:          errors.New("error"),
		},
		"error adding user in tournament": {
			mockingStop:  3,
			tournamentID: factory.NewString(),
			userID:       factory.NewString(),
			tournament: domain.Tournament{
				ID: factory.NewString(),
			},
			user: domain.User{
				ID: factory.NewString(),
			},
			err: errors.New("error"),
		},
		"error updating user balance": {
			mockingStop:  4,
			tournamentID: factory.NewString(),
			userID:       factory.NewString(),
			tournament: domain.Tournament{
				ID:      factory.NewString(),
				Deposit: 1,
			},
			user: domain.User{
				ID:      factory.NewString(),
				Balance: 2,
			},
			err: errors.New("error"),
		},
		"error user is participant": {
			tournamentID: factory.NewString(),
			userID:       factory.NewString(),
			tournament: domain.Tournament{
				ID: factory.NewString(),
				Participants: []string{
					dublicateUserID,
				},
			},
			user: domain.User{
				ID: dublicateUserID,
			},
			err: ErrParticipantExists,
		},
		"not enough points": {
			tournamentID: factory.NewString(),
			userID:       factory.NewString(),
			tournament: domain.Tournament{
				ID:      factory.NewString(),
				Deposit: 2,
			},
			user: domain.User{
				ID:      factory.NewString(),
				Balance: 1,
			},
			err: ErrNotEnoughPoints,
		},
	}

	mo := mockedRepository{}
	c := Controller{
		Repository: &mo,
		IDType:     myuuid.IDType{},
	}

	for name, tc := range testCases {
		if name == "error user is participant" {
			mo.On(methodNameGetTournamentByID, tc.tournamentID).Return(tc.tournament, nil)
			mo.On(methodNameGetUserByID, tc.userID).Return(tc.user, nil)
			continue
		}

		if name == "not enough points" {
			mo.On(methodNameGetTournamentByID, tc.tournamentID).Return(tc.tournament, nil)
			mo.On(methodNameGetUserByID, tc.userID).Return(tc.user, nil)
			continue
		}

		if tc.mockingStop == 1 {
			mo.On(methodNameGetTournamentByID, tc.tournamentID).Return(tc.tournament, tc.err)
			continue
		}
		
		mo.On(methodNameGetTournamentByID, tc.tournamentID).Return(tc.tournament, nil)

		if tc.mockingStop == 2 {
			mo.On(methodNameGetUserByID, tc.userID).Return(tc.user, tc.err)
			continue
		} 

		mo.On(methodNameGetUserByID, tc.userID).Return(tc.user, nil)

		if tc.mockingStop == 3 {
			mo.On(methodNameAddUserInTournament, tc.user.ID, tc.tournament.ID).Return(tc.err)
			continue
		} 

		mo.On(methodNameAddUserInTournament, tc.user.ID, tc.tournament.ID).Return(nil)

		mo.On(methodNameUpdateUserBalanceByID, tc.user.ID, tc.user.Balance-tc.tournament.Deposit).Return(tc.err)
	}

	for name, tc := range testCases {
		gotTC := tc
		gotTC.err = c.JoinTournament(tc.tournamentID, tc.userID)
		handleTestCase(name, tc, gotTC, t)
	}

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}

func mockForTestFinishTournament(tc testCase, name string) (*mockedRepository, *Controller){
	mo := mockedRepository{}
	c := Controller{
		Repository: &mo,
		IDType:     myuuid.IDType{},
	}

	if name == "error finished tournament" {
		mo.On(methodNameGetTournamentByID, tc.tournamentID).Return(tc.tournament, nil)
		return &mo, &c
	}

	if name == "error no participants"{
		mo.On(methodNameGetTournamentByID, tc.tournamentID).Return(tc.tournament, nil)
		mo.On(methodNameGetUsers).Return([]domain.User{tc.user}, nil)
		return &mo, &c
	}

	if tc.mockingStop == 1 {
		mo.On(methodNameGetTournamentByID, tc.tournamentID).Return(tc.tournament, tc.err)
		return &mo, &c
	}
	
	mo.On(methodNameGetTournamentByID, tc.tournamentID).Return(tc.tournament, nil)

	if tc.mockingStop == 2 {
		mo.On(methodNameGetUsers).Return([]domain.User{}, tc.err)
		return &mo, &c
	}
	
	mo.On(methodNameGetUsers).Return([]domain.User{{ID: "1", Balance: 1}, tc.user}, nil)

	if tc.mockingStop == 3 {
		mo.On(methodNameSetWinner, tc.user.ID, tc.tournament.ID).Return(tc.err)
		return &mo, &c
	}
	
	mo.On(methodNameSetWinner, tc.user.ID, tc.tournament.ID).Return(nil)

	mo.On(methodNameUpdateUserBalanceByID, tc.user.ID, tc.user.Balance+tc.tournament.Prize).Return(tc.err)
	return &mo, &c
}

func TestFinishTournament(t *testing.T) {
	factory := myuuid.IDFactory{}
	winnerID := factory.NewString()
	testCases := map[string]testCase{
		"everything ok": {
			tournamentID: factory.NewString(),
			tournament: domain.Tournament{
				ID:    factory.NewString(),
				Prize: 100,
				Participants: []string{
					winnerID,
				},
				WinnerID: myuuid.IDType{}.Null(),
			},
			user: domain.User{
				ID:      winnerID,
				Balance: 120,
			},
		},
		"error getting tournament": {
			mockingStop:  1,
			tournamentID: factory.NewString(),
			err:          errors.New("error"),
		},
		"error getting users": {
			mockingStop: 2,
			tournamentID: factory.NewString(),
			tournament: domain.Tournament{
				ID:       factory.NewString(),
				WinnerID: myuuid.IDType{}.Null(),
			},
			err:  errors.New("error"),
		},
		"error setting winner": {
			mockingStop:  3,
			tournamentID: factory.NewString(),
			tournament: domain.Tournament{
				ID: factory.NewString(),
				Participants: []string{
					winnerID,
				},
				WinnerID: myuuid.IDType{}.Null(),
			},
			user: domain.User{
				ID:      winnerID,
				Balance: 120,
			},
			err: errors.New("error"),
		},
		"error no participants": {
			tournamentID: factory.NewString(),
			tournament: domain.Tournament{
				ID:       factory.NewString(),
				WinnerID: myuuid.IDType{}.Null(),
			},
			err:  ErrNoParticipants,
		},
		"error finished tournament": {
			tournamentID: factory.NewString(),
			err:          ErrFinishedTournament,
		},
	}

	var mocks []*mockedRepository

	for name, tc := range testCases {
		mo, c := mockForTestFinishTournament(tc, name)
		mocks = append(mocks, mo)

		gotTC := tc
		gotTC.err = c.FinishTournament(tc.tournamentID)
		handleTestCase(name, tc, gotTC, t)
	}

	t.Logf("Asserted mocks expectations:")
	for _, mo := range mocks{
		mo.AssertExpectations(t)
	}
}
