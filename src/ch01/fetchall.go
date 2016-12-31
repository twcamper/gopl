package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from the channel
	}
	fmt.Printf("%.2fs  elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	response, error := http.Get(url)
	if error != nil {
		ch <- fmt.Sprint(error) // send to the channel
		return
	}

	byteCount, error := io.Copy(ioutil.Discard, response.Body)
	response.Body.Close()
	if error != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, error) // send to the channel
		return
	}
	seconds := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", seconds, byteCount, url)
}
