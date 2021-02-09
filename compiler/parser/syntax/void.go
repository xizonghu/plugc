package syntax

import (
	"plugc/compiler/ast/node"
	"plugc/compiler/token"
)

type Void struct {
	name string
	children []node.IAstNode
}

func (p *Void) GetType() string {
	return "void"
}

func (p *Void) Append(node node.IAstNode) {
	p.children = append(p.children, node)
}

func (p *Void) Children() []node.IAstNode {
	return p.children
}

func (p *Void) String() string {
	panic("implement me")
}

func (p *Void) Run(reader token.ITokenReader) (node node.IAstNode, err error) {
	var tk *token.Token
	if tk = reader.Read(); err != nil {
		return
	}
	p.name = tk.Val
	return
}

func VoidNew() *Void {
	return &Void{
	}
}

