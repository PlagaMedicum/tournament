package tournament

import (
	"testing"
	"tournament/pkg/controllers/repositories/postgresql"
	"tournament/pkg/controllers/repositories/postgresql/user"
	"tournament/pkg/domain/tournament"
	userDomain "tournament/pkg/domain/user"
	database "tournament/pkg/infrastructure/databases/postgresql"
)

func TestCRURDTournament(t *testing.T) {
	testCases := map[string]postgresql.TestCase{
		"everything ok": {
			Points: 200,
			User: userDomain.User{
				ID:      1,
				Name:    "Mehmede",
			},
			Tournament: tournament.Tournament{
				ID:      1,
				Name:    "Tournament",
				Deposit: 100,
				Prize: 1000,
			},
			UpdTournament: tournament.Tournament{
				ID:           1,
				Name:         "Tournament",
				Deposit:      100,
				Prize:        1000,
				Participants: []uint64{
					1,
				},
				WinnerID: 1,
			},
		},
	}

	db := database.DB{}
	db.InitNewPostgresDB()
	c := Controller{Database: db.Conn}
	userController := user.Controller{Database: db.Conn}

	for name, tc := range testCases {
		db.MigrateTablesDown()
		db.MigrateTablesUp()

		got := tc

		got.Tournament.ID, got.NilErr = c.InsertTournament(tc.Tournament.Name, tc.Tournament.Deposit, tc.Tournament.Prize)
		tc.Handle(name, got, t)

		got.Tournament, got.NilErr = c.GetTournamentByID(tc.Tournament.ID)
		tc.Handle(name, got, t)

		got.User.ID, got.NilErr = userController.InsertUser(tc.User.Name, tc.User.Balance)
		tc.Handle(name, got, t)

		got.NilErr = c.AddUserInTournament(tc.User.ID, tc.Tournament.ID)
		tc.Handle(name, got, t)

		got.NilErr = c.SetWinner(tc.User.ID, tc.Tournament.ID)
		tc.Handle(name, got, t)

		got.UpdTournament, got.NilErr = c.GetTournamentByID(tc.Tournament.ID)
		tc.Handle(name, got, t)

		got.NilErr = c.DeleteTournamentByID(tc.Tournament.ID)
		tc.Handle(name, got, t)
	}
}
