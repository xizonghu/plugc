package expr

import (
	"errors"
	"fmt"
	"plugc/compiler/ast/node"
	"plugc/compiler/token"
	"plugc/compiler/token/utils"
	"reflect"
)

type FormulaExpr struct {
	depth uint32
}

func (p *FormulaExpr) Check(reader token.ITokenReader) (ok bool) {
	tk := reader.Read()
	if utils.IsSymbol(tk) || utils.IsOperator(tk) || utils.IsParenLeft(tk) {
		ok = true
	}

	return
}

func (p *FormulaExpr) isEnd(tk *token.Token) bool {
	if utils.IsComma(tk) || utils.IsSemicolon(tk) {
		return true
	}
	return false
}

func (p *FormulaExpr) parseElementStart(reader token.ITokenReader)(err error) {
	tk := reader.Read()
	if !utils.IsParenLeft(tk) {
		return
	}

	p.depth++
	reader.Next()
	//nod = node.ExprNew()

	return
}

func (p *FormulaExpr) parseElementA(reader token.ITokenReader)(nod node.IAstNode, err error) {
	tk := reader.Read()
	exprA := node.DataElementNew(tk.Type, tk.Val)

	if tk = reader.Next(); utils.IsParenLeft(tk) {
		funcCall := node.FunctionCallNew()
		funcCall.Expr = exprA
		if funcCall.Input, err = CallArgs.Parse(reader); err != nil {
			return
		}
		nod = funcCall
	} else {
		nod = exprA
	}
	return
}

func (p *FormulaExpr) parseElementOp(reader token.ITokenReader)(op string, err error) {
	tk := reader.Read()
	if !utils.IsOperator(tk) {
		err = errors.New(fmt.Sprintf("expect a operator: %s", reader))
		return
	}

	for tk = reader.Read(); utils.IsOperator(tk); tk = reader.Next() {
		op += tk.Val
	}
	return
}

func (p *FormulaExpr) parseElementB(reader token.ITokenReader)(nod node.IAstNode, err error) {
	nod, err = p.parseExprNew(reader)
	return
}

func (p *FormulaExpr) parseElementEnd(reader token.ITokenReader)(err error) {
	tk := reader.Read()
	if utils.IsParenRight(tk) {
		p.depth--
		reader.Next()
		return
	} else if utils.IsSemicolon(tk) {
		return
	}

	err = errors.New(fmt.Sprintf("expect ')' or ';': %s", tk))
	return
}

func (p *FormulaExpr) parseExprNew(reader token.ITokenReader) (nod node.IAstNode, err error) {
	exprNew := node.ExprNew()
	defer func() {
		nod = exprNew
		if exprNew.Op == "" {
			if reflect.ValueOf(exprNew.ExprA).IsValid() {
				nod = exprNew.ExprA
			} else if reflect.ValueOf(exprNew.ExprB).IsValid() {
				nod = exprNew.ExprB
			} else {
				err = errors.New("expr op is nil, but expr has two parameter")
				return
			}
		}
	}()

	if tk := reader.Read(); utils.IsParenLeft(tk) {
		reader.Next()
		p.depth++
		exprNew.ExprA, err = p.parse(reader)
	} else {
		exprNew.ExprA, err = p.parseElementA(reader)
	}

	if err != nil {
		return
	}

	if tk := reader.Read(); utils.IsParenRight(tk) && p.depth != 0 {
		reader.Next()
		p.depth--
	}

	if tk := reader.Read(); utils.IsComma(tk) || utils.IsSemicolon(tk) || utils.IsParenRight(tk) {
		return
	}

	if exprNew.Op, err = p.parseElementOp(reader); err != nil {
		return
	}

	if tk := reader.Read(); utils.IsParenRight(tk) && p.depth != 0 {
		reader.Next()
		p.depth--
	}

	if tk := reader.Read(); utils.IsComma(tk) || utils.IsSemicolon(tk) || utils.IsParenRight(tk) {
		return
	}

	if tk := reader.Read(); utils.IsParenLeft(tk) {
		reader.Next()
		p.depth--
		exprNew.ExprB, err = p.parse(reader)
	} else {
		exprNew.ExprB, err = p.parseElementB(reader)
	}

	if err != nil {
		return
	}

	if tk := reader.Read(); utils.IsParenRight(tk) && p.depth != 0 {
		reader.Next()
		p.depth--
	}

	return
}

func (p *FormulaExpr) parse(reader token.ITokenReader) (nod node.IAstNode, err error) {
	//if err = p.parseElementStart(reader); err != nil {
	//	return
	//}

	if nod, err = p.parseExprNew(reader); err != nil {
		return
	}

	//if err = p.parseElementEnd(reader); err != nil {
	//	return
	//}

	return
}

func (p *FormulaExpr) Parse(reader token.ITokenReader) (nod node.IAstNode, err error) {
	if nod, err = p.parse(reader); err != nil {
		return
	}

	if p.depth != 0 {
		err = errors.New("the count of '(' dot not equal ')'")
	}
	return
}

var Formula = &FormulaExpr{}
