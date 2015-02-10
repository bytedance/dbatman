package sql

type Partition struct{}

func (*Partition) Statement() {}

type Check struct{}

func (*Check) Statement() {}

type CheckSum struct{}

func (*CheckSum) Statement() {}

type Repair struct{}

func (*Repair) Statement() {}

type Analyze struct{}

func (*Analyze) Statement() {}

type Optimize struct{}

func (*Optimize) Statement() {}

type RenameUser struct{}

func (*RenameUser) Statement() {}

type CacheIndex struct{}

func (*CacheIndex) Statement() {}

type LoadIndex struct{}

func (*LoadIndex) Statement() {}

type Binlog struct{}

func (*Binlog) Statement() {}
