package main

import (
	"fmt"
)

func main() {
	for v := range Filter([]int{1, 2, 3, 4}, func(v int) bool { return v%2 != 0 }) {
		fmt.Println(v)
	}

	// for _, v := range  {

	// }
	// slices.All()
	// for _, v := range interate() {
	// 	fmt.Println(v)
	// }

}

func Filter[Slice ~[]E, E any](s Slice, filter func(E) bool) func(yield func(E) bool) {
	return func(yield func(E) bool) {
		for _, v := range s {
			if filter(v) && !yield(v) {
				return
			}
		}
	}
}

func interate() func(yield func(int, string) bool) {
	return func(yield func(int, string) bool) {
		for i := 0; i < 10; i++ {
			if !yield(i, string(i+48)) {
				return
			}
		}
	}
}
