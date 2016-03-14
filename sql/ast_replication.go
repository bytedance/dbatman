package sql

type Change struct{}

func (*Change) IStatement() {}

type Purge struct{}

func (*Purge) IStatement() {}

type StartSlave struct{}

func (*StartSlave) IStatement() {}

type StopSlave struct{}

func (*StopSlave) IStatement() {}
