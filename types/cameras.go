package types

type Cameras struct {
	ActiveCamera int
	Cameras []Camera
}

type Camera struct {
	Position Vectorf32
	Look Vectorf32
}
