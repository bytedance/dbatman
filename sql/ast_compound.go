package sql

type Signal struct{}

func (*Signal) Statement() {}

type Resignal struct{}

func (*Resignal) Statement() {}

type Diagnostics struct{}

func (*Diagnostics) Statement() {}
