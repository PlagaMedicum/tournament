package usecases

import (
	"errors"
	"testing"
	"tournament/pkg/domain"
	"tournament/pkg/usecases/mocks"
)

const (
	methodNameInsertTournament     = "InsertTournament"
	methodNameGetTournamentByID    = "GetTournamentByID"
	methodNameDeleteTournamentByID = "DeleteTournamentByID"
	methodNameAddUserInTournament  = "AddUserInTournament"
	methodNameSetWinner            = "SetWinner"
)

func TestCreateTournament(t *testing.T) {
	testCases := map[string]testCase{
		"everything ok": {
			requestName:    "Tournament",
			requestDeposit: 500,
			resultID:       1,
		},
	}

	mock := mocks.MockedRepository{}
	c := Controller{
		Repository: &mock,
	}

	name := "everything ok"
	tc := testCases[name]

	mock.On(methodNameInsertTournament, tc.requestName, tc.requestDeposit, defaultPrize).Return(tc.resultID, tc.err)

	gotTC := tc
	gotTC.resultID, gotTC.err = c.CreateTournament(tc.requestName, tc.requestDeposit)
	handleTestCase(name, tc, gotTC, t)

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}

func TestGetTournament(t *testing.T) {
	testCases := map[string]testCase{
		"everything ok": {
			tournamentID: 1,
			tournament:   domain.Tournament{ID: 1},
		},
	}

	mock := mocks.MockedRepository{}
	c := Controller{
		Repository: &mock,
	}

	name := "everything ok"
	tc := testCases[name]

	mock.On(methodNameGetTournamentByID, tc.tournamentID).Return(tc.tournament, tc.err)

	gotTC := tc
	gotTC.tournament, gotTC.err = c.GetTournament(tc.tournamentID)
	handleTestCase(name, tc, gotTC, t)

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}

func TestDeleteTournament(t *testing.T) {
	testCases := map[string]testCase{
		"everything ok": {
			tournamentID: 1,
		},
	}

	mock := mocks.MockedRepository{}
	c := Controller{
		Repository: &mock,
	}

	name := "everything ok"
	tc := testCases[name]

	mock.On(methodNameDeleteTournamentByID, tc.tournamentID).Return(tc.err)

	gotTC := tc
	gotTC.err = c.DeleteTournament(tc.tournamentID)
	handleTestCase(name, tc, gotTC, t)

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}

func TestJoinTournament(t *testing.T) {
	var dublicateUserID uint64 = 33
	testCases := map[string]testCase{
		"everything ok": {
			tournamentID: 1,
			userID:       2,
			tournament: domain.Tournament{
				ID:      1,
				Deposit: 100,
			},
			user: domain.User{
				ID:      2,
				Balance: 120,
			},
		},
		"error getting tournament": {
			mockingStop:  1,
			tournamentID: 3,
			userID:       4,
			tournament:   domain.Tournament{},
			err:          errors.New("error"),
		},
		"error getting user": {
			mockingStop:  2,
			tournamentID: 5,
			userID:       6,
			tournament:   domain.Tournament{},
			user:         domain.User{},
			err:          errors.New("error"),
		},
		"error adding user in tournament": {
			mockingStop:  3,
			tournamentID: 7,
			userID:       8,
			tournament: domain.Tournament{
				ID: 7,
			},
			user: domain.User{
				ID: 8,
			},
			err: errors.New("error"),
		},
		"error updating user balance": {
			mockingStop:  4,
			tournamentID: 9,
			userID:       10,
			tournament: domain.Tournament{
				ID:      9,
				Deposit: 1,
			},
			user: domain.User{
				ID:      10,
				Balance: 2,
			},
			err: errors.New("error"),
		},
		"error user is participant": {
			tournamentID: 11,
			userID:       12,
			tournament: domain.Tournament{
				ID: 11,
				Participants: []uint64{
					dublicateUserID,
				},
			},
			user: domain.User{
				ID: dublicateUserID,
			},
			err: ErrParticipantExists,
		},
		"not enough points": {
			tournamentID: 13,
			userID:       14,
			tournament: domain.Tournament{
				ID:      13,
				Deposit: 2,
			},
			user: domain.User{
				ID:      14,
				Balance: 1,
			},
			err: ErrNotEnoughPoints,
		},
	}

	mock := mocks.MockedRepository{}
	c := Controller{
		Repository: &mock,
	}

	for name, tc := range testCases {
		if name == "error user is participant" {
			mock.On(methodNameGetTournamentByID, tc.tournamentID).Return(tc.tournament, nil)
			mock.On(methodNameGetUserByID, tc.userID).Return(tc.user, nil)
			continue
		}

		if name == "not enough points" {
			mock.On(methodNameGetTournamentByID, tc.tournamentID).Return(tc.tournament, nil)
			mock.On(methodNameGetUserByID, tc.userID).Return(tc.user, nil)
			continue
		}

		if tc.mockingStop == 1 {
			mock.On(methodNameGetTournamentByID, tc.tournamentID).Return(tc.tournament, tc.err)
			continue
		}
		
		mock.On(methodNameGetTournamentByID, tc.tournamentID).Return(tc.tournament, nil)

		if tc.mockingStop == 2 {
			mock.On(methodNameGetUserByID, tc.userID).Return(tc.user, tc.err)
			continue
		} 

		mock.On(methodNameGetUserByID, tc.userID).Return(tc.user, nil)

		if tc.mockingStop == 3 {
			mock.On(methodNameAddUserInTournament, tc.user.ID, tc.tournament.ID).Return(tc.err)
			continue
		} 

		mock.On(methodNameAddUserInTournament, tc.user.ID, tc.tournament.ID).Return(nil)

		mock.On(methodNameUpdateUserBalanceByID, tc.user.ID, tc.user.Balance-tc.tournament.Deposit).Return(tc.err)
	}

	for name, tc := range testCases {
		gotTC := tc
		gotTC.err = c.JoinTournament(tc.tournamentID, tc.userID)
		handleTestCase(name, tc, gotTC, t)
	}

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}

