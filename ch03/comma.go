package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Println(comma(arg))
	}
}

func comma(s string) string {
	var buffer bytes.Buffer
	var to int = 3
	if (r := len(s) % 3) > 0 {
		to = r
	}
	for from := 0; to <= len(s); from, to = to, to+3 {
		buffer.WriteString(s[from:to] + ",")
	}
	withCommas := buffer.String()
	return withCommas[:len(withCommas)-1]
}
