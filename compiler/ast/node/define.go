package node

type Void struct {
	WithName
}

func VoidNew(name string) *Void {
	return &Void{
		WithName:WithNameInit("void", name),
	}
}

type String struct {
	WithValue
}

func StringNew(name string, val string) *String {
	return &String{
		WithValue: WithValueInit("string", name, val),
	}
}