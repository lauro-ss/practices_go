package main

import (
	"fmt"
)

type queue struct {
	head *node
	size int
}

type node struct {
	value int
	next  *node
	prev  *node
}

func createQueue() *queue {
	return &queue{}
}

func (q *queue) add(i int) {
	n := &node{value: i}

	if q.head == nil {
		q.head = n
		q.size++
		return
	}

	tail := getTail(q.head, n.value)
	if tail == nil {
		return
	}
	tail.next = n
	n.prev = tail
	q.size++
}

func (q *queue) get() int {
	if q.head == nil {
		return 0
	}

	n := q.head
	q.head = q.head.next
	q.size--
	return n.value
}

func getTail(n *node, v int) *node {
	if n.value == v {
		return nil
	}
	if n.next != nil {
		return getTail(n.next, v)
	}
	return n
}

func (q *queue) String() string {
	return stringNode(q.head)
}

func stringNode(n *node) string {
	if n != nil {
		return fmt.Sprint(n) + stringNode(n.next)
	}

	return ""
}

func main() {
	fmt.Println(testQueue())
	q := createQueue()

	for v := range 10 {
		q.add(v)
		q.add(v)
	}

	fmt.Println(q.get())
	fmt.Println(q)
	fmt.Println(q.get())
	fmt.Println(q)
}

func testQueue() bool {
	q := createQueue()

	for v := range 10 {
		q.add(v)
		q.add(v)
	}

	if q.size != 10 {
		return false
	}

	q.get()
	if q.size != 9 {
		return false
	}

	return testPoint(q.head)
}

func testPoint(n *node) bool {
	if n.next != nil {
		if fmt.Sprintf("%p", n) == fmt.Sprintf("%p", n.next.prev) {
			return testPoint(n.next)
		} else {
			return false
		}
	}

	return true
}
