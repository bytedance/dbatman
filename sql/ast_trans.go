package sql

type Start struct{}

func (*Start) Statement() {}

type Lock struct{}

func (*Lock) Statement() {}

type Unlock struct{}

func (*Unlock) Statement() {}

type Begin struct{}

func (*Begin) Statement() {}

type Commit struct{}

func (*Commit) Statement() {}

type Rollback struct{}

func (*Rollback) Statement() {}

type XA struct{}

func (*XA) Statement() {}

type SavePoint struct{}

func (*SavePoint) Statement() {}

type Release struct{}

func (*Release) Statement() {}
