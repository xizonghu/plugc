package stmt

import (
	"fmt"
	"plugc/compiler/ast/node"
	"plugc/compiler/errors"
	"plugc/compiler/token"
	"plugc/compiler/token/utils"
)

type NoneStmt struct {
}

func (p *NoneStmt) Check(reader token.ITokenReader) (ok bool) {
	if tk := reader.Read(); utils.IsSemicolon(tk) {
		ok = true
	}
	return
}

func (p *NoneStmt) Parse(reader token.ITokenReader) (n node.IAstNode, err error) {
	if tk := reader.Read(); !utils.IsSemicolon(tk) {
		err = errors.New(fmt.Sprintf("expect ';': %s", tk))
	}

	return
}

func NoneStmtNew() *NoneStmt {
	return &NoneStmt{
	}
}
