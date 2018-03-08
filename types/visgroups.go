package types

type VisGroup struct {
	Name string
	Id int
	Color RGB
	Children []VisGroup
}