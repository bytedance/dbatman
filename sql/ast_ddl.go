package sql

type RenameTable struct{}

func (*RenameTable) Statement() {}

type TruncateTable struct{}

func (*TruncateTable) Statement() {}
