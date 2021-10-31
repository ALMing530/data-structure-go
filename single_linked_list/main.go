package main

import (
	"fmt"
)

func main() {
	test_singly_linked_list()
	// time.Sleep(10 * time.Minute)
}

func test_singly_linked_list() {
	fmt.Println("singly_linked_list")
	sl := new(single_linked_list)
	// sl.insertAfter(1)
	// sl.insertAfter(2)
	// sl.insertAfter(4)
	sl.insertBefore(1)
	sl.insertBefore(2)
	sl.insertBefore(3)
	sl.insertBefore(4)
	sl.insert(3, 8)
	sl.delete(8)
	// sl.deleteIndex(2)
	print(sl)
	fmt.Println(sl.size)
}
func print(sl *single_linked_list) {
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
