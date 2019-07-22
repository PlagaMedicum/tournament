package pkg

import (
	"net/http"
	database "tournament/database/model"
	"tournament/env/myhandler"
)

const (
	// UUIDRegex is a regular expression for uuid.
	UUIDRegex = "[0-9+a-z]{8}-[0-9+a-z]{4}-[0-9+a-z]{4}-[0-9+a-z]{4}-[0-9+a-z]{12}"
)

// DB is main instance of database using in the application
// for processing queries.
var DB database.DB

type IDInterface interface {
	String() string
	FromString(string)
}

// InitServer initialises server.
func InitServer(h *myhandler.Handler) *http.Server {
	s := &http.Server{
		Addr:    ":8080",
		Handler: h,
	}

	return s
}
