package vmf

import "github.com/galaco/vmf/types"

type Cameras struct {
	ActiveCamera int
	Cameras []Camera
}

type Camera struct {
	Position types.Vectorf32
	Look types.Vectorf32
}
