package usecases

import (
	"testing"
	"tournament/pkg/domain"
)

type testCase struct {
	caseName            string
	mockingStop         int
	requestUserID       string
	requestTournamentID string
	requestName         string
	requestDeposit      int
	requestPoints       int
	resultID            string
	resultUser          domain.User
	resultTournament    domain.Tournament
	resultErr           error
}

func handleTestCase(expectedTC testCase, gotTC testCase, t *testing.T) {
	if expectedTC.resultID != gotTC.resultID {
		t.Errorf("FAIL! Test case: %s.\n"+
			"resultID:\n\texpected: \"%s\"\n\tgot: \"%s\"\n",
			expectedTC.caseName,
			expectedTC.resultID, gotTC.resultID)

		return
	}

	if expectedTC.resultUser != gotTC.resultUser {
		t.Errorf("FAIL! Test case: %s.\n"+
			"resultUser:\n\texpected: %v\n\tgot: %v\n",
			expectedTC.caseName,
			expectedTC.resultUser, gotTC.resultUser)

		return
	}

	if expectedTC.resultTournament.ID != gotTC.resultTournament.ID {
		t.Errorf("FAIL! Test case: %s.\n"+
			"resultTournament ID:\n\texpected: %v\n\tgot: %v\n",
			expectedTC.caseName,
			expectedTC.resultTournament.ID, gotTC.resultTournament.ID)

		return
	}

	t.Logf("PASSED. Test case: %s", expectedTC.caseName)
	return
}
