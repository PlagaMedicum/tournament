package routers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"tournament/pkg/domain/tournament"
	"tournament/pkg/domain/user"
	"tournament/pkg/infrastructure/handler"
)

type TestCase struct {
	CaseName          string
	Path              string
	Method            string
	NoMock            bool
	ResultErr         error
	ResultID          uint64
	ResultTournament  tournament.Tournament
	RequestTournament tournament.Tournament
	ResultUser        user.User
	RequestUser       user.User
	RequestID         uint64
	RequestBody       string
	ExpectedStatus    int
}

func HandleTestCase(t *testing.T, h *handler.Handler, tc TestCase) {
	r, err := http.NewRequest(
		tc.Method, tc.Path, bytes.NewBuffer([]byte(tc.RequestBody)))
	if err != nil {
		t.Errorf("Cannot encode request: '%v'", err)
	}

	w := httptest.NewRecorder()

	h.ServeHTTP(w, r)

	if w.Code != tc.ExpectedStatus {
		t.Errorf("FAIL! Test case: %s. Wrong status code.\nURL-Path: %v\nExpected: %v\nGot: %v\nResponse body is: \n'''\n%s\n'''\nTester info: %+v",
			tc.CaseName, r.URL.Path, tc.ExpectedStatus, w.Code, w.Body, tc)
		return
	}

	t.Logf("PASSED. Test case: %s", tc.CaseName)
}
