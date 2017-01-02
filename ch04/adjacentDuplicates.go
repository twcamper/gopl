// Exercise 4.5: Write an in-place function to eliminate adjacent duplicates in a []string slice.
package main

import (
	"fmt"
	"os"
)

func main() {
	input := make([]string, len(os.Args[1:]))
	copy(input, os.Args[1:])
	fmt.Println(input)
	elimnateAdjacentDuplicatesInPlace(input)
	fmt.Println("In Place: %s", input)
	fmt.Println("Via Copy: %s", elimnateAdjacentDuplicates(os.Args[1:]))
}

// does not shrink original slice, but replaces trailing elements with empty strings
func elimnateAdjacentDuplicatesInPlace(strings []string) {
	i := eliminate(strings)
	blank := make([]string, len(strings)-i)
	copy(strings[i:], blank)
}

func elimnateAdjacentDuplicates(strings []string) []string {
	i := eliminate(strings)
	return strings[:i]
}

func eliminate(strings []string) int {
	newIndex := 1
	for previous, s := range strings[1:] {
		if s != strings[previous] {
			strings[newIndex] = s
			newIndex++
		}
	}
	return newIndex
}
