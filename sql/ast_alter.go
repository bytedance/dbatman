package sql

type AlterTable struct{}

func (*AlterTable) Statement() {}

type AlterDatabase struct{}

func (*AlterDatabase) Statement() {}

type AlterProcedure struct{}

func (*AlterProcedure) Statement() {}

type AlterFunction struct{}

func (*AlterFunction) Statement() {}

type AlterView struct{}

func (*AlterView) Statement() {}

type AlterEvent struct{}

func (*AlterEvent) Statement() {}

type AlterTablespace struct{}

func (*AlterTablespace) Statement() {}

type AlterLogfile struct{}

func (*AlterLogfile) Statement() {}

type AlterServer struct{}

func (*AlterServer) Statement() {}

type AlterUser struct{}

func (*AlterUser) Statement() {}
