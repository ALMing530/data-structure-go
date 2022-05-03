package binary_tree

import (
	"fmt"
)

type node struct {
	val   interface{}
	left  *node
	right *node
}

type BinaryTree struct {
	root *node
}

func New(root interface{}) *BinaryTree {
	rootNode := &node{
		val: root,
	}
	return &BinaryTree{
		root: rootNode,
	}
}

func (tree *BinaryTree) GetRoot() *node {
	return tree.root
}

func (tree *BinaryTree) InsertLeft(parent *node, element interface{}) {
	node := &node{
		val: element,
	}
	parent.left = node
}
func (tree *BinaryTree) InsertRight(parent *node, element interface{}) {
	node := &node{
		val: element,
	}
	parent.right = node
}

// 前序遍历
func (tree *BinaryTree) preOrder(node *node) {
	if node == nil {
		return
	}
	fmt.Println(node.val)
	tree.preOrder(node.left)
	tree.preOrder(node.right)
}

// 后续遍历
func (tree *BinaryTree) postOrder(node *node) {
	if node == nil {
		return
	}
	tree.postOrder(node.left)
	tree.postOrder(node.right)
	fmt.Println(node.val)
}

// 中序遍历
func (tree *BinaryTree) inOrder(node *node) {
	if node == nil {
		return
	}
	tree.inOrder(node.left)
	fmt.Println(node.val)
	tree.inOrder(node.right)
}
