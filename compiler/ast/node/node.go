package node

import (
	"encoding/json"
	"fmt"
)

const (
	Define				= "define"
	//Segment				= "segment"
	Control				= "control"
	Operate				= "operate"
)

type IAstNode interface {
	GetType() string
	Children() []IAstNode
	String() string
}

type Location struct {
	File string
	Line uint32
	Col uint32
}

func (p *Location) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"Location": fmt.Sprintf("%s:%d:%d", p.File, p.Line, p.Col),
	})
}

type Node struct {
	Kind string
	//Locate Location
	off uint32
	children []IAstNode `json:"children"`
}

func (p *Node) GetType() string {
	return p.Kind
}


func (p *Node) Children() []IAstNode {
	return p.children
}

func (p *Node) String() string {
	b, _ := json.Marshal(p)
	return string(b)
}

func NodeInit(typ string) Node {
	return Node{
		Kind:      typ,
		off:      0,
		children: nil,
	}
}

type WithName struct {
	Node
	name string `json:name`
}

func WithNameInit(typ string, name string) WithName {
	return WithName{
		Node: NodeInit(typ),
		name: name,
	}
}

type WithValue struct {
	WithName
	value interface{} `json:value`
}

func WithValueInit(typ string, name string, value interface{}) WithValue {
	return WithValue{
		WithName: WithNameInit(typ, name),
		value:    value,
	}
}