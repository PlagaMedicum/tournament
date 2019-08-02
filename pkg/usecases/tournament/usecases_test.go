package tournament

import (
	"errors"
	"testing"
	tournamentDomain"tournament/pkg/domain/tournament"
	userDomain "tournament/pkg/domain/user"
	"tournament/pkg/usecases"
	"tournament/pkg/usecases/tournament/mocks"
)

const (
	methodNameGetUserByID           = "GetUserByID"
	methodNameGetUsers				= "GetUsers"
	methodNameUpdateUserBalanceByID = "UpdateUserBalanceByID"
	methodNameInsertTournament     = "InsertTournament"
	methodNameGetTournamentByID    = "GetTournamentByID"
	methodNameDeleteTournamentByID = "DeleteTournamentByID"
	methodNameAddUserInTournament  = "AddUserInTournament"
	methodNameSetWinner            = "SetWinner"
)

func TestCreateTournament(t *testing.T) {
	testCases := map[string]usecases.TestCase{
		"everything ok": {
			RequestName:    "Tournament",
			RequestDeposit: 500,
			ResultID:       1,
		},
	}

	mock := mocks.MockedRepository{}
	c := Controller{
		Repository: &mock,
	}

	name := "everything ok"
	tc := testCases[name]

	mock.On(methodNameInsertTournament, tc.RequestName, tc.RequestDeposit, defaultPrize).Return(tc.ResultID, tc.Err)

	gotTC := tc
	gotTC.ResultID, gotTC.Err = c.CreateTournament(tc.RequestName, tc.RequestDeposit)
	usecases.HandleTestCase(name, tc, gotTC, t)

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}

func TestGetTournament(t *testing.T) {
	testCases := map[string]usecases.TestCase{
		"everything ok": {
			TournamentID: 1,
			Tournament:   tournamentDomain.Tournament{ID: 1},
		},
	}

	mock := mocks.MockedRepository{}
	c := Controller{
		Repository: &mock,
	}

	name := "everything ok"
	tc := testCases[name]

	mock.On(methodNameGetTournamentByID, tc.TournamentID).Return(tc.Tournament, tc.Err)

	gotTC := tc
	gotTC.Tournament, gotTC.Err = c.GetTournament(tc.TournamentID)
	usecases.HandleTestCase(name, tc, gotTC, t)

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}

func TestDeleteTournament(t *testing.T) {
	testCases := map[string]usecases.TestCase{
		"everything ok": {
			TournamentID: 1,
		},
	}

	mock := mocks.MockedRepository{}
	c := Controller{
		Repository: &mock,
	}

	name := "everything ok"
	tc := testCases[name]

	mock.On(methodNameDeleteTournamentByID, tc.TournamentID).Return(tc.Err)

	gotTC := tc
	gotTC.Err = c.DeleteTournament(tc.TournamentID)
	usecases.HandleTestCase(name, tc, gotTC, t)

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}

func TestJoinTournament(t *testing.T) {
	var dublicateUserID uint64 = 33
	testCases := map[string]usecases.TestCase{
		"everything ok": {
			TournamentID: 1,
			UserID:       2,
			Tournament: tournamentDomain.Tournament{
				ID:      1,
				Deposit: 100,
			},
			User: userDomain.User{
				ID:      2,
				Balance: 120,
			},
		},
		"error getting tournament": {
			MockingStop:  1,
			TournamentID: 3,
			UserID:       4,
			Tournament:   tournamentDomain.Tournament{},
			Err:          errors.New("error"),
		},
		"error getting user": {
			MockingStop:  2,
			TournamentID: 5,
			UserID:       6,
			Tournament:   tournamentDomain.Tournament{},
			User:         userDomain.User{},
			Err:          errors.New("error"),
		},
		"error adding user in tournament": {
			MockingStop:  3,
			TournamentID: 7,
			UserID:       8,
			Tournament: tournamentDomain.Tournament{
				ID: 7,
			},
			User: userDomain.User{
				ID: 8,
			},
			Err: errors.New("error"),
		},
		"error updating user balance": {
			MockingStop:  4,
			TournamentID: 9,
			UserID:       10,
			Tournament: tournamentDomain.Tournament{
				ID:      9,
				Deposit: 1,
			},
			User: userDomain.User{
				ID:      10,
				Balance: 2,
			},
			Err: errors.New("error"),
		},
		"error user is participant": {
			TournamentID: 11,
			UserID:       12,
			Tournament: tournamentDomain.Tournament{
				ID: 11,
				Participants: []uint64{
					dublicateUserID,
				},
			},
			User: userDomain.User{
				ID: dublicateUserID,
			},
			Err: usecases.ErrParticipantExists,
		},
		"not enough points": {
			TournamentID: 13,
			UserID:       14,
			Tournament: tournamentDomain.Tournament{
				ID:      13,
				Deposit: 2,
			},
			User: userDomain.User{
				ID:      14,
				Balance: 1,
			},
			Err: usecases.ErrNotEnoughPoints,
		},
	}

	mock := mocks.MockedRepository{}
	c := Controller{
		Repository: &mock,
	}

	for name, tc := range testCases {
		if name == "error user is participant" {
			mock.On(methodNameGetTournamentByID, tc.TournamentID).Return(tc.Tournament, nil)
			mock.On(methodNameGetUserByID, tc.UserID).Return(tc.User, nil)
			continue
		}

		if name == "not enough points" {
			mock.On(methodNameGetTournamentByID, tc.TournamentID).Return(tc.Tournament, nil)
			mock.On(methodNameGetUserByID, tc.UserID).Return(tc.User, nil)
			continue
		}

		if tc.MockingStop == 1 {
			mock.On(methodNameGetTournamentByID, tc.TournamentID).Return(tc.Tournament, tc.Err)
			continue
		}
		
		mock.On(methodNameGetTournamentByID, tc.TournamentID).Return(tc.Tournament, nil)

		if tc.MockingStop == 2 {
			mock.On(methodNameGetUserByID, tc.UserID).Return(tc.User, tc.Err)
			continue
		} 

		mock.On(methodNameGetUserByID, tc.UserID).Return(tc.User, nil)

		if tc.MockingStop == 3 {
			mock.On(methodNameAddUserInTournament, tc.User.ID, tc.Tournament.ID).Return(tc.Err)
			continue
		} 

		mock.On(methodNameAddUserInTournament, tc.User.ID, tc.Tournament.ID).Return(nil)

		mock.On(methodNameUpdateUserBalanceByID, tc.User.ID, tc.User.Balance-tc.Tournament.Deposit).Return(tc.Err)
	}

	for name, tc := range testCases {
		gotTC := tc
		gotTC.Err = c.JoinTournament(tc.TournamentID, tc.UserID)
		usecases.HandleTestCase(name, tc, gotTC, t)
	}

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}

