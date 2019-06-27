package errproc

import (
	"fmt"
	"github.com/jackc/pgx"
	"os"
)

func FprintErr(message string, a ...interface{}) (n int) {
	if a != nil {
		var err error
		n, err = fmt.Fprintf(os.Stderr, message, a)
		if err != nil {
			fmt.Printf("Unexpected error trying to write an error: %v\n", err)
		}
		return
	}
	return 0
}

func HandleJSONErr(err error) {
	FprintErr("Unexpected json error: %v\n", err)
}

func HandleSQLErr(action string, err error) {
	if pgerr, ok := err.(pgx.PgError); ok {
		FprintErr("Unexpected postgresql error trying to %s: %v\n", action, pgerr)
	} else {
		FprintErr("Unexpected error trying to %s: %v\n", action, err)
	}
}
