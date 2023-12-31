package env

import (
	"bytes"
	"os"
)

func SourceEnv() (map[string]string, error) {
	b, err := os.ReadFile("./.env")
	if err != nil {
		return nil, err
	}
	i := 0
	end := len(b)
	envs := make(map[string]string, 5)
	for i < end {
		//35 == #
		if b[i] == 35 {
			skipLine(b, &i, &end)
		} else {
			getEnv(envs, readLine(b, &i, &end))
			//jumps only if the next caractere isn't #
			if i < end && b[i] != 35 {
				i++
			}
		}
	}
	return envs, nil
}

func readLine(b []byte, i *int, end *int) []byte {
	s := make([]byte, 20) //make
	c := 0
	//10 == \n
	for *i < *end && b[*i] != 10 {
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
	value := b[bytes.IndexByte(b, byte(61))+1:]
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