func mockForTestFinishTournament(tc testCase, name string) (*mocks.MockedRepository, *Controller){
	mock := mocks.MockedRepository{}
	c := Controller{
		Repository: &mock,
	}

	if name == "error no participants"{
		mock.On(methodNameGetTournamentByID, tc.tournamentID).Return(tc.tournament, nil)
		mock.On(methodNameGetUsers).Return([]domain.User{tc.user}, nil)
		return &mock, &c
	}

	if name == "error finished tournament" {
		mock.On(methodNameGetTournamentByID, tc.tournamentID).Return(tc.tournament, nil)
		return &mock, &c
	}

	if tc.mockingStop == 1 {
		mock.On(methodNameGetTournamentByID, tc.tournamentID).Return(tc.tournament, tc.err)
		return &mock, &c
	}
	
	mock.On(methodNameGetTournamentByID, tc.tournamentID).Return(tc.tournament, nil)

	if tc.mockingStop == 2 {
		mock.On(methodNameGetUsers).Return([]domain.User{}, tc.err)
		return &mock, &c
	}
	
	mock.On(methodNameGetUsers).Return([]domain.User{{ID: 99, Balance: 1}, tc.user}, nil)

	if tc.mockingStop == 3 {
		mock.On(methodNameSetWinner, tc.user.ID, tc.tournament.ID).Return(tc.err)
		return &mock, &c
	}
	
	mock.On(methodNameSetWinner, tc.user.ID, tc.tournament.ID).Return(nil)

	mock.On(methodNameUpdateUserBalanceByID, tc.user.ID, tc.user.Balance+tc.tournament.Prize).Return(tc.err)
	return &mock, &c
}

func TestFinishTournament(t *testing.T) {
	var winnerID uint64 = 33
	testCases := map[string]testCase{
		"everything ok": {
			tournamentID: 1,
			tournament: domain.Tournament{
				ID:    1,
				Prize: 100,
				Participants: []uint64{
					winnerID,
				},
				WinnerID: 0,
			},
			user: domain.User{
				ID:      winnerID,
				Balance: 120,
			},
		},
		"error getting tournament": {
			mockingStop:  1,
			tournamentID: 2,
			err:          errors.New("error"),
		},
		"error getting users": {
			mockingStop: 2,
			tournamentID: 3,
			tournament: domain.Tournament{
				ID:       3,
				WinnerID: 0,
			},
			err:  errors.New("error"),
		},
		"error setting winner": {
			mockingStop:  3,
			tournamentID: 4,
			tournament: domain.Tournament{
				ID: 4,
				Participants: []uint64{
					winnerID,
				},
				WinnerID: 0,
			},
			user: domain.User{
				ID:      winnerID,
				Balance: 120,
			},
			err: errors.New("error"),
		},
		"error no participants": {
			tournamentID: 5,
			tournament: domain.Tournament{
				ID:       5,
				WinnerID: 0,
			},
			err:  ErrNoParticipants,
		},
		"error finished tournament": {
			tournamentID: 6,
			tournament: domain.Tournament{
				WinnerID: 7,
			},
			err:          ErrFinishedTournament,
		},
	}

	var mockList []*mocks.MockedRepository

	for name, tc := range testCases {
		mock, c := mockForTestFinishTournament(tc, name)
		mockList = append(mockList, mock)

		gotTC := tc
		gotTC.err = c.FinishTournament(tc.tournamentID)
		handleTestCase(name, tc, gotTC, t)
	}

	t.Logf("Asserted mocks expectations:")
	for _, mock := range mockList{
		mock.AssertExpectations(t)
	}
}
