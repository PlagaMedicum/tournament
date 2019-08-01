package http

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"tournament/pkg/domain"
	"tournament/pkg/infrastructure/myhandler"
)

type testCase struct {
	caseName          string
	path              string
	method            string
	noMock			  bool
	resultErr         error
	resultID          uint64
	resultTournament  domain.Tournament
	requestTournament domain.Tournament
	resultUser        domain.User
	requestUser       domain.User
	requestID         uint64
	requestBody       string
	expectedStatus    int
}

func handleTestCase(t *testing.T, h *myhandler.Handler, tc testCase) {
	r, err := http.NewRequest(
		tc.method, tc.path, bytes.NewBuffer([]byte(tc.requestBody)))
	if err != nil {
		t.Errorf("Cannot encode request: '%v'", err)
	}

	w := httptest.NewRecorder()

	h.ServeHTTP(w, r)

	if w.Code != tc.expectedStatus {
		t.Errorf("FAIL! Test case: %s. Wrong status code.\nURL-Path: %v\nExpected: %v\nGot: %v\nResponse body is: \n'''\n%s\n'''\nTester info: %+v",
			tc.caseName, r.URL.Path, tc.expectedStatus, w.Code, w.Body, tc)
		return
	}

	t.Logf("PASSED. Test case: %s", tc.caseName)
}
