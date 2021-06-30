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
		Title	 string 	`json:"title"`
		URL		 string 	`json:"url"`
		Create	 time.Time  `json:"create"`
		Close 	 time.Time  `json:"close"`
	}
	Issues []Issue
)

func main() {
	filepath := flag.String("f", "", "file path string")
	due := flag.Bool("d", false, "detail duration per items")
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

	list := map[int]map[int]Issues{}
	for _, issue := range issues {
		list[issue.Close.Year()] = map[int]Issues{}
	}
	for _, issue := range issues {
		list[issue.Close.Year()][int(issue.Close.Month())] = append(list[issue.Close.Year()][int(issue.Close.Month())], issue)
	}

	if !*due {
		for key, months := range list {
			for month, issues := range months {
				fmt.Printf("%d,%d,%d\n", key, month, len(issues))
			}
		}
	} else {
		for key, months := range list {
			for month, issues := range months {
				for _, issue := range issues {
					duration := issue.Close.Sub(issue.Create)
					day := duration.Hours() / 24
					fmt.Printf("%d,%d,%s,%s,%d\n", key, month, issue.Title, issue.URL, int(day))
				}
			}
		}
	}
}