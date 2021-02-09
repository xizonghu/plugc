package ast

//abstract syntax tree

const (
	TranslationUnit = ""
	//Function        = "func"
	//Declaration     = "declaration"

	TypeName             = ""
	Specifiers           = ""
	Token                = ""
	TypedefName          = "typedef"
	EnumSpecifier        = ""
	Enumerator           = ""
	StructSpecifier      = "struct"
	UnionSpecifier       = "union"
	StructDeclaration    = ""
	StructDeclarator     = ""
	PointerDeclarator    = "point"
	ArrayDeclarator      = "array"
	FunctionDeclarator   = ""
	ParameterTypeList    = ""
	ParameterDeclaration = ""
	NameDeclarator       = ""
	InitDeclarator       = ""
	Initializer          = ""

	Expression = ""

	ExpressionStatement = ""
	LabelStatement      = ""
	CaseStatement       = "case"
	DefaultStatement    = "default"
	IfStatement         = "if"
	SwitchStatement     = "switch"
	WhileStatement      = "while"
	DoStatement         = "do"
	ForStatement        = "for"
	GotoStatement       = "goto"
	BreakStatement      = "break"
	ContinueStatement   = "continue"
	ReturnStatement     = "return"
	CompoundStatement   = ""
)
