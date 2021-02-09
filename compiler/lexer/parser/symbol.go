package parser

import (
	"plugc/compiler/token"
)

type Symbol struct {

}

func (p *Symbol) GetKey() string {
	return token.Symbol
}

func (p *Symbol) Process(reader IStringListReader) (tk *token.Token, err error) {
	var str string
	if str, err = reader.Read(); err != nil {
		return
	}

	ch := str[0]

	if ch == '_' || 'a' <= ch || ch <= 'z' || 'A' <= ch || ch <= 'Z' {
		tk = & token.Token{
			Type: p.GetKey(),
			Val:  str,
		}
		return
	}
	reader.UnRead()
	return
}

func SymbolNew() *Symbol {
	return &Symbol{}
}