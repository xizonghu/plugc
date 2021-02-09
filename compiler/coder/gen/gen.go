package gen

import "plugc/compiler/ast/node"

type IBuffer interface {
	Write(b []byte)
	Seek(offset uint32)
	Offset() uint32
}

type IVarTable interface {
}

type ITables interface {
}

type IGenerator interface {
	Parse(table *Table, node node.IAstNode, buffer IBuffer) error
}
