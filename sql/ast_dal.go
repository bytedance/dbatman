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

type Flush struct{}

func (*Flush) Statement() {}

type Kill struct{}

func (*Kill) Statement() {}

type Reset struct{}

func (*Reset) Statement() {}

type Install struct{}

func (*Install) Statement() {}

type Uninstall struct{}

func (*Uninstall) Statement() {}

type Revoke struct{}

func (*Revoke) Statement() {}

type Grant struct{}

func (*Grant) Statement() {}
