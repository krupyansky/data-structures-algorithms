package main

import "fmt"

type Table struct {
	first *Node
	count int32
}

func newTable() *Table {
	return &Table{}
}

func (t *Table) tryGet(key interface{}) interface{} {
	for x := t.first; x != nil; x = x.next {
		if x.key == key {
			return x.value
		}
	}

	return nil
}

func (t *Table) contains(key interface{}) bool {
	for x := t.first; x != nil; x = x.next {
		if x.key == key {
			return true
		}
	}

	return false
}

func (t *Table) keys() []interface{} {
	var keys []interface{}

	for x := t.first; x != nil; x = x.next {
		keys = append(keys, x.key)
	}

	return keys
}

func (t *Table) add(key, value interface{}) {
	if key == nil {
		panic("key is nil")
	}

	t.count++

	for x := t.first; x != nil; x = x.next {
		if x.key == key {
			x.value = value
			return
		}
	}

	t.first = newNode(key, value, t.first)
}

func (t *Table) delete(key interface{}) bool {
	if key == nil {
		panic("key is nil")
	}

	if !t.contains(key) {
		return false
	}

	prev := t.first

	for x := t.first; x != nil; x = x.next {
		if x.key == key {
			if t.first == x {
				t.first = x.next
			} else {
				prev.next = x.next
			}
		}

		prev = x
	}

	t.count--
	return true
}

func main() {
	table := newTable()

	nodeAsKey := newNode(654654, 5435, nil)
	table.first = newNode(nodeAsKey, 1123, nil)

	val := table.tryGet(nodeAsKey)

	fmt.Println(val)

	table.add("some-key", "some-value")

	fmt.Println(table.tryGet("some-key"))

	fmt.Println(table.contains("some-key"))
	fmt.Println(table.contains("fake-key"))

	fmt.Println(table.keys())

	table.add("some-key-new", "some-value-new")

	fmt.Println(table.tryGet("some-key-new"))

	fmt.Println(table.delete("uga-buga"))

	fmt.Println(table.count)

	fmt.Println(table.contains("some-key-new"))
	fmt.Println(table.delete("some-key-new"))
	fmt.Println(table.contains("some-key-new"))

	fmt.Println(table.count)

	fmt.Println(table.contains(nodeAsKey))
	fmt.Println(table.delete(nodeAsKey))
	fmt.Println(table.contains(nodeAsKey))
}
