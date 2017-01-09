package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var stubFile string

func main() {
	stubFile = os.Args[1]
	http.HandleFunc("/search/issues", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile(stubFile)
	if err != nil {
		log.Fatalf("%s: %v\n", os.Args[0], err)
	}

	fmt.Fprintf(w, "%s\n", data)
}
