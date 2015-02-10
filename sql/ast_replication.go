package sql

type Change struct{}

func (*Change) Statement() {}

type StartSlave struct{}

func (*StartSlave) Statement() {}

type StopSlave struct{}

func (*StopSlave) Statement() {}
