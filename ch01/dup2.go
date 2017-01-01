package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var counts map[string]int
var files map[string][]string

func init() {
	counts = make(map[string]int)
	files = make(map[string][]string)
}

func main() {
	inputFiles := os.Args[1:]

	for _, arg := range inputFiles {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		analyzeFile(f)
		f.Close()
	}
	for key, val := range counts {
		if val > 1 {
			fmt.Printf("%d\t'%s'\t%v\n", val, key, strings.Join(files[key], ", "))
		}
	}
}

func analyzeFile(f *os.File) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		analyzeLine(input.Text(), f.Name())
	}
	// NOTE: ignoring potential errors from input.Error
}

func analyzeLine(line, fileName string) {
	counts[line]++
	if !contains(files[line], fileName) {
		files[line] = append(files[line], fileName)
	}
}

func contains(files []string, path string) bool {
	for _, item := range files {
		if item == path {
			return true
		}
	}
	return false
}
