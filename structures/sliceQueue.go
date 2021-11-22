package main

import "fmt"

type sliceQueue struct {
	Queueable
	items []interface{}
}

func newSliceQueue(cap int) sliceQueue {
	return sliceQueue{items: make([]interface{}, 0, cap)}
}

type Queueable interface {
	enqueue(value interface{})
	dequeue()
	peek() interface{}
}

func (q *sliceQueue) enqueue(value interface{}) {
	q.items = append(q.items, value)
}

func (q *sliceQueue) dequeue() {
	if len(q.items) == 0 {
		panic("invalid operation `dequeue` for empty queue")
	}
	q.items = q.items[1:]
}

func (q *sliceQueue) peek() interface{} {
	if len(q.items) == 0 {
		panic("invalid operation `peek` for empty queue")
	}
	return q.items[0]
}

func main() {
	stack := newSliceQueue(8)

	stack.enqueue(6)
	stack.enqueue(10)

	fmt.Println(stack.peek())

	stack.dequeue()

	fmt.Println(stack.peek())
}
