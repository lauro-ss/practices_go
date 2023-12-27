package main

import (
	"fmt"
	"os"
)

func main() {
	// db, err := sql.Open("postgres", "")
	GetEnv("t")
}

func GetEnv(key string) string {
	b, _ := os.ReadFile(".env")
	i := 0
	end := len(b) - 1
	for i < end {
		if b[i] == 35 {
			skip(b, &i, &end)
		} else {
			// fmt.Println(b[i], string(b[i]), i)
			fmt.Println(read(b, &i, &end))
			i++
		}

	}
	return ""
}

func read(b []byte, i *int, end *int) string {
	s := make([]byte, 100) //make
	c := 0
	for b[*i] != 10 && *i < *end {
		s[c] = b[*i]
		*i++
		c++
	}
	return string(s)
}

func skip(b []byte, i *int, end *int) *int {
	for b[*i] != 10 && *i < *end {
		// fmt.Println(b[i], string(b[i]), i)
		*i++
	}
	//skips the \n
	if *i != *end {
		*i++
	}
	return i
}
