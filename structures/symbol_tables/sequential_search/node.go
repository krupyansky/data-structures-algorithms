package main

type Node struct {
	key   interface{}
	value interface{}
	next  *Node
}

func newNode(key, value interface{}, next *Node) *Node {
	return &Node{
		key:   key,
		value: value,
		next:  next,
	}
}
