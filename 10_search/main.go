package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func binarySearch[V int64 | int32](s []V, v V) *V {
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

func linearSearch[V int64 | int32](s []V, v V) *V {
	for i := range s {
		if s[i] == v {
			return &s[i]
		}
	}
	return nil
}

func linearSearchGo[V int64 | int32](s []V, v V) *V {
	mid := len(s) / 2
	c := make(chan *V, 2)
	go linearSearchRoutine(s[:mid], &v, c)
	go linearSearchRoutine(s[mid:], &v, c)
	for value := range c {
		if value != nil {
			return value
		}
	}
	return nil
}

func linearSearchRoutine[V int64 | int32](s []V, v *V, c chan<- *V) {
	for i := range s {
		if s[i] == *v {
			c <- &s[i]
		}
	}
	c <- nil
}

const SIZE = math.MaxInt16
const TRIES = 1000

func main() {
	list := make([]int32, SIZE)
	for i := range list {
		list[i] = int32(i + 1)
	}
	value := rand.Int31n(SIZE)
	var find any
	t1 := time.Now()
	for i := 0; i < TRIES; i++ {
		find = binarySearch(list, value)
	}
	t2 := time.Now()
	fmt.Println("Takes:", t2.Sub(t1), "Find", find)

	t1 = time.Now()
	for i := 0; i < TRIES; i++ {
		find = linearSearch(list, value)
	}
	t2 = time.Now()
	fmt.Println("Takes:", t2.Sub(t1), "Find", find)

	t1 = time.Now()
	for i := 0; i < TRIES; i++ {
		find = linearSearchGo(list, value)
	}
	t2 = time.Now()
	fmt.Println("Takes:", t2.Sub(t1), "Find", find)
}
