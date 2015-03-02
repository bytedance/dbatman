package sql

func (*Set) IStatement() {}

type Set struct {
	VarList Vars
}

type Vars []*Variable

type Variable struct {
	Type  VarType
	Life  LifeType
	Name  string
	Value IExpr
}

type VarType int
type LifeType int

const (
	Type_Sys = 1
	Type_Usr = 2

	Life_Unknown = 0
	Life_Global  = 1
	Life_Local   = 2
	Life_Session = 3
)

type IAccountMgrStmt interface {
	IsAccountMgrStmt()
	IStatement
}

type Partition struct{}

func (*Partition) IStatement() {}

/*******************************
 * Table Maintenance Statements
 ******************************/
type ITableMtStmt interface {
	IStatement
	IsTableMtStmt()
	GetSchemas() []string
}

func (*Check) IStatement()       {}
func (*Check) IsTableMtStmt()    {}
func (*CheckSum) IStatement()    {}
func (*CheckSum) IsTableMtStmt() {}
func (*Repair) IStatement()      {}
func (*Repair) IsTableMtStmt()   {}
func (*Analyze) IStatement()     {}
func (*Analyze) IsTableMtStmt()  {}
func (*Optimize) IStatement()    {}
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

/****************************
 * Cache Index Statement
 ***************************/
func (*CacheIndex) IStatement() {}

type CacheIndex struct {
	TableIndexList TableIndexes
}

func (c *CacheIndex) GetSchemas() []string {
	if c.TableIndexList == nil || len(c.TableIndexList) == 0 {
		return nil
	}
	return c.TableIndexList.GetSchemas()
}

func (*LoadIndex) IStatement() {}

type LoadIndex struct {
	TableIndexList TableIndexes
}

func (l *LoadIndex) GetSchemas() []string {
	if l.TableIndexList == nil || len(l.TableIndexList) == 0 {
		return nil
	}
	return l.TableIndexList.GetSchemas()
}

type TableIndexes []*TableIndex

func (tis TableIndexes) GetSchemas() []string {
	var rt []string
	for _, v := range tis {
		if v == nil {
			continue
		}

		if r := v.Table.GetSchemas(); r != nil && len(r) != 0 {
			rt = append(rt, r...)
		}
	}

	if len(rt) == 0 {
		return nil
	}

	return rt
}

type TableIndex struct {
	Table ISimpleTable
}

type Binlog struct{}

func (*Binlog) IStatement() {}

func (*Flush) IStatement() {}

type Flush struct{}

func (*FlushTables) IStatement() {}

func (f *FlushTables) GetSchemas() []string {
	if f.Tables == nil {
		return nil
	}
	return f.Tables.GetSchemas()
}

type FlushTables struct {
	Tables ISimpleTables
}

type Kill struct{}

func (*Kill) IStatement() {}

type Reset struct{}

func (*Reset) IStatement() {}

/**********************************************
 * Plugin and User-Defined Function Statements
 *********************************************/
type IPluginAndUdf interface {
	IStatement
	IsPluginAndUdf()
}

func (*Install) IStatement()       {}
func (*Install) IsPluginAndUdf()   {}
func (*CreateUDF) IStatement()     {}
func (*CreateUDF) IDDLStatement()  {}
func (*CreateUDF) IsPluginAndUdf() {}
func (*Uninstall) IStatement()     {}
func (*Uninstall) IsPluginAndUdf() {}

type Install struct{}

type Uninstall struct{}

type CreateUDF struct {
	Function ISimpleTable
}

type udfTail struct {
	Function ISimpleTable
}

/**********************************
 * Account Management Statements
 *********************************/
func (*Grant) IStatement()       {}
func (*Grant) IsAccountMgrStmt() {}

type Grant struct{}

func (*SetPassword) IStatement()    {}
func (*SetPassword) IsAccountStmt() {}

type SetPassword struct{}

func (*RenameUser) IStatement()       {}
func (*RenameUser) IsAccountMgrStmt() {}

type RenameUser struct{}

func (*Revoke) IStatement()       {}
func (*Revoke) IsAccountMgrStmt() {}

type Revoke struct{}

func (*CreateUser) IStatement()       {}
func (*CreateUser) IDDLStatement()    {}
func (*CreateUser) IsAccountMgrStmt() {}

type CreateUser struct{}

func (*AlterUser) IStatement()       {}
func (*AlterUser) IDDLStatement()    {}
func (*AlterUser) IsAccountMgrStmt() {}

type AlterUser struct{}

func (*DropUser) IStatement()       {}
func (*DropUser) IDDLStatement()    {}
func (*DropUser) IsAccountMgrStmt() {}

type DropUser struct{}
