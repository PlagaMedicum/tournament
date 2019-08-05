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
			ResID:    1,
			ReqTournament: tournamentDomain.Tournament{
				Name:    "Unreal Tournament",
				Deposit: 10000,
			},
			ReqBody:   `{"name": "Unreal Tournament", "deposit": 10000}`,
			ResStatus: http.StatusCreated,
		},
		{
			CaseName:      "wrong body",
			NoMock:        true,
			ReqBody:       `i'm the wrong body"`,
			ReqTournament: tournamentDomain.Tournament{},
			ResStatus:     http.StatusBadRequest,
		},
		{
			CaseName: "wrong tournament error",
			ResErr:   errors.New("i'm the bad err"),
			ResID:    2,
			ReqTournament: tournamentDomain.Tournament{
				Name:    "test tour",
				Deposit: 1,
			},
			ReqBody:   `{"name": "test tour", "deposit": 1}`,
			ResStatus: http.StatusBadRequest,
		},
	}

	mock := mocks.MockedUsecases{}
	h := httpHandler.Handler{}
	RouteTournament(&h, tournament.Controller{Usecases: &mock})

	for _, tc := range testCases {
		tc.Path = httpHandlers.TournamentPath
		tc.Method = http.MethodPost

		if !tc.NoMock {
			mock.On(methodNameCreateTournament, tc.ReqTournament.Name, tc.ReqTournament.Deposit).Return(tc.ResID, tc.ResErr)
		}

		tc.Handle(t, &h)
	}

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}

// TestGetTournamentHandler tests getting of tournament's information.
func TestGetTournamentHandler(t *testing.T) {
	testCases := []routers.TestCase{
		{
			CaseName: "everything ok",
			ResTournament: tournamentDomain.Tournament{
				Name: "test tour",
			},
			ReqID:     1,
			ResStatus: http.StatusOK,
		},
		{
			CaseName:      "wrong tournament error",
			ResErr:        errors.New("i'm the bad err"),
			ResTournament: tournamentDomain.Tournament{},
			ReqID:         2,
			ResStatus:     http.StatusBadRequest,
		},
	}

	mock := mocks.MockedUsecases{}
	h := httpHandler.Handler{}
	RouteTournament(&h, tournament.Controller{Usecases: &mock})

	for _, tc := range testCases {
		tc.Path = httpHandlers.TournamentPath + "/" + strconv.FormatUint(tc.ReqID, 10)
		tc.Method = http.MethodGet

		if !tc.NoMock {
			mock.On(methodNameGetTournament, tc.ReqID).Return(tc.ResTournament, tc.ResErr)
		}

		tc.Handle(t, &h)
	}

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}

// TestGetTournamentHandler tests deleting of tournament.
func TestDeleteTournamentHandler(t *testing.T) {
	testCases := []routers.TestCase{
		{
			CaseName:  "everything ok",
			ReqID:     1,
			ResStatus: http.StatusOK,
		},
		{
			CaseName:  "wrong tournament error",
			ResErr:    errors.New("i'm the bad err"),
			ReqID:     2,
			ResStatus: http.StatusBadRequest,
		},
	}

	mock := mocks.MockedUsecases{}
	h := httpHandler.Handler{}
	RouteTournament(&h, tournament.Controller{Usecases: &mock})

	for _, tc := range testCases {
		tc.Path = httpHandlers.TournamentPath + "/" + strconv.FormatUint(tc.ReqID, 10)
		tc.Method = http.MethodDelete

		if !tc.NoMock {
			mock.On(methodNameDeleteTournament, tc.ReqID).Return(tc.ResErr)
		}

		tc.Handle(t, &h)
	}

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}

// TestJoinTournamentHandler tests joining tournament.
func TestJoinTournamentHandler(t *testing.T) {
	var requestUserID uint64 = 33
	testCases := []routers.TestCase{
		{
			CaseName:  "everything ok",
			ReqID:     1,
			ReqBody:   `{"userId": ` + strconv.FormatUint(requestUserID, 10) + `}`,
			ResStatus: http.StatusOK,
		},
		{
			CaseName:  "wrong body",
			NoMock:    true,
			ReqBody:   `i'm the wrong body"`,
			ReqID:     2,
			ResStatus: http.StatusBadRequest,
		},
		{
			CaseName:  "wrong tournament error",
			ResErr:    errors.New("i'm the bad err"),
			ReqID:     3,
			ReqBody:   `{"userId": ` + strconv.FormatUint(requestUserID, 10) + `}`,
			ResStatus: http.StatusBadRequest,
		},
	}

	mock := mocks.MockedUsecases{}
	h := httpHandler.Handler{}
	RouteTournament(&h, tournament.Controller{Usecases: &mock})

	for _, tc := range testCases {
		tc.Path = httpHandlers.TournamentPath + "/" + strconv.FormatUint(tc.ReqID, 10) + httpHandlers.JoinTournamentPath
		tc.Method = http.MethodPost

		if !tc.NoMock {
			mock.On(methodNameJoinTournament, tc.ReqID, requestUserID).Return(tc.ResErr)
		}

		tc.Handle(t, &h)
	}

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}

// TestFinishTournamentHandler tests joining tournament.
func TestFinishTournamentHandler(t *testing.T) {

	testCases := []routers.TestCase{
		{
			CaseName:  "everything ok",
			ReqID:     1,
			ResStatus: http.StatusOK,
		},
		{
			CaseName:  "wrong tournament error",
			ResErr:    errors.New("i'm the bad err"),
			ReqID:     2,
			ResStatus: http.StatusBadRequest,
		},
	}

	mock := mocks.MockedUsecases{}
	h := httpHandler.Handler{}
	RouteTournament(&h, tournament.Controller{Usecases: &mock})

	for _, tc := range testCases {
		tc.Path = httpHandlers.TournamentPath + "/" + strconv.FormatUint(tc.ReqID, 10) + httpHandlers.FinishTournamentPath
		tc.Method = http.MethodPost

		if !tc.NoMock {
			mock.On(methodNameFinishTournament, tc.ReqID).Return(tc.ResErr)
		}

		tc.Handle(t, &h)
	}

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}
