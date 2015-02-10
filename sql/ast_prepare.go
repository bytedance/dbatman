package sql

type Deallocate struct{}

func (*Deallocate) Statement() {}

type Prepare struct{}

func (*Prepare) Statement() {}

type Execute struct{}

func (*Execute) Statement() {}
