package tests

import (
	"bytes"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"net/http/httptest"
	"testing"
	"tournament/pkg/domain"
	"tournament/pkg/infrastructure/myhandler"
	user "tournament/pkg/user/model"
)

type tester struct {
	caseName          string
	path              string
	method            string
	resultErr         error
	resultID          uuid.UUID
	resultTournament  domain.Tournament
	requestTournament domain.Tournament
	resultUser        user.User
	requestUser       user.User
	requestID         uuid.UUID
	requestBody       []byte
	expectedStatus    int
}

func handleTester(t *testing.T, h *myhandler.Handler, tr tester) {
	r, err := http.NewRequest(
		tr.method, tr.path, bytes.NewBuffer(tr.requestBody))
	if err != nil {
		t.Errorf("Cannot encode request: '%v'", err)
	}

	w := httptest.NewRecorder()

	h.ServeHTTP(w, r)

	if w.Code != tr.expectedStatus {
		t.Errorf("Test case: %s. Wrong status code.\nExpected: %v\nGot: %v\nResponse body is: %s",
			tr.caseName, tr.expectedStatus, w.Code, w.Body)
		return
	}

	t.Logf("Test case PASSED: %s", tr.caseName)
}
