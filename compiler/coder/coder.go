package coder

import (
	"plugc/compiler/ast/node"
)

type ICoder interface {
	Run(root node.IAstNode) (c []byte, err error)
}

type Coder struct {
}

func (p *Coder) Run(root node.IAstNode) (c []byte, err error) {
	//buf := bytes.NewBuffer(c)

	return
}

func CoderNew() *Coder {
	return &Coder{}
}

type ICodeBuffer interface {
	Seek(offset uint32) bool
	Write(b []byte) int
	Offset() uint32
}