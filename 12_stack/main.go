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

func (s *Stack) Print() string {
	if s.Next != nil {
		return fmt.Sprintf("Prev: %v, Actual: %v, Next: %v \n", s.Prev, s.Data, s.Next.Data) + s.Next.Print()
	}
	return fmt.Sprintf("Prev: %v, Actual: %v, Next: %v", s.Prev.Data, s.Data, s.Next)
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
	s := NewStack()
	//s.String()
	s.Push("1")
	s.Push("2")
	//s.String()
	fmt.Println(s.Print())
}
