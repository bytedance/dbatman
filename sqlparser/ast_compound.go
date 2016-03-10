package sql

type Signal struct{}

func (*Signal) IStatement() {}

type Resignal struct{}

func (*Resignal) IStatement() {}

type Diagnostics struct{}

func (*Diagnostics) IStatement() {}
