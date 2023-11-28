package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/eiannone/keyboard"
)

const SIZE_PAGE = 400

type Page struct {
	Row    uint32
	Column uint32
	Value  []rune
}

func main() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()
	page := Page{Value: make([]rune, SIZE_PAGE)}
	clear()
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		page.handlerRune(char, key)
		page.Show()
		//fmt.Printf("You pressed: rune %q, key %X\r\n", char, key)

		// if key == keyboard.KeyEnter {
		// 	lines[i] += "\n"
		// 	fmt.Print(lines[i])
		// 	i++
		// } else {
		// 	lines[i] += string(char)
		// 	fmt.Print(lines[i])
		// }

		// if key == keyboard.KeyArrowUp {
		// 	if i > 0 {
		// 		i--
		// 	}
		// 	fmt.Print(lines[i])
		// }

		// if key == keyboard.KeyArrowDown {
		// 	if i < len(lines) {
		// 		i++
		// 	}
		// 	fmt.Print(i)
		// }

		// if key == keyboard.KeyEsc {
		// 	fmt.Print(lines)
		// 	break
		// }
	}
}

func (p *Page) handlerRune(char rune, key keyboard.Key) {
	switch key {
	case keyboard.KeyBackspace:
		if p.Column > 0 {
			p.Column--
			p.Value[p.Column] = 0
			return
		}
	case keyboard.KeySpace:
		p.Value[p.Column] = ' '
	default:
		p.Value[p.Column] = char
	}
	if p.Column < SIZE_PAGE {
		p.Column++
	}
}

func (p *Page) Show() {
	clear()
	for _, line := range p.Value[:p.Column] {
		fmt.Print(string(line))
	}
}

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
