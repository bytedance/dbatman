package sql

/***********************************
 * Select Clause
 ***********************************/

type ISelect interface {
	ISelect()
	IsLocked() bool
	GetSchemas() []string
	IStatement
}

func (*Select) ISelect()      {}
func (*ParenSelect) ISelect() {}
func (*Union) ISelect()       {}
func (*SubQuery) ISelect()    {}

func (*Select) IStatement()      {}
func (*ParenSelect) IStatement() {}
func (*Union) IStatement()       {}
func (*SubQuery) IStatement()    {}

type Union struct {
	Left, Right ISelect
}

func (u *Union) IsLocked() bool {
	return u.Left.IsLocked() || u.Right.IsLocked()
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

func (s *SubQuery) IsLocked() bool {
	return s.SelectStatement.IsLocked()
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

func (s *Select) IsLocked() bool {
	return s.LockType != LockType_NoLock
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

// ParenSelect ------
type ParenSelect struct {
	Select ISelect
}

func (p *ParenSelect) IsLocked() bool {
	return p.Select.IsLocked()
}

func (p *ParenSelect) GetSchemas() []string {
	return p.Select.GetSchemas()
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
func (*Insert) IStatement() {}
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
func (*Update) IStatement() {}
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
func (*Delete) IStatement() {}

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
func (*Replace) IStatement() {}
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

func (*Call) IStatement() {}

type Do struct{}

func (*Do) IStatement() {}

type Load struct{}

func (*Load) IStatement() {}

type Handler struct{}

func (*Handler) IStatement() {}
