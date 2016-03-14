package sql

func (*CreateTable) IStatement()    {}
func (*CreateTable) IDDLStatement() {}
func (*CreateTable) HasDDLSchemas() {}

func (c *CreateTable) GetSchemas() []string {
	return c.Table.GetSchemas()
}

type CreateTable struct {
	Table ISimpleTable
}

func (*CreateIndex) IStatement()    {}
func (*CreateIndex) IDDLStatement() {}

type CreateIndex struct{}

/****************************
 * Create Database Statement
 ***************************/
func (*CreateDatabase) IStatement()    {}
func (*CreateDatabase) IDDLStatement() {}

type CreateDatabase struct{}

func (*CreateView) IStatement()    {}
func (*CreateView) IDDLStatement() {}
func (*CreateView) HasDDLSchemas() {}

type CreateView struct {
	View ISimpleTable
	As   ISelect
}

func (c *CreateView) GetSchemas() []string {
	return GetSchemas(c.View.GetSchemas(), c.As.GetSchemas())
}

func (*CreateLog) IStatement()    {}
func (*CreateLog) IDDLStatement() {}

type CreateLog struct{}

func (*CreateTablespace) IStatement()    {}
func (*CreateTablespace) IDDLStatement() {}

type CreateTablespace struct{}

func (*CreateServer) IStatement()    {}
func (*CreateServer) IDDLStatement() {}

type CreateServer struct{}

/**********************
 * Create Event Statement
 * http://dev.mysql.com/doc/refman/5.7/en/create-event.html
 *********************/
func (*CreateEvent) IStatement()    {}
func (*CreateEvent) IDDLStatement() {}
func (*CreateEvent) HasDDLSchemas() {}

type CreateEvent struct {
	Event ISimpleTable
}

func (c *CreateEvent) GetSchemas() []string {
	return c.Event.GetSchemas()
}

type eventTail struct {
	Event ISimpleTable
}

func (*CreateProcedure) IStatement()    {}
func (*CreateProcedure) IDDLStatement() {}
func (*CreateProcedure) HasDDLSchemas() {}

type CreateProcedure struct {
	Procedure ISimpleTable
}

func (c *CreateProcedure) GetSchemas() []string {
	return c.Procedure.GetSchemas()
}

type spTail struct {
	Procedure ISimpleTable
}

func (*CreateFunction) IStatement()    {}
func (*CreateFunction) IDDLStatement() {}
func (*CreateFunction) HasDDLSchemas() {}

type CreateFunction struct {
	Function ISimpleTable
}
type sfTail struct {
	Function ISimpleTable
}

func (c *CreateFunction) GetSchemas() []string {
	return c.Function.GetSchemas()
}

func (*CreateTrigger) IStatement()    {}
func (*CreateTrigger) IDDLStatement() {}
func (*CreateTrigger) HasDDLSchemas() {}

type CreateTrigger struct {
	Trigger ISimpleTable
}
type triggerTail struct {
	Trigger ISimpleTable
}

func (c *CreateTrigger) GetSchemas() []string {
	return c.Trigger.GetSchemas()
}
