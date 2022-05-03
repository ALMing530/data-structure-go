package singly_linked_list

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	// test_singly_linked_list()
	// test_deleteFirst()
	test_deleteLast()
}
func test_deleteFirst() {
	list := initList()
	list.DeleteFirst()
	print(list)
}
func test_deleteLast() {
	list := initList()
	list.deleteLast()
	print(list)
}
func test_singly_linked_list() {
	fmt.Println("singly_linked_list")
	list := initList()
	list.Insert(3, 8)
	list.Delete(8)
	// sl.deleteIndex(2)
	print(list)
	fmt.Println(list.Size)
}
func initList() *SinglyLinkedList {
	list := new(SinglyLinkedList)
	list.InsertBefore(1)
	list.InsertBefore(2)
	list.InsertBefore(3)
	list.InsertBefore(4)
	return list
}
func print(sl *SinglyLinkedList) {
	current := sl.head
	for {
		if current != nil {
			fmt.Println(current.val)
			current = current.next
		} else {
			break
		}
	}
}
