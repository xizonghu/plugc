package parser

import (
	"bytes"
	"plugc/compile/token"
)

type Semicolon struct {

}

func (p *Semicolon) GetKey() string {
	return token.TokenTypeSemicolon
}

func (p *Semicolon) Process(text *bytes.Buffer) (tk *token.Token, err error) {
	if !bytes.HasPrefix(text.Bytes(), []byte(p.GetKey())) {
		return
	}
	text.Next(len(p.GetKey()))
	tk = &token.Token{
		Type: p.GetKey(),
	}
	return
}