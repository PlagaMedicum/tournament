package tests

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"testing"
	"tournament/env/myhandler"
	"tournament/pkg/router"
	tournament "tournament/pkg/tournament/model"
)

// TestCreateTournamentHandler tests creation of tournament.
func TestCreateTournamentHandler(t *testing.T) {
	trList := []tester{
		{
			caseName: "everything ok",
			resultErr: nil,
			resultID: uuid.NewV1(),
			requestTournament: tournament.Tournament{
				Name: "Unreal Tournament",
				Deposit: 10000,
			},
			requestBody: []byte(`{"name": "Unreal Tournament", "deposit": 10000}`),
			expectedStatus: http.StatusCreated,
		},
		{
			caseName: "wrong body",
			requestBody: []byte(`i'm the wrong body"`),
			requestTournament: tournament.Tournament{},
			expectedStatus: http.StatusBadRequest,
		},
		{
			caseName: "wrong tournament error",
			resultErr: errors.New("i'm the bad err"),
			resultID: uuid.NewV1(),
			requestTournament: tournament.Tournament{
				Name: "test tour",
				Deposit: 1,
			},
			requestBody: []byte(`{"name": "test tour", "deposit": 1}`),
			expectedStatus: http.StatusBadRequest,
		},
	}

	mo := mockedTournament{}
	h := myhandler.Handler{}
	router.RouteForTournament(&h, &mo)

	for _, tr := range trList {
		tr.path = router.TournamentPath
		tr.method = http.MethodPost

		if tr.requestTournament.Name != "" {
			mo.On("CreateTournament", tr.requestTournament).Return(tr.resultID, tr.resultErr)
		}

		handleTester(t, &h, tr)
	}

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}

// TestGetTournamentHandler tests getting of tournament's information.
func TestGetTournamentHandler(t *testing.T) {
	trList := []tester{
		{
			caseName: "everything ok",
			resultErr: nil,
			resultTournament: tournament.Tournament{
				Name: "test tour",
			},
			requestID: uuid.NewV1(),
			expectedStatus: http.StatusOK,
		},
		{
			caseName: "wrong tournament error",
			resultErr: errors.New("i'm the bad err"),
			resultTournament: tournament.Tournament{},
			requestID: uuid.NewV1(),
			expectedStatus: http.StatusBadRequest,
		},
	}

	mo := mockedTournament{}
	h := myhandler.Handler{}
	router.RouteForTournament(&h, &mo)

	for _, tr := range trList {
		tr.path = router.TournamentPath + "/" + tr.requestID.String()
		tr.method = http.MethodGet

		if tr.requestID != uuid.Nil {
			mo.On("GetTournament", tr.requestID).Return(tr.resultTournament, tr.resultErr)
		}

		handleTester(t, &h, tr)
	}

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}

// TestGetTournamentHandler tests deleting of tournament.
func TestDeleteTournamentHandler(t *testing.T) {
	trList := []tester{
		{
			caseName: "everything ok",
			resultErr: nil,
			requestID: uuid.NewV1(),
			expectedStatus: http.StatusOK,
		},
		{
			caseName: "wrong tournament error",
			resultErr: errors.New("i'm the bad err"),
			requestID: uuid.NewV1(),
			expectedStatus: http.StatusBadRequest,
		},
	}

	mo := mockedTournament{}
	h := myhandler.Handler{}
	router.RouteForTournament(&h, &mo)

	for _, tr := range trList {
		tr.path = router.TournamentPath + "/" + tr.requestID.String()
		tr.method = http.MethodDelete

		if tr.requestID != uuid.Nil {
			mo.On("DeleteTournament", tr.requestID).Return(tr.resultErr)
		}

		handleTester(t, &h, tr)
	}

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}

// TestJoinTournamentHandler tests joining tournament.
func TestJoinTournamentHandler(t *testing.T) {
	requestUserID := uuid.NewV1()
	trList := []tester{
		{
			caseName: "everything ok",
			resultErr: nil,
			requestID: uuid.NewV1(),
			requestBody: []byte(`{"userId": "`+requestUserID.String()+`"}`),
			expectedStatus: http.StatusOK,
		},
		{
			caseName: "wrong body",
			requestBody: []byte(`i'm the wrong body"`),
			requestID: uuid.Nil,
			expectedStatus: http.StatusBadRequest,
		},
		{
			caseName: "wrong tournament error",
			resultErr: errors.New("i'm the bad err"),
			requestID: uuid.NewV1(),
			requestBody: []byte(`{"userId": "`+requestUserID.String()+`"}`),
			expectedStatus: http.StatusBadRequest,
		},
	}

	mo := mockedTournament{}
	h := myhandler.Handler{}
	router.RouteForTournament(&h, &mo)

	for _, tr := range trList {
		tr.path = router.TournamentPath + "/" + tr.requestID.String() + router.JoinTournamentPath
		tr.method = http.MethodPost

		if tr.requestID != uuid.Nil {
			mo.On("JoinTournament", tr.requestID, requestUserID).Return(tr.resultErr)
		}

		handleTester(t, &h, tr)
	}

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}

// TestFinishTournamentHandler tests joining tournament.
func TestFinishTournamentHandler(t *testing.T) {
	trList := []tester{
		{
			caseName: "everything ok",
			resultErr: nil,
			requestID: uuid.NewV1(),
			expectedStatus: http.StatusOK,
		},
		{
			caseName: "wrong tournament error",
			resultErr: errors.New("i'm the bad err"),
			requestID: uuid.NewV1(),
			expectedStatus: http.StatusBadRequest,
		},
	}

	mo := mockedTournament{}
	h := myhandler.Handler{}
	router.RouteForTournament(&h, &mo)

	for _, tr := range trList {
		tr.path = router.TournamentPath + "/" + tr.requestID.String() + router.FinishTournamentPath
		tr.method = http.MethodPost

		if tr.requestID != uuid.Nil {
			mo.On("FinishTournament", tr.requestID).Return(tr.resultErr)
		}

		handleTester(t, &h, tr)
	}

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}
