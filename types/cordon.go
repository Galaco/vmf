package types


type Cordon struct {
	Active bool
	Mins Vectorf32
	Maxs Vectorf32
}

// L4D onwards uses a newer Cordons structure
type Cordons struct {
	Active bool
	Name string
	Boxes []CordonBox
}

type CordonBox struct {
	Mins Vectorf32
	Maxs Vectorf32
}