package node

import "encoding/json"

type FunctionDecl struct {
	Node
	Output IAstNode//*DataDecl
	Input []IAstNode//[]*DataDecl
	Seg *Segment
}

func (p *FunctionDecl) MarshalJSON()([]byte, error) {
	var its []interface{}
	its = append(its, map[string]interface{} {
			"Kind": p.Node.Kind,
			"Output": p.Output,
			"Input": p.Input,
		})
	for _, stmt := range p.Seg.Stmts {
		its = append(its, stmt)
	}
	return json.Marshal(its)
}

func FunctionDeclNew() *FunctionDecl {
	return &FunctionDecl{
		Node: NodeInit("FuncDecl"),
	}
}

type FunctionCall struct {
	Node
	Expr IAstNode
	Input []IAstNode//[]*Expr
}

func FunctionCallNew() *FunctionCall {
	return &FunctionCall{
		Node:  NodeInit("FuncCall"),
		Expr:  nil,
		Input: nil,
	}
}