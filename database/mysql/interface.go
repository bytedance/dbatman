package mysql

type Executor interface {
	Exec(query string, args ...interface{}) (Result, error)
	Query(query string, args ...interface{}) (Rows, error)
	Prepare(query string) (*Stmt, error)
}
