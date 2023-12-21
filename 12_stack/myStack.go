package main

import "fmt"

type MyStack struct {
	Data     any
	Prev     *MyStack
	Next     *MyStack
	consumed bool
}

func NewMyStack() *MyStack {
	return &MyStack{}
}

func (s *MyStack) Push(d any) {
	//Fill the first elemenet
	if s.Data == nil {
		s.Data = d
		return
	}
	if s.Next != nil {
		s.Next.Push(d)
		return
	}
	s.Next = &MyStack{}
	s.Next.Data = d
	s.Next.Next = nil
	s.Next.Prev = s
}

func (s *MyStack) Pop() *any {
	if s.Next != nil && s.Next.consumed != true {
		return s.Next.Pop()
	}
	//Uncomment for common MyStack
	// s.Prev.Next = nil
	// s.Prev = nil

	//Consume the elements and not remove from the MyStack
	s.consumed = true

	//Refil the next element
	//After consume the last MyStack element,
	//all orthers will be ready to consume
	if s.Next != nil {
		s.Next.consumed = false
	}
	return &s.Data
}

func (s *MyStack) String() string {
	return fmt.Sprintf("%v %v", s.Data, s.consumed)
}

func (s *MyStack) List() string {
	if s.Next != nil && s.Next.consumed != true {
		return fmt.Sprintf("Prev: %v, Actual: %v, Next: %v \n", s.Prev, s.Data, s.Next) + s.Next.List()
	}
	return fmt.Sprintf("Prev: %v, Actual: %v, Next: %v", s.Prev, s.Data, s.Next)
}

// func main() {
// 	s := NewMyStack()
// 	s.Push("1")
// 	s.Push("2")
// 	s.Push("3")
// 	fmt.Println(s.List())
// 	fmt.Println("Consumed:", *s.Pop())
// 	fmt.Println(s.List())
// 	fmt.Println("Consumed:", *s.Pop())
// 	fmt.Println(s.List())
// 	fmt.Println("Consumed:", *s.Pop())
// 	fmt.Println(s.List())
// }
