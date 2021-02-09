package token

import "encoding/json"

const (
	Keyword			= "key"
	Operator		= "op"
	Symbol			= "sym"
	Delimiter		= "div"
	Constant		= "const"
	End				= "end"
)

type Token struct {
	Type string
	Val string
	LineNumber uint32
	Index uint32
}

func (p *Token) SetVal(val string) {
	p.Val = val
}

func (p *Token) String() string {
	b, _ := json.Marshal(p)
	return string(b)
}

func TokenNew(Type string) *Token {
	index++
	return &Token {
		Index: index,
		Type: Type,
	}
}

var index = uint32(0)