package token

const (
	KEY_AUTO		= "auto"
	KEY_STATIC		= "static"

	KEY_SHORT		= "short"
	KEY_LONG		= "long"

	KEY_UNSIGNED	= "unsigned"
	KEY_SIGNED		= "signed"

	KEY_VOID		= "void"
	KEY_CHAR		= "char"
	KEY_INT			= "int"

	KEY_IF			= "if"
	KEY_ELSE		= "else"
)

var Decls = []string{
	KEY_AUTO, KEY_STATIC,
	KEY_SHORT, KEY_LONG,
	KEY_UNSIGNED, KEY_SIGNED,
	KEY_VOID, KEY_CHAR,
	KEY_INT,
}

var conds = []string{
	KEY_IF, KEY_ELSE,
}

const (
	LIM_BRACE_LEFT			= "{"
	LIM_BRACE_RIGHT			= "}"
	LIM_BRACK_LEFT			= "["
	LIM_BRACK_RIGHT			= "]"
	LIM_PAREN_LEFT			= "("
	LIM_PAREN_RIGHT			= ")"
	LIM_SEMICOLON			= ";"
)

const (
	OP_COMMA			= ","
	OP_DOT				= "."
	OP_POINTRT			= "->"
	OP_ADD				= "+"
	OP_ASSIGN			= "="
	OP_BOOL_NOT			= "!"
)

var OperatorArray = []string{
	OP_COMMA, OP_DOT, OP_POINTRT, OP_ASSIGN,
	OP_ADD,
	OP_BOOL_NOT,
}

var Keywords = []string{
	"auto",
	"extern",
	"register",
	"static",
	"typedef",
	"const",
	"volatile",
	"signed",
	"unsigned",
	"short",
	"long",
	"char",
	"int",
	"__int64",
	"float",
	"double",
	"enum",
	"struct",
	"union",
	"void",
	"if",
	"else",
	"for",
	"do",
	"while",
	"break",
	"continue",
	"return",
	"goto",
	"switch",
	"case",
	"default",
	"sizeof",
}

var Operators = "+-*/%&|^!<>=:?."

var Delimiters = ";()[]{}"





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
