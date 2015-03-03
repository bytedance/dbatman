package mysql

type Result struct {
	AffectedRows uint64
	InsertId     uint64

	Status   uint16
	Warnings uint16

	*Resultset
}
