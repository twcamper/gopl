package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

type ByteCounter int
type RuneCounter int
type WordCounter int
type LineCounter int

func (c *ByteCounter) Write(input bytes.Buffer) {
	scanner := bufio.NewScanner(&input)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		*c++
	}
}
func (c *RuneCounter) Write(input bytes.Buffer) {
	scanner := bufio.NewScanner(&input)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		*c++
	}
}
func (c *WordCounter) Write(input bytes.Buffer) {
	scanner := bufio.NewScanner(&input)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*c++
	}
}

func (c *LineCounter) Write(input bytes.Buffer) {
	scanner := bufio.NewScanner(&input)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		*c++
	}
}

var (
	input bytes.Buffer
	bc    ByteCounter
	rc    RuneCounter
	wc    WordCounter
	lc    LineCounter
)

func read() {
	input = *new(bytes.Buffer)
	_, err := io.Copy(&input, os.Stdin)
	if err != nil {
		panic(err)
	}
}

func main() {
	read()
	bc.Write(input)
	wc.Write(input)
	rc.Write(input)
	lc.Write(input)
	fmt.Printf("\t%d\t%d\t%d\t%d\n", lc, wc, rc, bc)

}
