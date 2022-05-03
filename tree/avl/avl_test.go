package avl

import (
	"testing"
)

func TestAVL(t *testing.T) {

}
func TestLeftHeight(t *testing.T) {
	avl := new(AVLTree)
	avl.Insert(12)
	avl.Insert(14)
	avl.Insert(4)
	avl.Insert(13)
	avl.Insert(15)
	avl.Insert(3)
	avl.Insert(16)
	// avl.Delete(3)
	avl.Delete(4)
}

func TestLeftHeight_1(t *testing.T) {
	avl := new(AVLTree)
	avl.Insert(12)
	avl.Insert(14)
	avl.Insert(4)
	avl.Insert(13)
	avl.Insert(15)
	avl.Insert(5)
	avl.Insert(16)
	// avl.Delete(3)
	avl.Delete(5)
}

func TestLeftHeight_2(t *testing.T) {
	avl := new(AVLTree)
	avl.Insert(12)
	avl.Insert(6)
	avl.Insert(14)
	avl.Insert(4)
	avl.Insert(9)
	avl.Insert(13)
	avl.Insert(15)
	avl.Insert(5)
	avl.Insert(8)
	avl.Insert(10)
	avl.Insert(7)
	avl.Delete(6)
}

func TestRightWeight(t *testing.T) {
	avl := new(AVLTree)
	avl.Insert(8)
	avl.Insert(10)
	avl.Insert(9)
}
