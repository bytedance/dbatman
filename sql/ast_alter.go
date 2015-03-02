package sql

type AlterTable struct {
	Table ISimpleTable
}

func (*AlterTable) IStatement()    {}
func (*AlterTable) IDDLStatement() {}

type AlterDatabase struct {
	Schema []byte
}

func (*AlterDatabase) IStatement()    {}
func (*AlterDatabase) IDDLStatement() {}

type AlterProcedure struct {
	Procedure *Spname
}

func (*AlterProcedure) IStatement()    {}
func (*AlterProcedure) IDDLStatement() {}

type AlterFunction struct {
	Function *Spname
}

func (*AlterFunction) IStatement()    {}
func (*AlterFunction) IDDLStatement() {}

/*************************
 * Alter View Statement
 *************************/
func (*AlterView) IStatement()    {}
func (*AlterView) IDDLStatement() {}

type AlterView struct {
	View ISimpleTable
	As   ISelect
}

type viewTail struct {
	View ISimpleTable
	As   ISelect
}

func (av *AlterView) GetSchemas() []string {
	d := av.View.GetSchemas()
	p := av.As.GetSchemas()
	if d != nil && p != nil {
		d = append(d, p...)
	}

	return d
}

/*************************
 * Alter Event Statement
 *************************/
func (*AlterEvent) IStatement()    {}
func (*AlterEvent) IDDLStatement() {}
func (*AlterEvent) HasDDLSchemas() {}
func (a *AlterEvent) GetSchemas() []string {
	if a.Rename == nil {
		return a.Event.GetSchemas()
	}

	return GetSchemas(a.Event.GetSchemas(), a.Rename.GetSchemas())
}

type AlterEvent struct {
	Event  *Spname
	Rename *Spname
}

type AlterTablespace struct{}

func (*AlterTablespace) IStatement()    {}
func (*AlterTablespace) IDDLStatement() {}

type AlterLogfile struct{}

func (*AlterLogfile) IStatement()    {}
func (*AlterLogfile) IDDLStatement() {}

type AlterServer struct{}

func (*AlterServer) IStatement()    {}
func (*AlterServer) IDDLStatement() {}
