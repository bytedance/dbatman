package sql

func (*StartTrans) IStatement() {}
func (*Lock) IStatement()       {}
func (*Unlock) IStatement()     {}
func (*Begin) IStatement()      {}
func (*Commit) IStatement()     {}
func (*Rollback) IStatement()   {}
func (*XA) IStatement()         {}
func (*SavePoint) IStatement()  {}
func (*Release) IStatement()    {}
func (*SetTrans) IStatement()   {}

type StartTrans struct{}

func (l *Lock) GetSchemas() []string {
	return l.Tables.GetSchemas()
}

type Lock struct {
	Tables ISimpleTables
}

type Unlock struct{}

type Begin struct{}

type Commit struct{}

type Rollback struct {
	Point []byte
}

type XA struct{}

type SavePoint struct {
	Point []byte
}

type Release struct {
	Point []byte
}

type SetTrans struct{}
