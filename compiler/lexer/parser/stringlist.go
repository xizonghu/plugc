package parser

import "errors"

type StringListReader struct {
	list []string
	pos uint32
}

func (p *StringListReader) Read() (string, error) {
	if p.Pos() >= uint32(len(p.list)) {
		return "", errors.New("overflow")
	}
	str := p.list[p.pos]
	p.pos++
	return str, nil
}

func (p *StringListReader) UnRead() {
	p.pos--
}

func (p *StringListReader) Reset() {
	p.pos = 0
}

func (p *StringListReader) Pos() uint32 {
	return p.pos
}

func (p *StringListReader) Len() uint32 {
	return uint32(len(p.list))
}

func StringListReaderNew(strs []string) *StringListReader {
	return & StringListReader{
		list: strs,
		pos:  0,
	}
}

type StrSub struct {
	str string
	line uint32
}