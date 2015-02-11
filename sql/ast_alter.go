package sql

type AlterTable struct {
	Table ISimpleTable
}

func (*AlterTable) Statement() {}

type AlterDatabase struct {
	Schema []byte
}

func (*AlterDatabase) Statement() {}

type AlterProcedure struct {
	Spname *Spname
}

func (*AlterProcedure) Statement() {}

type AlterFunction struct {
	FuncName *Spname
}

func (*AlterFunction) Statement() {}

type AlterView struct{}

func (*AlterView) Statement() {}

type AlterEvent struct {
	EventName *Spname
	Rename    *Spname
}

func (*AlterEvent) Statement() {}

type AlterTablespace struct{}

func (*AlterTablespace) Statement() {}

type AlterLogfile struct{}

func (*AlterLogfile) Statement() {}

type AlterServer struct{}

func (*AlterServer) Statement() {}

type AlterUser struct{}

func (*AlterUser) Statement() {}
