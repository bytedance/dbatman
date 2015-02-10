package sql

type Insert struct {
	Table *TableInfo
}

func (*Insert) Statement() {}

type Replace struct {
	Table *TableInfo
}

func (*Replace) Statement() {}

type Call struct {
	SpName *TableInfo
}

func (*Call) Statement() {}
