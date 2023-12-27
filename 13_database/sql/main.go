package main

import (
	"bytes"
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
	envs := make(map[string]string, 1)
	for i < end {
		//35 == #
		if b[i] == 35 {
			skipLine(b, &i, &end)
		} else {
			getEnv(envs, readLine(b, &i, &end))
			//jumps only if the next caractere isn't #
			if b[i] != 35 {
				i++
			}
		}
	}
	return ""
}

func readLine(b []byte, i *int, end *int) []byte {
	s := make([]byte, 20) //make
	c := 0
	//10 == \n
	for b[*i] != 10 && *i < *end {
		// comment
		if b[*i] == 35 {
			skipLine(b, i, end)
			break
		} else {
			s[c] = b[*i]
			*i++
			c++
		}
	}
	return s
}

func getEnv(envs map[string]string, b []byte) {
	key := b[:bytes.IndexByte(b, byte(61))]
	value := b[bytes.IndexByte(b, byte(61)):]
	envs[string(key)] = string(value)
}

func skipLine(b []byte, i *int, end *int) *int {
	for b[*i] != 10 && *i < *end {
		*i++
	}
	//skipLines the \n
	if *i != *end {
		*i++
	}
	return i
}
