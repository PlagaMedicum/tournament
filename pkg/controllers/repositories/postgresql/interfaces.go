package postgresql

// Row is interface for postgresql query row.
type Row interface {
	Scan(...interface{}) error
}

// Rows is interface for postgresql query rows.
type Rows interface {
	Next() bool
	Scan(...interface{}) error
}

// Database is interface for querying postgresql server.
type Database interface {
	QueryRow(string, ...interface{}) Row
	Query(string, ...interface{}) (Rows, error)
	Exec(string, ...interface{}) (interface{}, error)
	ErrNoRows() error
}
