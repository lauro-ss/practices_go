package main

import "fmt"

type Stack struct {
	Func I
	Prev *Stack
	Next *Stack
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(d func(string)) {
	//Fill the first elemenet
	if s.Func == nil {
		s.Func = d
		return
	}
	if s.Next != nil {
		s.Next.Push(d)
		return
	}
	s.Next = &Stack{}
	s.Next.Func = d
	s.Next.Next = nil
	s.Next.Prev = s
}

func (s *Stack) Pop() func(string) {
	if s.Next != nil {
		return s.Next.Pop()
	}
	//Uncomment for common Stack
	s.Prev.Next = nil
	s.Prev = nil
	return s.Func
}

// func (s *Stack) String() string {
// 	return fmt.Sprintf("%v", s.Func)
// }

// func (s *Stack) List() string {
// 	if s.Next != nil {
// 		return fmt.Sprintf("Prev: %v, Actual: %v, Next: %v \n", s.Prev, s.Func, s.Next) + s.Next.List()
// 	}
// 	return fmt.Sprintf("Prev: %v, Actual: %v, Next: %v", s.Prev, s.Func, s.Next)
// }

func (s *Stack) String() string {
	return fmt.Sprintf("%v", s.Func)
}

func (s *Stack) List() string {
	if s.Next != nil {
		return fmt.Sprintf("Prev: %v, Actual: %v, Next: %v \n", s.Prev, s.Func, s.Next) + s.Next.List()
	}
	return fmt.Sprintf("Prev: %v, Actual: %v, Next: %v", s.Prev, s.Func, s.Next)
}

func main() {
	s := NewStack()
	s.Push(Call1)
	s.Push(Call2)
	s.Push(func(s string) { fmt.Println(s) })
	// fmt.Println(s.List())
	// fmt.Println("Consumed:", *s.Pop())
	// fmt.Println(s.List())
	// fmt.Println("Consumed:", *s.Pop())
	// fmt.Println(s.List())
	// fmt.Println("Consumed:", *s.Pop())
	// fmt.Println(s.List())
}
