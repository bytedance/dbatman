package sql

type Select struct{}

func (*Select) Statement() {}

type Insert struct {
	Table *TableInfo
}

func (*Insert) Statement() {}

type Update struct{}

func (*Update) Statement() {}

type Delete struct{}

func (*Delete) Statement() {}

type Replace struct {
	Table *TableInfo
}

func (*Replace) Statement() {}

type Call struct {
	SpName *TableInfo
}

func (*Call) Statement() {}

type Do struct{}

func (*Do) Statement() {}

type Load struct{}

func (*Load) Statement() {}

type Handler struct{}

func (*Handler) Statement() {}
