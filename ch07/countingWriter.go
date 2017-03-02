package main

import (
	"fmt"
	"io"
	"os"
)

type countingWriter struct {
	c             *int64
	wrappedWriter io.Writer
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := countingWriter{new(int64), w}
	return cw, cw.c
}

func (cw countingWriter) Write(b []byte) (int, error) {
	cInt, err := cw.wrappedWriter.Write(b)
	*cw.c += int64(cInt)
	return cInt, err
}

func main() {
	myWriter, runningCount := CountingWriter(os.Stdout)
	_, err := myWriter.Write([]byte("0123456789\n"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d\n", *runningCount)
	_, err = myWriter.Write([]byte("Holy crap what a lot of crap!\n"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d\n", *runningCount)
}
