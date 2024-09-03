package main

import (
	"fmt"
	"iterators"
)

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	filter := func(v int) bool { return v%2 == 0 }

	for v := range iterators.FilterYield(slice, filter) {
		fmt.Println("Working on", v)
	}

	fmt.Println("FilterSlice")
	for _, v := range iterators.FilterSlice(slice, filter) {
		fmt.Println("Working on", v)
	}

	fmt.Println("FilterNext")
	iterator := iterators.Iterator{Slice: slice}
	for iterator.Next() {
		v := iterator.FilterIterator(filter)
		if v != nil {
			fmt.Println("Working on", *v)
		}
	}
}
