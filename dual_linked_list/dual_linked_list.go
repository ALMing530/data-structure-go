package main

type node struct {
	val  interface{}
	pre  *node
	next *node
}

type dual_linked_list struct {
	head *node
	tail *node
	size int64
}

func (list *dual_linked_list) insertAfter(value interface{}) {
	node := &node{
		val: value,
	}
	if list.head == nil {
		list.head = node
		list.tail = node
		list.size++
	} else {
		list.tail.next = node
		node.pre = list.tail
		list.tail = node
		list.size++
	}
}

func (list *dual_linked_list) insertBefore(value interface{}) {
	node := &node{
		val: value,
	}
	if list.head == nil {
		list.head = node
		list.tail = node
		list.size++
	} else {
		node.next = list.head
		list.head.pre = node
		list.head = node
		list.size++
	}
}

func (list *dual_linked_list) delete(value interface{}) {

}
