package expr

import (
	"fmt"
	"plugc/compiler/ast/node"
	"plugc/compiler/errors"
	"plugc/compiler/token"
	"plugc/compiler/token/utils"
)

type VarDeclExpr struct {

}

func (p *VarDeclExpr) Check(reader token.ITokenReader) (ok bool) {
	if tk := reader.Read(); !utils.IsDecl(tk) {
		return
	}

	for tk := reader.Read(); utils.IsDecl(tk); tk = reader.Next() {
	}

	if tk := reader.Read(); !utils.IsSymbol(tk) {
		return
	}

	reader.Next()
	ok = true

	return
}

func (p *VarDeclExpr) Parse(reader token.ITokenReader) (n node.IAstNode, err error) {
	if tk := reader.Read(); !utils.IsDecl(tk) {
		err = errors.New(fmt.Sprintf("expect a type keyword: %s", reader))
		return
	}

	var varType, varName string
	if varType, varName, err = utils.MatchDecl(reader); err != nil {
		return
	}
	if varType == "" || varName == "" {
		err = errors.New(fmt.Sprintf("var decl error: %s", reader))
		return
	}

	n = node.DataDeclNew(varType, varName)
	return
}

var VarDecl = &VarDeclExpr{}