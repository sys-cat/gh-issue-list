package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type (
	Issue struct {
		Title string `json:"title"`
		URL string `json:"url"`
		Create time.Time `json:"create"`
		Close time.Time `json:"close"`
	}
	Issues []Issue
)

func main() {
	filepath := flag.String("f", "", "file path string")
	flag.Parse()
	if *filepath == "" {
		fmt.Println("-f is required")
		os.Exit(0)
	}

	file, err := ioutil.ReadFile(*filepath)
	if err != nil {
		fmt.Println("file not exists.")
		os.Exit(0)
	}

	var issues Issues
	err = json.Unmarshal(file, &issues)
	if err != nil {
		fmt.Println("cant parse json")
		os.Exit(0)
	}

	for _, issue := range issues {
		fmt.Println(issue.Title)
	}
}