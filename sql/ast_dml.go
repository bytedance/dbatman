package sql

/***********************************
 * Select Clause
 ***********************************/

type ISelect interface {
	IsSelect()
	GetSchemas() []string
	IStatement
}

func (*Select) IsSelect()   {}
func (*Union) IsSelect()    {}
func (*SubQuery) IsSelect() {}

func (*Select) Statement()   {}
func (*Union) Statement()    {}
func (*SubQuery) Statement() {}

type Union struct {
	Left, Right ISelect
}

func (u *Union) GetSchemas() []string {
	if u.Left == nil {
		panic("union must have left select statement")
	}

	if u.Right == nil {
		panic("union must have right select statement")
	}

	l := u.Left.GetSchemas()
	r := u.Right.GetSchemas()

	if l == nil && r == nil {
		return nil
	} else if l == nil {
		return r
	} else if r == nil {
		return l
	}
	return append(l, r...)
}

// SubQuery ---------
type SubQuery struct {
	SelectStatement ISelect
}

func (s *SubQuery) GetSchemas() []string {
	if s.SelectStatement == nil {
		panic("subquery has no content")
	}

	return s.SelectStatement.GetSchemas()
}

// Select -----------
type Select struct {
	From     ITables
	LockType LockType
}

func (s *Select) GetSchemas() []string {
	if s.From == nil {
		return nil
	}

	ret := make([]string, 0, 8)
	for _, v := range s.From {
		r := v.GetSchemas()
		if r != nil || len(r) != 0 {
			ret = append(ret, r...)
		}
	}

	return ret
}

type LockType int

const (
	LockType_NoLock = iota
	LockType_ForUpdate
	LockType_LockInShareMode
)

/*********************************
 * Insert Clause
 ********************************/
type Insert struct {
	Table ISimpleTable
}

func (*Insert) Statement() {}

type Update struct{}

func (*Update) Statement() {}

/*********************************
 * Delete Clause
 ********************************/
type IDelete interface {
	Delete()
	IStatement
}

type SingleTableDelete struct {
	Table ISimpleTable
}

func (*SingleTableDelete) Statement() {}
func (*SingleTableDelete) Delete()    {}

type MultiTableDelete struct {
	Tables []ISimpleTable
}

func (*MultiTableDelete) Statement() {}
func (*MultiTableDelete) Delete()    {}

/***********************************************
 * Replace Clause
 **********************************************/
type Replace struct {
	Table ISimpleTable
}

func (*Replace) Statement() {}

type Call struct {
	Spname *Spname
}

func (*Call) Statement() {}

type Do struct{}

func (*Do) Statement() {}

type Load struct{}

func (*Load) Statement() {}

type Handler struct{}

func (*Handler) Statement() {}
