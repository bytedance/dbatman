package sql

type Create interface {
	Statement
}

type CreateTable struct{}
type CreateIndex struct{}
type CreateDatabase struct{}
type CreateView struct{}
type CreateUser struct{}
type CreateLog struct{}
type CreateTablespace struct{}
type CreateServer struct{}
