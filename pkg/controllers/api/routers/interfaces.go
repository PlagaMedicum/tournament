package routers

import "net/http"

type Handler interface {
	Handle(string, http.Handler, string)
	HandleFunc(string, func(http.ResponseWriter, *http.Request), string)
	http.Handler
}