func mockForTestFinishTournament(tc usecases.TestCase, name string) (*mocks.MockedRepository, *Controller){
	mock := mocks.MockedRepository{}
	c := Controller{
		Repository: &mock,
	}

	if name == "error no participants"{
		mock.On(methodNameGetTournamentByID, tc.TournamentID).Return(tc.Tournament, nil)
		mock.On(methodNameGetUsers).Return([]userDomain.User{tc.User}, nil)
		return &mock, &c
	}

	if name == "error finished tournament" {
		mock.On(methodNameGetTournamentByID, tc.TournamentID).Return(tc.Tournament, nil)
		return &mock, &c
	}

	if tc.MockingStop == 1 {
		mock.On(methodNameGetTournamentByID, tc.TournamentID).Return(tc.Tournament, tc.Err)
		return &mock, &c
	}
	
	mock.On(methodNameGetTournamentByID, tc.TournamentID).Return(tc.Tournament, nil)

	if tc.MockingStop == 2 {
		mock.On(methodNameGetUsers).Return([]userDomain.User{}, tc.Err)
		return &mock, &c
	}
	
	mock.On(methodNameGetUsers).Return([]userDomain.User{{ID: 99, Balance: 1}, tc.User}, nil)

	if tc.MockingStop == 3 {
		mock.On(methodNameSetWinner, tc.User.ID, tc.Tournament.ID).Return(tc.Err)
		return &mock, &c
	}
	
	mock.On(methodNameSetWinner, tc.User.ID, tc.Tournament.ID).Return(nil)

	mock.On(methodNameUpdateUserBalanceByID, tc.User.ID, tc.User.Balance+tc.Tournament.Prize).Return(tc.Err)
	return &mock, &c
}

func TestFinishTournament(t *testing.T) {
	var winnerID uint64 = 33
	testCases := map[string]usecases.TestCase{
		"everything ok": {
			TournamentID: 1,
			Tournament: tournamentDomain.Tournament{
				ID:    1,
				Prize: 100,
				Participants: []uint64{
					winnerID,
				},
				WinnerID: 0,
			},
			User: userDomain.User{
				ID:      winnerID,
				Balance: 120,
			},
		},
		"error getting tournament": {
			MockingStop:  1,
			TournamentID: 2,
			Err:          errors.New("error"),
		},
		"error getting users": {
			MockingStop:  2,
			TournamentID: 3,
			Tournament: tournamentDomain.Tournament{
				ID:       3,
				WinnerID: 0,
			},
			Err: errors.New("error"),
		},
		"error setting winner": {
			MockingStop:  3,
			TournamentID: 4,
			Tournament: tournamentDomain.Tournament{
				ID: 4,
				Participants: []uint64{
					winnerID,
				},
				WinnerID: 0,
			},
			User: userDomain.User{
				ID:      winnerID,
				Balance: 120,
			},
			Err: errors.New("error"),
		},
		"error no participants": {
			TournamentID: 5,
			Tournament: tournamentDomain.Tournament{
				ID:       5,
				WinnerID: 0,
			},
			Err: usecases.ErrNoParticipants,
		},
		"error finished tournament": {
			TournamentID: 6,
			Tournament: tournamentDomain.Tournament{
				WinnerID: 7,
			},
			Err: usecases.ErrFinishedTournament,
		},
	}

	var mockList []*mocks.MockedRepository

	for name, tc := range testCases {
		mock, c := mockForTestFinishTournament(tc, name)
		mockList = append(mockList, mock)

		gotTC := tc
		gotTC.Err = c.FinishTournament(tc.TournamentID)
		usecases.HandleTestCase(name, tc, gotTC, t)
	}

	t.Logf("Asserted mocks expectations:")
	for _, mock := range mockList{
		mock.AssertExpectations(t)
	}
}
