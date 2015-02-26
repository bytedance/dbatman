package sql

type IAccountMgrStmt interface {
	IsAccountMgrStmt()
	IStatement
}

type Partition struct{}

func (*Partition) Statement() {}

type Check struct{}

func (*Check) Statement() {}

type CheckSum struct{}

func (*CheckSum) Statement() {}

type Repair struct{}

func (*Repair) Statement() {}

type Analyze struct{}

func (*Analyze) Statement() {}

type Optimize struct{}

func (*Optimize) Statement() {}

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
