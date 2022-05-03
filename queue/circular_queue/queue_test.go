package circular_queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := New(3)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Dequeue()
	queue.Enqueue(100)
	fmt.Print()
}
