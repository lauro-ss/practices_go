package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
		m.large += uint64(s / 100)
		m.small = uint8(s % 100)
	} else {
		m.small = s
	}
	m.large += large
}

func (m *Money) Remove(large uint64, small uint8) {
	sf := (large * 100) + uint64(small)
	sm := (m.large * 100) + uint64(m.small)

	if sm < sf {
		m.large = 0
		m.small = 0
	} else {
		sm -= sf
		m.large = sm / 100
		m.small = uint8(sm % 100)
	}
}

func (m Money) String() string {
	return fmt.Sprintf("%v,%v", m.large, m.small)
}

func (p Person) String() string {
	return fmt.Sprintf("%v %v - %v", p.name, p.lastName, p.money)
}

func main() {
	var values [4]string
	reader := bufio.NewReader(os.Stdin)

	i := 0
	for {
		switch i {
		case 0:
			fmt.Print("Digite seu nome: ")
			break
		case 1:
			fmt.Print("Digite seu sobrenome: ")
			break
		case 2:
			fmt.Print("Digite seu dinheiro ex(2,80): ")
			break
		}
		v, _ := reader.ReadString('\n')
		v = string(v[0 : len(v)-2])

		fmt.Printf("Você digitou %q - Deseja refazer operação? (s/n)", v)
		values[i] = v
		v, _ = reader.ReadString('\n')
		v = string(v[0 : len(v)-2])
		switch v {
		case "s":
			values[i] = ""
			break
		case "n":
			i++
			break
		}
		if i > 2 {
			break
		}
	}
	m := strings.Split(values[2], ",")
	ml, _ := strconv.ParseUint(m[0], 10, 64)
	ms, _ := strconv.ParseUint(m[0], 10, 8)
	p := Person{name: values[0], lastName: values[1], money: Money{large: ml, small: uint8(ms)}}
	fmt.Print(p)
}
