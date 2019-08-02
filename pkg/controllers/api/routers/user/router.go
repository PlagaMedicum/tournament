package user

import (
	"net/http"
	"tournament/pkg/controllers/api/http_handlers"
	"tournament/pkg/controllers/api/routers"
)

// RouteUser connects user endpoints with handling functions.
func RouteUser(h routers.Handler, c httpController) {
	h.HandleFunc("^"+http_handlers.UserPath+"$", c.CreateUserHandler, http.MethodPost)
	h.HandleFunc("^"+http_handlers.UserPath+routers.Regex+"$", c.GetUserHandler, http.MethodGet)
	h.HandleFunc("^"+http_handlers.UserPath+routers.Regex+"$", c.DeleteUserHandler, http.MethodDelete)
	h.HandleFunc("^"+http_handlers.UserPath+routers.Regex+http_handlers.TakingPointsPath+"$", c.TakePointsHandler, http.MethodPost)
	h.HandleFunc("^"+http_handlers.UserPath+routers.Regex+http_handlers.GivingPointsPath+"$", c.GivePointsHandler, http.MethodPost)
}
