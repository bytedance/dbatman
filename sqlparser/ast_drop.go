package sql

func (*DropTables) IStatement()    {}
func (*DropTables) IDDLStatement() {}
func (*DropTables) HasDDLSchemas() {}
func (d *DropTables) GetSchemas() []string {
	return d.Tables.GetSchemas()
}

type DropTables struct {
	Tables ISimpleTables
}

func (*DropIndex) IStatement()    {}
func (*DropIndex) IDDLStatement() {}
func (*DropIndex) HasDDLSchemas() {}
func (d *DropIndex) GetSchemas() []string {
	return d.On.GetSchemas()
}

type DropIndex struct {
	On ISimpleTable
}

type DropDatabase struct{}

func (*DropDatabase) IStatement()    {}
func (*DropDatabase) IDDLStatement() {}

func (*DropFunction) IStatement()    {}
func (*DropFunction) IDDLStatement() {}
func (*DropFunction) HasDDLSchemas() {}
func (d *DropFunction) GetSchemas() []string {
	return d.Function.GetSchemas()
}

type DropFunction struct {
	Function *Spname
}

func (*DropProcedure) IStatement()    {}
func (*DropProcedure) IDDLStatement() {}
func (*DropProcedure) HasDDLSchemas() {}
func (d *DropProcedure) GetSchemas() []string {
	return d.Procedure.GetSchemas()
}

type DropProcedure struct {
	Procedure *Spname
}

type DropView struct{}

func (*DropView) IStatement()    {}
func (*DropView) IDDLStatement() {}

func (*DropTrigger) IStatement()    {}
func (*DropTrigger) IDDLStatement() {}
func (*DropTrigger) HasDDLSchemas() {}
func (d *DropTrigger) GetSchemas() []string {
	return d.Trigger.GetSchemas()
}

type DropTrigger struct {
	Trigger *Spname
}

func (*DropTablespace) IStatement()    {}
func (*DropTablespace) IDDLStatement() {}

type DropTablespace struct{}

func (*DropLogfile) IStatement()    {}
func (*DropLogfile) IDDLStatement() {}

type DropLogfile struct{}

func (*DropServer) IStatement()    {}
func (*DropServer) IDDLStatement() {}

type DropServer struct{}

func (*DropEvent) IStatement()    {}
func (*DropEvent) IDDLStatement() {}
func (*DropEvent) HasDDLSchemas() {}
func (d *DropEvent) GetSchemas() []string {
	return d.Event.GetSchemas()
}

type DropEvent struct {
	Event *Spname
}
