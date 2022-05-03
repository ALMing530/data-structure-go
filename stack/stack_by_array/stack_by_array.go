package stack_by_array

type Stack struct {
	element []interface{}
	index   int
	capcity int
}

func New(capcity int) *Stack {
	return &Stack{
		element: make([]interface{}, capcity),
		index:   -1,
		capcity: capcity,
	}
}

func (stack *Stack) Push(element interface{}) {
	if stack.index == stack.capcity-1 {
		panic("Stack overflow")
	}
	stack.index++
	stack.element[stack.index] = element
}
func (stack *Stack) Pop() interface{} {
	if stack.index == -1 {
		panic("Stack is already empty")
	}
	element := stack.element[stack.index]
	stack.index--
	return element
}
