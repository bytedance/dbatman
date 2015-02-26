package sql

func (*DropTables) Statement()     {}
func (*DropTables) HasDDLSchemas() {}
func (d *DropTables) GetSchemas() []string {
	return d.Tables.GetSchemas()
}

type DropTables struct {
	Tables ISimpleTables
}

func (*DropIndex) Statement()     {}
func (*DropIndex) HasDDLSchemas() {}
func (d *DropIndex) GetSchemas() []string {
	return d.On.GetSchemas()
}

type DropIndex struct {
	On ISimpleTable
}

type DropDatabase struct{}

func (*DropDatabase) Statement() {}

func (*DropFunction) Statement()     {}
func (*DropFunction) HasDDLSchemas() {}
func (d *DropFunction) GetSchemas() []string {
	return d.Function.GetSchemas()
}

type DropFunction struct {
	Function *Spname
}

func (*DropProcedure) Statement()     {}
func (*DropProcedure) HasDDLSchemas() {}
func (d *DropProcedure) GetSchemas() []string {
	return d.Procedure.GetSchemas()
}

type DropProcedure struct {
	Procedure *Spname
}

type DropView struct{}

func (*DropView) Statement() {}

func (*DropTrigger) Statement()     {}
func (*DropTrigger) HasDDLSchemas() {}
func (d *DropTrigger) GetSchemas() []string {
	return d.Trigger.GetSchemas()
}

type DropTrigger struct {
	Trigger *Spname
}

type DropTablespace struct{}

func (*DropTablespace) Statement() {}

type DropLogfile struct{}

func (*DropLogfile) Statement() {}

type DropServer struct{}

func (*DropServer) Statement() {}

func (*DropEvent) Statement()     {}
func (*DropEvent) HasDDLSchemas() {}
func (d *DropEvent) GetSchemas() []string {
	return d.Event.GetSchemas()
}

type DropEvent struct {
	Event *Spname
}
