package models

type product interface {
	GetReagentCost() float64
	GetListPrice() int
}

type Test struct {
	ID          string
	Name        string
	ReagentCost float64
	ListPrice   int
}

func (t Test) GetReagentCost() float64 {
	return t.ReagentCost
}

func (t Test) GetListPrice() int {
	return t.ListPrice
}
