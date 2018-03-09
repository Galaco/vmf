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

// Get node key
func (node *Node) GetKey() *string {
	return &node.key
}

// Get all node values
func (node *Node) GetAllValues() *[]interface{} {
	return &node.value
}

// Check if Node has a property defined
func (node *Node) HasProperty(name string) bool {
	v := node.GetProperty(name)
	return v != ""
}

// Return the value of a property from its name
// Note: A property is a string:string, not a string:[]Node
func (node *Node) GetProperty(name string) string {
	for _,child := range node.value {
		n,_ := child.(Node)
		if n.key == name {
			return (n.value[0]).(string)
		}
	}
	return ""
}

// Return all children of a given type for a node.
// This is different from properties, as a property is a string:string
func (node *Node) GetChildrenByKey(name string) (children []Node) {
	for _,child := range node.value {
		n,_ := child.(Node)
		if n.key == name {
			children = append(children, n)
		}
	}

	return children
}