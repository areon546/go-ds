package node

type Node struct {
	data any
}

func NewNode(data any) Node {
	return Node{data}
}

func (n Node) Value() any {
	return n.data
}
