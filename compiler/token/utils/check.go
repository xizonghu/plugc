package utils

import (
	"plugc/compiler/token"
)

func CheckType(reader token.ITokenReader) (checked bool) {
	var tk *token.Token
	for tk = reader.Read(); ; tk = reader.Next() {
		if tk != nil || IsSemicolon(tk) {
			return
		} else if IsKeyword(tk) {
			break
		}
	}
	return
}

func CheckVarDecl(reader token.ITokenReader) (checked bool) {
	if typ, name, err := MatchDecl(reader); typ == "" || name == "" || err != nil {
		return
	}
	checked = true
	return
}