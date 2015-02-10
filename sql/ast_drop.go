package sql

type DropTables struct{}

func (*DropTables) Statement() {}

type DropIndex struct{}

func (*DropIndex) Statement() {}

type DropDatabase struct{}

func (*DropDatabase) Statement() {}

type DropFunction struct{}

func (*DropFunction) Statement() {}

type DropProcedure struct{}

func (*DropProcedure) Statement() {}

type DropView struct{}

func (*DropView) Statement() {}

type DropTrigger struct{}

func (*DropTrigger) Statement() {}

type DropTablespace struct{}

func (*DropTablespace) Statement() {}

type DropLogfile struct{}

func (*DropLogfile) Statement() {}

type DropServer struct{}

func (*DropServer) Statement() {}

type DropEvent struct{}

func (*DropEvent) Statement() {}
