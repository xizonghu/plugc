package utils

import "plugc/compiler/token"

func IsDecl(tk *token.Token) (checked bool) {
	for _, key := range token.Decls {
		if key == tk.Val {
			checked = true
			return
		}
	}
	return
}

func IsParenLeft(tk *token.Token) (checked bool) {
	if tk.Val == token.LIM_PAREN_LEFT {
		checked = true
	}
	return
}

func IsParenRight(tk *token.Token) (checked bool) {
	if tk.Val == token.LIM_PAREN_RIGHT {
		checked = true
	}
	return
}

func IsComma(tk *token.Token) (checked bool) {
	if tk.Val == token.OP_COMMA {
		checked = true
	}
	return
}

func IsSemicolon(tk *token.Token) (checked bool) {
	if tk.Val == token.LIM_SEMICOLON {
		checked = true
	}
	return
}

func IsSymbol(tk *token.Token) (checked bool) {
	if tk.Type == token.Symbol {
		checked = true
	}
	return
}

func IsKeyword(tk *token.Token) (checked bool) {
	if tk.Type == token.Keyword {
		checked = true
	}
	return
}

func IsBraceLeft(tk *token.Token) (checked bool) {
	if tk.Val == token.LIM_BRACE_LEFT {
		checked = true
	}
	return
}

func IsBraceRight(tk *token.Token) (checked bool) {
	if tk.Val == token.LIM_BRACE_RIGHT {
		checked = true
	}
	return
}

func IsOperator(tk *token.Token) (checked bool) {
	for _, op := range token.OperatorArray {
		if tk.Val == op {
			checked = true
			break
		}
	}
	return
}

func IsConstant(tk *token.Token) (checked bool) {
	if tk.Type == token.Constant {
		checked = true
	}
	return
}



