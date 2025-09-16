package tree

import (
	"github.com/areon546/ds/node"
)

type Tree struct {
	root *TreeNode
}

func NewTree(node *TreeNode) Tree {
	return Tree{root: node}
}

func (tr Tree) Root() *TreeNode {
	return tr.root
}

type TreeNode struct {
	node.Node

	parent   *TreeNode
	children []*TreeNode
}

func NewTreeNode(data any) *TreeNode {
	node := node.NewNode(data)

	return &TreeNode{Node: node, parent: nil, children: []*TreeNode{}}
}

func (tn *TreeNode) Parent() *TreeNode {
	if tn.parent != nil {
		return tn.parent
	}

	return nil
}

func (tn *TreeNode) AddParent(n *TreeNode) {
	tn.parent = n
}

func (tn *TreeNode) Children() []*TreeNode {
	return tn.children
}

func (tn *TreeNode) AddChild(kid *TreeNode) {
	tn.children = append(tn.children, kid)
}
