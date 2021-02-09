package syntax

import "plugc/compiler/ast/node"

type Program struct {
	children []node.IAstNode
}

func (p *Program) GetType() string {
	return "program"
}

func (p *Program) Append(node node.IAstNode) {
	p.children = append(p.children, node)
}

func (p *Program) Children() []node.IAstNode {
	return p.children
}

func (p *Program) String() string {
	panic("implement me")
}

func ProgramNew() *Program {
	return &Program{
	}
}