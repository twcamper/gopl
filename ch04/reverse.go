// Exercise 4.3: Rewrite reverse to use an array pointer instead of a slice.

package main

import (
	"fmt"
)

func main() {
	ints := [10]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Println(ints)
	reverse(&ints)
	fmt.Println(ints)
}

func reverse(s *[10]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
