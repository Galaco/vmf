package types

type Vectorf32 struct {
	X float32
	Y float32
	Z float32
}

type Vectori32 struct {
	X int
	Y int
	Z int
}

type RGB Vectori32

type KeyValues struct {
	keyValues []KeyValue
}

type KeyValue struct {
	Key string
	Value string
}

type Editor struct {
	Color RGB
	VisgroupId int
	GroupId int
	VisgroupShown bool
	VisgroupAutoShown bool
	Comments string
}