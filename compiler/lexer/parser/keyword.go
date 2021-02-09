package parser

import (
	"plugc/compiler/token"
)

type Keyword struct {

}

func (p *Keyword) GetKey() string {
	return token.Keyword
}

func (p *Keyword) Process(reader IStringListReader) (tk *token.Token, err error) {
	var str string
	if str, err = reader.Read(); err != nil {
		return
	}

	for _, keyword := range token.Keywords {
		if keyword == str {
			tk = & token.Token{
				Type: p.GetKey(),
				Val:  str,
			}
			return
		}
	}
	reader.UnRead()

	return
}

func KeywordNew() *Keyword {
	return &Keyword{}
}

