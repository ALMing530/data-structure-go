package queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := New(5)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)
	// queue.Enqueue(6)
	element := queue.Dequeue()
	fmt.Print(element)

}
