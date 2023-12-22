package main

import "fmt"

type I interface {
	Call(I, string)
}

type Obs string

func (o *Obs) Call(i I, s string) {
	fmt.Println(s)
	i.Call(i, "Next")
}

func Call1(s string) {
	fmt.Println(s)
}

func Call2(s string) {
	fmt.Println(s)
}
