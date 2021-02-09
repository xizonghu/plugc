package node

//regex: [(]{0,1} A{0,1} Op{1,} B{1} [)]{0,1}

type Expr struct {
	Node
	Op string
	ExprA IAstNode //*Expr
	ExprB IAstNode //*Expr
}

func ExprNew() *Expr {
	return & Expr{
		Node: NodeInit("Expr"),
		Op: "",
		ExprA:  nil,
		ExprB:  nil,
	}
}