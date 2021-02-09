package node

type Operator struct {
	Node
	Name string
	A *Node
	B *Node
}