package vmf

type Vmf struct {
	VersionInfo Node
	ViewSettings Node
	VisGroup Node
	World Node
	Entities Node
	Cameras Node
	Cordon Node // Pre-L4D only
	Cordons Node // Post-L4D only
}

// A KeyValue object, that may hold multiple Values
type Node struct {
	key string
	value []interface{}
}
func (node *Node) HasProperty(name string) bool {
	v := node.GetProperty(name)
	return v != ""
}

func (node *Node) GetProperty(name string) string {
	for _,child := range node.value {
		n,_ := child.(Node)
		if n.key == name {
			return (n.value[0]).(string)
		}
	}
	return ""
}

func (node *Node) GetChildrenOfType(name string) (children []Node) {
	for _,child := range node.value {
		n,_ := child.(Node)
		if n.key == name {
			children = append(children, n)
		}
	}

	return children
}