package vmf

import "github.com/galaco/vmf/types"

type Entity struct {
	KeyValues types.KeyValues
	Connections types.KeyValues
	Solid []Solid
	Editor types.Editor
	//Hidden
}
