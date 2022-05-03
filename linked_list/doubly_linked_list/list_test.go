package doubly_linked_list

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	ls := new(DoublyLinkedList)
	ls.InsertBefore(1)
	ls.InsertBefore(2)
	ls.InsertBefore(3)
	ls.InsertBefore(4)
	// ls.delete(1)
	// ls.insert(1, 6)
	ls.DeleteIndex(3)
	fmt.Println()
}
