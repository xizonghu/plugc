package lexer

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"plugc/compiler/lexer/parser"
	"plugc/compiler/token"
)

type Lexer struct {
	parsers []parser.IParser
}

func (p *Lexer) split(text *bytes.Buffer) (strs []string, lines []uint32) {
	//remove all blank
	b := text.Bytes()

	newText := bytes.NewBuffer([]byte{})

	//add blank to split string
	lastType := ChTypeUnkown
	line := uint32(1)
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
		} else if lastType == ChTypeChar {
				if currType == ChTypeChar {
					newB := newText.Bytes()
					lastCh := newB[len(newB)-1]
					if lastCh != '\\' {
						lastType = ChTypeUnkown
					}
				}
				newText.WriteByte(ch)
				continue
		} else if currType == ChTypeBlank || currType == ChTypeEndOfLine {
			lastType = currType
			if ch == '\n' {
				line++
			}
			continue
		} else if currType == ChTypeBracket || currType == ChTypeSymbol || currType == ChTypeChar || currType == ChTypeString || lastType != currType {
			if "" != newText.String() {
				strs = append(strs, newText.String())
				lines = append(lines, line)
			}
			newText.Reset()
			newText.WriteByte(ch)
			lastType = currType
			continue
		}
		newText.WriteByte(ch)
	}
	if len(newText.Bytes()) != 0 {
		strs = append(strs, newText.String())
		lines = append(lines, line)
	}
	return
}

func (p *Lexer) parse(reader parser.IStringListReader) (token *token.Token, err error) {
	for _, parser := range p.parsers {
		if token, err = parser.Process(reader); err != nil {
			return
		}
		if token != nil {
			return
		}
	}
	return nil, errors.New("can not parse this string")
}

func (p *Lexer) Run(text *bytes.Buffer) (tokens []*token.Token, err error) {
	strs, lines := p.split(text)
	p.printStrlines(strs, lines)

	reader := parser.StringListReaderNew(strs)
	index := uint32(0)
	for {
		if reader.Pos() >= reader.Len() {
			break
		}
		var token *token.Token
		if token, err = p.parse(reader); err != nil {
			return
		}
		if token != nil {
			index++
			token.LineNumber = lines[reader.Pos()-1]
			token.Index = index
			tokens = append(tokens, token)
		}
	}
	tokens = append(tokens, &token.Token{Type:token.End})
	return
}

func (p *Lexer) printStrlines(strs []string, lines []uint32) {
	f, _ := os.Create("build/test.strlines.text")
	defer f.Close()
	for i, str := range strs {
		//fmt.Fprintf(f, "i=%5d, line=%5d, str=%5.3s\n", i, lines[i], str)
		fmt.Fprintf(f, "i=%5d, line=%5d, str=%5s\n", i, lines[i], str)
	}
}

func (p *Lexer) Print(tokens []*token.Token) string {
	var str string
	for _, token := range tokens {
		str += fmt.Sprintf("token: index=%5d, line=%5d, type=%5s, val= %s\n",token.Index, token.LineNumber, token.Type, token.Val)
	}
	return str
}

func LexerNew() *Lexer {
	return &Lexer{
		parsers: parser.Parsers,
	}
}

type ChType uint8

func ChTypeNew(ch byte) ChType {
	if bytes.Contains([]byte("()[]{}"), []byte{ch}) {
		return ChTypeBracket
	} else if bytes.Contains([]byte(",.;:=+-*/!^&|<>?\\"), []byte{ch}) {
		return ChTypeSymbol
	} else if bytes.Contains([]byte("\""), []byte{ch}) {
		return ChTypeString
	} else if bytes.Contains([]byte("'"), []byte{ch}) {
		return ChTypeChar
	} else if bytes.Contains([]byte("\r\n"), []byte{ch}) {
		return ChTypeEndOfLine
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
	ChTypeEndOfLine
	ChTypeChar
)
