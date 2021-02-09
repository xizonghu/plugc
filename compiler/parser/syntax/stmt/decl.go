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

type IDataType interface {
	Storage() string
	Type() string
	Value() interface{}
	UnSigned() bool
	Size() uint8
}

type VarDeclStmt struct {

}

func (p *VarDeclStmt) Check(reader token.ITokenReader) (ok bool) {
	pos := reader.Pos()
	defer reader.Seek(pos)
	ok = expr.VarDecl.Check(reader)
	return
}

func (p *VarDeclStmt) Parse(reader token.ITokenReader) (n node.IAstNode, err error) {
	log.Debugf("parse var decl stmt at: line %d", reader.Read().LineNumber)
	if n, err = expr.VarDecl.Parse(reader); err != nil {
		return
	}

	if tk := reader.Read(); !utils.IsSemicolon(tk) {
		err = errors.New(fmt.Sprintf("expect ';': %s", tk))
		return
	}

	reader.Next()

	return
}

func VarDeclStmtNew() *VarDeclStmt {
	return &VarDeclStmt{}
}