package routers

import "net/http"

// Handler contains generic set of http handler functions.
type Handler interface {
	Handle(string, http.Handler, string)
	HandleFunc(string, func(http.ResponseWriter, *http.Request), string)
	http.Handler
}
