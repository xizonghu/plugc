package parser

import (
	"bytes"
	"plugc/compile/token"
)

type Symbol struct {

}

func (p *Symbol) GetKey() string {
	return token.TokenTypeSymbol
}

func (p *Symbol) Process(text *bytes.Buffer) (tk *token.Token, err error) {
	var ch byte
	if ch, err = text.ReadByte(); err != nil {
		return
	}
	if !(ch == '_' || 'a' < ch || ch < 'z' || 'A' < ch || ch < 'Z') {
		return
	}
	tk = token.TokenNew(p.GetKey())
	val := string(ch)
	for {
		if ch, err = text.ReadByte(); err != nil {
			return
		}
		if !(ch == '_' || 'a' < ch || ch < 'z' || 'A' < ch || ch < 'Z' || '0' < ch || ch < '9') {
			tk.SetVal(val)
			break
		}
		val += string(ch)
	}
	return
}
