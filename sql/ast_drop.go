package sql

func (*DropTables) IStatement()    {}
func (*DropTables) HasDDLSchemas() {}
func (d *DropTables) GetSchemas() []string {
	return d.Tables.GetSchemas()
}

type DropTables struct {
	Tables ISimpleTables
}

func (*DropIndex) IStatement()    {}
func (*DropIndex) HasDDLSchemas() {}
func (d *DropIndex) GetSchemas() []string {
	return d.On.GetSchemas()
}

type DropIndex struct {
	On ISimpleTable
}

type DropDatabase struct{}

func (*DropDatabase) IStatement() {}

func (*DropFunction) IStatement()    {}
func (*DropFunction) HasDDLSchemas() {}
func (d *DropFunction) GetSchemas() []string {
	return d.Function.GetSchemas()
}

type DropFunction struct {
	Function *Spname
}

func (*DropProcedure) IStatement()    {}
func (*DropProcedure) HasDDLSchemas() {}
func (d *DropProcedure) GetSchemas() []string {
	return d.Procedure.GetSchemas()
}

type DropProcedure struct {
	Procedure *Spname
}

type DropView struct{}

func (*DropView) IStatement() {}

func (*DropTrigger) IStatement()    {}
func (*DropTrigger) HasDDLSchemas() {}
func (d *DropTrigger) GetSchemas() []string {
	return d.Trigger.GetSchemas()
}

type DropTrigger struct {
	Trigger *Spname
}

type DropTablespace struct{}

func (*DropTablespace) IStatement() {}

type DropLogfile struct{}

func (*DropLogfile) IStatement() {}

type DropServer struct{}

func (*DropServer) IStatement() {}

func (*DropEvent) IStatement()    {}
func (*DropEvent) HasDDLSchemas() {}
func (d *DropEvent) GetSchemas() []string {
	return d.Event.GetSchemas()
}

type DropEvent struct {
	Event *Spname
}
