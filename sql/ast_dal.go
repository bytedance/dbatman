package sql

type IAccountMgrStmt interface {
	IsAccountMgrStmt()
	IStatement
}

type Partition struct{}

func (*Partition) Statement() {}

/*******************************
 * Table Maintenance Statements
 ******************************/
type ITableMtStmt interface {
	IStatement
	IsTableMtStmt()
	GetSchemas() []string
}

func (*Check) Statement()        {}
func (*Check) IsTableMtStmt()    {}
func (*CheckSum) Statement()     {}
func (*CheckSum) IsTableMtStmt() {}
func (*Repair) Statement()       {}
func (*Repair) IsTableMtStmt()   {}
func (*Analyze) Statement()      {}
func (*Analyze) IsTableMtStmt()  {}
func (*Optimize) Statement()     {}
func (*Optimize) IsTableMtStmt() {}

func (c *Check) GetSchemas() []string {
	return c.Tables.GetSchemas()
}

func (c *CheckSum) GetSchemas() []string {
	return c.Tables.GetSchemas()
}

func (r *Repair) GetSchemas() []string {
	return r.Tables.GetSchemas()
}

func (a *Analyze) GetSchemas() []string {
	return a.Tables.GetSchemas()
}

func (o *Optimize) GetSchemas() []string {
	return o.Tables.GetSchemas()
}

type Check struct {
	Tables ISimpleTables
}

type CheckSum struct {
	Tables ISimpleTables
}

type Repair struct {
	Tables ISimpleTables
}

type Analyze struct {
	Tables ISimpleTables
}

type Optimize struct {
	Tables ISimpleTables
}

type CacheIndex struct{}

func (*CacheIndex) Statement() {}

type LoadIndex struct{}

func (*LoadIndex) Statement() {}

type Binlog struct{}

func (*Binlog) Statement() {}

type Flush struct{}

func (*Flush) Statement() {}

type Kill struct{}

func (*Kill) Statement() {}

type Reset struct{}

func (*Reset) Statement() {}

type Install struct{}

func (*Install) Statement() {}

type Uninstall struct{}

func (*Uninstall) Statement() {}

/**********************************
 * Account Management Statements
 *********************************/
func (*Grant) Statement()        {}
func (*Grant) IsAccountMgrStmt() {}

type Grant struct{}

func (*SetPassword) Statement()     {}
func (*SetPassword) IsAccountStmt() {}

type SetPassword struct{}

func (*RenameUser) Statement()        {}
func (*RenameUser) IsAccountMgrStmt() {}

type RenameUser struct{}

func (*Revoke) Statement()        {}
func (*Revoke) IsAccountMgrStmt() {}

type Revoke struct{}

func (*CreateUser) Statement()        {}
func (*CreateUser) IsAccountMgrStmt() {}

type CreateUser struct{}

func (*AlterUser) Statement()        {}
func (*AlterUser) IsAccountMgrStmt() {}

type AlterUser struct{}

func (*DropUser) Statement()        {}
func (*DropUser) IsAccountMgrStmt() {}

type DropUser struct{}
