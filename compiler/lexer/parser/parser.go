package parser

import (
	"plugc/compiler/token"
)

type IStringListReader interface {
	Read() (string, error)
	UnRead()
	Reset()
	Pos() uint32
	Len() uint32
}

type IParser interface {
	GetKey() string
	Process(reader IStringListReader) (*token.Token, error)
}

var Parsers = []IParser{ConstantNew(), DelimiterNew(), KeywordNew(), OperatorNew(), SymbolNew()}