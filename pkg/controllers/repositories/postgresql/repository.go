package postgresql

type Row interface {
	Scan(...interface{}) error
}

type Rows interface {
	Next() bool
	Scan(...interface{}) error
}

type handler interface{
	// TODO: Get rid of QueryRow?
	QueryRow(string, ...interface{}) Row
	Query(string, ...interface{}) (Rows, error)
	Exec(string, ...interface{}) (interface{}, error)
}

type ID interface {
	String() string
	UnmarshalText([]byte) error
}

type NullableID interface{
	IsValid() bool
	GetPointer() interface{}
	GetNotNullPointer() ID
	ID
}

type IDFactory interface {
	NewNullable() NullableID
	New() ID
}

type PSQLController struct {
	Handler   handler
	IDFactory IDFactory
}
