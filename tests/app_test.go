package tests

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	app "tournament/pkg"
)

func TestCreateUser(t *testing.T) {
	r, _ := http.NewRequest("POST", "/users", nil)
	w := httptest.NewRecorder()
	app.Handler.ServeHTTP(w, r)
	// b := json.Unmarshal(w.Body.Bytes())
	if b, _ := regexp.MatchString(w.Body.String(), app.UUIDRegex); !b {
		t.Errorf("Test failed, response body doesn't correspond UUID regexp.\nResponse body: '%s'", w.Body.String())
	}
}
