package tree

import (
	"testing"

	"github.com/areon546/go-helpers/helpers"
)

func TestNewTree(t *testing.T) {
	root := NewTreeNode("ROOT")
	tree := NewTree(root)

	// Root has no parents
	helpers.AssertEqualsObject(t, root, tree.root)
}

func TestRoot(t *testing.T) {
	root := NewTreeNode("ROOT")
	tree := NewTree(root)

	helpers.AssertEqualsObject(t, root, tree.Root())
	helpers.AssertEqualsObject(t, "ROOT", tree.Root().Value())
}

func TestParent(t *testing.T) {
	root := NewTreeNode("ROOT")
	tree := NewTree(root)

	// Parent
	t.Run("DEFAULT: Parent is nil", func(t *testing.T) {
		helpers.AssertEqualsBool(t, true, nil == tree.Root().Parent())
	})

	// AddParent
	t.Run("Add Parent overwrites the parent", func(t *testing.T) {
		parent := NewTreeNode("PARENT")
		root.AddParent(parent)
		helpers.AssertEqualsObject(t, parent, root.Parent())
	})
}

func TestChildren(t *testing.T) {
	root := NewTreeNode("ROOT")
	tree := NewTree(root)
	childCount := 0

	// Children
	t.Run("DEFAULT: 0 children", func(t *testing.T) {
		children := tree.Root().Children()
		helpers.AssertEqualsInt(t, childCount, len(children))
	})

	// AddChild
	t.Run("Add Child appends to children", func(t *testing.T) {
		child := NewTreeNode("CHILD")
		root.AddChild(child)
		childCount++

		// Updated in tree node
		children := root.Children()
		helpers.AssertEqualsInt(t, childCount, len(children)) //
		helpers.AssertEqualsObject(t, child, children[0])

		// Updated in tree
		children = tree.Root().Children()
		helpers.AssertEqualsInt(t, childCount, len(children)) //
		helpers.AssertEqualsObject(t, child, children[0])
	})
}
