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
	CaseName      string
	Path          string
	Method        string
	NoMock        bool
	ReqBody       string
	ReqTournament tournament.Tournament
	ResTournament tournament.Tournament
	ReqUser       user.User
	ResUser       user.User
	ReqID         uint64
	ResID         uint64
	ResErr        error
	ResStatus     int
}

func (tc *TestCase) Handle(t *testing.T, h *handler.Handler) {
	r, err := http.NewRequest(
		tc.Method, tc.Path, bytes.NewBuffer([]byte(tc.ReqBody)))
	if err != nil {
		t.Errorf("Cannot encode request: '%v'", err)
	}

	w := httptest.NewRecorder()

	h.ServeHTTP(w, r)

	if w.Code != tc.ResStatus {
		t.Errorf("FAIL! Test case: %s. Wrong status code.\nURL-Path: %v\nExpected: %v\nGot: %v\nResponse body is: \n'''\n%s\n'''\nTester info: %+v",
			tc.CaseName, r.URL.Path, tc.ResStatus, w.Code, w.Body, tc)
		return
	}

	t.Logf("PASSED. Test case: %s", tc.CaseName)
}
