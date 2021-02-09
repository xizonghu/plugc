package node

type Segment struct {
	Node
	Stmts []IAstNode
}

func SegmentNew() *Segment {
	return &Segment{
		Node: NodeInit("Segment"),
	}
}
