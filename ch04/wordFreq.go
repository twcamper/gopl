package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var wordCounts map[string]int

func init() {
	wordCounts = make(map[string]int)
}

func main() {
	count()
	report()
}

func count() {
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := strings.Trim(input.Text(), ",.'\";:!?")
		wordCounts[word]++
	}
	if err := input.Err(); err != nil {

		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
	}
}

func report() {
	for word, count := range wordCounts {
		fmt.Printf("%s: %d\n", word, count)
	}
}
