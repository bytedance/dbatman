package sql

type Start struct{}

func (*Start) Statement() {}

type Lock struct{}

func (*Lock) Statement() {}

type Unlock struct{}

func (*Unlock) Statement() {}

type Begin struct{}

func (*Begin) Statement() {}
