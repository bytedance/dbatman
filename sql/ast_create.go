package sql

func (*CreateTable) IStatement()    {}
func (*CreateTable) HasDDLSchemas() {}

func (c *CreateTable) GetSchemas() []string {
	return c.Table.GetSchemas()
}

type CreateTable struct {
	Table ISimpleTable
}

type CreateIndex struct{}

func (*CreateIndex) IStatement() {}

/****************************
 * Create Database Statement
 ***************************/
func (*CreateDatabase) IStatement() {}

type CreateDatabase struct{}

type CreateView struct {
	View ISimpleTable
	As   ISelect
}

func (*CreateView) IStatement()    {}
func (*CreateView) HasDDLSchemas() {}

func (c *CreateView) GetSchemas() []string {
	return GetSchemas(c.View.GetSchemas(), c.As.GetSchemas())
}

type CreateLog struct{}

func (*CreateLog) IStatement() {}

type CreateTablespace struct{}

func (*CreateTablespace) IStatement() {}

type CreateServer struct{}

func (*CreateServer) IStatement() {}

/**********************
 * Create Event Statement
 * http://dev.mysql.com/doc/refman/5.7/en/create-event.html
 *********************/
func (*CreateEvent) IStatement()    {}
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
