package main

import "fmt"

type Stack struct {
	Data any
	Prev *Stack
	Next *Stack
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(d any) {
	//Fill the first elemenet
	if s.Data == nil {
		s.Data = d
		return
	}
	if s.Next != nil {
		s.Next.Push(d)
		return
	}
	s.Next = &Stack{}
	s.Next.Data = d
	s.Next.Next = nil
	s.Next.Prev = s
}

func (s *Stack) Pop() *any {
	if s.Next != nil {
		return s.Next.Pop()
	}
	//Uncomment for common Stack
	s.Prev.Next = nil
	s.Prev = nil
	return &s.Data
}

func (s *Stack) String() string {
	return fmt.Sprintf("%v", s.Data)
}

func (s *Stack) List() string {
	if s.Next != nil {
		return fmt.Sprintf("Prev: %v, Actual: %v, Next: %v \n", s.Prev, s.Data, s.Next) + s.Next.List()
	}
	return fmt.Sprintf("Prev: %v, Actual: %v, Next: %v", s.Prev, s.Data, s.Next)
}

// func main() {
// 	s := NewStack()
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
