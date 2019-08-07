package usecases

import (
	"testing"
	"tournament/pkg/domain/tournament"
	"tournament/pkg/domain/user"
)

type TestCase struct {
	MockingStop  int
	ReqName      string
	ReqDeposit   int
	ReqPoints    int
	ResID        uint64
	UserID       uint64
	TournamentID uint64
	User         user.User
	Tournament   tournament.Tournament
	Err          error
}

func (tc *TestCase) Handle(caseName string, got TestCase, t *testing.T) {
	if tc.Err != got.Err {
		t.Errorf("FAIL! Test case: %s.\n"+
			"err:\n\texpected: \"%v\"\n\tgot: \"%v\"\n",
			caseName,
			tc.Err, got.Err)
		return
	}

	if tc.ResID != got.ResID {
		t.Errorf("FAIL! Test case: %s.\n"+
			"resultID:\n\texpected: \"%d\"\n\tgot: \"%d\"\n",
			caseName,
			tc.ResID, got.ResID)
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
