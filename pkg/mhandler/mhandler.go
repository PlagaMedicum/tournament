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

func (mHandler *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range mHandler.routes {
		if b, _ := regexp.MatchString(route.pattern, r.URL.Path); b && r.Method == route.method {
			w.Header().Set("Content-Type", "application/json")
			route.handler.ServeHTTP(w, r)
			return
		}
	}
	http.NotFound(w, r)
}
