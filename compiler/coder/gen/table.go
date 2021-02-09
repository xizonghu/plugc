package gen

type Variable struct {
	Addr uint32
}

type Table struct {
	Symbols map[string]*SymbolInfo
	GlobalVars map[string]*VarInfo
	LocalVar *LocalVarNode
	Functions map[string]*FunctionInfo
	Consts map[string]*VarInfo
}

type SymbolInfo struct {
	Name string
}

type VarInfo struct {
	Name string
	Size uint32
	Addr uint32
}

type FunctionInfo struct {
	Name string
	Addr uint32
}

type LocalVarNode struct {
	Vars map[string]*VarInfo
	Father *LocalVarNode
	Child *LocalVarNode
}