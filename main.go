package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Make and sort an slice.
	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)
	slice := makeRandomSlice(numItems, max)
	printSlice(slice, 40)

	for {
		// Get the target as a string.
		var targetString string
		fmt.Printf("Target: ")
		fmt.Scanln(&targetString)

		// If the target string is blank, break out of the loop.
		if len(targetString) == 0 {
			break
		}

		// Convert to int and find it.
		target, _ := strconv.Atoi(targetString)
		index, numTests := linearSearch(slice, target)
		if index < 0 || index >= len(slice) {
			fmt.Printf("Target %d not found, %d tests\n", target, numTests)
		} else {
			fmt.Printf("slice[%d] = %d, %d tests\n", index, slice[index], numTests)
		}
	}
}
