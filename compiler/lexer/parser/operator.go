package parser

import (
	"plugc/compiler/token"
	"strings"
)

type Operator struct {

}

func (p *Operator) GetKey() string {
	return token.Operator
}

func (p *Operator) Process(reader IStringListReader) (tk *token.Token, err error) {
	var op string
	if op, err = reader.Read(); err != nil {
		return
	}

	if strings.Contains(token.Operators, string(op[0])) {
		tk = & token.Token{
			Type: p.GetKey(),
			Val:  op,
		}
		if op, err = reader.Read(); err != nil {
			return
		}

		if strings.Contains(token.Operators, op) {
			tk.SetVal(tk.Val + op)
			return
		} else {
			reader.UnRead()
		}
	} else {
		reader.UnRead()
	}

	return
}

func OperatorNew() *Operator {
	return &Operator{}
}
