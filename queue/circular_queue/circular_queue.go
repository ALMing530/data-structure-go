package circular_queue

/**
 * 循环队列
 */

type CicularQueue struct {
	element []interface{}
	head    int
	tail    int
	size    int
	capcity int
}

func New(capcity int) *CicularQueue {
	return &CicularQueue{
		element: make([]interface{}, capcity),
		head:    0,
		tail:    0,
		size:    0,
		capcity: capcity,
	}
}

func (queue *CicularQueue) Enqueue(element interface{}) {
	if queue.size == queue.capcity {
		panic("Queue is already full")
	}
	queue.element[queue.tail] = element
	queue.tail++
	queue.tail = queue.tail % queue.capcity
	queue.size++
}
func (queue *CicularQueue) Dequeue() interface{} {
	if queue.size == 0 {
		panic("Queue is already empty")
	}
	element := queue.element[queue.head]
	queue.head++
	queue.head = queue.head % queue.capcity
	queue.size--
	return element
}
