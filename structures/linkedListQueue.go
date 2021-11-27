package main

import "fmt"

type linkedListQueue struct {
	linkedList *singlyLinkedList
}

func newLinkedListQueue(linkedList singlyLinkedList) *linkedListQueue {
	return &linkedListQueue{linkedList: &linkedList}
}

func (q *linkedListQueue) enqueue(value interface{}) {
	q.linkedList.addLast(value)
}

func (q *linkedListQueue) dequeue() {
	q.linkedList.removeFirst()
}

func (q *linkedListQueue) peek() interface{} {
	return q.linkedList.head.value
}

func main() {
	list := newSinglyLinkedList()
	queue := newLinkedListQueue(list)

	queue.enqueue(6)
	queue.enqueue(10)

	fmt.Println(queue.peek())

	queue.dequeue()

	fmt.Println(queue.peek())
}
