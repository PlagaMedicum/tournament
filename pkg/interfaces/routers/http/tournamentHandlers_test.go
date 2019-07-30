package http

import (
	"errors"
	"net/http"
	"testing"
	"tournament/pkg/domain"
	"tournament/pkg/infrastructure/myhandler"
	"tournament/pkg/infrastructure/myuuid"
	handler "tournament/pkg/interfaces/handlers/http"
)

const(
	methodNameCreateTournament = "CreateTournament"
	methodNameGetTournament    = "GetTournament"
	methodNameDeleteTournament = "DeleteTournament"
	methodNameJoinTournament   = "JoinTournament"
	methodNameFinishTournament = "FinishTournament"
)

// TestCreateTournamentHandler tests creation of tournament.
func TestCreateTournamentHandler(t *testing.T) {
	idFactory := myuuid.IDFactory{}
	testCases := []testCase{
		{
			caseName:  "everything ok",
			resultID:  idFactory.NewString(),
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
			resultID:  idFactory.NewString(),
			requestTournament: domain.Tournament{
				Name:    "test tour",
				Deposit: 1,
			},
			requestBody:    `{"name": "test tour", "deposit": 1}`,
			expectedStatus: http.StatusBadRequest,
		},
	}

	idType := myuuid.IDType{}
	mo := mockedRepositoryInteractor{}
	h := myhandler.Handler{}
	r := Router{IDType: idType}
	r.Route(&h, &mo)

	for _, tc := range testCases {
		tc.path = handler.TournamentPath
		tc.method = http.MethodPost

		if !tc.noMock {
			mo.On(methodNameCreateTournament, tc.requestTournament.Name, tc.requestTournament.Deposit).Return(tc.resultID, tc.resultErr)
		}

		handleTestCase(t, &h, tc)
	}

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}

// TestGetTournamentHandler tests getting of tournament's information.
func TestGetTournamentHandler(t *testing.T) {
	idFactory := myuuid.IDFactory{}
	testCases := []testCase{
		{
			caseName:  "everything ok",
			resultTournament: domain.Tournament{
				Name: "test tour",
			},
			requestID:      idFactory.NewString(),
			expectedStatus: http.StatusOK,
		},
		{
			caseName:         "wrong tournament error",
			resultErr:        errors.New("i'm the bad err"),
			resultTournament: domain.Tournament{},
			requestID:        idFactory.NewString(),
			expectedStatus:   http.StatusBadRequest,
		},
	}

	idType := myuuid.IDType{}
	mo := mockedRepositoryInteractor{}
	h := myhandler.Handler{}
	r := Router{IDType: idType}
	r.Route(&h, &mo)

	for _, tc := range testCases {
		tc.path = handler.TournamentPath + "/" + tc.requestID
		tc.method = http.MethodGet

		if !tc.noMock {
			mo.On(methodNameGetTournament, tc.requestID).Return(tc.resultTournament, tc.resultErr)
		}

		handleTestCase(t, &h, tc)
	}

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}

// TestGetTournamentHandler tests deleting of tournament.
func TestDeleteTournamentHandler(t *testing.T) {
	idFactory := myuuid.IDFactory{}
	testCases := []testCase{
		{
			caseName:       "everything ok",
			requestID:      idFactory.NewString(),
			expectedStatus: http.StatusOK,
		},
		{
			caseName:       "wrong tournament error",
			resultErr:      errors.New("i'm the bad err"),
			requestID:      idFactory.NewString(),
			expectedStatus: http.StatusBadRequest,
		},
	}

	idType := myuuid.IDType{}
	mo := mockedRepositoryInteractor{}
	h := myhandler.Handler{}
	r := Router{IDType: idType}
	r.Route(&h, &mo)

	for _, tc := range testCases {
		tc.path = handler.TournamentPath + "/" + tc.requestID
		tc.method = http.MethodDelete

		if !tc.noMock {
			mo.On(methodNameDeleteTournament, tc.requestID).Return(tc.resultErr)
		}

		handleTestCase(t, &h, tc)
	}

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}

// TestJoinTournamentHandler tests joining tournament.
func TestJoinTournamentHandler(t *testing.T) {
	idFactory := myuuid.IDFactory{}
	requestUserID := idFactory.NewString()
	testCases := []testCase{
		{
			caseName:       "everything ok",
			requestID:      idFactory.NewString(),
			requestBody:    `{"userId": "` + requestUserID + `"}`,
			expectedStatus: http.StatusOK,
		},
		{
			caseName:       "wrong body",
			noMock:         true,
			requestBody:    `i'm the wrong body"`,
			requestID:      idFactory.NewString(),
			expectedStatus: http.StatusBadRequest,
		},
		{
			caseName:       "wrong tournament error",
			resultErr:      errors.New("i'm the bad err"),
			requestID:      idFactory.NewString(),
			requestBody:    `{"userId": "` + requestUserID + `"}`,
			expectedStatus: http.StatusBadRequest,
		},
	}

	idType := myuuid.IDType{}
	mo := mockedRepositoryInteractor{}
	h := myhandler.Handler{}
	r := Router{IDType: idType}
	r.Route(&h, &mo)

	for _, tc := range testCases {
		tc.path = handler.TournamentPath + "/" + tc.requestID + handler.JoinTournamentPath
		tc.method = http.MethodPost

		if !tc.noMock {
			mo.On(methodNameJoinTournament, tc.requestID, requestUserID).Return(tc.resultErr)
		}

		handleTestCase(t, &h, tc)
	}

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}

// TestFinishTournamentHandler tests joining tournament.
func TestFinishTournamentHandler(t *testing.T) {
	idFactory := myuuid.IDFactory{}
	testCases := []testCase{
		{
			caseName:       "everything ok",
			requestID:      idFactory.NewString(),
			expectedStatus: http.StatusOK,
		},
		{
			caseName:       "wrong tournament error",
			resultErr:      errors.New("i'm the bad err"),
			requestID:      idFactory.NewString(),
			expectedStatus: http.StatusBadRequest,
		},
	}

	idType := myuuid.IDType{}
	mo := mockedRepositoryInteractor{}
	h := myhandler.Handler{}
	r := Router{IDType: idType}
	r.Route(&h, &mo)

	for _, tc := range testCases {
		tc.path = handler.TournamentPath + "/" + tc.requestID + handler.FinishTournamentPath
		tc.method = http.MethodPost

		if !tc.noMock {
			mo.On(methodNameFinishTournament, tc.requestID).Return(tc.resultErr)
		}

		handleTestCase(t, &h, tc)
	}

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}
