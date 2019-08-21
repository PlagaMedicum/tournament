package postgresql

import (
	"testing"
	"tournament/pkg/domain/tournament"
	"tournament/pkg/domain/user"
)

// TestCase contains all necessary information for test cases.
type TestCase struct {
	Stop          int
	Points        int
	UserID        uint64
	TournamentID  uint64
	User          user.User
	UpdUser       user.User
	Tournament    tournament.Tournament
	UpdTournament tournament.Tournament
	NilErr        error
	Err           error
}

// Handle analyses and prints results of running a test case.
func (tc *TestCase) Handle(caseName string, got TestCase, t *testing.T) {
	if tc.NilErr != got.NilErr {
		t.Errorf("FAIL! Test case: %s.\n"+
			"unexpectd error: \"%v\"\n",
			caseName,
			got.NilErr)
		return
	}

	if tc.Err != got.Err {
		t.Errorf("FAIL! Test case: %s.\n"+
			"err:\n\texpected: \"%v\"\n\tgot: \"%v\"\n",
			caseName,
			tc.Err, got.Err)
		return
	}

	if tc.User != got.User {
		t.Errorf("FAIL! Test case: %s.\n"+
			"user:\n\texpected: %v\n\tgot: %v\n",
			caseName,
			tc.User, got.User)
		return
	}

	if tc.Tournament.ID != got.Tournament.ID {
		t.Errorf("FAIL! Test case: %s.\n"+
			"tournament ID:\n\texpected: %v\n\tgot: %v\n",
			caseName,
			tc.Tournament.ID, got.Tournament.ID)
		return
	}

	t.Logf("PASSED. Test case: %s", caseName)
}
