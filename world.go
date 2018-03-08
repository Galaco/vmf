package vmf

import "github.com/galaco/vmf/types"

type World struct {
	KeyValues types.KeyValues
	Solids []Solid
}

type Solid struct {
	Id int
	Sides []Side
	Editor types.Editor
}

type Side struct {
	Id int
	Plane [2]float32
	Material string
	UAxis string
	VAxis string
	Rotation float32
	LightmapScale int
	SmoothingGroups int
	DispInfo DispInfo
}

type DispInfo struct {
	Power int
	StartPosition types.Vectorf32
	Elevation float32
	SubDiv bool
	Normals []Normal
	Distances []Distance
	Offsets []Offset
	OffsetNormals []Normal
	Alphas []Alpha
	TriangleTags []TriangleTag
	AllowedVerts AllowedVert
}

type Normal struct {
	Id string
	Values []types.Vectorf32
}

type Distance struct {
	Id string
	Values []float32
}

type Offset struct {
	Id string
	Values []types.Vectorf32
}

type Alpha struct {
	Id string
	Values []float32
}

type TriangleTag struct {
	Id string
	Values []float32
}

type AllowedVert struct {
	Id int
	Flags []int
}