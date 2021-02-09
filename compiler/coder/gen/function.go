package gen

import (
	"plugc/compiler/ast/node"
	"plugc/compiler/ins"
	"reflect"
	"unsafe"
)

type FunctionGen struct {

}

func (p *FunctionGen) Parse(table *Table, nd node.IAstNode, buffer IBuffer) (err error) {
	funcNode := (*node.FunctionDecl)(unsafe.Pointer(reflect.ValueOf(nd).Pointer()))
	funcName := (*node.DataDecl)(unsafe.Pointer(reflect.ValueOf(funcNode.Output).Pointer()))
	addr := buffer.Offset()
	table.Functions[funcName.Name] = &FunctionInfo{
		Name: funcName.Name,
		Addr: addr,
	}

	buffer.Write(ins.FuncEntryNew().Bytes())
	return
}