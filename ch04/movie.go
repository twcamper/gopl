package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

type Title struct{ Title string }

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman", "Claude Raines"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman", "George Kennedy"}},
	{Title: "Bullitt", Year: 1968, Color: true, Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

func main() {
	data := marshal()
	fmt.Printf("%s\n", data)
	for _, t := range unmarshal(data) {
		fmt.Printf("\"%s\"\n", t.Title)
	}
}

func marshal() []byte {
	data, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	return data
}

func unmarshal(j []byte) []Title {
	var titles []Title
	if err := json.Unmarshal(j, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	return titles
}
