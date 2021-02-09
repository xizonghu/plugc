package parser

import (
	"encoding/json"
	"fmt"
	"plugc/compiler/ast/node"
	"plugc/compiler/errors"
	"plugc/compiler/parser/syntax"
	"plugc/compiler/parser/syntax/stmt"
	"plugc/compiler/token"
	"reflect"
	"time"
)

type IParser interface {
	Run(tokens []*token.Token) (node.IAstNode, error)
	Print(node node.IAstNode) string
}

type CParser struct {
	syntaxs map[string]syntax.ISyntaxParser
	stmts []stmt.IStmt
}

func (p *CParser) parse(reader token.ITokenReader) (node node.IAstNode, err error) {
	for _, statement := range p.stmts {
		if !statement.Check(reader) {
			continue
		}
		node, err = statement.Parse(reader)
	}
	return
}

func (p *CParser) Run(tokens []*token.Token) (root node.IAstNode, err error) {
	reader := token.TokenReaderNew(tokens)

	seg := node.SegmentNew()
	root = seg
	for reader.Read().Type != token.End {
		var statement stmt.IStmt
		for _, item := range stmt.Stmts {
			if item.Check(reader) {
				statement = item
				break
			}
		}
		if !reflect.ValueOf(statement).IsValid() {
			err = errors.New(fmt.Sprintf("can not parse: %s", reader))
			return
		}

		var nod node.IAstNode
		if nod, err = statement.Parse(reader); err != nil {
			return
		}
		seg.Stmts = append(seg.Stmts, nod)
		//if reflect.TypeOf(nod) == reflect.TypeOf(&node.FunctionDecl{}) {
		//	ty := reflect.ValueOf(nod)
		//	funcDecl := (*node.FunctionDecl)(unsafe.Pointer(ty.Pointer()))
		//	ty = reflect.ValueOf(funcDecl.Output)
		//	data := (*node.DataDecl)(unsafe.Pointer(ty.Pointer()))
		//}
		time.Sleep(100*time.Millisecond)
	}

	return
	//tree := syntax.ProgramNew()
	//for {
	//	if reader.Pos() >= reader.Len() {
	//		break
	//	}
	//	var node node.IAstNode
	//	if node, err = p.parse(reader); err != nil {
	//		return
	//	}
	//	tree.Append(node)
	//}
	//return tree, nil
}

func (p *CParser) Print(root node.IAstNode) string {
	b, _ := json.Marshal(root)
	return string(b)
}

//func (p *CParser) addSyntaxParser(sp syntax.ISyntaxParser) {
//	p.syntaxs[sp.]
//}

func CParserNew() *CParser {
	parser := & CParser{
		syntaxs: make(map[string]syntax.ISyntaxParser),
	}

	parser.syntaxs[token.Keyword] = KeywordParserNew()
	return parser
}