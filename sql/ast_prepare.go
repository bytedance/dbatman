package sql

type Deallocate struct{}

func (*Deallocate) IStatement() {}

type Prepare struct{}

func (*Prepare) IStatement() {}

type Execute struct{}

func (*Execute) IStatement() {}
