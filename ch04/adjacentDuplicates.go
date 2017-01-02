// Exercise 4.5: Write an in-place function to eliminate adjacent duplicates in a []string slice.
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[1:])
	fmt.Println(eliminate(os.Args[1:]))
	fmt.Println(os.Args[1:])
}

func eliminate(strings []string) []string {
	i := 1
	for previous, s := range strings[1:] {
		if s != strings[previous] {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}
