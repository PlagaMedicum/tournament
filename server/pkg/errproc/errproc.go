package errproc

import (
	"fmt"
	"github.com/jackc/pgx"
	"os"
)

func FprintErr(message string, err error) int {
	if err != nil {
		n, e := fmt.Fprintf(os.Stderr, message, err)
		if e != nil {
			fmt.Printf("Unexpected error trying to write an error: %v\n", err)
		}
		return n
	}
	return 0
}

func HandleJSONErr(err error) {
	FprintErr("Unexpected json error: %v\n", err)
}

func HandleSQLErr(action string, err error) {
	if pgerr, ok := err.(pgx.PgError); ok {
		_, e := fmt.Fprintf(os.Stderr, "Unexpected postgresql error trying to " + action + ": %v\n", pgerr)
		if e != nil {
			fmt.Printf("Unexpected error trying to write an error: %v\n", e)
		}
	} else {
		FprintErr("Unexpected error trying to " + action + ": %v\n", err)
	}
}
