package parser

import (
	"bytes"
	"plugc/compile/token"
)

type Viod struct {

}

func (p *Viod) GetKey() string {
	return token.TokenTypeSymbol
}

func (p *Viod) Process(text *bytes.Buffer) (tk *token.Token, err error) {
	if !bytes.HasPrefix(text.Bytes(), []byte(p.GetKey())) {
		return
	}
	text.Next(len(p.GetKey()))
	tk = token.TokenNew(p.GetKey())
	return
}
