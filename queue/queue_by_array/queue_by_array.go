package queue_by_array

type Queue struct {
	element []interface{}
	index   int
	capcity int
}

func New(capcity int) *Queue {
	return &Queue{
		element: make([]interface{}, capcity),
		index:   -1,
		capcity: capcity,
	}
}

func (queue *Queue) Enqueue(element interface{}) {
	if queue.index >= queue.capcity-1 {
		panic("Queue is already full")
	}
	queue.index++
	queue.element[queue.index] = element
}
func (queue *Queue) Dequeue() interface{} {
	if queue.index <= -1 {
		panic("Queue is already empty")
	}
	element := queue.element[0]
	for i := 0; i <= queue.index-1; i++ {
		queue.element[i] = queue.element[i+1]
	}
	queue.index--
	return element
}
