package utils

import (
	"plugc/compiler/token"
)

func MatchType(reader token.ITokenReader) (str string, err error) {
	var tk *token.Token
	for tk = reader.Read(); tk != nil; tk = reader.Next() {
		if tk == nil || IsSemicolon(tk) {
			return
		} else if IsDecl(tk) {
			str += tk.Val
		} else {
			break
		}
	}
	return
}

func MatchSymbol(reader token.ITokenReader) (str string, err error) {
	var tk *token.Token

	if tk = reader.Read(); tk == nil || tk.Type != token.Symbol {
		return
	}

	str = tk.Val
	reader.Next()
	return
}

func MatchParenLeft(reader token.ITokenReader) (str string, err error) {
	tk := reader.Read()
	if IsParenLeft(tk) {
		str = tk.Val
		reader.Next()
	}
	return
}

func MatchParenRight(reader token.ITokenReader) (str string, err error) {
	tk := reader.Read()
	if IsParenRight(tk) {
		str = tk.Val
		reader.Next()
	}
	return
}

func MatchDecl(reader token.ITokenReader) (typ string, name string, err error) {
	if typ, err = MatchType(reader); err != nil || typ == "" {
		return
	}
	if name, err = MatchSymbol(reader); err != nil || name == "" {
		return
	}
	return
}

func MatchComma(reader token.ITokenReader) (str string, err error) {
	tk := reader.Read()
	if IsComma(tk) {
		str = tk.Val
		reader.Next()
	}
	return
}

func MatchFuncCall(reader token.ITokenReader) (str string, err error) {
	tk := reader.Read()
	if !IsSymbol(tk) {
		return
	}
	funName := tk.Val
	tk = reader.Next()
	if !IsParenLeft(tk) {
		return
	}
	str = funName
	return
}