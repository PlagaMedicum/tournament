package main

import (
	"log"
	"net/http"
	tournamentHandlers "tournament/pkg/controllers/api/http_handlers/tournament"
	userHandlers "tournament/pkg/controllers/api/http_handlers/user"
	tournamentRouter "tournament/pkg/controllers/api/routers/tournament"
	userRouter "tournament/pkg/controllers/api/routers/user"
	tournamentRepository "tournament/pkg/controllers/repositories/postgresql/tournament"
	userRepository "tournament/pkg/controllers/repositories/postgresql/user"
	"tournament/pkg/infrastructure/databases/postgresql"
	_ "tournament/pkg/infrastructure/databases/postgresql"
	"tournament/pkg/infrastructure/handler"
	tournamentUsecases "tournament/pkg/usecases/tournament"
	userUsecases "tournament/pkg/usecases/user"
)

func main() {
	db := postgresql.DB{}
	db.InitNewPostgresDB()
	db.MigrateTablesUp()

	h := handler.Handler{}

	userRouter.RouteUser(&h, &userHandlers.Controller{
		Usecases: &userUsecases.Controller{
			Repository: &userRepository.Controller{
				Database: db.Conn,
			},
		},
	})

	tournamentRouter.RouteTournament(&h, &tournamentHandlers.Controller{
		Usecases: &tournamentUsecases.Controller{
			Repository: &tournamentRepository.Controller{
				Database: db.Conn,
			},
		},
	})

	s := http.Server{
		Addr:    ":8080",
		Handler: &h,
	}

	err := s.ListenAndServe()
	if err != nil {
		log.Printf("Unexpected http server error: " + err.Error())
	}
}
