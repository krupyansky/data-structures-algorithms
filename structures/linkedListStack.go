package main

import "fmt"

type linkedListStack struct {
	linkedList *singlyLinkedList
}

func newLinkedListStack(linkedList singlyLinkedList) *linkedListStack {
	return &linkedListStack{linkedList: &linkedList}
}

func (s *linkedListStack) peek() interface{} {
	return s.linkedList.head.value
}

func (s *linkedListStack) pop() {
	s.linkedList.removeFirst()
}

func (s *linkedListStack) push(value interface{}) {
	s.linkedList.addFirst(value)
}

func main() {
	list := newSinglyLinkedList()
	stack := newLinkedListStack(list)

	stack.push(6)
	stack.push(10)

	fmt.Println(stack.peek())

	stack.pop()

	fmt.Println(stack.peek())
}
