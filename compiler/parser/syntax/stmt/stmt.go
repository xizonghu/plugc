package stmt

import (
	"plugc/compiler/ast/node"
	"plugc/compiler/token"
)

type IStmt interface {
	Check(reader token.ITokenReader) (ok bool)
	Parse(reader token.ITokenReader) (node node.IAstNode, err error)
}

var Stmts = []IStmt{
	SegmentStmtNew(),
	FunctionDeclStmtNew(),
	ExprStmtNew(),
	VarDeclStmtNew(),
	//NoneStmtNew(),
}
