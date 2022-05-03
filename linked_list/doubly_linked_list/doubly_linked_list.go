package doubly_linked_list

type node struct {
	val  interface{}
	pre  *node
	next *node
}

type DoublyLinkedList struct {
	head *node
	tail *node
	size int
}

func (list *DoublyLinkedList) InsertAfter(value interface{}) {
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

func (list *DoublyLinkedList) InsertBefore(value interface{}) {
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

func (list *DoublyLinkedList) Insert(index int, value interface{}) {
	node := &node{
		val: value,
	}
	el := list.position(index)
	if el != nil {
		el.next.pre = node
		node.next = el.next
		el.next = node
		node.pre = el
		list.size++
		if el.next == nil {
			list.tail = node
		}
	}

}

func (list *DoublyLinkedList) Delete(value interface{}) {
	current := list.head
	for {
		if current == nil {
			break
		}
		if current.val == value {
			current.pre.next = current.next
			list.size--
			//如果 next 为 nil 说明是 tail 更新 tail
			if current.next == nil {
				list.tail = current.pre
			}
		}
		current = current.next
	}
}
func (list *DoublyLinkedList) DeleteIndex(index int) {
	if index == list.size-1 {
		list.tail = list.tail.pre
		list.size--
		return
	}
	el := list.position(index)
	if el != nil {
		el.pre.next = el.next
		el = nil
		//如果 next 为 nil 说明是 tail 更新 tail
		if el.next == nil {
			list.tail = el.pre
		}
	}
}
func (list *DoublyLinkedList) position(index int) *node {
	len := list.size
	if index >= len {
		panic("index out of bound")
	}
	current := list.head
	for i := 0; i < len; i++ {
		if index == i {
			return current
		}
		current = current.next
	}
	return nil
}
