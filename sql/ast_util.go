package sql

type Help struct{}

func (*Help) Statement() {}

type Describe struct{}

func (*Describe) Statement() {}

type Use struct{}

func (*Use) Statement() {}
