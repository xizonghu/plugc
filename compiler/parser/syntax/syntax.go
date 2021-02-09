package syntax

import (
	"plugc/compiler/ast/node"
	"plugc/compiler/token"
)

type ISyntaxParser interface {
	Run(reader token.ITokenReader) (node node.IAstNode, err error)
}
