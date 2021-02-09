package parser

import (
	"plugc/compiler/token"
)

type Constant struct {

}

func (p *Constant) GetKey() string {
	return token.Constant
}

func (p *Constant) Process(reader IStringListReader) (tk *token.Token, err error) {

	if tk, err = p.ParseAsChar(reader); err != nil {
		return
	}
	if tk != nil {
		return
	}

	if tk, err = p.ParseAsString(reader); err != nil {
		return
	}
	if tk != nil {
		return
	}

	if tk, err = p.ParseAsNumber(reader); err != nil {
		return
	}
	if tk != nil {
		return
	}
	return
}

func (p *Constant) ParseAsChar(reader IStringListReader) (tk *token.Token, err error) {
	var constant string
	if constant, err = reader.Read(); err != nil {
		return
	}

	if constant[0] == '\'' {
		tk = & token.Token{
			Type: p.GetKey(),
			Val:  constant,
		}
		return
	}
	reader.UnRead()
	return
}

func (p *Constant) ParseAsString(reader IStringListReader) (tk *token.Token, err error) {
	var constant string
	if constant, err = reader.Read(); err != nil {
		return
	}

	if constant[0] == '"' {
		tk = & token.Token{
			Type: p.GetKey(),
			Val:  constant,
		}
		return
	}
	reader.UnRead()
	return
}

func (p *Constant) ParseAsNumber(reader IStringListReader) (tk *token.Token, err error) {
	var constant string
	if constant, err = reader.Read(); err != nil {
		return
	}

	firstByte := constant[0]

	if '0' <= firstByte && firstByte <= '9' {
		tk = & token.Token{
			Type: p.GetKey(),
			Val:  constant,
		}
		return
	}
	reader.UnRead()
	return
}

func ConstantNew() *Constant {
	return & Constant{}
}