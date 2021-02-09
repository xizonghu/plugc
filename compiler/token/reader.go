package token

import "fmt"

type ITokenReader interface {
	Read() *Token
	Next() *Token
	Prev() *Token
	Seek(pos uint32)
	Pos() uint32
	Len() uint32
}

type TokenReader struct {
	list []*Token
	pos uint32
}

func (p *TokenReader) Read() *Token {
	//if p.Pos() >= p.Len() {
	//	return nil
	//}
	return p.list[p.pos]
}

func (p *TokenReader) Next() *Token {
	if p.Pos() < p.Len() {
		p.pos++
	}
	return p.Read()
}

func (p *TokenReader) Prev() *Token {
	if p.Pos() == 0 {
		return nil
	}
	p.pos--
	return p.Read()
}

func (p *TokenReader) Seek(pos uint32) {
	p.pos = pos
}

func (p *TokenReader) Pos() uint32 {
	return p.pos
}

func (p *TokenReader) Len() uint32 {
	return uint32(len(p.list))
}

func (p *TokenReader) String() string {
	line := p.Read().LineNumber
	str := fmt.Sprintf("index=%d, line=%d, stmt=[", p.Read().Index, line)
	for tk := p.Read(); tk.LineNumber == line; tk = p.Next() {
		str += " " + tk.Val
	}
	str += "]"
	return str
}

func TokenReaderNew(list []*Token) *TokenReader {
	return & TokenReader{
		list: list,
		pos:  0,
	}
}
