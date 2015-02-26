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

/****************************
 * Cache Index Statement
 ***************************/
func (*CacheIndex) Statement() {}

type CacheIndex struct {
	TableIndexList TableIndexes
}

func (c *CacheIndex) GetSchemas() []string {
	if c.TableIndexList == nil || len(c.TableIndexList) == 0 {
		return nil
	}
	return c.TableIndexList.GetSchemas()
}

func (*LoadIndex) Statement() {}

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

func (*Binlog) Statement() {}

func (*Flush) Statement() {}

type Flush struct{}

func (*FlushTables) Statement() {}

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

func (*Kill) Statement() {}

type Reset struct{}

func (*Reset) Statement() {}

/**********************************************
 * Plugin and User-Defined Function Statements
 *********************************************/
type IPluginAndUdf interface {
	IStatement
	IsPluginAndUdf()
}

func (*Install) Statement()        {}
func (*Install) IsPluginAndUdf()   {}
func (*CreateUDF) Statement()      {}
func (*CreateUDF) IsPluginAndUdf() {}
func (*Uninstall) Statement()      {}
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
