package main

import "fmt"

func main() {
	ls := new(dual_linked_list)
	ls.insertBefore(1)
	ls.insertBefore(2)
	ls.insertBefore(3)
	ls.insertBefore(4)

	fmt.Println()
}
