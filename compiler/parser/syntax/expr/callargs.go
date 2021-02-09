package expr

import (
	"fmt"
	"plugc/compiler/ast/node"
	"plugc/compiler/errors"
	"plugc/compiler/token"
	"plugc/compiler/token/utils"
)

type CallArgsExpr struct {

}

func (p *CallArgsExpr) Check(reader token.ITokenReader) (ok bool) {
	if tk := reader.Read(); !utils.IsParenLeft(tk) {
		return
	}

	reader.Next()

	for tk := reader.Read(); ; tk=reader.Next() {
		if tk == nil {
			return
		} else if utils.IsParenRight(tk) {
			ok = true
			return
		}
	}

	return
}

func (p *CallArgsExpr) Parse(reader token.ITokenReader) (nodes []node.IAstNode, err error) {
	if tk := reader.Read(); !utils.IsParenLeft(tk) {
		err = errors.New("expect '('")
		return
	}

	reader.Next()
	if tk := reader.Read(); utils.IsParenRight(tk) {
		reader.Next()
		return
	}

	for {
		var exprArg node.IAstNode
		if exprArg, err = Formula.Parse(reader); err != nil {
			return
		}
		nodes = append(nodes, exprArg)
		if tk := reader.Read(); utils.IsParenRight(tk) {
			reader.Next()
			return
		}
		if tk := reader.Read(); !utils.IsComma(tk) {
			err = errors.New(fmt.Sprintf("expect ',': %s", tk))
			return
		}
		reader.Next()
	}

	//if tk := reader.Read(); !utils.IsParenRight(tk) {
	//	err = errors.New(fmt.Sprintf("expect ')': %s", tk))
	//	return
	//}

	return
}

var CallArgs = &CallArgsExpr{}

