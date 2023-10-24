package main

import "fmt"

type Money struct {
	large uint64
	small uint8
}

type Person struct {
	name     string
	lastName string
	money    Money
}

func (m *Money) Add(large uint64, small uint8) {
	s := m.small + small

	if s > 100 {
		m.large = uint64(s / 100)
		m.small = uint8(s % 100)
	} else {
		m.small = small
	}

	m.large = large
}

func (m Money) String() string {
	return fmt.Sprintf("%v,%v", m.large, m.small)
}

func (p Person) String() string {
	return fmt.Sprintf("%v %v - %v", p.name, p.lastName, p.money)
}

func main() {
	p := Person{name: "Lauro", lastName: "Santana", money: Money{10, 20}}
	fmt.Print(p)
}
