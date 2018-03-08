package vmf

import (
	"github.com/galaco/vmf/types"
)

type Cordon struct {
	Active bool
	Mins types.Vectorf32
	Maxs types.Vectorf32
}

// L4D onwards uses a newer Cordons structure
type Cordons struct {
	Active bool
	Name string
	Boxes []CordonBox
}

type CordonBox struct {
	Mins types.Vectorf32
	Maxs types.Vectorf32
}