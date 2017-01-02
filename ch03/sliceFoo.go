package main

import "fmt"

func main() {
	var x, y []int
	for i := 0; i < 17; i++ {
		y = append(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}

func appendInt(willGrow []int, value int) []int {
	var tempSlice []int
	newLength := len(willGrow) + 1
	if newLength <= cap(willGrow) {
		tempSlice = willGrow[:newLength]
	} else {
		newCapacity := newLength
		if newCapacity < 2*len(willGrow) {
			newCapacity = 2 * len(willGrow)
		}
		tempSlice = make([]int, newLength, newCapacity)
		copy(tempSlice, willGrow)
	}
	tempSlice[len(willGrow)] = value
	return tempSlice
}
