package queue

import "ds/linked_list/singly_linked_list"

type Queue struct {
	element *singly_linked_list.SinglyLinkedList
	capcity int
}

func New(capcity int) *Queue {
	return &Queue{
		element: new(singly_linked_list.SinglyLinkedList),
		capcity: capcity,
	}
}

func (queue *Queue) Enqueue(element interface{}) {
	if queue.element.Size >= queue.capcity {
		panic("Queue is already full")
	}
	queue.element.InsertAfter(element)
}

func (queue Queue) Dequeue() interface{} {
	if queue.element.Size <= 0 {
		panic("Queue is already empty")
	}
	element := queue.element.First()
	queue.element.DeleteFirst()
	return element
}
