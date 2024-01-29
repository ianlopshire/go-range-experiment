package main

import (
	"fmt"

	rangex "github.com/ianlopshire/go-range-experiment"
)

func main() {

	// Declare and initialize an array of integers.
	s := []int{1, 2, 3, 4, 5}
	fmt.Println("slice:", s)

	// Iterate over the array forward.
	fmt.Print("\t- forward: ")
	for _, c := range s {
		fmt.Print(c)
	}
	fmt.Println()

	// Iterate over the array backwards.
	fmt.Print("\t- forward: ")
	for _, c := range rangex.Backward(s) {
		fmt.Print(c)
	}
	fmt.Println()

	// Iterate over the array in chuncks.
	fmt.Print("\t- forward: ")
	for _, c := range rangex.Chunck(s, 2) {
		fmt.Print(c)
	}
	fmt.Println()

	// Iterate over the array with a filter (>2).
	fmt.Print("\t- filter(>2): ")
	filter := func(e int) bool { return e > 2 }
	for i, c := range rangex.Filter(s, filter) {
		fmt.Print(i, c, " ")
	}
	fmt.Println()

}
