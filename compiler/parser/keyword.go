package parser

import (
	"plugc/compiler/ast/node"
	"plugc/compiler/parser/syntax"
	"plugc/compiler/token"
)

type KeywordParser struct {
	parsers map[string] syntax.ISyntaxParser
}

func (p *KeywordParser) Run(reader token.ITokenReader) (node node.IAstNode, err error) {
	return
}

func KeywordParserNew() *KeywordParser {
	keyword := &KeywordParser{
		parsers: make(map[string]syntax.ISyntaxParser),
	}
	keyword.parsers[token.KEY_VOID] = syntax.VoidNew()
	return keyword
}
