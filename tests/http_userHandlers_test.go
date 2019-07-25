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

// TestCreateUserHandler tests creation of user.
func TestCreateUserHandler(t *testing.T) {
	idFabric := myuuid.IDFabric{}
	trList := []tester{
		{
			caseName:  "everything ok",
			resultErr: nil,
			resultID:  uuid.NewV1(),
			requestUser: domain.User{
				Name: "Daniil Dankovskij",
			},
			requestBody: []byte(`{"name": "Daniil Dankovskij"}`),
			expectedStatus: http.StatusCreated,
		},
		{
			caseName: "wrong body",
			requestBody: []byte(`i'm the wrong body"`),
			requestUser: domain.User{},
			expectedStatus: http.StatusBadRequest,
		},
		{
			caseName: "wrong user error",
			resultErr: errors.New("i'm the bad err"),
			resultID: uuid.NewV1(),
			requestUser: domain.User{
				Name: "Artemij Burah",
			},
			requestBody: []byte(`{"name": "Artemij Burah"}`),
			expectedStatus: http.StatusBadRequest,
		},
	}

	idType := myuuid.IDType{}
	mo := mockedRepositoryInteractor{}
	h := myhandler.Handler{}
	r := router.Router{IDType: idType}
	r.Route(&h, &mo)

	for _, tr := range trList {
		tr.path = handler.UserPath
		tr.method = http.MethodPost

		if tr.requestUser != (domain.User{}) {
			mo.On("CreateUser", tr.requestUser.Name).Return(tr.resultID, tr.resultErr)
		}

		handleTester(t, &h, tr)
	}

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}

// TestGetUserHandler tests getting of user's information.
func TestGetUserHandler(t *testing.T) {
	idFabric := myuuid.IDFabric{}
	trList := []tester{
		{
			caseName:  "everything ok",
			resultErr: nil,
			resultUser: domain.User{
				Name: "Anna Angel",
			},
			requestID: uuid.NewV1(),
			expectedStatus: http.StatusOK,
		},
		{
			caseName: "wrong user error",
			resultErr: errors.New("i'm the bad err"),
			resultUser: domain.User{},
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
		tr.path = handler.UserPath + "/" + tr.requestID
		tr.method = http.MethodGet

		if tr.requestID != idType.Null(){
			mo.On("GetUser", tr.requestID).Return(tr.resultUser, tr.resultErr)
		}

		handleTester(t, &h, tr)
	}

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}

// TestDeleteUserHandler tests deleting of user.
func TestDeleteUserHandler(t *testing.T) {
	idFabric := myuuid.IDFabric{}
	trList := []tester{
		{
			caseName: "everything ok",
			resultErr: nil,
			requestID: uuid.NewV1(),
			expectedStatus: http.StatusOK,
		},
		{
			caseName: "wrong user error",
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
		tr.path = handler.UserPath + "/" + tr.requestID
		tr.method = http.MethodDelete

		if tr.requestID != idType.Null(){
			mo.On("DeleteUser", tr.requestID).Return(tr.resultErr)
		}

		handleTester(t, &h, tr)
	}

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}

// TestTakePointsHandler tests taking points from user.
func TestTakePointsHandler(t *testing.T) {
	idFabric := myuuid.IDFabric{}
	trList := []tester{
		{
			caseName: "everything ok",
			resultErr: nil,
			requestID: uuid.NewV1(),
			requestBody: []byte(`{"points": 1}`),
			expectedStatus: http.StatusOK,
		},
		{
			caseName: "wrong body",
			requestBody: []byte(`i'm the wrong body"`),
			requestID: uuid.Nil,
			expectedStatus: http.StatusBadRequest,
		},
		{
			caseName: "wrong user error",
			resultErr: errors.New("i'm the bad err"),
			requestID: uuid.NewV1(),
			requestBody: []byte(`{"points": 1}`),
			expectedStatus: http.StatusBadRequest,
		},
	}

	idType := myuuid.IDType{}
	mo := mockedRepositoryInteractor{}
	h := myhandler.Handler{}
	r := router.Router{IDType: idType}
	r.Route(&h, &mo)

	for _, tr := range trList {
		tr.path = handler.UserPath + "/" + tr.requestID + handler.TakingPointsPath
		tr.method = http.MethodPost

		if tr.requestID != uuid.Nil{
			mo.On("FundUser", tr.requestID, -1).Return(tr.resultErr)
		}

		handleTester(t, &h, tr)
	}

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}

// TestGivePointsHandler tests giving points to user.
func TestGivePointsHandler(t *testing.T) {
	idFabric := myuuid.IDFabric{}
	trList := []tester{
		{
			caseName: "everything ok",
			resultErr: nil,
			requestID: uuid.NewV1(),
			requestBody: []byte(`{"points": 1}`),
			expectedStatus: http.StatusOK,
		},
		{
			caseName: "wrong body",
			requestBody: []byte(`i'm the wrong body"`),
			requestID: uuid.Nil,
			expectedStatus: http.StatusBadRequest,
		},
		{
			caseName: "wrong user error",
			resultErr: errors.New("i'm the bad err"),
			requestID: uuid.NewV1(),
			requestBody: []byte(`{"points": 1}`),
			expectedStatus: http.StatusBadRequest,
		},
	}

	idType := myuuid.IDType{}
	mo := mockedRepositoryInteractor{}
	h := myhandler.Handler{}
	r := router.Router{IDType: idType}
	r.Route(&h, &mo)

	for _, tr := range trList {
		tr.path = handler.UserPath + "/" + tr.requestID + handler.GivingPointsPath
		tr.method = http.MethodPost

		if tr.requestID != uuid.Nil{
			mo.On("FundUser", tr.requestID, 1).Return(tr.resultErr)
		}

		handleTester(t, &h, tr)
	}

	t.Logf("Asserted mocks expectations:")
	mo.AssertExpectations(t)
}
