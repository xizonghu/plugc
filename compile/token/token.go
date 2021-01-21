package token

//type TokenType uint16

//constant
//TOKEN(TK_INTCONST,     "int")
//TOKEN(TK_UINTCONST,    "unsigned int")
//TOKEN(TK_LONGCONST,    "long")
//TOKEN(TK_ULONGCONST,   "unsigned long")
//TOKEN(TK_LLONGCONST,   "long long")
//TOKEN(TK_ULLONGCONST,  "unsigned long long")
//TOKEN(TK_FLOATCONST,   "float")
//TOKEN(TK_DOUBLECONST,  "double")
//TOKEN(TK_LDOUBLECONST, "long double")
//TOKEN(TK_STRING,       "STR")
//TOKEN(TK_WIDESTRING,   "WSTR")

const (
	//keywords
	TokenTypeAuto		= "auto"
	TokenTypeExtern		= "extern"
	TokenTypeRegister	= "register"
	TokenTypeStatic		= "static"
	TokenTypeTypeDef	= "typedef"
	TokenTypeConst		= "const"
	TokenTypeVolatile	= "volatile"
	TokenTypeSigned		= "signed"
	TokenTypeUnSigned	= "unsigned"
	TokenTypeShort		= "short"
	TokenTypeLong		= "long"
	TokenTypeChar		= "char"
	TokenTypeInt		= "int"
	TokenTypeInt64		= "__int64"
	TokenTypeFloat		= "float"
	TokenTypeDouble		= "double"
	TokenTypeEnum		= "enum"
	TokenTypeStruct		= "struct"
	TokenTypeUnion		= "union"
	TokenTypeVoid		= "void"
	TokenTypeIf			= "if"
	TokenTypeElse		= "else"
	TokenTypeFor		= "for"
	TokenTypeDo			= "do"
	TokenTypeWhile		= "while"
	TokenTypeBreak		= "break"
	TokenTypeContinue	= "continue"
	TokenTypeReturn		= "return"
	TokenTypeGoto		= "goto"
	TokenTypeSwitch		= "switch"
	TokenTypeCase		= "case"
	TokenTypeDefault	= "default"
	TokenTypeSizeof		= "sizeof"

	//identifier
	TokenTypeSymbol		= "sym"

	//constant
	TokenTypeString		= "str"
	TokenTypeNumber		= "num"

	//operators
	TokenTypeComma		= ","
	TokenTypeQuestion	= "?"
	TokenTypeColon		= ":"
	TokenTypeAssign		= "="
	TokenTypeBoolOr		= "||"
	TokenTypeBoolAnd	= "&&"
	TokenTypeBitsReverse		= "~"
	TokenTypeBoolEqual			= "=="
	TokenTypeBoolUnEqual		= "!="
	TokenTypeBoolGreater		= ">"
	TokenTypeBoolLess			= "<"
	TokenTypeBoolGreaterEqual	= ">="
	TokenTypeBoolLessEqual		= "<="
	TokenTypeBoolNot			= "!"
	TokenTypeBitsOr		= "|"
	TokenTypeBitsXor	= "^"
	TokenTypeBitsAnd	= "&"
	TokenTypeShiftLeft	= "<<"
	TokenTypeShiftRight	= ">>"
	TokenTypeAdd		= "+"
	TokenTypeSub		= "-"
	TokenTypeMul		= "*"
	TokenTypeDiv		= "/"
	TokenTypeMod		= "%"
	TokenTypeInc		= "++"
	TokenTypeDec		= "--"
	TokenTypeDot		= "."
	TokenTypePointer	= "->"
	TokenTypeParenLeft	= "("
	TokenTypeParenRight	= ")"
	TokenTypeBracketLeft		= "["
	TokenTypeBracketRight		= "]"

	//punctuators
	TokenTypeBraceLeft	= "{"
	TokenTypeBraceRight	= "}"
	TokenTypeSemicolon	= ";"
	TokenTypeEllipse	= "..."
	TokenTypePound		= "#"
	TokenTypeNewline	= "\n"
	TokenTypeEnd		= "EOF"
)

type Token struct {
	Type string
	Val interface{}
}

func (p *Token) SetVal(val interface{}) {
	p.Val = val
}

func TokenNew(Type string) *Token {
	return &Token {
		Type: Type,
	}
}