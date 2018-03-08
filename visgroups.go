package vmf

import "github.com/galaco/vmf/types"

type VisGroup struct {
	Name string
	Id int
	Color types.RGB
	Children []VisGroup
}