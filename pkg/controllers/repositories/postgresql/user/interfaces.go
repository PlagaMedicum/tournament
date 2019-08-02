package user

import(
	"tournament/pkg/controllers/repositories/postgresql"
)

type Controller struct {
	postgresql.Database
}
