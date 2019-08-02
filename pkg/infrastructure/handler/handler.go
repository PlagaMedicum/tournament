package handler

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

func (h *Handler) Handle(pattern string, handler http.Handler, method string) {
	h.routes = append(h.routes, &route{pattern, handler, method})
}

func (h *Handler) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request), method string) {
	h.routes = append(h.routes, &route{pattern, http.HandlerFunc(handler), method})
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range h.routes {
		if b, _ := regexp.MatchString(route.pattern, r.URL.Path); b && r.Method == route.method {
			w.Header().Set("Content-Type", "application/json")
			route.handler.ServeHTTP(w, r)

			return
		}
	}
	http.NotFound(w, r)
}
