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
	Rows  uint32
	Lines []Line
	Cursor
	Writer       io.Writer
	ReloadCursor bool
}

type Line struct {
	Columns uint32
	Value   []rune
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
	clear(&page)
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		page.Lines[page.Cursor.Row].handlerRune(char, key, &page)
		page.Show()
	}
}

func (l Line) String() string {
	str := ""
	for _, char := range l.Value {
		str += string(char)
	}
	return str
}

func (p *Page) Show() {
	clear(p)
	fmt.Printf("Row: %v Column: %v \n", p.Cursor.Row, p.Cursor.Column)
	for _, line := range p.Lines {
		p.Writer.Write([]byte(fmt.Sprint(line)))
	}
	if p.ReloadCursor {
		fmt.Fprintf(p.Writer, "\x1b[%dA", 1)
		fmt.Fprintf(p.Writer, "\x1b[%dC", p.Cursor.Column)
		p.ReloadCursor = false
	}
}

func handlerEnter(p *Page) {
	if p.Row < NUM_LINES {
		p.Rows++
		p.Cursor.Row++
		p.Cursor.Column = 0
		p.Lines[p.Row].Value = make([]rune, NUM_COLUMNS)
	}
}

func (l *Line) handlerRune(char rune, key keyboard.Key, p *Page) {
	switch key {
	case keyboard.KeyBackspace:
		// replace the current rune with blank space
		if p.Cursor.Column > 0 {
			p.Cursor.Column--
			l.Value[p.Cursor.Column] = 0
			return
		} else {
			if p.Cursor.Row > 0 {
				p.Cursor.Row--
				p.Cursor.Column = p.Lines[p.Cursor.Row].Columns - 1
				// put the cursor one line up
				p.ReloadCursor = true
			}
		}
	case keyboard.KeySpace:
		l.Value[p.Cursor.Column] = ' '
	case keyboard.KeyEnter:
		l.Value[p.Cursor.Column] = '\n'
		handlerEnter(p)
		return
	default:
		l.Value[p.Cursor.Column] = char
	}
	if l.Columns < NUM_COLUMNS {
		l.Columns++
		p.Cursor.Column++
	}
}

func clear(p *Page) {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = p.Writer
	cmd.Run()
}
