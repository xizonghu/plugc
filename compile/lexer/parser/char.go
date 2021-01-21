package parser

import (
	"bytes"
	"plugc/compile/token"
)

type Char struct {

}

func (p *Char) GetKey() string {
	return token.TokenTypeChar
}

func (p *Char) Process(text *bytes.Buffer) (tk *token.Token, err error) {
	if !bytes.HasPrefix(text.Bytes(), []byte(p.GetKey())) {
		return
	}
	text.Next(len(p.GetKey()))
	tk = &token.Token{
		Type: p.GetKey(),
	}
	return
}