package stack_by_array

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	stack := New(3)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	el := stack.Pop()
	fmt.Println(el)
	stack.Push(100)
}
