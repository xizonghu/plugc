package stmt

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"plugc/compiler/ast/node"
	"plugc/compiler/errors"
	"plugc/compiler/parser/syntax/expr"
	"plugc/compiler/token"
	"plugc/compiler/token/utils"
)

type IExpr interface {
}

type ExprStmt struct {
}

func (p *ExprStmt) Check(reader token.ITokenReader) (ok bool) {
	pos := reader.Pos()
	defer reader.Seek(pos)
	ok = expr.Formula.Check(reader)
	return
}

func (p *ExprStmt) Parse(reader token.ITokenReader) (nod node.IAstNode, err error) {
	log.Debugf("parse expr stmt at: line %d", reader.Read().LineNumber)

	nod , err = expr.Formula.Parse(reader)
	if tk := reader.Read(); !utils.IsSemicolon(tk) {
		err = errors.New(fmt.Sprintf("expect ';': %s", reader))
		return
	}
	reader.Next()

	return
}

func ExprStmtNew() *ExprStmt {
	return &ExprStmt{
	}
}