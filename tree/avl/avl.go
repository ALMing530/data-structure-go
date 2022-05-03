package avl

import (
	"math"
)

type node struct {
	val    int
	parent *node
	left   *node
	right  *node
	//平衡因子 （左子树高度-右子树高度）
	blance int
}

type AVLTree struct {
	root *node
}

func (tree *AVLTree) Insert(element int) {
	node := &node{
		val: element,
	}
	if tree.root == nil {
		tree.root = node
		return
	}
	current := tree.root
	for {
		if node.val > current.val {
			if current.right == nil {
				current.right = node
				node.parent = current
				//插入节点后重新计算其父节点的平衡因子（左子树树高-右子树树高）
				//且在根节点插入时不需要计算，即 current.parent == nil
				if current.parent != nil {
					current.parent.blance = calcBlance(current.parent)
					if math.Abs(float64(current.parent.blance)) == 2 {
						//TODO 重平衡
						tree.reblance(current.parent)
					}
				}
				break
			}
			current = current.right
		} else {
			if current.left == nil {

				current.left = node
				node.parent = current
				//插入节点后重新计算其父节点的平衡因子（左子树树高-右子树树高）
				//且在根节点插入时不需要计算，即 current.parent == nil
				if current.parent != nil {
					current.parent.blance = calcBlance(current.parent)
					if math.Abs(float64(current.parent.blance)) == 2 {
						//TODO 重平衡
						tree.reblance(current.parent)
					}
				}
				break
			}
			current = current.left
		}
	}
}

/**
	假设使用左后继节点补位
	被删除节点：
	1 无子节点，直接删除然后重平衡
	2 仅有一个左/右子节点，由左/右子节点替换被删除节点然后重平衡
	3 左右子节点均存在，由左子节点替换被删除节点，然后重平衡
	注： 2，3情况使用得替换节点实际为被删除节点得后继节点，即按中
		序遍历紧随被删除节点的节点
**/
func (tree *AVLTree) Delete(element int) {
	current := tree.root

	for {
		if current == nil {
			break
		}
		if element > current.val {
			current = current.right
		}
		if element < current.val {
			current = current.left
		}
		if element == current.val {
			parent := current.parent
			//被删除节点是叶子节点
			if current.left == nil && current.right == nil {

				if parent.left == current {
					parent.left = nil
				}
				if parent.right == current {
					parent.right = nil
				}
				tree.reblanceUpward(parent)
			} else {
				if current.left != nil {
					//找到前驱节点
					preNode := tree.findPrecursorNode(current.left)
					preNodeParent := preNode.parent
					//将找到的前驱节点从其父节点中剥离
					if preNodeParent.left == preNode {
						preNodeParent.left = nil
					}
					if preNodeParent.right == preNode {
						preNodeParent.right = nil
					}
					//使用前驱节点替换当前节点
					preNode.left = current.left
					preNode.right = current.right
					if parent.left == current {
						parent.left = preNode
						preNode.parent = parent
					}
					if parent.right == current {
						parent.right = preNode
						preNode.parent = parent
					}
					tree.reblanceUpward(preNode)
					current = nil
				} else {
					current.val = current.right.val
					current.left = current.right.left
					current.right = current.right.right
					// 如果current没有左节点那么右节点高最高为 1 即一个节点，删除该节点 current不会失衡
					// 直接由其父节点逐级向上平衡
					tree.reblanceUpward(current.parent)
				}
			}
			break
		}
	}
}

func (tree *AVLTree) findPrecursorNode(node *node) *node {
	if node.left == nil && node.right == nil {
		return node
	}
	if node.right == nil {
		return tree.findPrecursorNode(node.left)
	}
	return tree.findPrecursorNode(node.right)
}

func (tree *AVLTree) reblance(node *node) {
	//左树高
	if node.blance > 0 {
		//新增节点在左子树左节点
		if node.left.left != nil {
			tree.rightRoate(node)
		} else {
			//新增节点在左子树右节点
			tree.leftRotate(node.left)
			tree.rightRoate(node)
		}
		//右树高
	} else {
		//新增节点在右子树右节点
		if node.right.right != nil {
			tree.leftRotate(node)
		} else {
			//新增节点在右子树左节点
			tree.rightRoate(node.right)
			tree.leftRotate(node)
		}
	}
	node.blance = 0
}

func (tree *AVLTree) reblanceUpward(node *node) {
	current := node
	for {
		if current == nil {
			break
		}
		current.blance = calcBlance(current)
		if math.Abs(float64(current.blance)) == 2 {
			tree.reblance(current)
			// 如果触发重平衡那么，current将不再是当前子树的父节点。详细见左右
			// 旋代码 node.parent = xxx
			current = current.parent.parent
		} else {
			current = current.parent
		}
	}
}
func (tree *AVLTree) leftRotate(node *node) {
	parent := node.parent
	right := node.right
	node.right = right.left
	right.left = node
	node.parent = right
	//重平衡子树父节点为root
	if node == tree.root {
		tree.root = right
		tree.root.parent = nil
	} else {
		if parent.left == node {
			parent.left = right
			right.parent = parent
		} else {
			parent.right = right
			right.parent = parent
		}
	}
}

func (tree *AVLTree) rightRoate(node *node) {
	parent := node.parent
	left := node.left
	node.left = left.right
	left.right = node
	node.parent = left
	//重平衡子树父节点为root
	if node == tree.root {
		tree.root = left
		tree.root.parent = nil
	} else {
		if parent.left == node {
			parent.left = left
			left.parent = parent
		} else {
			parent.right = left
			left.parent = parent
		}
	}
}

func calcBlance(node *node) int {
	return height(node.left) - height(node.right)
}

func height(node *node) int {
	if node == nil {
		return 0
	}
	return max(height(node.left), height(node.right)) + 1

}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
