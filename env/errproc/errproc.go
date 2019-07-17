package errproc

import (
	"encoding/json"
	"errors"
	"github.com/jackc/pgx"
	"log"
	"net/http"
)

var(
	NoTournamentWithID = errors.New("no tournament with target id found")
	NoUserWithID = errors.New("no user with target id found")
	NotEnoughPoints = errors.New("user have not enough points to join the tournament")
)

// HandleJSONErr handles json errors.
func HandleJSONErr(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)

	err = json.NewEncoder(w).Encode("Unexpected json sending error: "+err.Error())
	if err != nil {
		log.Printf("Unable to encode and send error in json.")
	}
}

func encodeErrInJSON(err error, w http.ResponseWriter) {
	err = json.NewEncoder(w).Encode(err.Error())
	if err != nil {
		HandleJSONErr(err, w)
	}
}

func writePSQLErr(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	encodeErrInJSON(errors.New("Unexpected postgresql error: "+err.Error()), w)
}

func writeNotAcceptable(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotAcceptable)
	encodeErrInJSON(err, w)

}

func writeNotFound(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	encodeErrInJSON(err, w)
}

func writeBadRequest(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	encodeErrInJSON(err, w)
}

// HandleErr handles other types of errors.
func HandleErr(err error, w http.ResponseWriter) {
	if pgerr, ok := err.(pgx.PgError); ok {
		writePSQLErr(pgerr, w)

		return
	}

	if err == NotEnoughPoints {
		writeNotAcceptable(err, w)

		return
	}

	if (err == NoUserWithID) || (err == NoTournamentWithID) {
		writeNotFound(err, w)

		return
	}

	writeBadRequest(err, w)
}
