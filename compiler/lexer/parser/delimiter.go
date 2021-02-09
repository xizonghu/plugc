package parser

import (
	"plugc/compiler/token"
	"strings"
)

type Delimiter struct {

}

func (p *Delimiter) GetKey() string {
	return token.Delimiter
}

func (p *Delimiter) Process(reader IStringListReader) (tk *token.Token, err error) {
	var delimiter string
	if delimiter, err = reader.Read(); err != nil {
		return
	}

	if strings.Contains(token.Delimiters, string(delimiter[0])) {
		tk = & token.Token{
			Type: p.GetKey(),
			Val:  delimiter,
		}
		return
	}
	reader.UnRead()
	return
}

func DelimiterNew() *Delimiter {
	return &Delimiter{}
}
