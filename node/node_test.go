package node_test

import (
	"testing"

	"github.com/areon546/go-ds/node"
	"github.com/areon546/go-helpers/helpers"
)

func TestValue(t *testing.T) {
	t.Run("Test String", func(t *testing.T) {
		strNode := node.NewNode("STRING")

		helpers.AssertEqualsObject(t, "STRING", strNode.Value())
	})

	t.Run("Test Int", func(t *testing.T) {
		intNode := node.NewNode(1)
		helpers.AssertEqualsObject(t, 1, intNode.Value())
	})
}
