package lexer

import (
	"bytes"
	"fmt"
	"plugc/compiler/token"
	"testing"
)

const (
	sample1 = "void  main()   {char cc=123;cc +=12;printf(\"a  b, \\\"w'o'r\\\"ld\n\");}"
)

func TestFormat(t *testing.T) {
	text := bytes.NewBufferString(sample1)
	fmt.Printf("b=%s\n", text.String())
	lexer := LexerNew()
	result := lexer.split(text)
	fmt.Printf("====\n")
	for i, str := range result {
		fmt.Printf("i=%d, str=%s\n", i, str)
	}
}

func TestLexer_Run(t *testing.T) {
	text := bytes.NewBufferString(sample1)
	lexer := LexerNew()
	var tokens []*token.Token
	var err error
	if tokens, err = lexer.Run(text); err != nil {
		t.Error(err)
		//return
	}
	for i, token := range tokens {
		fmt.Printf("token: i=%d type=%s, val=%s\n", i, token.Type, token.Val)
	}
}
