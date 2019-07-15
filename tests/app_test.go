package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	app "tournament/pkg"
)

func TestCreateUser(t *testing.T) {
	body := []byte(`{"name": "Johnny"}`)
	r, err := http.NewRequest("POST", "localhost:8080/user", bytes.NewBuffer(body))
	if err != nil {
		t.Errorf("Test failed! Cannot encode request: '%v'", err)
	}
	w := httptest.NewRecorder()
	app.Handler.ServeHTTP(w, r)
	var respStr string
	err = json.Unmarshal(w.Body.Bytes(), respStr)
	if err != nil {
		t.Errorf("Test failed! Cannot unmarshal response body: %v\nResponse body: '%s'", err, w.Body.String())
	}
	if b, _ := regexp.MatchString(respStr, app.UUIDRegex); !b {
		t.Errorf("Test failed! Response string doesn't correspond UUID regexp.\nResponse string: '%s'", respStr)
	}
	if w.Code != http.StatusCreated {
		t.Errorf("Test failed! Wrong status code.\nExpected: %v\nGot: %v", http.StatusCreated, w.Code)
	}
}
