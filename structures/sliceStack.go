package main

import "fmt"

type sliceStack struct {
	Stackable
	items []interface{}
}

func newSliceStack(cap int) sliceStack {
	return sliceStack{items: make([]interface{}, 0, cap)}
}

type Stackable interface {
	peek() interface{}
	pop()
	push(value interface{})
}

func (s *sliceStack) peek() interface{} {
	if len(s.items) == 0 {
		panic("invalid operation `peek` for empty stack")
	}
	return s.items[len(s.items)-1]
}

func (s *sliceStack) pop() {
	if len(s.items) == 0 {
		panic("invalid operation `pop` for empty stack")
	}
	s.items = s.items[:len(s.items)-1]
}

func (s *sliceStack) push(value interface{}) {
	s.items = append(s.items, value)
}

func main() {
	stack := newSliceStack(8)

	stack.push(6)
	stack.push(10)

	fmt.Println(stack.peek())

	stack.pop()

	fmt.Println(stack.peek())
}
