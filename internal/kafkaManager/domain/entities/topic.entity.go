package entities

type Topic struct {
	ID         int64
	CompanyID  int
	Name       string
	Partitions int
}
