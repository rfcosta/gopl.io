// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 108.

// Movie prints Movies as JSON.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//!+
type VizItemImpl struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
type VizTmMetric struct {
	ValueType  string      `json:"valueType"`
	MetricType string      `json:"metricType"`
	Bucket     int32       `json:"bucket"`
	Tags       []string    `json:"tags"`
	ItemImpl   VizItemImpl `json:"itemImpl"`
}

type VizChart struct {
	Title     string        `json:"title"`
	TmMetrics []VizTmMetric `json:"tmMetrics"`
	X         int           `json:"x"`
	Y         int           `json:"y"`
	ViewType  string        `json:"viewType"`
}
type VizView struct {
	Charts          []VizChart `json:"charts"`
	EditableOcmTags []string   `json:"editableOcmTags"`
	DefaultTimeSpan int32      `json:"defaultTimeSpan"`
	Granularity     int        `json:"granularity"`
}

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

type Title struct {
	Title string
}

type Entry struct {
	Title string
	Year  int  `json:"year of release"`
	Color bool `json:"color,omitempty"`
}

var movies = []Movie{
	{Title: "Casablanca",
		Year:   1942,
		Color:  false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke",
		Year:  1967,
		Color: true,
		Actors: []string{
			"Paul Newman"}},
	{Title: "Bullitt",
		Year:  1968,
		Color: true,
		Actors: []string{
			"Steve McQueen",
			"Jacqueline Bisset"}},
	// ...
}

//!-

func main() {
	{
		//!+Marshal
		data, err := json.Marshal(movies)
		if err != nil {
			log.Fatalf("JSON marshaling failed: %s", err)
		}
		fmt.Printf("%s\n", data)
		//!-Marshal
	}

	{
		//!+MarshalIndent
		data, err := json.MarshalIndent(movies, "", "    ")
		if err != nil {
			log.Fatalf("JSON marshaling failed: %s", err)
		}
		fmt.Printf("%s\n", data)
		//!-MarshalIndent

		//!+Unmarshal

		//var titles []struct{ Title string }
		var titles []Title

		if err := json.Unmarshal(data, &titles); err != nil {
			log.Fatalf("JSON unmarshaling failed: %s", err)
		}
		fmt.Println(titles) // "[{Casablanca} {Cool Hand Luke} {Bullitt}]"
		//!-Unmarshal

		var movies_entries []Entry
		if err := json.Unmarshal(data, &movies_entries); err != nil {
			log.Fatalf("JSON unmarshaling failed: %s", err)
		}
		fmt.Println(movies_entries) // "[{Casablanca} {Cool Hand Luke} {Bullitt}]"

		dat, er := json.MarshalIndent(movies_entries, "", "    ")
		if er != nil {
			log.Fatalf("JSON marshaling failed: %s", er)
		}
		fmt.Printf("%s\n", dat)

		for sq, movEntry := range movies_entries {
			ent, err := json.MarshalIndent(movEntry, "", "    ")
			if err != nil {
				log.Fatalf("JSON marshaling failed: %s", err)
			}
			fmt.Printf("Seq: %d Entry: %s\n", sq, ent)

		}

	}

	{
		//!+File
		for _, filename := range os.Args[1:] {
			fmt.Printf("Processing file '%s'\n", filename)
			data, err := ioutil.ReadFile(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
				continue
			}

			vizView := VizView{}

			if err := json.Unmarshal([]byte(data), &vizView); err != nil {
				log.Fatalf("JSON unmarshaling failed: %s", err)
			}
			fmt.Println(vizView)
		}
		//!-File

	}

}

/*
//!+output
[{"Title":"Casablanca","released":1942,"Actors":["Humphrey Bogart","Ingr
id Bergman"]},{"Title":"Cool Hand Luke","released":1967,"color":true,"Ac
tors":["Paul Newman"]},{"Title":"Bullitt","released":1968,"color":true,"
Actors":["Steve McQueen","Jacqueline Bisset"]}]
//!-output
*/

/*
//!+indented
[
    {
        "Title": "Casablanca",
        "released": 1942,
        "Actors": [
            "Humphrey Bogart",
            "Ingrid Bergman"
        ]
    },
    {
        "Title": "Cool Hand Luke",
        "released": 1967,
        "color": true,
        "Actors": [
            "Paul Newman"
        ]
    },
    {
        "Title": "Bullitt",
        "released": 1968,
        "color": true,
        "Actors": [
            "Steve McQueen",
            "Jacqueline Bisset"
        ]
    }
]
//!-indented
*/
