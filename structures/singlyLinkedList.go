package main

import "fmt"

type singlyLinkedList struct {
	singlyLinkedListable
	head  *node
	tail  *node
	count int
}

type node struct {
	value interface{}
	next  *node
}

func newSinglyLinkedList() singlyLinkedList {
	return singlyLinkedList{}
}

type singlyLinkedListable interface {
	addFirst(value interface{})
	addLast(value interface{})
	removeFirst()
	removeLast()
}

func (l *singlyLinkedList) addFirst(value interface{}) {
	newHeadNode := node{value: value}

	tmpNode := l.head

	l.head = &newHeadNode

	l.head.next = tmpNode

	l.count++

	if l.count == 1 {
		l.tail = l.head
	}
}

func (l *singlyLinkedList) addLast(value interface{}) {
	newTailNode := node{value: value}

	if l.count == 0 {
		l.head = &newTailNode
	} else {
		l.tail.next = &newTailNode
	}

	l.tail = &newTailNode

	l.count++
}

func (l *singlyLinkedList) removeFirst() {
	if l.count == 0 {
		panic("invalid operation `removeFirst` for empty singly linked list")
	}

	l.head = l.head.next

	if l.count == 1 {
		l.tail = nil
	}

	l.count--
}

func (l *singlyLinkedList) removeLast() {
	if l.count == 0 {
		panic("invalid operation `removeLast` for empty singly linked list")
	}

	if l.count == 1 {
		l.head = nil
		l.tail = nil
	} else {
		penultimateNode := l.head
		for penultimateNode.next != l.tail {
			penultimateNode = penultimateNode.next
		}
		l.tail = penultimateNode
		l.tail.next = nil
	}

	l.count--
}

func main() {
	list := newSinglyLinkedList()

	list.addFirst(5)
	list.addFirst(8)
	list.addLast(13)
	list.removeLast()
	list.removeFirst()

	fmt.Println(list.head.value)
	fmt.Println(list.tail.value)
	fmt.Println(list.count)
}
