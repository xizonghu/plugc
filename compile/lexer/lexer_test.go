package lexer

import (
	"bytes"
	"fmt"
	"testing"
)

func TestFormat(t *testing.T) {
	text := bytes.NewBufferString("void  main()   {char cc=123;cc +=12;printf(\"a  b, \\\"w'o'r\\\"ld\n\");}")
	fmt.Printf("b=%s\n", text.String())
	lexer := LexerNew()
	result := lexer.format(text)
	fmt.Printf("result=%s\n", result.String())
}
