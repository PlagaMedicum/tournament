package tests

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"testing"
	"tournament/pkg/domain"
	"tournament/pkg/infrastructure/myhandler"
	"tournament/pkg/infrastructure/myuuid"
	handler "tournament/pkg/interfaces/handlers/http"
	router "tournament/pkg/interfaces/routers/http"
)

// TestCreateTournamentHandler tests creation of tournament.
func TestCreateTournamentHandler(t *testing.T) {
	idFabric := myuuid.IDFabric{}
	trList := []tester{
		{
			caseName:  "everything ok",
			resultErr: nil,
			resultID:  idFabric.New(),
			requestTournament: domain.Tournament{
				Name: "Unreal Tournament",
				Deposit: 10000,
			},
			requestBody: []byte(`{"name": "Unreal Tournament", "deposit": 10000}`),
			expectedStatus: http.StatusCreated,
		},
		{
			caseName:          "wrong body",
			requestBody:       []byte(`i'm the wrong body"`),
			requestTournament: domain.Tournament{},
			expectedStatus:    http.StatusBadRequest,
		},
		{
			caseName:  "wrong tournament error",
			resultErr: errors.New("i'm the bad err"),
			resultID:  idFabric.New(),
			requestTournament: domain.Tournament{
				Name: "test tour",
				Deposit: 1,
			},
			requestBody: []byte(`{"name": "test tour", "deposit": 1}`),
			expectedStatus: http.StatusBadRequest,
		},
	}

	idType := myuuid.IDType{}
	mo := mockedRepositoryInteractor{}
	h := myhandler.Handler{}
	r := router.Router{IDType: idType}
	r.Route(&h, &mo)

	for _, tr := range trList {
		tr.path = handler.TournamentPath
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
	idFabric := myuuid.IDFabric{}
	trList := []tester{
		{
			caseName: "everything ok",
			resultErr: nil,
			resultTournament: domain.Tournament{
				Name: "test tour",
			},
			requestID: uuid.NewV1(),
			expectedStatus: http.StatusOK,
		},
		{
			caseName:         "wrong tournament error",
			resultErr:        errors.New("i'm the bad err"),
			resultTournament: domain.Tournament{},
			requestID:        uuid.NewV1(),
			expectedStatus:   http.StatusBadRequest,
		},
	}

	idType := myuuid.IDType{}
	mo := mockedRepositoryInteractor{}
	h := myhandler.Handler{}
	r := router.Router{IDType: idType}
	r.Route(&h, &mo)

	for _, tr := range trList {
		tr.path = handler.TournamentPath + "/" + tr.requestID
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
	idFabric := myuuid.IDFabric{}
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

	idType := myuuid.IDType{}
	mo := mockedRepositoryInteractor{}
	h := myhandler.Handler{}
	r := router.Router{IDType: idType}
	r.Route(&h, &mo)

	for _, tr := range trList {
		tr.path = handler.TournamentPath + "/" + tr.requestID
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
	idFabric := myuuid.IDFabric{}
	requestUserID := idFabric.New()
	trList := []tester{
		{
			caseName: "everything ok",
			resultErr: nil,
			requestID: uuid.NewV1(),
			requestBody: []byte(`{"userId": "`+requestUserID+`"}`),
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
			requestBody: []byte(`{"userId": "`+requestUserID+`"}`),
			expectedStatus: http.StatusBadRequest,
		},
	}

	idType := myuuid.IDType{}
	mo := mockedRepositoryInteractor{}
	h := myhandler.Handler{}
	r := router.Router{IDType: idType}
	r.Route(&h, &mo)

	for _, tr := range trList {
		tr.path = handler.TournamentPath + "/" + tr.requestID + handler.JoinTournamentPath
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
	idFabric := myuuid.IDFabric{}
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

	idType := myuuid.IDType{}
	mo := mockedRepositoryInteractor{}
	h := myhandler.Handler{}
	r := router.Router{IDType: idType}
	r.Route(&h, &mo)

	for _, tr := range trList {
		tr.path = handler.TournamentPath + "/" + tr.requestID + handler.FinishTournamentPath
		tr.method = http.MethodPost

		if tr.requestID != idType.Null() {
			mo.On("FinishTournament", tr.requestID).Return(tr.resultErr)
		}

		handleTester(t, &h, tr)
	}

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}
