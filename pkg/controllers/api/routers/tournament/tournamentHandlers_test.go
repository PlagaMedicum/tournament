package tournament

import (
	"errors"
	"net/http"
	"strconv"
	"testing"
	httpHandlers "tournament/pkg/controllers/api/http_handlers"
	"tournament/pkg/controllers/api/http_handlers/tournament"
	"tournament/pkg/controllers/api/routers"
	"tournament/pkg/controllers/api/routers/tournament/mocks"
	tournamentDomain "tournament/pkg/domain/tournament"
	httpHandler "tournament/pkg/infrastructure/handler"
)

const (
	methodNameCreateTournament = "CreateTournament"
	methodNameGetTournament    = "GetTournament"
	methodNameDeleteTournament = "DeleteTournament"
	methodNameJoinTournament   = "JoinTournament"
	methodNameFinishTournament = "FinishTournament"
)

// TestCreateTournamentHandler tests creation of tournament.
func TestCreateTournamentHandler(t *testing.T) {
	testCases := []routers.TestCase{
		{
			CaseName: "everything ok",
			ResultID: 1,
			RequestTournament: tournamentDomain.Tournament{
				Name:    "Unreal Tournament",
				Deposit: 10000,
			},
			RequestBody:    `{"name": "Unreal Tournament", "deposit": 10000}`,
			ExpectedStatus: http.StatusCreated,
		},
		{
			CaseName:          "wrong body",
			NoMock:            true,
			RequestBody:       `i'm the wrong body"`,
			RequestTournament: tournamentDomain.Tournament{},
			ExpectedStatus:    http.StatusBadRequest,
		},
		{
			CaseName:  "wrong tournament error",
			ResultErr: errors.New("i'm the bad err"),
			ResultID:  2,
			RequestTournament: tournamentDomain.Tournament{
				Name:    "test tour",
				Deposit: 1,
			},
			RequestBody:    `{"name": "test tour", "deposit": 1}`,
			ExpectedStatus: http.StatusBadRequest,
		},
	}

	mock := mocks.MockedUsecases{}
	h := httpHandler.Handler{}
	RouteTournament(&h, tournament.Controller{Usecases: &mock})

	for _, tc := range testCases {
		tc.Path = httpHandlers.TournamentPath
		tc.Method = http.MethodPost

		if !tc.NoMock {
			mock.On(methodNameCreateTournament, tc.RequestTournament.Name, tc.RequestTournament.Deposit).Return(tc.ResultID, tc.ResultErr)
		}

		routers.HandleTestCase(t, &h, tc)
	}

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}

// TestGetTournamentHandler tests getting of tournament's information.
func TestGetTournamentHandler(t *testing.T) {
	testCases := []routers.TestCase{
		{
			CaseName: "everything ok",
			ResultTournament: tournamentDomain.Tournament{
				Name: "test tour",
			},
			RequestID:      1,
			ExpectedStatus: http.StatusOK,
		},
		{
			CaseName:         "wrong tournament error",
			ResultErr:        errors.New("i'm the bad err"),
			ResultTournament: tournamentDomain.Tournament{},
			RequestID:        2,
			ExpectedStatus:   http.StatusBadRequest,
		},
	}

	mock := mocks.MockedUsecases{}
	h := httpHandler.Handler{}
	RouteTournament(&h, tournament.Controller{Usecases: &mock})

	for _, tc := range testCases {
		tc.Path = httpHandlers.TournamentPath + "/" + strconv.FormatUint(tc.RequestID, 10)
		tc.Method = http.MethodGet

		if !tc.NoMock {
			mock.On(methodNameGetTournament, tc.RequestID).Return(tc.ResultTournament, tc.ResultErr)
		}

		routers.HandleTestCase(t, &h, tc)
	}

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}

// TestGetTournamentHandler tests deleting of tournament.
func TestDeleteTournamentHandler(t *testing.T) {
	testCases := []routers.TestCase{
		{
			CaseName:       "everything ok",
			RequestID:      1,
			ExpectedStatus: http.StatusOK,
		},
		{
			CaseName:       "wrong tournament error",
			ResultErr:      errors.New("i'm the bad err"),
			RequestID:      2,
			ExpectedStatus: http.StatusBadRequest,
		},
	}

	mock := mocks.MockedUsecases{}
	h := httpHandler.Handler{}
	RouteTournament(&h, tournament.Controller{Usecases: &mock})

	for _, tc := range testCases {
		tc.Path = httpHandlers.TournamentPath + "/" + strconv.FormatUint(tc.RequestID, 10)
		tc.Method = http.MethodDelete

		if !tc.NoMock {
			mock.On(methodNameDeleteTournament, tc.RequestID).Return(tc.ResultErr)
		}

		routers.HandleTestCase(t, &h, tc)
	}

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}

// TestJoinTournamentHandler tests joining tournament.
func TestJoinTournamentHandler(t *testing.T) {
	var requestUserID uint64 = 33
	testCases := []routers.TestCase{
		{
			CaseName:       "everything ok",
			RequestID:      1,
			RequestBody:    `{"userId": ` + strconv.FormatUint(requestUserID, 10) + `}`,
			ExpectedStatus: http.StatusOK,
		},
		{
			CaseName:       "wrong body",
			NoMock:         true,
			RequestBody:    `i'm the wrong body"`,
			RequestID:      2,
			ExpectedStatus: http.StatusBadRequest,
		},
		{
			CaseName:       "wrong tournament error",
			ResultErr:      errors.New("i'm the bad err"),
			RequestID:      3,
			RequestBody:    `{"userId": ` + strconv.FormatUint(requestUserID, 10) + `}`,
			ExpectedStatus: http.StatusBadRequest,
		},
	}

	mock := mocks.MockedUsecases{}
	h := httpHandler.Handler{}
	RouteTournament(&h, tournament.Controller{Usecases: &mock})

	for _, tc := range testCases {
		tc.Path = httpHandlers.TournamentPath + "/" + strconv.FormatUint(tc.RequestID, 10) + httpHandlers.JoinTournamentPath
		tc.Method = http.MethodPost

		if !tc.NoMock {
			mock.On(methodNameJoinTournament, tc.RequestID, requestUserID).Return(tc.ResultErr)
		}

		routers.HandleTestCase(t, &h, tc)
	}

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}

// TestFinishTournamentHandler tests joining tournament.
func TestFinishTournamentHandler(t *testing.T) {

	testCases := []routers.TestCase{
		{
			CaseName:       "everything ok",
			RequestID:      1,
			ExpectedStatus: http.StatusOK,
		},
		{
			CaseName:       "wrong tournament error",
			ResultErr:      errors.New("i'm the bad err"),
			RequestID:      2,
			ExpectedStatus: http.StatusBadRequest,
		},
	}

	mock := mocks.MockedUsecases{}
	h := httpHandler.Handler{}
	RouteTournament(&h, tournament.Controller{Usecases: &mock})

	for _, tc := range testCases {
		tc.Path = httpHandlers.TournamentPath + "/" + strconv.FormatUint(tc.RequestID, 10) + httpHandlers.FinishTournamentPath
		tc.Method = http.MethodPost

		if !tc.NoMock {
			mock.On(methodNameFinishTournament, tc.RequestID).Return(tc.ResultErr)
		}

		routers.HandleTestCase(t, &h, tc)
	}

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}
