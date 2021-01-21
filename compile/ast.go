package compile

type AstNode struct {
	kind int
	next *AstNode
}
