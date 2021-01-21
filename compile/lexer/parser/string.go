package parser

import (
	"bytes"
	"plugc/compile/token"
)

type String struct {

}

func (p *String) GetKey() string {
	return token.TokenTypeString
}

func (p *String) Process(text *bytes.Buffer) (tk *token.Token, err error) {
	if !bytes.HasPrefix(text.Bytes(), []byte(p.GetKey())) {
		return
	}
	text.Next(len(p.GetKey()))
	tk = token.TokenNew(p.GetKey())
	return
}
