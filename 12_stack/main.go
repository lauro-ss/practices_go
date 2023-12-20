package main

import "fmt"

type Stack struct {
	Data any
	Prev *Stack
	Next *Stack
}

func NewStack(d any) *Stack {
	return &Stack{Data: d}
}

func (s *Stack) Push(d any) {
	if s.Next != nil {
		s.Next.Push(d)
		return
	}
	s.Next = &Stack{}
	s.Next.Data = d
	s.Next.Next = nil
	s.Next.Prev = s
}

func (s *Stack) String() {
	//fmt.Println("#", s.Prev, s.Data, s.Next)
	if s.Next != nil {
		fmt.Println("#", s.Prev, s.Data, s.Next)
		s.Next.String()
		return
	}
	fmt.Println("!", s.Prev, s.Data, s.Next)
}

func main() {
	s := NewStack("0")
	//s.String()
	s.Push("1")
	s.String()
	// s.Push("2")
	// s.String()
}
