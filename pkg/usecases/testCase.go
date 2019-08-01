package usecases

import (
	"testing"
	"tournament/pkg/domain"
)

type testCase struct {
	mockingStop    int
	userID         uint64
	tournamentID   uint64
	requestName    string
	requestDeposit int
	requestPoints  int
	resultID       uint64
	user           domain.User
	tournament     domain.Tournament
	err            error
}

func handleTestCase(caseName string, expectedTC testCase, gotTC testCase, t *testing.T) {
	if expectedTC.err != gotTC.err {
		t.Errorf("FAIL! Test case: %s.\n"+
			"err:\n\texpected: \"%v\"\n\tgot: \"%v\"\n",
			caseName,
			expectedTC.err, gotTC.err)

		return
	}

	if expectedTC.resultID != gotTC.resultID {
		t.Errorf("FAIL! Test case: %s.\n"+
			"resultID:\n\texpected: \"%d\"\n\tgot: \"%d\"\n",
			caseName,
			expectedTC.resultID, gotTC.resultID)

		return
	}

	if expectedTC.user != gotTC.user {
		t.Errorf("FAIL! Test case: %s.\n"+
			"user:\n\texpected: %v\n\tgot: %v\n",
			caseName,
			expectedTC.user, gotTC.user)

		return
	}

	if expectedTC.tournament.ID != gotTC.tournament.ID {
		t.Errorf("FAIL! Test case: %s.\n"+
			"tournament ID:\n\texpected: %v\n\tgot: %v\n",
			caseName,
			expectedTC.tournament.ID, gotTC.tournament.ID)

		return
	}

	t.Logf("PASSED. Test case: %s", caseName)
	return
}
