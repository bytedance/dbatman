package sql

type IShow interface {
	IShow()
	IStatement
}

func (*Show) IStatement() {}
func (*Show) IsShow()     {}

type Show struct{}
