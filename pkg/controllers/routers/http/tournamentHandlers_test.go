package http

import (
	"errors"
	"net/http"
	"strconv"
	"testing"
	httpHandler "tournament/pkg/controllers/handlers/http"
	"tournament/pkg/controllers/routers/http/mocks"
	"tournament/pkg/domain"
	"tournament/pkg/infrastructure/myhandler"
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
	testCases := []testCase{
		{
			caseName: "everything ok",
			resultID: 1,
			requestTournament: domain.Tournament{
				Name:    "Unreal Tournament",
				Deposit: 10000,
			},
			requestBody:    `{"name": "Unreal Tournament", "deposit": 10000}`,
			expectedStatus: http.StatusCreated,
		},
		{
			caseName:          "wrong body",
			noMock:            true,
			requestBody:       `i'm the wrong body"`,
			requestTournament: domain.Tournament{},
			expectedStatus:    http.StatusBadRequest,
		},
		{
			caseName:  "wrong tournament error",
			resultErr: errors.New("i'm the bad err"),
			resultID:  2,
			requestTournament: domain.Tournament{
				Name:    "test tour",
				Deposit: 1,
			},
			requestBody:    `{"name": "test tour", "deposit": 1}`,
			expectedStatus: http.StatusBadRequest,
		},
	}

	mock := mocks.MockedUsecases{}
	h := myhandler.Handler{}
	Route(&h, httpHandler.Controller{Usecases: &mock})

	for _, tc := range testCases {
		tc.path = httpHandler.TournamentPath
		tc.method = http.MethodPost

		if !tc.noMock {
			mock.On(methodNameCreateTournament, tc.requestTournament.Name, tc.requestTournament.Deposit).Return(tc.resultID, tc.resultErr)
		}

		handleTestCase(t, &h, tc)
	}

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}

// TestGetTournamentHandler tests getting of tournament's information.
func TestGetTournamentHandler(t *testing.T) {
	testCases := []testCase{
		{
			caseName: "everything ok",
			resultTournament: domain.Tournament{
				Name: "test tour",
			},
			requestID:      1,
			expectedStatus: http.StatusOK,
		},
		{
			caseName:         "wrong tournament error",
			resultErr:        errors.New("i'm the bad err"),
			resultTournament: domain.Tournament{},
			requestID:        2,
			expectedStatus:   http.StatusBadRequest,
		},
	}

	mock := mocks.MockedUsecases{}
	h := myhandler.Handler{}
	Route(&h, httpHandler.Controller{Usecases: &mock})

	for _, tc := range testCases {
		tc.path = httpHandler.TournamentPath + "/" + strconv.FormatUint(tc.requestID, 10)
		tc.method = http.MethodGet

		if !tc.noMock {
			mock.On(methodNameGetTournament, tc.requestID).Return(tc.resultTournament, tc.resultErr)
		}

		handleTestCase(t, &h, tc)
	}

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}

// TestGetTournamentHandler tests deleting of tournament.
func TestDeleteTournamentHandler(t *testing.T) {
	testCases := []testCase{
		{
			caseName:       "everything ok",
			requestID:      1,
			expectedStatus: http.StatusOK,
		},
		{
			caseName:       "wrong tournament error",
			resultErr:      errors.New("i'm the bad err"),
			requestID:      2,
			expectedStatus: http.StatusBadRequest,
		},
	}

	mock := mocks.MockedUsecases{}
	h := myhandler.Handler{}
	Route(&h, httpHandler.Controller{Usecases: &mock})

	for _, tc := range testCases {
		tc.path = httpHandler.TournamentPath + "/" + strconv.FormatUint(tc.requestID, 10)
		tc.method = http.MethodDelete

		if !tc.noMock {
			mock.On(methodNameDeleteTournament, tc.requestID).Return(tc.resultErr)
		}

		handleTestCase(t, &h, tc)
	}

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}

// TestJoinTournamentHandler tests joining tournament.
func TestJoinTournamentHandler(t *testing.T) {
	var requestUserID uint64 = 33
	testCases := []testCase{
		{
			caseName:       "everything ok",
			requestID:      1,
			requestBody:    `{"userId": ` + strconv.FormatUint(requestUserID, 10) + `}`,
			expectedStatus: http.StatusOK,
		},
		{
			caseName:       "wrong body",
			noMock:         true,
			requestBody:    `i'm the wrong body"`,
			requestID:      2,
			expectedStatus: http.StatusBadRequest,
		},
		{
			caseName:       "wrong tournament error",
			resultErr:      errors.New("i'm the bad err"),
			requestID:      3,
			requestBody:    `{"userId": ` + strconv.FormatUint(requestUserID, 10) + `}`,
			expectedStatus: http.StatusBadRequest,
		},
	}

	mock := mocks.MockedUsecases{}
	h := myhandler.Handler{}
	Route(&h, httpHandler.Controller{Usecases: &mock})

	for _, tc := range testCases {
		tc.path = httpHandler.TournamentPath + "/" + strconv.FormatUint(tc.requestID, 10) + httpHandler.JoinTournamentPath
		tc.method = http.MethodPost

		if !tc.noMock {
			mock.On(methodNameJoinTournament, tc.requestID, requestUserID).Return(tc.resultErr)
		}

		handleTestCase(t, &h, tc)
	}

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}

// TestFinishTournamentHandler tests joining tournament.
func TestFinishTournamentHandler(t *testing.T) {

	testCases := []testCase{
		{
			caseName:       "everything ok",
			requestID:      1,
			expectedStatus: http.StatusOK,
		},
		{
			caseName:       "wrong tournament error",
			resultErr:      errors.New("i'm the bad err"),
			requestID:      2,
			expectedStatus: http.StatusBadRequest,
		},
	}

	mock := mocks.MockedUsecases{}
	h := myhandler.Handler{}
	Route(&h, httpHandler.Controller{Usecases: &mock})

	for _, tc := range testCases {
		tc.path = httpHandler.TournamentPath + "/" + strconv.FormatUint(tc.requestID, 10) + httpHandler.FinishTournamentPath
		tc.method = http.MethodPost

		if !tc.noMock {
			mock.On(methodNameFinishTournament, tc.requestID).Return(tc.resultErr)
		}

		handleTestCase(t, &h, tc)
	}

	t.Logf("Asserted mocks expectations:")
	mock.AssertExpectations(t)
}
