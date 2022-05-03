package binary_tree

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	/**
	*
	*				1
	*			2		3
	*		5		7
	*
	 */
	tree := New(1)
	tree.InsertLeft(tree.GetRoot(), 2)
	tree.InsertRight(tree.GetRoot(), 3)
	tree.InsertLeft(tree.GetRoot().left, 5)
	tree.InsertRight(tree.GetRoot().left, 7)
	tree.preOrder(tree.root)
	fmt.Println("--------------------")
	tree.postOrder(tree.root)
	fmt.Println("--------------------")
	tree.inOrder(tree.root)
}
