package main

import "fmt"

type doublyLinkedList struct {
	head  *doublyNode
	tail  *doublyNode
	count int
}

type doublyNode struct {
	value interface{}
	next  *doublyNode
	prev  *doublyNode
}

func newDoublyLinkedList() doublyLinkedList {
	return doublyLinkedList{}
}

func (l *doublyLinkedList) addFirst(value interface{}) {
	newHeadNode := &doublyNode{value: value}

	tmpNode := l.head

	l.head = newHeadNode

	l.head.next = tmpNode

	l.count++

	if l.count == 1 {
		l.tail = l.head
	} else {
		l.head.next.prev = newHeadNode
	}
}

func (l *doublyLinkedList) addLast(value interface{}) {
	newTailNode := &doublyNode{value: value}

	if l.count == 0 {
		l.head = newTailNode
	} else {
		newTailNode.prev = l.tail
		l.tail.next = newTailNode
	}

	l.tail = newTailNode

	l.count++
}

func (l *doublyLinkedList) removeFirst() {
	if l.count == 0 {
		panic("invalid operation `removeFirst` for empty singly linked list")
	}

	l.head = l.head.next

	if l.count == 1 {
		l.tail = nil
	} else {
		l.head.prev = nil
	}

	l.count--
}

func (l *doublyLinkedList) removeLast() {
	if l.count == 0 {
		panic("invalid operation `removeLast` for empty singly linked list")
	}

	l.tail = l.tail.prev

	if l.count == 1 {
		l.head = nil
	} else {
		l.tail.next = nil
	}

	l.count--
}

func main() {
	list := newDoublyLinkedList()

	list.addFirst(5)
	list.addFirst(8)
	list.addLast(13)
	list.removeLast()
	list.removeFirst()

	fmt.Println(list.head.prev)
	fmt.Println(list.tail.next)
	fmt.Println(list.tail.value)
	fmt.Println(list.count)
}
