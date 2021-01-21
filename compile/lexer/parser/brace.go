package parser

import (
	"bytes"
	"plugc/compile/token"
)

type BraceLeft struct {

}

func (p *BraceLeft) GetKey() string {
	return token.TokenTypeBraceLeft
}

func (p *BraceLeft) Process(text *bytes.Buffer) (tk *token.Token, err error) {
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

type BraceRight struct {

}

func (p *BraceRight) GetKey() string {
	return token.TokenTypeBraceRight
}

func (p *BraceRight) Process(text *bytes.Buffer) (tk *token.Token, err error) {
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

