package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/twcamper/gopl/ch04/github"
)

func main() {
	http.HandleFunc("/search/issues", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(responseData())
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Fprintf(w, "%s\n", data)
}

func responseData() *github.IssuesSearchResult {
	eaigner := github.User{Login: "eaigner"}
	gopherbot := github.User{Login: "gopherbot"}
	item5680 := github.Issue{Number: 5680, User: &eaigner, Title: "encoding/json: set key converter on en/decoder"}
	item6050 := github.Issue{Number: 6050, User: &gopherbot, Title: "encoding/json: provide tokenizer"}
	item8658 := github.Issue{Number: 8658, User: &gopherbot, Title: "encoding/json: use bufio"}
	items := []*github.Issue{&item5680, &item6050, &item8658}
	return &github.IssuesSearchResult{len(items), items}
}
