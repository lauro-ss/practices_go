package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/eiannone/keyboard"
)

const (
	NUM_LINES   = 10
	NUM_COLUMNS = 400
)

type Cursor struct {
	Row    uint32
	Column uint32
}

type Page struct {
	Row          uint32
	Lines        []Line
	Cursor       Cursor
	Writer       io.Writer
	ReloadCursor bool
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

	page := Page{Lines: make([]Line, NUM_LINES), Writer: os.Stdout}
	page.Lines[page.Row].Value = make([]rune, NUM_COLUMNS)
	clear()
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		page.Lines[page.Cursor.Row].handlerRune(char, key, &page)
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

// func (p *Page) handlerRune(char rune, key keyboard.Key) {
// 	switch key {
// 	case keyboard.KeyBackspace:
// 		if p.Lines[p.Row].Column > 0 {
// 			p.Lines[p.Row].Column--
// 			p.Lines[p.Row].Value[p.Lines[p.Row].Column] = 0
// 			return
// 		}
// 	case keyboard.KeySpace:
// 		p.Lines[p.Row].Value[p.Lines[p.Row].Column] = ' '
// 	default:
// 		p.Lines[p.Row].Value[p.Lines[p.Row].Column] = char
// 	}
// 	if p.Lines[p.Row].Column < NUM_COLUMNS {
// 		p.Lines[p.Row].Column++
// 	}
// }

func (l Line) String() string {
	str := ""
	for _, char := range l.Value {
		str += string(char)
	}
	return str
}

func (p *Page) Show() {
	clear()
	for _, line := range p.Lines {
		os.Stdout.WriteString(fmt.Sprint(line))
	}
	if p.ReloadCursor {
		fmt.Fprintf(p.Writer, "\x1b[%dA", 1)
		p.ReloadCursor = false
	}
}

func handlerEnter(p *Page) {
	if p.Row < NUM_LINES {
		p.Row++
		p.Cursor.Row++
		p.Lines[p.Row].Value = make([]rune, NUM_COLUMNS)
	}
}

func (l *Line) handlerRune(char rune, key keyboard.Key, p *Page) {
	switch key {
	case keyboard.KeyBackspace:
		// replace the current rune with blank space
		if l.Column > 0 {
			l.Column--
			l.Value[l.Column] = 0
			return
		} else {
			p.Cursor.Row--
			// put the cursor one line up
			p.ReloadCursor = true
		}
	case keyboard.KeySpace:
		l.Value[l.Column] = ' '
	case keyboard.KeyEnter:
		l.Value[l.Column] = '\n'
		handlerEnter(p)
		return
	default:
		l.Value[l.Column] = char
	}
	if l.Column < NUM_COLUMNS {
		l.Column++
		p.Cursor.Column++
	}
}

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
