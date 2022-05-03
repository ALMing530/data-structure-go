package stack_by_linked_list

import "ds/linked_list/singly_linked_list"

type Stack struct {
	element *singly_linked_list.SinglyLinkedList
	capcity int
}

func New(capcity int) *Stack {
	return &Stack{
		element: new(singly_linked_list.SinglyLinkedList),
		capcity: capcity,
	}
}

func (stack *Stack) Push(element interface{}) {
	if stack.element.Size == stack.capcity {
		panic("Stack overflow")
	}
	stack.element.InsertBefore(element)
}

func (stack *Stack) Pop() interface{} {
	if stack.element.Size == 0 {
		panic("Stack is already empty")
	}
	element := stack.element.First()
	stack.element.DeleteFirst()
	return element
}
