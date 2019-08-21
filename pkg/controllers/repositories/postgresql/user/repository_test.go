package user

import (
	"github.com/jackc/pgx"
	"testing"
	"tournament/pkg/controllers/repositories/postgresql"
	"tournament/pkg/domain/user"
	_ "tournament/pkg/infrastructure/databases/postgresql"
	database "tournament/pkg/infrastructure/databases/postgresql"
)

func TestCRURDUser(t *testing.T) {
	testCases := map[string]postgresql.TestCase{
		"everything ok": {
			Points: 200,
			User: user.User{
				ID:      1,
				Name:    "Aleksandar Makedonski",
				Balance: 1500,
			},
			UpdUser: user.User{
				ID:      1,
				Name:    "Aleksandar Makedonski",
				Balance: 1700,
			},
		},
		"error getting user by ID": {
			Stop: 2,
			UserID: 99,
			User: user.User{
				ID:      1,
				Name:    "Aleksandar Makedonski",
				Balance: 1500,
			},
			Err: pgx.ErrNoRows,
		},
	}

	db := database.DB{}
	db.InitNewPostgresDB()
	c := Controller{Database: db.Conn}

	for name, tc := range testCases {
		db.MigrateTablesDown()
		db.MigrateTablesUp()

		got := tc

		got.User.ID, got.NilErr = c.InsertUser(tc.User.Name, tc.User.Balance)
		tc.Handle(name, got, t)

		if tc.Stop == 2 {
			_, got.Err = c.GetUserByID(tc.UserID)
			tc.Handle(name, got, t)
			continue
		}
		got.User, got.NilErr = c.GetUserByID(tc.User.ID)
		tc.Handle(name, got, t)

		got.NilErr = c.UpdateUserBalanceByID(tc.User.ID, tc.Points)
		tc.Handle(name, got, t)

		got.UpdUser, got.NilErr = c.GetUserByID(tc.User.ID)
		tc.Handle(name, got, t)

		got.NilErr = c.DeleteUserByID(tc.User.ID)
		tc.Handle(name, got, t)
	}
}
