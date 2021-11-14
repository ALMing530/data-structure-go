package single_linked_list

import (
	"fmt"
	"testing"
)

func Test_list(t *testing.T) {
	test_singly_linked_list()
}

func test_singly_linked_list() {
	fmt.Println("singly_linked_list")
	sl := new(Single_linked_list)
	// sl.insertAfter(1)
	// sl.insertAfter(2)
	// sl.insertAfter(4)
	sl.InsertBefore(1)
	sl.InsertBefore(2)
	sl.InsertBefore(3)
	sl.InsertBefore(4)
	sl.Insert(3, 8)
	sl.Delete(8)
	// sl.deleteIndex(2)
	print(sl)
	fmt.Println(sl.size)
}
func print(sl *Single_linked_list) {
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
