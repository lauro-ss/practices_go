package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/eiannone/keyboard"
)

const (
	NUM_LINES   = 10
	NUM_COLUMNS = 400
)

type Page struct {
	Row   uint32
	Lines []Line
}

type Line struct {
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
	page := Page{Lines: make([]Line, NUM_LINES)}
	page.Lines[page.Row].Value = make([]rune, NUM_COLUMNS)
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
		if p.Lines[p.Row].Column > 0 {
			p.Lines[p.Row].Column--
			p.Lines[p.Row].Value[p.Lines[p.Row].Column] = 0
			return
		}
	case keyboard.KeySpace:
		p.Lines[p.Row].Value[p.Lines[p.Row].Column] = ' '
	default:
		p.Lines[p.Row].Value[p.Lines[p.Row].Column] = char
	}
	if p.Lines[p.Row].Column < NUM_COLUMNS {
		p.Lines[p.Row].Column++
	}
}

func (p *Page) Show() {
	clear()
	for _, line := range p.Lines[p.Row].Value[:p.Lines[p.Row].Column] {
		fmt.Print(string(line))
	}
}

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
