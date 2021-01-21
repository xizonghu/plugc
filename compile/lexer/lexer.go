package lexer

import (
	"bytes"
	"plugc/compile/token"
	"strings"
)

type Lexer struct {
	parserTable map[byte][]IParser
}

func (p *Lexer) format(text *bytes.Buffer) *bytes.Buffer {
	//remove all blank
	b := text.Bytes()

	newText := bytes.NewBuffer([]byte{})

	//add blank to split string
	lastType := ChTypeUnkown
	for _, ch := range b {
		currType := ChTypeNew(ch)
		if lastType == ChTypeString {
			if currType == ChTypeString {
				newB := newText.Bytes()
				lastCh := newB[len(newB)-1]
				if lastCh != '\\' {
					lastType = ChTypeUnkown
				}
			}
			newText.WriteByte(ch)
			continue
		} else if currType == ChTypeBlank {
			lastType = ChTypeBlank
			continue
		} else if currType == ChTypeBracket || currType == ChTypeSymbol || lastType != currType {
			newText.WriteByte(' ')
			newText.WriteByte(ch)
			lastType = currType
			continue
		}
		newText.WriteByte(ch)
	}
	return newText
}

func (p *Lexer) split(text *bytes.Buffer) []string {
	return strings.Split(text.String(), " ")
}

func (p *Lexer) Run(text *bytes.Buffer) (tokens []*token.Token, err error) {
	for {
		var ch byte
		if ch, err = text.ReadByte(); err != nil {
			return
		}
		parsers := p.parserTable[ch]
		var token *token.Token
		for _, parser := range parsers {
			if token, err = parser.Process(text); err != nil {
				return
			}
			if token != nil {
				tokens = append(tokens, token)
				break
			}
		}
	}
	return
}

func LexerNew() *Lexer {
	return &Lexer{}
}

type ChType uint8

func ChTypeNew(ch byte) ChType {
	if bytes.Contains([]byte("()[]{}"), []byte{ch}) {
		return ChTypeBracket
	} else if bytes.Contains([]byte(",.;:=+-*/!^&|<>?'\\"), []byte{ch}) {
		return ChTypeSymbol
	} else if bytes.Contains([]byte("\""), []byte{ch}) {
		return ChTypeString
	} else if '0' <= ch && ch <= '9' || 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' {
		return ChTypeText
	} else if ch == ' ' {
		return ChTypeBlank
	}
	return ChTypeUnkown
}

const (
	ChTypeUnkown ChType = iota
	ChTypeSymbol
	ChTypeText
	ChTypeString
	ChTypeBracket
	ChTypeBlank
)
