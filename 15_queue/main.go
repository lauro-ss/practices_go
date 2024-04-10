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
		return
	}

	tail := getTail(q.head, n.value)
	if tail == nil {
		return
	}
	tail.next = n
	n.prev = tail
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
	q := createQueue()

	q.add(2)
	q.add(1)
	q.add(1)
	q.add(3)
	q.add(3)
	q.add(3)
	q.add(4)

	fmt.Println(q)
}
