package sql

type Executor interface {
	Exec(query string, args ...interface{}) (Result, error)
	Query(query string, args ...interface{}) (Rows, error)
}
