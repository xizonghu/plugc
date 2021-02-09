package parser

import (
	"bytes"
	"plugc/compiler/token"
)

type ParenLeft struct {

}

func (p *ParenLeft) GetKey() string {
	return token.TokenTypeParenLeft
}

func (p *ParenLeft) Process(text *bytes.Buffer) (tk *token.Token, err error) {
	var ch byte
	if ch, err = text.ReadByte(); err != nil {
		return
	}
	if ch != '(' {
		return
	}
	tk = token.TokenNew(p.GetKey())
	return
}

type ParenRight struct {

}

func (p *ParenRight) GetKey() string {
	return token.TokenTypeParenRight
}

func (p *ParenRight) Process(text *bytes.Buffer) (tk *token.Token, err error) {
	var ch byte
	if ch, err = text.ReadByte(); err != nil {
		return
	}
	if ch != ')' {
		return
	}
	tk = token.TokenNew(p.GetKey())
	return
}

