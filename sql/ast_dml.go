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
 * - http://dev.mysql.com/doc/refman/5.7/en/insert.html
 ********************************/
func (*Insert) Statement() {}
func (i *Insert) HasISelect() bool {
	if i.InsertFields == nil {
		return false
	}

	if _, ok := i.InsertFields.(ISelect); !ok {
		return false
	}

	return true
}

func (i *Insert) GetSchemas() []string {
	ret := i.Table.GetSchemas()
	var s []string = nil
	if i.HasISelect() {
		s = i.InsertFields.(*Select).GetSchemas()
	}

	if ret == nil || len(ret) == 0 {
		return s
	}

	if s == nil || len(s) == 0 {
		return ret
	}

	return append(ret, s...)
}

type Insert struct {
	Table ISimpleTable
	// can be `values(x,y,z)` list or `select` statement
	InsertFields interface{}
}

/*********************************
 * Update Clause
 * - http://dev.mysql.com/doc/refman/5.7/en/update.html
 ********************************/
func (*Update) Statement() {}
func (u *Update) GetSchemas() []string {
	if u.Tables == nil {
		panic("update must have table identifier")
	}

	return u.Tables.GetSchemas()
}

type Update struct {
	Tables ITables
}

/*********************************
 * Delete Clause
 ********************************/
func (*Delete) Statement() {}

type Delete struct {
	Tables ITables
}

func (d *Delete) GetSchemas() []string {
	if d.Tables == nil || len(d.Tables) == 0 {
		return nil
	}
	return d.Tables.GetSchemas()
}

/***********************************************
 * Replace Clause
 **********************************************/
func (*Replace) Statement() {}
func (r *Replace) HasISelect() bool {
	if r.ReplaceFields == nil {
		return false
	}

	if _, ok := r.ReplaceFields.(ISelect); !ok {
		return false
	}

	return true
}
func (r *Replace) GetSchemas() []string {
	ret := r.Table.GetSchemas()
	var s []string = nil
	if r.HasISelect() {
		s = r.ReplaceFields.(*Select).GetSchemas()
	}

	if ret == nil || len(ret) == 0 {
		return s
	}

	if s == nil || len(s) == 0 {
		return ret
	}

	return append(ret, s...)
}

type Replace struct {
	Table ITable
	// can be `values(x,y,z)` list or `select` statement
	ReplaceFields interface{}
}

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
