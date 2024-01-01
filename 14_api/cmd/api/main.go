package main

import (
	"fmt"

	"github.com/lauro-ss/practices_go/14_api/pkg/env"
)

func main() {
	fmt.Println(env.SourceEnv())
}
