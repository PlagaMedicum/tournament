package user

import "net/http"

type httpController interface {
	CreateUserHandler(http.ResponseWriter, *http.Request)
	GetUserHandler(http.ResponseWriter, *http.Request)
	DeleteUserHandler(http.ResponseWriter, *http.Request)
	TakePointsHandler(http.ResponseWriter, *http.Request)
	GivePointsHandler(http.ResponseWriter, *http.Request)
}
