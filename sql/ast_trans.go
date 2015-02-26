package sql

func (*StartTrans) Statement() {}
func (*Lock) Statement()       {}
func (*Unlock) Statement()     {}
func (*Begin) Statement()      {}
func (*Commit) Statement()     {}
func (*Rollback) Statement()   {}
func (*XA) Statement()         {}
func (*SavePoint) Statement()  {}
func (*Release) Statement()    {}

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
