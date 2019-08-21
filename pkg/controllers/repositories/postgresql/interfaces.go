package postgresql

type Row interface {
	Scan(...interface{}) error
}

type Rows interface {
	Next() bool
	Scan(...interface{}) error
}

type Database interface {
	QueryRow(string, ...interface{}) Row
	Query(string, ...interface{}) (Rows, error)
	Exec(string, ...interface{}) (interface{}, error)
	ErrNoRows() error
}
