package sql

type IShow interface {
	IShow()
	IStatement
}

type IShowSchemas interface {
	IShow
	GetSchemas() []string
}

func (*ShowLogs) IStatement()      {}
func (*ShowLogs) IShow()           {}
func (*ShowLogEvents) IStatement() {}
func (*ShowLogEvents) IShow()      {}
func (*ShowCharset) IStatement()   {}
func (*ShowCharset) IShow()        {}
func (*ShowCollation) IStatement() {}
func (*ShowCollation) IShow()      {}

// SHOW CREATE [event|procedure|table|trigger|view]
func (*ShowCreate) IStatement()         {}
func (*ShowCreate) IShow()              {}
func (*ShowCreateDatabase) IStatement() {}
func (*ShowCreateDatabase) IShow()      {}

func (*ShowColumns) IStatement() {}
func (*ShowColumns) IShow()      {}

func (*ShowDatabases) IStatement() {}
func (*ShowDatabases) IShow()      {}

func (*ShowEngines) IStatement() {}
func (*ShowEngines) IShow()      {}

func (*ShowErrors) IStatement()   {}
func (*ShowErrors) IShow()        {}
func (*ShowWarnings) IStatement() {}
func (*ShowWarnings) IShow()      {}

func (*ShowEvents) IStatement() {}
func (*ShowEvents) IShow()      {}

func (*ShowFunction) IStatement() {}
func (*ShowFunction) IShow()      {}

func (*ShowGrants) IStatement() {}
func (*ShowGrants) IShow()      {}

func (*ShowIndex) IStatement() {}
func (*ShowIndex) IShow()      {}

func (*ShowStatus) IStatement() {}
func (*ShowStatus) IShow()      {}

func (*ShowOpenTables) IStatement()  {}
func (*ShowOpenTables) IShow()       {}
func (*ShowTables) IStatement()      {}
func (*ShowTables) IShow()           {}
func (*ShowTableStatus) IStatement() {}
func (*ShowTableStatus) IShow()      {}

func (*ShowPlugins) IStatement() {}
func (*ShowPlugins) IShow()      {}

func (*ShowPrivileges) IStatement() {}
func (*ShowPrivileges) IShow()      {}

func (*ShowProcedure) IStatement() {}
func (*ShowProcedure) IShow()      {}

func (*ShowProcessList) IStatement() {}
func (*ShowProcessList) IShow()      {}

func (*ShowProfiles) IStatement() {}
func (*ShowProfiles) IShow()      {}

func (*ShowSlaveHosts) IStatement()   {}
func (*ShowSlaveHosts) IShow()        {}
func (*ShowSlaveStatus) IStatement()  {}
func (*ShowSlaveStatus) IShow()       {}
func (*ShowMasterStatus) IStatement() {}
func (*ShowMasterStatus) IShow()      {}

func (*ShowTriggers) IStatement() {}
func (*ShowTriggers) IShow()      {}

func (*ShowVariables) IStatement() {}
func (*ShowVariables) IShow()      {}

// currently we use only like for `show databases` syntax
type LikeOrWhere struct {
	Like string
}

type ShowDatabases struct {
	LikeOrWhere *LikeOrWhere
}

func (s *ShowTables) GetSchemas() []string {
	if s.From == nil || len(s.From) == 0 {
		return nil
	}

	return []string{string(s.From)}
}

type ShowTables struct {
	From []byte
}

func (s *ShowTriggers) GetSchemas() []string {
	if s.From == nil || len(s.From) == 0 {
		return nil
	}

	return []string{string(s.From)}
}

type ShowTriggers struct {
	From []byte
}

func (s *ShowEvents) GetSchemas() []string {
	if s.From == nil || len(s.From) == 0 {
		return nil
	}

	return []string{string(s.From)}
}

type ShowEvents struct {
	From []byte
}

func (s *ShowTableStatus) GetSchemas() []string {
	if s.From == nil || len(s.From) == 0 {
		return nil
	}

	return []string{string(s.From)}
}

type ShowTableStatus struct {
	From []byte
}

func (s *ShowOpenTables) GetSchemas() []string {
	if s.From == nil || len(s.From) == 0 {
		return nil
	}

	return []string{string(s.From)}
}

type ShowOpenTables struct {
	From []byte
}

func (s *ShowColumns) GetSchemas() []string {
	if s.From == nil || len(s.From) == 0 {
		return s.Table.GetSchemas()
	}

	return []string{string(s.From)}
}

type ShowColumns struct {
	Table ISimpleTable
	From  []byte
}

func (s *ShowIndex) GetSchemas() []string {
	if s.From == nil || len(s.From) == 0 {
		return s.Table.GetSchemas()
	}

	return []string{string(s.From)}
}

type ShowIndex struct {
	Table ISimpleTable
	From  []byte
}

func (s *ShowProcedure) GetSchemas() []string {
	return s.Procedure.GetSchemas()
}

type ShowProcedure struct {
	Procedure *Spname
}

func (s *ShowFunction) GetSchemas() []string {
	return s.Function.GetSchemas()
}

type ShowFunction struct {
	Function *Spname
}

func (s *ShowCreate) GetSchemas() []string {
	return s.Table.GetSchemas()
}

type ShowCreate struct {
	Prefix []byte
	Table  ISimpleTable
}

func (s *ShowCreateDatabase) GetSchemas() []string {
	if s.Schema == nil || len(s.Schema) == 0 {
		return nil
	}

	return []string{string(s.Schema)}
}

type ShowCreateDatabase struct {
	Schema []byte
}

type ShowGrants struct{}
type ShowCollation struct{}
type ShowCharset struct{}
type ShowVariables struct{}
type ShowProcessList struct{}
type ShowStatus struct{}
type ShowProfiles struct{}
type ShowPrivileges struct{}
type ShowWarnings struct{}
type ShowErrors struct{}
type ShowLogEvents struct{}
type ShowSlaveHosts struct{}
type ShowSlaveStatus struct{}
type ShowMasterStatus struct{}
type ShowLogs struct{}
type ShowPlugins struct{}
type ShowEngines struct{}
