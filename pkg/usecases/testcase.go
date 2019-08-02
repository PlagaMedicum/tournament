package usecases

import (
	"testing"
	"tournament/pkg/domain/tournament"
	"tournament/pkg/domain/user"
)

type TestCase struct {
	MockingStop    int
	UserID         uint64
	TournamentID   uint64
	RequestName    string
	RequestDeposit int
	RequestPoints  int
	ResultID       uint64
	User           user.User
	Tournament     tournament.Tournament
	Err            error
}

func HandleTestCase(caseName string, expectedTC TestCase, gotTC TestCase, t *testing.T) {
	if expectedTC.Err != gotTC.Err {
		t.Errorf("FAIL! Test case: %s.\n"+
			"err:\n\texpected: \"%v\"\n\tgot: \"%v\"\n",
			caseName,
			expectedTC.Err, gotTC.Err)

		return
	}

	if expectedTC.ResultID != gotTC.ResultID {
		t.Errorf("FAIL! Test case: %s.\n"+
			"resultID:\n\texpected: \"%d\"\n\tgot: \"%d\"\n",
			caseName,
			expectedTC.ResultID, gotTC.ResultID)

		return
	}

	if expectedTC.User != gotTC.User {
		t.Errorf("FAIL! Test case: %s.\n"+
			"user:\n\texpected: %v\n\tgot: %v\n",
			caseName,
			expectedTC.User, gotTC.User)

		return
	}

	if expectedTC.Tournament.ID != gotTC.Tournament.ID {
		t.Errorf("FAIL! Test case: %s.\n"+
			"tournament ID:\n\texpected: %v\n\tgot: %v\n",
			caseName,
			expectedTC.Tournament.ID, gotTC.Tournament.ID)

		return
	}

	t.Logf("PASSED. Test case: %s", caseName)
	return
}
