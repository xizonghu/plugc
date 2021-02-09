package node

type DataElement struct {
	Node
	Name string
	Type string
	Value string
	Signed bool
	Storage string
	Size uint32
	Load *Segment
}

func DataElementNew(typ string, name string) *DataElement {
	return &DataElement{
		Node: NodeInit("DataElement"),
		Type: typ,
		Name: name,
	}
}

type DataDecl struct {
	Node
	Name string
	Type string
}

func DataDeclNew(typ string, name string) *DataDecl {
	return &DataDecl{
		Node: NodeInit("DataDecl"),
		Name: name,
		Type: typ,
	}
}