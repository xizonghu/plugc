package node

type Var struct {
	Node
	Name string
	Value interface{}
	Size uint32
	Offset *Expr
}
