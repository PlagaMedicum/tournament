package postgresql

type Row interface {
	Scan(...interface{}) error
}

type Rows interface {
	Next() bool
	Scan(...interface{}) error
}

type database interface{
	// TODO: Get rid of QueryRow?
	QueryRow(string, ...interface{}) Row
	Query(string, ...interface{}) (Rows, error)
	Exec(string, ...interface{}) (interface{}, error)
}

type PSQLController struct {
	Database database
}
