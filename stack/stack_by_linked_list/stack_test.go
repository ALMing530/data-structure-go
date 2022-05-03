package stack_by_linked_list

import "testing"

func TestStack(t *testing.T) {
	stack := New(3)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Pop()
}
