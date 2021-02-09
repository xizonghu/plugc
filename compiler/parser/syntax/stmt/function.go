package stmt

import (
	log "github.com/sirupsen/logrus"
	"plugc/compiler/ast/node"
	"plugc/compiler/parser/syntax/expr"
	"plugc/compiler/token"
	"plugc/compiler/token/utils"
	"reflect"
)

type FunctionDeclStmt struct {
	FuncDeclNode *node.FunctionDecl
}

func (p *FunctionDeclStmt) Check(reader token.ITokenReader) (ok bool) {
	pos := reader.Pos()
	defer func() {
		reader.Seek(pos)
	}()

	if !expr.VarDecl.Check(reader) {
		return
	}
	if !utils.IsParenLeft(reader.Read()) {
		return
	}
	ok = true
	return
}

func (p *FunctionDeclStmt) parseFuncInput(reader token.ITokenReader) (err error) {
	if ch, _ := utils.MatchParenLeft(reader); ch == "" {
		return
	}

	var paramTypes, paramNames []string
	for {
		var paramType, paramName string
		if paramType, paramName, err = utils.MatchDecl(reader); err != nil || paramType == "" || paramName == "" {
			break
		}
		paramTypes = append(paramTypes, paramType)
		paramNames = append(paramNames, paramName)
		if ch, _ := utils.MatchComma(reader); ch == "" {
			break
		}
	}

	if str, _ := utils.MatchParenRight(reader); str == "" {
		return
	}

	for i := 0; i < len(paramTypes); i++ {
		p.FuncDeclNode.Input = append(p.FuncDeclNode.Input, node.DataDeclNew(paramTypes[i], paramNames[i]))
	}

	return
}

func (p *FunctionDeclStmt) parseFuncOutput(reader token.ITokenReader) (err error) {
	var funcType, funcName string

	if funcType, funcName, err = utils.MatchDecl(reader); err != nil || funcType == "" || funcName == "" {
		return
	}

	p.FuncDeclNode.Output = node.DataDeclNew(funcType, funcName)
	return
}

func (p *FunctionDeclStmt) Parse(reader token.ITokenReader) (n node.IAstNode, err error) {
	log.Debugf("parse function decl stmt at: line %d", reader.Read().LineNumber)

	p.FuncDeclNode = node.FunctionDeclNew()

	if err = p.parseFuncOutput(reader); err != nil {
		return
	}

	if err = p.parseFuncInput(reader); err != nil {
		return
	}

	segStmt := SegmentStmtNew()
	var tmpNode node.IAstNode
	if tmpNode, err = segStmt.Parse(reader); err != nil || reflect.TypeOf(tmpNode) != reflect.TypeOf(&node.Segment{}) {
		return
	}
	p.FuncDeclNode.Seg = tmpNode.(*node.Segment)
	n = p.FuncDeclNode
	return
}

func FunctionDeclStmtNew() *FunctionDeclStmt {
	return &FunctionDeclStmt{
	}
}

type FuncCallStmt struct {
	Name string
	Args [] *node.Expr
}

func (p *FuncCallStmt) Check(reader token.ITokenReader) (ok bool) {
	return
}

func (p *FuncCallStmt) Parse(reader token.ITokenReader) (node node.IAstNode, err error) {
	return
}

func FunCallNew() *FuncCallStmt {
	return &FuncCallStmt{}
}