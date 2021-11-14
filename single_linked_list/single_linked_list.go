package single_linked_list

type node struct {
	val  interface{}
	next *node
}

type Single_linked_list struct {
	head *node
	tail *node
	size int
}

//在链表尾添加
func (list *Single_linked_list) InsertAfter(value interface{}) {
	node := &node{
		val:  value,
		next: nil,
	}
	if list.head == nil {
		list.head = node
		list.tail = node
		list.size++
	} else {
		list.tail.next = node
		list.tail = node
		list.size++
	}
}

//在链表头添加
func (list *Single_linked_list) InsertBefore(value interface{}) {
	node := &node{
		val: value,
	}
	if list.head == nil {
		list.head = node
		list.tail = node
		list.size++
	} else {
		node.next = list.head
		list.head = node
		list.size++
	}
}

//在指定索引处添加
func (list *Single_linked_list) Insert(index int, value interface{}) {
	if index >= list.size {
		panic("index out of bounds")
	}
	node := &node{
		val: value,
	}
	//如果索引是链表长度-1则说明插入位置是最后一个元素，直接操作tail
	if index == list.size-1 {
		list.tail.next = node
		list.tail = node
		list.size++
		return
	}
	_, el := list.position(index)
	node.next = el.next
	el.next = node

}

//通过值删除，如果有重复值只会删除第一个
func (list *Single_linked_list) Delete(value interface{}) {
	var pre *node
	current := list.head
	for {
		//current 为 nil 表示遍历到最后一个节点还没找到相同值
		if current == nil {
			break
		}
		if current.val == value {
			pre.next = current.next
			list.size--
			break
		}
		pre = current
		current = current.next
	}
}

//删除指定索引处元素
func (list *Single_linked_list) DeleteIndex(index int) {
	pre, current := list.position(index)
	pre.next = current.next
	current = nil
}

//根据索引定位元素以及该元素的前置元素，删除时需要前置元素。（将前置元素与该元素后置元素链接）
//因为是单向链表所以仅获取索引所在节点无法获得其前置节点
func (list *Single_linked_list) position(index int) (pre *node, current *node) {
	size := list.size
	pre = list.head
	for i := 0; i < size; i++ {
		//定位到索引元素的前一个元素，那么next就是目标元素
		if i == index-1 {
			return pre, pre.next
		}
		pre = pre.next
	}
	return nil, nil
}
