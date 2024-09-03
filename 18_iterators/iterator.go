package iterators

import (
	"fmt"
)

type Iterator struct {
	Slice []int
	index int
}

func (i *Iterator) FilterIterator(filter func(int) bool) *int {
	if filter(i.Slice[i.index]) {
		fmt.Println("Filter for", i.Slice[i.index])
		return &i.Slice[i.index]
	}
	return nil
}

func (i *Iterator) Next() bool {
	i.index++
	return i.index < len(i.Slice)
}

func FilterYield[Slice ~[]E, E any](s Slice, filter func(E) bool) func(yield func(E) bool) {
	return func(yield func(E) bool) {
		for _, v := range s {
			if filter(v) && func() bool { fmt.Println("Filter for", v); return true }() && !yield(v) {
				return
			}
		}
	}
}

func FilterSlice(s []int, filter func(int) bool) []int {
	sl := []int{}
	for _, v := range s {
		if filter(v) && func() bool { fmt.Println("Filter for", v); return true }() {
			sl = append(sl, v)
		}
	}
	return sl
}
