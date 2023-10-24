package main

import (
	"fmt"
)

func main() {
	// arrays
	println("####### Arrays #######")
	list := [2]string{"1", "2"}
	var listInt = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	partial := [4]int{1: 10, 3: 40}
	listInt[0] = -1
	fmt.Println(partial)
	fmt.Printf("%q \n", list)
	fmt.Printf("%#v \n", list)
	fmt.Print(listInt[0:8], "\n")

	// slices
	println("####### Slices #######")
	var (
		slice          = []string{"1", "2", "3", "5"}
		sliceFromArray = list[:]
	)

	sliceMake := make([]string, 10, 100)

	sliceFromArray = append(sliceFromArray, "3")

	body := fmt.Sprintf("Slice Value: %v Slice Type: %T, Slice Capacity: %v Slice Lenght: %v", sliceFromArray, sliceFromArray, cap(sliceFromArray), len(sliceFromArray))
	fmt.Println(body)
	fmt.Println(slice)
	body = fmt.Sprintf("Slice Value: %v Slice Type: %T, Slice Capacity: %v Slice Lenght: %v", sliceMake, sliceMake, cap(sliceMake), len(sliceMake))
	fmt.Println(body)

	for i, v := range sliceFromArray {
		fmt.Printf("%v \t %v \n", i, v)
	}

	for _, v := range sliceFromArray {
		fmt.Print(v)
	}
}
