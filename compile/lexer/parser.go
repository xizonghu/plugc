package lexer

import (
	"bytes"
	"plugc/compile/token"
)

type IParser interface {
	GetKey() string
	Process(text *bytes.Buffer) (*token.Token, error)
}
