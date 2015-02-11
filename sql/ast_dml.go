package sql

type ISelect interface {
	Select()
}

func (*Select) ISelect() {}
func (*Union) ISelect()  {}

func (*Select) Statement() {}
func (*Union) Statement()  {}

type Union struct {
	Left, Right ISelect
}

type Select struct {
}

type Insert struct {
	Table ISimpleTable
}

func (*Insert) Statement() {}

type Update struct{}

func (*Update) Statement() {}

/***********************************************
 * Delete Clause
 **********************************************/
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
