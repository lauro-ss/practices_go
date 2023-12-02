package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func binarySearch(s []int32, v int32) *int32 {
	//Gets the middle size of the slice
	mid := len(s) / 2

	if v == s[mid] {
		return &s[mid]
	}

	// Cheks if the value is less or bigger then them middle value
	// Case the value is bigger, it's only passed mid > data from the array
	// Ex: Look for 5, and the slice is [1,2,3,4,5], then, it's passed [4,5]
	if v > s[mid] && mid != 1 {
		return binarySearch(s[mid:], v)
	} else if mid != 1 {
		return binarySearch(s[:mid], v)
	}

	return nil
}

//func

const SIZE = math.MaxInt16

func main() {
	list := make([]int32, SIZE)
	for i := range list {
		list[i] = int32(i + 1)
	}
	t1 := time.Now()
	for i := 0; i < 1; i++ {
		fmt.Println(*binarySearch(list, rand.Int31n(SIZE)))
	}
	t2 := time.Now()
	fmt.Println("Takes:", t2.Sub(t1))
}
