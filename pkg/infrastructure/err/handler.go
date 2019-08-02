package err

import (
	"encoding/json"
	"errors"
	"github.com/jackc/pgx"
	"log"
	"net/http"
	"tournament/pkg/usecases"
)

// HandleJSONErr handles json errors.
func HandleJSONErr(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)

	err = json.NewEncoder(w).Encode("Unexpected json encoding/decoding error: "+err.Error())
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

func writeStatus(err error, w http.ResponseWriter) {
	if err == pgx.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if pgerr, ok := err.(pgx.PgError); ok {
		w.WriteHeader(http.StatusBadRequest)
		encodeErrInJSON(errors.New("Unexpected postgresql error: "+pgerr.Error()), w)
		return
	}

	if (err == usecases.ErrNotEnoughPoints) || (err == usecases.ErrFinishedTournament) || (err == usecases.ErrParticipantExists) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}

// HandleErr handles other types of errors.
func HandleErr(err error, w http.ResponseWriter) {
	writeStatus(err, w)
	encodeErrInJSON(err, w)
}
