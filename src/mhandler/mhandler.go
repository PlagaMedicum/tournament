package mhandler

import (
	"net/http"
	"regexp"
)

type route struct {
	pattern string
	handler http.Handler
	method  string
}

type Handler struct {
	routes []*route
}

func (mHandler *Handler) Handle(pattern string, handler http.Handler, method string) {
	mHandler.routes = append(mHandler.routes, &route{pattern, handler, method})
}

func (mHandler *Handler) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request), method string) {
	mHandler.routes = append(mHandler.routes, &route{pattern, http.HandlerFunc(handler), method})
}

func (mHandler *Handler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	for _, route := range mHandler.routes {
		if b, _ := regexp.MatchString(route.pattern, request.URL.Path); b && request.Method == route.method {
			response.Header().Set("Content-Type", "application/json")
			route.handler.ServeHTTP(response, request)
			return
		}
	}
	http.NotFound(response, request)
}
